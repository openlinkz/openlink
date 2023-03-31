package gateway

import (
	"container/list"
	"github.com/openlinkz/openlink/api/protocol"
	"sync"
)

func NewUserConnMapping() *UserConnMapping {
	return &UserConnMapping{userConns: make(map[string]*list.List)}
}

type UserConnMapping struct {
	userConns map[string]*list.List

	mu sync.RWMutex
}

func (m *UserConnMapping) Add(uid string, conn *Session) {
	m.mu.Lock()
	if m.userConns[uid] == nil {
		m.userConns[uid] = list.New()
	}
	m.userConns[uid].PushFront(conn)
	m.mu.Unlock()
}

func (m *UserConnMapping) Delete(uid string, conn *Session) {
	// todo
	m.mu.Lock()
	conns, has := m.userConns[uid]
	if !has {
		return
	}
	conns.Remove(&list.Element{Value: conn})
	m.mu.Unlock()
}

func (m *UserConnMapping) Get(uid string, platform protocol.Platform) (userConns []*Session) {
	m.mu.RLock()
	conns, has := m.userConns[uid]
	m.mu.RUnlock()
	if !has {
		return
	}
	for i := conns.Front(); i != nil; i = i.Next() {
		uc := i.Value.(*Session)
		if uc.Platform == platform {
			userConns = append(userConns, uc)
		}
	}
	return
}

func NewSIDConnMapping() *SidConnMapping {
	return &SidConnMapping{sidConn: make(map[string]*Session)}
}

type SidConnMapping struct {
	sidConn map[string]*Session

	mu sync.RWMutex
}

func (mapping *SidConnMapping) Add(sid string, userConn *Session) {
	mapping.mu.Lock()
	mapping.sidConn[sid] = userConn
	mapping.mu.Unlock()
}

func (mapping *SidConnMapping) Delete(sid string) {
	mapping.mu.Lock()
	delete(mapping.sidConn, sid)
	mapping.mu.Unlock()
}

func (mapping *SidConnMapping) Get(sid string) *Session {
	mapping.mu.RLock()
	conn, _ := mapping.sidConn[sid]
	mapping.mu.RUnlock()
	return conn
}
