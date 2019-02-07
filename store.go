package main

import (
	"github.com/minio/minio-go"
)

type Store struct {
	Client *minio.Client
}

func NewStore() *Store {
	minioClient, err := minio.New("localhost:9000", "minio", "minio123", false)

	if err != nil {
		panic(err)
	}

	return &Store{
		Client: minioClient,
	}
}
