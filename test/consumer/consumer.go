package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	pConsumer, err := consumer.ConsumePartition("IM", 0, 0)
	if err != nil {
		panic(err)
	}
	for {
		msg := <-pConsumer.Messages()
		fmt.Println(msg.Topic + string(msg.Key) + string(msg.Value))
	}
}
