// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/openlinkz/openlink/app/msg_api/internal/biz"
	"github.com/openlinkz/openlink/app/msg_api/internal/config"
	"github.com/openlinkz/openlink/app/msg_api/internal/data"
	"github.com/openlinkz/openlink/app/msg_api/internal/server"
	"github.com/openlinkz/openlink/app/msg_api/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func initApp(configServer *config.Server, registry *config.Registry, redis *config.Redis, kafka *config.Kafka, logger log.Logger) (*kratos.App, error) {
	syncProducer := data.NewKafkaSyncProducer(kafka)
	msgRepo := data.NewMsgRepo(syncProducer)
	msgExchangeBiz := biz.NewMsgExchangeBiz(msgRepo)
	client := data.NewRedisClient(redis)
	userStatusRepo := data.NewUserStatusRepo(client)
	userStatusBiz := biz.NewUserStatusBiz(userStatusRepo)
	msgExchangeService := service.NewMsgExchangeService(msgExchangeBiz, userStatusBiz)
	v := server.NewServers(configServer, msgExchangeService)
	registrar := server.NewRegistry(registry)
	app := newApp(logger, v, registrar)
	return app, nil
}
