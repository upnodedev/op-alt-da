package file

import (
	"alt-da/common"
	"context"
	"encoding/hex"
	"os"
	"path"
)

type Store struct {
	directory string
}

func NewFileStore(cfg Config) (*Store, error) {
	if _, err := os.Stat(cfg.Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg.Directory, 0755); err != nil {
			return nil, err
		}
	}
	return &Store{
		directory: cfg.Directory,
	}, nil
}

func (s *Store) Get(_ context.Context, key []byte) ([]byte, error) {
	data, err := os.ReadFile(s.fileName(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (s *Store) Put(_ context.Context, key []byte, value []byte) error {
	return os.WriteFile(s.fileName(key), value, 0600)
}

func (s *Store) fileName(key []byte) string {
	return path.Join(s.directory, hex.EncodeToString(key))
}
