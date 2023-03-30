package biz

import (
	"context"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

func NewMsgExchangeBiz(msgRepo domain.MsgRepo) *MsgExchangeBiz {
	return &MsgExchangeBiz{msgRepo: msgRepo}
}

type MsgExchangeBiz struct {
	msgRepo domain.MsgRepo
}

func (biz *MsgExchangeBiz) Send(ctx context.Context, msg *domain.Msg) error {
	return biz.msgRepo.SendMsg(ctx, msg)
}
