package da

import (
	"context"
	"encoding/hex"
	"os"
	"path"
	"plasma/common"
)

const (
	DaFile = "file"
	DaPath = "filestore.path"
)

type FileStoreCfg struct {
	Directory string
}

func DefaultFileStoreCfg() FileStoreCfg {
	return FileStoreCfg{
		Directory: "plasm-hub/data",
	}
}

func NewFileStoreCfg() FileStoreCfg {
	cfg := DefaultFileStoreCfg()

	if daPath := os.Getenv(DaPath); daPath != "" {
		cfg.Directory = daPath
	}

	return cfg
}

type FileStore struct {
	directory string
}

func NewFileStore(directory string) (*FileStore, error) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.MkdirAll(directory, 0755); err != nil {
			return nil, err
		}
	}
	return &FileStore{
		directory: directory,
	}, nil
}

func (s *FileStore) Get(_ context.Context, key []byte) ([]byte, error) {
	data, err := os.ReadFile(s.fileName(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (s *FileStore) Put(_ context.Context, key []byte, value []byte) error {
	return os.WriteFile(s.fileName(key), value, 0600)
}

func (s *FileStore) fileName(key []byte) string {
	return path.Join(s.directory, hex.EncodeToString(key))
}
