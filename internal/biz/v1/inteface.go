package v1

import (
	"context"
)

type BuyerRepo interface {
	New(context.Context, string, string) (int64, error)
	GetByEmail(context.Context, string) (int64, string, error)
}

type SellerRepo interface {
	New(context.Context, string, string) (int64, error)
	GetByEmail(context.Context, string) (int64, string, error)
}
