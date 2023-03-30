//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/openlinkz/openlink/app/msg_api/internal/biz"
	"github.com/openlinkz/openlink/app/msg_api/internal/config"
	"github.com/openlinkz/openlink/app/msg_api/internal/data"
	"github.com/openlinkz/openlink/app/msg_api/internal/server"
	"github.com/openlinkz/openlink/app/msg_api/internal/service"
)

func initApp(*config.Server, *config.Registry, log.Logger) (*kratos.App, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
