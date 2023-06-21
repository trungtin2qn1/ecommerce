package v1

import (
	"context"

	pb "ecommerce/api/v1/auth"
	biz "ecommerce/internal/biz/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc  *biz.AuthUsecase
	log *log.Helper
}

func NewAuthService(uc *biz.AuthUsecase, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return s.uc.Login(ctx, req)
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return s.uc.Register(ctx, req)
}
