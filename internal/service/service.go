package service

import (
	v1 "ecommerce/internal/service/v1"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(v1.NewAuthService, v1.NewBuyerService)
