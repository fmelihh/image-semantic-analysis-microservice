package kafkaProducer

import "github.com/IBM/sarama"

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
	_, err := s.ConnectProducer()
	if err != nil {
		return err
	}
	return nil
}
