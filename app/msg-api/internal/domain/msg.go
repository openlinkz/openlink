package domain

import "context"

type MsgRepo interface {
	SendMsg(ctx context.Context, msg *Msg) error
}

type Msg struct {
	Type    string
	Payload []byte
}
