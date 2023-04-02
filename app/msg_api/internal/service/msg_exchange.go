package service

import (
	"context"
	"github.com/google/wire"
	"github.com/openlinkz/openlink/api/msg_api"
	"github.com/openlinkz/openlink/app/msg_api/internal/biz"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
	"time"
)

var ProviderSet = wire.NewSet(NewMsgExchangeService)

func NewMsgExchangeService(msgExchangeBiz *biz.MsgExchangeBiz, userStatusBiz *biz.UserStatusBiz) *MsgExchangeService {
	svc := &MsgExchangeService{msgExchangeBiz: msgExchangeBiz, userStatusBiz: userStatusBiz}
	return svc
}

type MsgExchangeService struct {
	msg_api.UnimplementedMsgExchangeServiceServer
	msgExchangeBiz *biz.MsgExchangeBiz
	userStatusBiz  *biz.UserStatusBiz
}

func (svc *MsgExchangeService) SendMsg(ctx context.Context, msg *msg_api.Msg) (*msg_api.SendMsgReply, error) {
	return nil, svc.msgExchangeBiz.Send(ctx, &domain.Message{
		ServerIP: msg.Server,
		SID:      msg.SID,
		UID:      msg.UID,
		Platform: domain.PlatformFrom(int(msg.Platform)),
		Body:     domain.MsgBody{Type: msg.Protocol.Type, Payload: msg.Protocol.Payload},
	})
}

func (svc *MsgExchangeService) Connect(ctx context.Context, req *msg_api.ConnectRequest) (*msg_api.ConnectReply, error) {
	err := svc.userStatusBiz.Connect(ctx, &domain.UserStatus{
		Server:         req.Server,
		SID:            req.SID,
		UID:            req.UID,
		Platform:       domain.PlatformFrom(int(req.Platform.Number())),
		ExpireDuration: 60 * time.Minute,
	})
	return &msg_api.ConnectReply{}, err
}

func (svc *MsgExchangeService) Disconnect(ctx context.Context, req *msg_api.DisconnectRequest) (*msg_api.DisconnectReply, error) {
	err := svc.userStatusBiz.Disconnect(ctx, &domain.UserStatus{
		Server:   req.Server,
		SID:      req.SID,
		UID:      req.UID,
		Platform: domain.PlatformFrom(int(req.Platform.Number())),
	})
	return &msg_api.DisconnectReply{}, err
}

func (svc *MsgExchangeService) KeepAlive(ctx context.Context, req *msg_api.KeepAliveRequest) (*msg_api.KeepAliveReply, error) {
	err := svc.userStatusBiz.KeepAlive(ctx, &domain.UserStatus{
		Server:         req.Server,
		SID:            req.SID,
		UID:            req.UID,
		Platform:       domain.PlatformFrom(int(req.Platform.Number())),
		ExpireDuration: 0,
	})
	return &msg_api.KeepAliveReply{}, err
}
