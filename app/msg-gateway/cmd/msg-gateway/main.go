package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/openlinkz/openlink/api/msg_api"
	appconfig "github.com/openlinkz/openlink/app/msg-gateway/internal/config"
	"github.com/openlinkz/openlink/app/msg-gateway/internal/gateway"
	"github.com/openlinkz/openlink/app/msg-gateway/internal/server"
	"github.com/openlinkz/openlink/pkg/client"
	"github.com/openlinkz/openlink/pkg/env"
	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	_ "net/http/pprof"
)

var (
	configPath = flag.String("config", "", "config file path of project")
)

func newApp(logger log.Logger, svrs []transport.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(env.AppName),
		kratos.Version(env.AppVersion),
		kratos.Logger(logger),
		kratos.Server(svrs...),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	c := config.New(config.WithSource(file.NewSource(*configPath)))
	if err := c.Load(); err != nil {
		panic(errors.WithStack(err))
	}
	var bc appconfig.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(errors.WithStack(err))
	}

	r := server.NewETCDRegistry(bc.Registry)

	wsAPIClient, err := bc.Client.WsAPI.BuildGRPCClient(context.Background(), client.GRPCOptionClientOptions(grpc.WithDiscovery(r)))
	if err != nil {
		panic(errors.WithStack(err))
	}
	exc := msg_api.NewMsgExchangeServiceClient(wsAPIClient)

	websocketGateway := gateway.NewMsgGateway(
		gateway.MsgGatewayOptionMsgExchangeServiceClient(exc),
	)

	app, err := initApp(bc.Server, r, websocketGateway, log.DefaultLogger)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
