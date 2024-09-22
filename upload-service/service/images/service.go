package images

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"upload-service/types"

	"github.com/minio/minio-go/v7"
)

type Service struct {
	minioClient *minio.Client
}

func NewService(minioClient *minio.Client) *Service {
	return &Service{minioClient: minioClient}
}

func (s *Service) SaveImage(f multipart.File, h *multipart.FileHeader) (types.ImageMetadata, error) {
	defer f.Close()

	fileName := strings.TrimSuffix(h.Filename, filepath.Ext(h.Filename))
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return types.ImageMetadata{}, err
	}

	ctx := context.Background()
	imageIOReader := bytes.NewReader(fileBytes)
	s.minioClient.PutObject(ctx, "image", h.Filename, imageIOReader, h.Size, minio.PutObjectOptions{})

	locationUrl := fmt.Sprintf("http://localhost:9000/image/%s", h.Filename)
	imageMetadata := types.ImageMetadata{
		Name:        fileName,
		MimeType:    h.Header.Get("Content-Type"),
		LocationUrl: locationUrl,
	}

	return imageMetadata, nil
}
