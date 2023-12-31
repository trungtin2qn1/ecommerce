// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ecommerce/internal/biz/v1"
	"ecommerce/internal/conf"
	"ecommerce/internal/data"
	"ecommerce/internal/server"
	v1_2 "ecommerce/internal/service/v1"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	buyerRepo := data.NewBuyerRepo(dataData, logger)
	sellerRepo := data.NewSellerRepo(dataData, logger)
	authUsecase := v1.NewAuthUsecase(buyerRepo, sellerRepo, auth, logger)
	authService := v1_2.NewAuthService(authUsecase, logger)
	buyerUsecase := v1.NewBuyerUsecase(buyerRepo, logger)
	buyerService := v1_2.NewBuyerService(buyerUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, auth, authService, buyerService, logger)
	httpServer := server.NewHTTPServer(confServer, auth, authService, buyerService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
