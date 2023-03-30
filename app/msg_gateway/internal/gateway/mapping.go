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

func (m *UserConnMapping) Add(uid string, conn *UserConnection) {
	m.mu.Lock()
	if m.userConns[uid] == nil {
		m.userConns[uid] = list.New()
	}
	m.userConns[uid].PushFront(conn)
	m.mu.Unlock()
}

func (m *UserConnMapping) Delete(uid string, conn *UserConnection) {
	// todo
	m.mu.Lock()
	conns, has := m.userConns[uid]
	if !has {
		return
	}
	conns.Remove(&list.Element{Value: conn})
	m.mu.Unlock()
}

func (m *UserConnMapping) Get(uid string, platform protocol.Platform) (userConns []*UserConnection) {
	m.mu.RLock()
	conns, has := m.userConns[uid]
	m.mu.RUnlock()
	if !has {
		return
	}
	for i := conns.Front(); i != nil; i = i.Next() {
		uc := i.Value.(*UserConnection)
		if uc.platform == platform {
			userConns = append(userConns, uc)
		}
	}
	return
}

func NewSIDConnMapping() *SidConnMapping {
	return &SidConnMapping{sidConn: make(map[string]*UserConnection)}
}

type SidConnMapping struct {
	sidConn map[string]*UserConnection

	mu sync.RWMutex
}

func (mapping *SidConnMapping) Add(sid string, userConn *UserConnection) {
	mapping.mu.Lock()
	mapping.sidConn[sid] = userConn
	mapping.mu.Unlock()
}

func (mapping *SidConnMapping) Delete(sid string) {
	mapping.mu.Lock()
	delete(mapping.sidConn, sid)
	mapping.mu.Unlock()
}

func (mapping *SidConnMapping) Get(sid string) *UserConnection {
	mapping.mu.RLock()
	conn, _ := mapping.sidConn[sid]
	mapping.mu.RUnlock()
	return conn
}
