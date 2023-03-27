package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	appconfig "github.com/openlinkz/openlink/app/msg-api/internal/config"
)

var ProviderSet = wire.NewSet(NewUserStatusRepo, NewMsgRepo)

func NewRedisClient(c *appconfig.Data_Redis) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: c.Addr, Password: c.Password, DB: int(c.Db)})
}
