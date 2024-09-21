package types

import (
	"mime/multipart"
)

type ImageUploadService interface {
	SaveImage(multipart.File, *multipart.FileHeader) (ImageMetadata, error)
}

type ImageMetadata struct {
	Name     string
	MimeType string
	Bytes    []byte
}
