package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/openlinkz/openlink/api/msg_gateway"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/config"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/gateway"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewRegistry, NewServers)

func NewServers(conf *config.Server, gateway *gateway.MsgGateway, _ log.Logger) []transport.Server {
	httpServer := conf.Http.BuildHTTPServer()
	msg_gateway.RegisterMsgGatewayServiceHTTPServer(httpServer, gateway)
	_WebsocketGateway_Connect_Handler(gateway, httpServer)
	_PromHTTP_Metrics_Handler(httpServer)

	grpcServer := conf.Grpc.BuildGRPCServer()
	msg_gateway.RegisterMsgGatewayServiceServer(grpcServer, gateway)
	return []transport.Server{httpServer, grpcServer}
}

func NewETCDRegistry(conf *config.Registry) *etcd.Registry {
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{conf.Etcd.Addr}, DialTimeout: conf.Etcd.DialTimeout.AsDuration()})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

func NewRegistry(r *etcd.Registry) registry.Registrar {
	return r
}

// 注册 http websocket 路由
func _WebsocketGateway_Connect_Handler(gateway *gateway.MsgGateway, srv *http.Server) {
	srv.Handle("/gateway/connect", gateway.WebsocketConnectHandler())
}

// 注册 prometheus metrics 路由
func _PromHTTP_Metrics_Handler(srv *http.Server) {
	srv.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true}))
}
