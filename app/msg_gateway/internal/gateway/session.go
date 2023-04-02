package gateway

import (
	"github.com/gorilla/websocket"
	"github.com/openlinkz/openlink/api/protocol"
	"github.com/pkg/errors"
	"sync"
	"sync/atomic"
)

type Session struct {
	*websocket.Conn
	mu       sync.Mutex
	closed   atomic.Int32
	ServerIP string
	SID      string
	UID      string
	Platform protocol.Platform
}

func (sess *Session) Close() error {
	if !sess.closed.CompareAndSwap(0, 1) {
		return errors.New("session has closed")
	}
	return sess.Conn.Close()
}
