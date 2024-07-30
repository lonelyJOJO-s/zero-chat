package kafka

import (
	"zero-chat/app/chat/cmd/api/internal/config"
	"zero-chat/app/chat/cmd/api/internal/ws"

	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

func StartConsume(c config.Config) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(c.KqConsumerConf.Brokers, config)
	if err != nil {
		logx.Error(err)
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		logx.Error(err)
	}
	partitionConsumer, err := consumer.ConsumePartition(c.KqConsumerConf.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		logx.Errorf("ConsumePartition error:%s", err.Error())
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		ws.WsServer.Broadcast <- msg.Value
	}
}
