package gateway

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/openlinkz/openlink/api/msg_api"
	"github.com/openlinkz/openlink/api/msg_gateway"
	"github.com/openlinkz/openlink/api/protocol"
	metrics "github.com/openlinkz/openlink/app/msg_gateway/internal/metric"
	"github.com/openlinkz/openlink/app/msg_gateway/internal/socketid"
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
			_, _ = w.Write([]byte("websocket no UID"))
			return
		}
		sid, _ := gateway.sidGenerator.NextSid()

		session := &Session{
			Conn:     conn,
			SID:      sid,
			UID:      uid,
			Platform: protocol.Platform(platform),
		}

		// 用户连接
		if err = gateway.connectSession(r.Context(), session); err != nil {
			// todo
			return
		}
		// 用户断开
		defer func() { _ = gateway.disconnectSession(r.Context(), session) }()

		// proto 序列化
		codec := encoding.GetCodec(proto.Name)

		for {
			_, data, err := session.ReadMessage()
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
				if err = gateway.keepAliveSession(r.Context(), session); err != nil {
					log.Errorf("conn keepalive failed. SID: %s. UID: %s", session.SID, session.UID)
					continue
				}
			}

			if err = gateway.SendMsg(r.Context(), &msg_api.Msg{Server: session.ServerIP, SID: session.SID, Protocol: protoMsg}); err != nil {
				log.Errorf("send msg err: %s", err.Error())
				continue
			}
		}
	}
}

func (gateway *MsgGateway) connectSession(ctx context.Context, session *Session) (err error) {
	if _, err = gateway.msgExchangeClient.Connect(ctx, &msg_api.ConnectRequest{
		Server:   session.ServerIP,
		SID:      session.SID,
		UID:      session.UID,
		Platform: session.Platform.String(),
	}); err != nil {
		return
	}
	log.Infof("user connected\n")

	gateway.sidConnMapping.Add(session.SID, session)
	gateway.userConnMapping.Add(session.UID, session)

	metrics.GatewayOnlineTotals.Inc()
	return
}

func (gateway *MsgGateway) disconnectSession(ctx context.Context, session *Session) (err error) {
	_ = session.Close()

	metrics.GatewayOnlineTotals.Dec()

	gateway.sidConnMapping.Delete(session.SID)
	gateway.userConnMapping.Delete(session.UID, session)

	// 用户断开连接
	_, err = gateway.msgExchangeClient.Disconnect(ctx, &msg_api.DisconnectRequest{Server: gateway.serverIP, SID: session.SID})
	if err != nil {
		log.Errorf("conn closed. SID: %d\n", session.SID)
		return
	}
	log.Infof("conn closed. SID: %d\n", session.SID)
	return
}

func (gateway *MsgGateway) keepAliveSession(ctx context.Context, session *Session) (err error) {
	_, err = gateway.msgExchangeClient.KeepAlive(ctx, &msg_api.KeepAliveRequest{Server: gateway.serverIP, SID: session.SID, UID: session.UID})
	return
}

func (gateway *MsgGateway) SendMsg(ctx context.Context, msg *msg_api.Msg) (err error) {
	_, err = gateway.msgExchangeClient.SendMsg(ctx, msg)
	return
}

func (gateway *MsgGateway) PushMsg(ctx context.Context, pushMsgReq *msg_gateway.PushMsgReq) (*msg_gateway.PushMsgReply, error) {
	var (
		successTotal, failedTotal int32
		err                       error
	)
	for _, sid := range pushMsgReq.SIDs {
		if err = gateway.pushOne(ctx, sid, pushMsgReq.Protocol); err != nil {
			failedTotal++
			continue
		}
		successTotal++
	}
	return &msg_gateway.PushMsgReply{SuccessTotal: successTotal, FailedTotal: failedTotal}, nil
}

func (gateway *MsgGateway) pushOne(ctx context.Context, sid string, protocol *protocol.Protocol) (err error) {
	session := gateway.sidConnMapping.Get(sid)
	if session == nil {
		err = fmt.Errorf("sid not found")
		return
	}
	body, err := encoding.GetCodec("proto").Marshal(protocol)
	if err != nil {
		err = fmt.Errorf("proto marshal failed")
		return
	}
	err = session.WriteMessage(websocket.TextMessage, body)
	return
}
