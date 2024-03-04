package db

import (
	"ApscIM/pkg/common/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	client *minio.Client
}

func NewMinioService() (*MinioService, error) {
	endpoint := config.Conf.Minio.Host
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}
	return &MinioService{
		client: client,
	}, nil
}

func (m *MinioService) AddObjByKey() {
}

func (m *MinioService) DleObjByKey() {
}
