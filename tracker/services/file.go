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
	members, err := s.rdb.SMembers(ctx, fileName).Result()
	if err != nil {
		return []string{}, err
	}

	peers, err := LookupPeersWithConsul()
	if err != nil {
		return []string{}, err
	}

	var result, removedPeers []string
	for _, m := range members {
		isMember := false
		for _, p := range peers {
			if m == p {
				isMember = true
				break
			}
		}
		if isMember {
			result = append(result, m)
		} else {
			removedPeers = append(removedPeers, m)
		}
	}
	err = s.rdb.SRem(ctx, fileName, removedPeers).Err()
	if err != nil {
		return []string{}, nil
	}

	return result, nil
}

func (s *FileService) AddFileToPeer(fileName string, peerAddr string) error {
	return s.rdb.SAdd(ctx, fileName, peerAddr).Err()
}
