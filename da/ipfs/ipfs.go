package ipfs

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
	"path"
	"plasma/common"
)

const (
	DefaultIpfsMapPath = ".plasma-da/data/ipfs"
	DaIpfs             = "ipfs"
)

type MappingCID struct {
	Path string `json:"path"`
}

type Store struct {
	Shel *shell.Shell

	// temporary file mapping
	mappingPath string
}

func NewIpfsStore(cfg Config, homeDir string) (*Store, error) {
	sh := shell.NewShell(cfg.Url)

	mapPath := path.Join(homeDir, DefaultIpfsMapPath)
	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		if err := os.MkdirAll(mapPath, 0755); err != nil {
			return nil, err
		}
	}

	return &Store{
		Shel:        sh,
		mappingPath: mapPath,
	}, nil
}

func (s *Store) Get(_ context.Context, key []byte) ([]byte, error) {
	// get path from data map
	dataRead, err := s.readFile(key)
	if err != nil {
		return nil, err
	}

	var dataMap MappingCID
	if err := json.Unmarshal(dataRead, &dataMap); err != nil {
		return nil, err
	}

	data, err := s.Shel.Cat(dataMap.Path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Store) Put(_ context.Context, key []byte, value []byte) error {
	cid, err := s.Shel.Add(bytes.NewReader(value))
	if err != nil {
		return err
	}
	println(cid)

	dataMap := MappingCID{
		Path: cid,
	}

	// save path to data map
	dataWrite, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return s.writeFile(key, dataWrite)
}

func (s *Store) readFile(key []byte) ([]byte, error) {
	data, err := os.ReadFile(s.fileName(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (s *Store) writeFile(key []byte, value []byte) error {
	return os.WriteFile(s.fileName(key), value, 0600)
}

func (s *Store) fileName(key []byte) string {
	return path.Join(s.mappingPath, hex.EncodeToString(key))
}
