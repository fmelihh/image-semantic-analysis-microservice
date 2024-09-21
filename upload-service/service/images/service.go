package images

import (
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"upload-service/types"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SaveImage(f multipart.File, h *multipart.FileHeader) (types.ImageMetadata, error) {
	defer f.Close()

	fileName := strings.TrimSuffix(h.Filename, filepath.Ext(h.Filename))
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return types.ImageMetadata{}, err
	}

	imageMetadata := types.ImageMetadata{
		Name:     fileName,
		MimeType: h.Header.Get("Content-Type"),
		Bytes:    fileBytes,
	}

	return imageMetadata, nil
}
