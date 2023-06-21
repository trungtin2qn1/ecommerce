package v1

import (
	"context"
	"database/sql"
	"fmt"

	"ecommerce/api/errors_reason"
	v1 "ecommerce/api/v1/auth"
	"ecommerce/internal/conf"
	"ecommerce/internal/utils"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lib/pq"
)

// BuyerUsecase is a Buyer usecase.
type AuthUsecase struct {
	config     *conf.Auth
	buyerRepo  BuyerRepo
	sellerRepo SellerRepo
	log        *log.Helper
}

func NewAuthUsecase(buyerRepo BuyerRepo, sellerRepo SellerRepo, conf *conf.Auth, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{buyerRepo: buyerRepo, sellerRepo: sellerRepo, config: conf, log: log.NewHelper(logger)}
}

// CreateBuyer creates a Buyer, and returns the new Buyer.
func (uc *AuthUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	uc.log.WithContext(ctx).Infof("Register email: %s", req.Email)

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.InternalServer(errors_reason.ErrorReason_INTERNAL_SERVER_ERROR.String(), "cannot hash password")
	}

	var idUser int64
	switch req.Type {
	case 1:
		idUser, err = uc.buyerRepo.New(ctx, req.Email, password)
	case 2:
		idUser, err = uc.sellerRepo.New(ctx, req.Email, password)
	default:
		return nil, errors.BadRequest(errors_reason.ErrorReason_BAD_REQUEST.String(), fmt.Sprintf("invalid request type: %d", req.Type))
	}

	pgErr, ok := err.(*pq.Error)
	if ok {
		if pgErr.Code == utils.ErrUniqueViolationCode {
			return nil, errors.Unauthorized(errors_reason.ErrorReason_BAD_REQUEST.String(), fmt.Sprintf("email %s has been registered", req.Email))
		}
	}
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := uc.generateAccessTokens(idUser)
	if err != nil {
		return nil, errors.InternalServer(errors_reason.ErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf("Fail to generate access token %v", err))
	}

	return &v1.RegisterReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GetBuyerByIdUser get buyer info by id user.
func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	uc.log.WithContext(ctx).Infof("Login email: %s", req.Email)
	var err error
	var idUser int64
	var password string

	switch req.Type {
	case 1:
		idUser, password, err = uc.buyerRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
	case 2:
		idUser, password, err = uc.sellerRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.BadRequest(errors_reason.ErrorReason_BAD_REQUEST.String(), fmt.Sprintf("invalid request type: %d", req.Type))
	}

	if err == sql.ErrNoRows || idUser == 0 {
		return nil, errors.NotFound(errors_reason.ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf("User with email %s has not registered yet", req.Email))
	}
	if !utils.CheckPasswordHash(req.Password, password) {
		return nil, errors.Unauthorized(errors_reason.ErrorReason_UNAUTHORIZED.String(), "wrong password")
	}

	accessToken, refreshToken, err := uc.generateAccessTokens(idUser)
	if err != nil {
		return nil, errors.InternalServer(errors_reason.ErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf("Fail to generate access token %v", err))
	}

	return &v1.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
