//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/config"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/gateway"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/server"
)

func initApp(*config.Server, *etcd.Registry, *gateway.MsgGateway, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
