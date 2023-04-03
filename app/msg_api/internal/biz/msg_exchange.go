package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

func NewMsgExchangeBiz(msgRepo domain.MsgRepo) *MsgExchangeBiz {
	return &MsgExchangeBiz{msgRepo: msgRepo}
}

type MsgExchangeBiz struct {
	msgRepo domain.MsgRepo
}

func (biz *MsgExchangeBiz) Send(ctx context.Context, msg *domain.Message) error {
	log.Infof("revc msg: %s", msg.String())
	return biz.msgRepo.SendMsg(ctx, msg)
}
