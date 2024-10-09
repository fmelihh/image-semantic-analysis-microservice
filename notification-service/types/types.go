package types

import "github.com/IBM/sarama"

type KafkaConsumer interface {
	ConnectConsumer([]string) (sarama.Consumer, error)
	SubscribeTopic(sarama.Consumer, string) (sarama.PartitionConsumer, error)
}
