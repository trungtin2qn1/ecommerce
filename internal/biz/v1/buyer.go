package v1

import (
	"context"

	pb "ecommerce/api/v1/buyer"

	"github.com/go-kratos/kratos/v2/log"
)

type BuyerUsecase struct {
	buyerRepo BuyerRepo
	log       *log.Helper
}

func NewBuyerUsecase(buyerRepo BuyerRepo, logger log.Logger) *BuyerUsecase {
	return &BuyerUsecase{buyerRepo: buyerRepo, log: log.NewHelper(logger)}
}

func (uc *BuyerUsecase) Self(ctx context.Context, idBuyer int64) (*pb.SelfReply, error) {
	return &pb.SelfReply{}, nil
}
