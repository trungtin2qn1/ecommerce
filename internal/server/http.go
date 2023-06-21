package server

import (
	"context"
	auth "ecommerce/api/v1/auth"
	buyer "ecommerce/api/v1/buyer"
	"ecommerce/internal/conf"
	v1 "ecommerce/internal/service/v1"
	"ecommerce/internal/utils"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

func AuthListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return strings.HasPrefix(operation, "/api.v1.auth")
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	ac *conf.Auth,
	authService *v1.AuthService,
	buyerService *v1.BuyerService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(
					func(token *jwt2.Token) (interface{}, error) {
						return []byte(ac.JwtKey), nil
					},
					jwt.WithClaims(func() jwt2.Claims {
						return &utils.Claims{}
					}),
				),
			).
				Match(AuthListMatcher()).
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	auth.RegisterAuthHTTPServer(srv, authService)
	buyer.RegisterBuyerHTTPServer(srv, buyerService)
	return srv
}
