package service

import (
	"context"
	"github.com/openlinkz/openlink/api/msg_api"
	"github.com/openlinkz/openlink/app/msg_api/internal/biz"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

func NewMsgAPI(msgAPIBiz *biz.MsgAPIBiz) *MsgAPIService {
	return &MsgAPIService{msgAPIBiz: msgAPIBiz}
}

type MsgAPIService struct {
	msg_api.UnimplementedMsgAPIServiceServer
	msgAPIBiz *biz.MsgAPIBiz
}

func (api *MsgAPIService) PushMsg(ctx context.Context, bizMsg *msg_api.BizMsg) (*msg_api.PushBizMsgReply, error) {
	err := api.msgAPIBiz.PushMsg(ctx, &domain.BizMsg{
		UID:      bizMsg.Uid,
		Platform: domain.PlatformFrom(int(bizMsg.Platform.Number())),
		Body:     domain.MsgBody{Payload: bizMsg.Protocol.Payload, Type: bizMsg.Protocol.Type},
	})
	if err != nil {
		return nil, err
	}
	return &msg_api.PushBizMsgReply{}, nil
}
