package da

import (
	"context"
)

type FileStore struct {
	directory string
}

func NewFileStore(directory string) *FileStore {
	return &FileStore{
		directory: directory,
	}
}

func (s *FileStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	panic("implement me")
}

func (s *FileStore) Put(ctx context.Context, key []byte, value []byte) error {
	panic("implement me")
}
