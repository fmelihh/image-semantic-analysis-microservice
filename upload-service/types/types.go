package types

import (
	"mime/multipart"

	"github.com/IBM/sarama"
)

type ImageUploadService interface {
	SaveImage(multipart.File, *multipart.FileHeader) (ImageMetadata, error)
}

type KafkaProducerService interface {
	ConnectProducer() (sarama.SyncProducer, error)
	PushMessage(topic string, message []byte) error
}

type ImageMetadata struct {
	Name     string
	MimeType string
	Bytes    []byte
}
