package gateway

import (
	"context"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	msg_api "github.com/openlinkz/openlink/api/msg_api"
	msg_gateway "github.com/openlinkz/openlink/api/msg_gateway"
	"github.com/openlinkz/openlink/api/protocol"
	metrics "github.com/openlinkz/openlink/app/msg-gateway/internal/metric"
	"github.com/openlinkz/openlink/app/msg-gateway/internal/socketid"
	"github.com/openlinkz/openlink/pkg/encoding/proto"
	"github.com/openlinkz/openlink/pkg/netutil"
	"github.com/spf13/cast"
	"net/http"
	"sync"
)

type MsgGatewayOption func(*MsgGateway)

func MsgGatewayOptionUpgrader(upgrader *websocket.Upgrader) MsgGatewayOption {
	return func(gateway *MsgGateway) { gateway.upgrader = upgrader }
}

func MsgGatewayOptionSidGenerator(g socketid.Generator) MsgGatewayOption {
	return func(gateway *MsgGateway) { gateway.sidGenerator = g }
}

func MsgGatewayOptionMsgExchangeServiceClient(c msg_api.MsgExchangeServiceClient) MsgGatewayOption {
	return func(gateway *MsgGateway) { gateway.msgExchangeClient = c }
}

func NewMsgGateway(opts ...MsgGatewayOption) *MsgGateway {
	gateway := &MsgGateway{
		sidGenerator:    socketid.NewAtomicGenerator(0),
		upgrader:        &websocket.Upgrader{},
		sidConnMapping:  NewSIDConnMapping(),
		userConnMapping: NewUserConnMapping(),
		serverIP:        netutil.LocalIPString(),
	}

	for _, opt := range opts {
		opt(gateway)
	}

	return gateway
}

type MsgGateway struct {
	msgExchangeClient msg_api.MsgExchangeServiceClient
	upgrader          *websocket.Upgrader
	sidGenerator      socketid.Generator
	sidConnMapping    *SidConnMapping
	userConnMapping   *UserConnMapping
	serverIP          string

	mu sync.Mutex
}

func (gateway *MsgGateway) WebsocketConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := gateway.upgrader.Upgrade(w, r, nil)
		if err != nil {
			// todo
			_, _ = w.Write([]byte("websocket upgrade err: " + err.Error()))
			return
		}

		uid := r.Header.Get("uid")
		platform := cast.ToInt32(r.Header.Get("platform"))
		if uid == "" {
			// todo
			_, _ = w.Write([]byte("websocket no uid"))
			return
		}
		sid, _ := gateway.sidGenerator.NextSid()

		userConn := &UserConnection{
			conn:     conn,
			sid:      sid,
			uid:      uid,
			platform: protocol.Platform(platform),
		}

		gateway.sidConnMapping.Add(sid, userConn)
		gateway.userConnMapping.Add(uid, userConn)

		// 用户连接
		_, err = gateway.msgExchangeClient.Connect(r.Context(), &msg_api.ConnectRequest{
			Server:   gateway.serverIP,
			SID:      sid,
			UID:      uid,
			Platform: protocol.Platform_name[platform],
		})
		if err != nil {
			// todo
			return
		}
		log.Infof("user connected\n")

		metrics.GatewayOnlineTotals.Inc()
		clear := func() {
			metrics.GatewayOnlineTotals.Dec()

			gateway.sidConnMapping.Delete(sid)
			gateway.userConnMapping.Delete(uid, userConn)

			// 用户断开连接
			_, _ = gateway.msgExchangeClient.Disconnect(r.Context(), &msg_api.DisconnectRequest{Server: gateway.serverIP, SID: sid})
			_ = conn.Close()

			log.Infof("conn closed. sid: %d\n", sid)
		}
		defer clear()

		log.Infof("conn connect. sid: %d", sid)
		// proto 序列化
		codec := encoding.GetCodec(proto.Name)

		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("read message err: %s", err.Error())
				break
			}

			metrics.GatewayInputBytes.Add(float64(len(data)))

			protoMsg := &protocol.Protocol{}
			if err := codec.Unmarshal(data, &protoMsg); err != nil {
				log.Errorf("unmarshal protocol err: %s", err.Error())
				continue
			}

			if protoMsg.Type == protocol.Type_name[int32(protocol.Type_HEARTBEAT)] {
				_, _ = gateway.msgExchangeClient.KeepAlive(context.Background(), &msg_api.KeepAliveRequest{Server: gateway.serverIP, SID: sid, UID: ""})
			}

			if _, err = gateway.msgExchangeClient.SendMsg(context.Background(), &msg_api.Msg{
				Server:   gateway.serverIP,
				SID:      sid,
				Protocol: protoMsg,
			}); err != nil {
				log.Errorf("send msghandler err: %s", err.Error())
			}
		}
	}
}

func (gateway *MsgGateway) PushMsg(ctx context.Context, protocol *msg_gateway.PushMsgReq) (*msg_gateway.PushMsgReply, error) {
	return nil, nil
}
