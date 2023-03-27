package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/openlinkz/openlink/app/msg-api/internal/domain"
)

var _ domain.MsgRepo = (*msgKafkaMQRepo)(nil)

func NewMsgRepo() domain.MsgRepo {
	return &msgKafkaMQRepo{}
}

type msgKafkaMQRepo struct {
	rdb *redis.Client
}

func (m *msgKafkaMQRepo) SendMsg(ctx context.Context, msg *domain.Msg) error {
	return nil
}
