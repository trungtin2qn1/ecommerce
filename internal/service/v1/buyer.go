package v1

import (
	"context"

	pb "ecommerce/api/v1/buyer"
	biz "ecommerce/internal/biz/v1"
	"ecommerce/internal/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type BuyerService struct {
	pb.UnimplementedBuyerServer

	uc  *biz.BuyerUsecase
	log *log.Helper
}

func NewBuyerService(uc *biz.BuyerUsecase, logger log.Logger) *BuyerService {
	return &BuyerService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BuyerService) Self(ctx context.Context, req *pb.SelfRequest) (*pb.SelfReply, error) {
	idUser, err := utils.GetIdUserByJwt(ctx)
	if err != nil {
		return nil, err
	}
	return s.uc.Self(ctx, idUser)
}
