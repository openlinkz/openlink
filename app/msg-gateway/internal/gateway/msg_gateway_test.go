package gateway

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/openlinkz/openlink/api/msg-gateway"
	"math/rand"
	"testing"
)

func TestWebsocketProtocol(t *testing.T) {
	v := randStr(1024 * 5)

	payload := map[string]string{
		"key": v,
	}
	payloadContent, _ := jsoniter.Marshal(payload)

	fmt.Printf("payload len: %d\n", len(payloadContent))

	protoMsg := &msg_gateway.Protocol{
		Type:    "message",
		Payload: payloadContent,
	}

	protoMsgContent, _ := jsoniter.Marshal(protoMsg)

	fmt.Printf("payload len: %d\n", len(protoMsgContent))
	fmt.Println(string(protoMsgContent))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
