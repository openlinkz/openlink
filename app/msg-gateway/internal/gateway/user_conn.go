package gateway

import (
	"github.com/gorilla/websocket"
	"github.com/openlinkz/openlink/api/protocol"
	"sync"
)

type UserConnection struct {
	mu       sync.Mutex
	conn     *websocket.Conn
	sid      string
	uid      string
	platform protocol.Platform
}
