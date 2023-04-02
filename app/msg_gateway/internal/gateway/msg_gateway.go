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
	"github.com/openlinkz/openlink/pkg/encoding/json"
	"github.com/openlinkz/openlink/pkg/netutil"
	"github.com/spf13/cast"
	"net/http"
	"sync"
)

const _defaultMaxConnection = 50000

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
		sidGenerator:    socketid.NewUUIDGenerator(),
		upgrader:        &websocket.Upgrader{},
		sidConnMapping:  NewSIDConnMapping(),
		userConnMapping: NewUserConnMapping(),
		serverIP:        netutil.LocalIPString(),
		maxConnection:   _defaultMaxConnection,
	}

	for _, opt := range opts {
		opt(gateway)
	}

	return gateway
}

type MsgGateway struct {
	msg_gateway.UnimplementedMsgGatewayServiceServer
	msgExchangeClient msg_api.MsgExchangeServiceClient
	upgrader          *websocket.Upgrader
	sidGenerator      socketid.Generator
	sidConnMapping    *SidConnMapping
	userConnMapping   *UserConnMapping
	serverIP          string
	maxConnection     int

	mu sync.Mutex
}

func (gateway *MsgGateway) WebsocketConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := gateway.upgrader.Upgrade(w, r, nil)
		if err != nil {
			renderJSON(w, &ConnectResponse{Code: -1, Message: err.Error()})
			return
		}
		uid, platform, err := gateway.parseMetadataFromRequest(r)
		if err != nil {
			renderJSON(w, &ConnectResponse{Code: -1, Message: err.Error()})
			return
		}
		sid, _ := gateway.sidGenerator.NextSid()

		session := &Session{
			Conn:     conn,
			SID:      sid,
			UID:      uid,
			Platform: platform,
		}

		// 用户连接
		if err = gateway.connectSession(r.Context(), session); err != nil {
			renderJSON(w, &ConnectResponse{Code: -1, Message: err.Error()})
			return
		}
		// 用户断开
		defer func() { _ = gateway.disconnectSession(r.Context(), session) }()

		// proto 序列化
		codec := encoding.GetCodec(json.Name)

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
		Platform: session.Platform,
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
	log.Infof("conn closed. SID: %s\n", session.SID)
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

func (gateway *MsgGateway) PushBatchMsg(ctx context.Context, pushMsgReq *msg_gateway.PushBatchMsgReq) (*msg_gateway.PushBatchMsgReply, error) {
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
	return &msg_gateway.PushBatchMsgReply{SuccessTotal: successTotal, FailedTotal: failedTotal}, nil
}

func (gateway *MsgGateway) PushMsg(ctx context.Context, pushMsgReq *msg_gateway.PushMsgReq) (*msg_gateway.PushMsgReply, error) {
	if err := gateway.pushOne(ctx, pushMsgReq.SID, pushMsgReq.Protocol); err != nil {
		return nil, err
	}
	return &msg_gateway.PushMsgReply{Success: true}, nil
}

func (gateway *MsgGateway) pushOne(_ context.Context, sid string, protocol *protocol.Protocol) (err error) {
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

func (gateway *MsgGateway) parseMetadataFromRequest(r *http.Request) (string, protocol.Platform, error) {
	_ = r.ParseForm()
	uid := r.Form.Get("uid")
	if uid == "" {
		return "", 0, fmt.Errorf("uid must not empty")
	}

	platform, err := cast.ToInt32E(r.Form.Get("platform"))
	if err != nil {
		return "", 0, fmt.Errorf("platform is invalid. %s", err.Error())
	}
	return uid, protocol.Platform(platform), nil
}
