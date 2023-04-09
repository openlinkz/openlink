package domain

import (
	"context"
	jsoniter "github.com/json-iterator/go"
)

type MsgRepo interface {
	SendMsg(ctx context.Context, msg *Msg) error
	SendBatchMsg(ctx context.Context, msg []*Msg) error
}

type (
	Msg struct {
		ServerIP string
		SID      string
		UID      string
		Platform Platform
		Body     MsgBody
	}

	BizMsg struct {
		UID      string
		Platform Platform
		Body     MsgBody
	}

	MsgBody struct {
		Type    string
		Payload []byte
	}
)

func (msg *Msg) BizTopic() string {
	return msg.Body.Type
}

func (msg *Msg) BizKey() string {
	return msg.SID
}

func (msg *Msg) String() string {
	body, _ := jsoniter.Marshal(msg)
	return string(body)
}
