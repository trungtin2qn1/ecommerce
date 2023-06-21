package server

import (
	auth "ecommerce/api/v1/auth"
	buyer "ecommerce/api/v1/buyer"
	"ecommerce/internal/conf"
	v1 "ecommerce/internal/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	ac *conf.Auth,
	authService *v1.AuthService,
	buyerService *v1.BuyerService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.JwtKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
					return &jwt2.MapClaims{}
				})),
			).
				Regex(`/v[1-9]+/auth`).
				Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	auth.RegisterAuthServer(srv, authService)
	buyer.RegisterBuyerServer(srv, buyerService)
	return srv
}
