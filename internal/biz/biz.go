package biz

import (
	v1 "ecommerce/internal/biz/v1"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(v1.NewAuthUsecase, v1.NewBuyerUsecase)
