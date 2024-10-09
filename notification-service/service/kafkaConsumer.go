package service

import "github.com/IBM/sarama"

type KafkaConsumerService struct {
}

func NewKafkaConsumerService() *KafkaConsumerService {
	return &KafkaConsumerService{}
}

func (k *KafkaConsumerService) ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (k *KafkaConsumerService) SubscribeTopic(conn sarama.Consumer, topic string) (sarama.PartitionConsumer, error) {
	consumer, err := conn.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
