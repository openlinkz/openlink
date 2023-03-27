package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/openlinkz/openlink/app/msg-api/internal/domain"
)

var _ domain.UserStatusRepo = (*userStatusRepo)(nil)

func NewUserStatusRepo() domain.UserStatusRepo {
	return &userStatusRepo{}
}

const _userOnlinePrefix string = "USER:ONLINE:"

type userStatusRepo struct {
	rdb *redis.Client
}

func (repo *userStatusRepo) Connect(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Set(ctx, _userOnlinePrefix+userStatus.UID, userStatus.SID, userStatus.ExpireDuration).Err()
}

func (repo *userStatusRepo) Disconnect(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Del(ctx, _userOnlinePrefix+userStatus.UID).Err()
}

func (repo *userStatusRepo) KeepAlive(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Expire(ctx, _userOnlinePrefix+userStatus.UID, userStatus.ExpireDuration).Err()
}
