package biz

import (
	"context"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

func NewUserStatusBiz(usRepo domain.UserStatusRepo) *UserStatusBiz {
	return &UserStatusBiz{userStatusRepo: usRepo}
}

type UserStatusBiz struct {
	userStatusRepo domain.UserStatusRepo
}

func (biz *UserStatusBiz) Connect(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.Connect(ctx, status)
}

func (biz *UserStatusBiz) Disconnect(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.Disconnect(ctx, status)
}

func (biz *UserStatusBiz) KeepAlive(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.KeepAlive(ctx, status)
}
