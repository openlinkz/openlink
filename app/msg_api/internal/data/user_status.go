package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
)

var _ domain.UserStatusRepo = (*userStatusRepo)(nil)

const _userOnlinePrefix string = "USER:ONLINE:"

func NewUserStatusRepo(rdb *redis.Client) domain.UserStatusRepo {
	return &userStatusRepo{rdb: rdb}
}

type userStatusRepo struct {
	rdb *redis.Client
}

func (repo *userStatusRepo) Connect(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Set(ctx, userKey(userStatus), userStatus.SID, userStatus.ExpireDuration).Err()
}

func (repo *userStatusRepo) Disconnect(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Del(ctx, userKey(userStatus)).Err()
}

func (repo *userStatusRepo) KeepAlive(ctx context.Context, userStatus *domain.UserStatus) error {
	return repo.rdb.Expire(ctx, userKey(userStatus), userStatus.ExpireDuration).Err()
}

func userKey(u *domain.UserStatus) string {
	return _userOnlinePrefix + domain.PlatformText(u.Platform) + ":" + u.UID
}
