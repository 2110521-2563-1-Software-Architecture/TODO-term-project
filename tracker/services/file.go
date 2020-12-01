package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type FileService struct {
	rdb *redis.Client
}

func NewFileService(rdb *redis.Client) *FileService {
	return &FileService{
		rdb: rdb,
	}
}

func (s *FileService) GetAllFileNames() ([]string, error) {
	return s.rdb.Keys(ctx, "*").Result()
}

func (s *FileService) GetPeersWithFile(fileName string) ([]string, error) {
	return s.rdb.SMembers(ctx, fileName).Result()
}

func (s *FileService) AddFileToPeer(fileName string, peerAddr string) error {
	return s.rdb.SAdd(ctx, fileName, peerAddr).Err()
}
