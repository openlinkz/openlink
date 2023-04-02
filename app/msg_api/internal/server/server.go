package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/openlinkz/openlink/api/msg_api"
	"github.com/openlinkz/openlink/app/msg_api/internal/config"
	"github.com/openlinkz/openlink/app/msg_api/internal/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewServers, NewRegistry)

func NewServers(conf *config.Server, msgExchangeService *service.MsgExchangeService) []transport.Server {
	httpServer := conf.Http.BuildHTTPServer()
	_PromHTTP_Metrics_Handler(httpServer)
	msg_api.RegisterMsgExchangeServiceHTTPServer(httpServer, msgExchangeService)

	grpcServer := conf.Grpc.BuildGRPCServer()
	msg_api.RegisterMsgExchangeServiceServer(grpcServer, msgExchangeService)
	//channelzservice.RegisterChannelzServiceToServer(grpcServer)
	return []transport.Server{httpServer, grpcServer}
}

func NewRegistry(conf *config.Registry) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{conf.Etcd.Addr}, DialTimeout: conf.Etcd.DialTimeout.AsDuration()})
	if err != nil {
		panic(err)
	}
	// todo connect async. check client healthy before use.
	return etcd.New(client)
}

// 注册 prometheus metrics 路由
func _PromHTTP_Metrics_Handler(srv *http.Server) {
	srv.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true}))
}
