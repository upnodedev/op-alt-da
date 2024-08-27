package celestia

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"os"
	"path"
	"plasma/common"
	"time"
)

// ErrBlobTooLarge is returned when the blob is too large.
var ErrBlobTooLarge = errors.New("blob too large")

// ErrFailedToSubmit is returned when the blob submission fails.
var ErrFailedToSubmit = errors.New("failed to submit blob")

var ErrBlobNotFound = errors.New("blob: not found")

const (
	DefaultDataMapPath = ".plasma-da/data/celestia"
	DaCelestia         = "celestia"
)

type MappingCommitment struct {
	// key is the key of the blob
	Key []da.ID `json:"key"`
}

type Store struct {
	Client     da.DA
	GetTimeout time.Duration
	Namespace  da.Namespace
	cfg        Config

	// temporary file mapping
	mappingPath string
}

// NewCelestiaStore creates a new CelestiaStore.
func NewCelestiaStore(cfg Config, homeDir string) (*Store, error) {
	client, err := proxy.NewClient(cfg.Rpc, cfg.AuthToken)
	if err != nil {
		return nil, err
	}

	ns, err := hex.DecodeString(cfg.Namespace)
	if err != nil {
		return nil, err
	}

	mapPath := path.Join(homeDir, DefaultDataMapPath)

	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		if err := os.MkdirAll(mapPath, 0755); err != nil {
			return nil, err
		}
	}

	return &Store{
		Client:      client,
		GetTimeout:  time.Minute,
		Namespace:   ns,
		cfg:         cfg,
		mappingPath: mapPath,
	}, nil
}

func (c *Store) Get(ctx context.Context, key []byte) ([]byte, error) {
	// get ids from data map
	dataRead, err := c.readFile(key)
	if err != nil {
		return nil, err
	}

	var dataMap MappingCommitment
	if err := json.Unmarshal(dataRead, &dataMap); err != nil {
		return nil, err
	}
	dataBlob, err := c.Client.Get(ctx, dataMap.Key, c.Namespace)
	if err != nil {
		return nil, err
	}
	if len(dataBlob) == 0 {
		return nil, ErrBlobNotFound
	}
	return dataBlob[0], nil
}

func (c *Store) Put(ctx context.Context, key []byte, value []byte) error {
	if uint64(len(value)) > c.cfg.MaxBlobSize {
		return ErrBlobTooLarge
	}

	var blobs [][]byte
	blobs = append(blobs, value)

	ids, err := c.Client.Submit(ctx, blobs, c.cfg.GasPrice, c.Namespace)
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		return ErrFailedToSubmit
	}

	dataMap := MappingCommitment{
		Key: ids,
	}
	dataWrite, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.writeFile(key, dataWrite)
}

func (c *Store) readFile(key []byte) ([]byte, error) {
	data, err := os.ReadFile(c.fileName(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (c *Store) writeFile(key []byte, value []byte) error {
	return os.WriteFile(c.fileName(key), value, 0600)
}

func (c *Store) fileName(key []byte) string {
	return path.Join(c.mappingPath, hex.EncodeToString(key))
}
