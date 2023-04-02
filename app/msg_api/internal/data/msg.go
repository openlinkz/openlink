package data

import (
	"context"
	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"
	"github.com/openlinkz/openlink/app/msg_api/internal/domain"
	"github.com/pkg/errors"
)

var _ domain.MsgRepo = (*msgKafkaMQRepo)(nil)

func NewMsgRepo(producer sarama.SyncProducer) domain.MsgRepo {
	return &msgKafkaMQRepo{producer: producer}
}

type msgKafkaMQRepo struct {
	producer sarama.SyncProducer
}

func (repo *msgKafkaMQRepo) SendMsg(ctx context.Context, msg *domain.Message) error {
	topic, key := msg.BizTopic(), msg.BizKey()
	if topic == "" || key == "" {
		return errors.New("invalid msg")
	}

	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = topic
	kMsg.Key = sarama.StringEncoder(key)
	bMsg, err := jsoniter.Marshal(msg)
	if err != nil {
		return err
	}
	kMsg.Value = sarama.ByteEncoder(bMsg)

	_, _, err = repo.producer.SendMessage(kMsg)
	return err
}

func (repo *msgKafkaMQRepo) SendBatchMsg(ctx context.Context, messages []*domain.Message) error {
	pMessage := make([]*sarama.ProducerMessage, 0, len(messages))

	for _, msg := range messages {
		msg, err := convertToProducerMsg(msg)
		if err != nil {
			continue
		}
		pMessage = append(pMessage, msg)
	}

	return repo.producer.SendMessages(pMessage)
}

func convertToProducerMsg(msg *domain.Message) (*sarama.ProducerMessage, error) {
	topic, key := msg.BizTopic(), msg.BizKey()
	if topic == "" || key == "" {
		return nil, errors.New("invalid msg")
	}

	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = msg.BizTopic()
	kMsg.Key = sarama.StringEncoder(key)
	bMsg, err := jsoniter.Marshal(msg)
	if err != nil {
		return nil, err
	}
	kMsg.Value = sarama.ByteEncoder(bMsg)

	return kMsg, nil
}
