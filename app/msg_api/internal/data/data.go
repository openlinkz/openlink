package data

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	appconfig "github.com/openlinkz/openlink/app/msg_api/internal/config"
	"github.com/pkg/errors"
)

var ProviderSet = wire.NewSet(NewKafkaSyncProducer, NewRedisClient, NewUserStatusRepo, NewMsgRepo)

func NewRedisClient(c *appconfig.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: c.Addr, Password: c.Password, DB: int(c.Db)})
}

func NewKafkaSyncProducer(c *appconfig.Kafka) sarama.SyncProducer {
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewSyncProducer(c.Addrs, producerConfig)
	if err != nil {
		panic(errors.Wrap(err, "build kafka producer failed"))
	}

	return producer
}
