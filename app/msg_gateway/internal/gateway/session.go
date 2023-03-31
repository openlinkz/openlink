package gateway

import (
	"github.com/gorilla/websocket"
	"github.com/openlinkz/openlink/api/protocol"
	"sync"
)

type Session struct {
	*websocket.Conn
	mu       sync.Mutex
	ServerIP string
	SID      string
	UID      string
	Platform protocol.Platform
}

func (sess *Session) Close() error {
	return sess.Close()
}
