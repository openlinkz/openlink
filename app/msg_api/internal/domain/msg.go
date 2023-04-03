package domain

import (
	"context"
	jsoniter "github.com/json-iterator/go"
)

type MsgRepo interface {
	SendMsg(ctx context.Context, msg *Message) error
	SendBatchMsg(ctx context.Context, msg []*Message) error
}

type MsgBody struct {
	Type    string
	Payload []byte
}

type Message struct {
	ServerIP string
	SID      string
	UID      string
	Platform Platform
	Body     MsgBody
}

func (msg *Message) BizTopic() string {
	return msg.Body.Type
}

func (msg *Message) BizKey() string {
	return msg.SID
}

func (msg *Message) String() string {
	body, _ := jsoniter.Marshal(msg)
	return string(body)
}
