package gateway

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/openlinkz/openlink/api/protocol"
	"github.com/stretchr/testify/require"
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

	protoMsg := &protocol.Protocol{
		Type:    "message",
		Payload: payloadContent,
	}

	protoMsgContent, _ := jsoniter.Marshal(protoMsg)

	fmt.Printf("payload len: %d\n", len(protoMsgContent))
	fmt.Println(string(protoMsgContent))
}

func TestHeartbeatMessage(t *testing.T) {
	v := randStr(1024 * 5)
	payload := map[string]string{
		"key": v,
	}
	payloadContent, _ := jsoniter.Marshal(payload)

	protoMsg := &protocol.Protocol{
		Type:    protocol.Type_HEARTBEAT.String(),
		Payload: payloadContent,
	}

	protoMsgContent, _ := jsoniter.Marshal(protoMsg)

	var p = new(protocol.Protocol)
	err := jsoniter.Unmarshal(protoMsgContent, &p)
	require.NoError(t, err)
	require.Equal(t, protoMsg.Type, protocol.Type_HEARTBEAT.String())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
