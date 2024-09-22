package kafkaProducer

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Service struct {
}

func NewKafkaProducerService() *Service {
	return &Service{}
}

func (s *Service) ConnectProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Service) PushMessage(topic string, message []byte) error {
	producer, err := s.ConnectProducer()
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
