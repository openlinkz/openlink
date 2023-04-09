package biz

import (
	"context"
	"github.com/openlinkz/openlink/api/msg_gateway"
	"github.com/openlinkz/openlink/api/protocol"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

func NewMsgAPIBiz(gatewayClient msg_gateway.MsgGatewayServiceClient, usRepo domain.UserStatusRepo) *MsgAPIBiz {
	return &MsgAPIBiz{gatewayClient: gatewayClient, userStatusRepo: usRepo}
}

type MsgAPIBiz struct {
	gatewayClient  msg_gateway.MsgGatewayServiceClient
	userStatusRepo domain.UserStatusRepo
}

func (api *MsgAPIBiz) PushMsg(ctx context.Context, bizMsg *domain.BizMsg) error {
	sid, err := api.userStatusRepo.GetSID(ctx, bizMsg.UID, bizMsg.Platform)
	if err != nil {
		return err
	}
	_, err = api.gatewayClient.PushMsg(ctx, &msg_gateway.PushMsgReq{
		SID: sid,
		Protocol: &protocol.Protocol{
			Type:    bizMsg.Body.Type,
			Payload: bizMsg.Body.Payload,
		},
	})
	return err
}
