package gateway

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/openlinkz/openlink/api/msg_api"
	"github.com/openlinkz/openlink/api/protocol"
	"github.com/openlinkz/openlink/pkg/client"
	"github.com/stretchr/testify/require"
	"log"
	"net/url"
	"testing"
	"time"
)

func connectGateway(uid string, platform protocol.Platform) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:28080", Path: "/gateway/connect", RawQuery: fmt.Sprintf("uid=%s&platform=%d", uid, platform)}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func TestConnect(t *testing.T) {

}

func TestPushMsg(t *testing.T) {
	uid := "1111"
	platform := protocol.Platform_PC
	wsConn, err := connectGateway(uid, platform)

	go func() {
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				p := &protocol.Protocol{
					Type:    "IM",
					Payload: []byte("AAAAAAAAAAAAAAAAAAAAAA"),
				}
				body, _ := jsoniter.Marshal(p)
				if err := wsConn.WriteMessage(websocket.TextMessage, body); err != nil {
					log.Println("write:", err)
					return
				}
			}
		}
	}()

	msgAPIgRPCConn := client.DefaultConfig().WithAddr("localhost:18081").WithTimeout("1h").BuildGRPCClientMust(context.Background())
	msgAPIClient := msg_api.NewMsgAPIServiceClient(msgAPIgRPCConn)

	_, err = msgAPIClient.PushMsg(context.Background(), &msg_api.BizMsg{
		Uid:      uid,
		Platform: protocol.Platform_PC,
		Protocol: nil,
	})

	require.NoError(t, err)
}
