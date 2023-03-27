package domain

import (
	"context"
	"time"
)

type UserStatusRepo interface {
	Connect(ctx context.Context, status *UserStatus) error
	Disconnect(ctx context.Context, status *UserStatus) error
	KeepAlive(ctx context.Context, status *UserStatus) error
}

type UserStatus struct {
	// 网关IP
	Server string
	// socketId
	SID string
	// 用户ID
	UID string
	// 平台
	Platform Platform
	// 过期时间
	ExpireDuration time.Duration
}
