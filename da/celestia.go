package da

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"os"
	"time"
)

// ErrBlobTooLarge is returned when the blob is too large.
var ErrBlobTooLarge = errors.New("blob too large")

// ErrFailedToSubmit is returned when the blob submission fails.
var ErrFailedToSubmit = errors.New("failed to submit blob")

var ErrBlobNotFound = errors.New("blob: not found")

var DefaultDataMapPath = "plasma-da/data/celestia"

type CelestiaCfg struct {
	Rpc                 string
	AuthToken           string
	Namespace           string
	EthFallbackDisabled bool
	MaxBlobSize         uint64
	GasPrice            float64
}

type CelestiaMap struct {
	// key is the key of the blob
	Key []da.ID `json:"key"`
}

type CelestiaStore struct {
	Client     da.DA
	GetTimeout time.Duration
	Namespace  da.Namespace
	cfg        CelestiaCfg

	// template save mapping key with ids into file
	fileStore FileStore
}

// NewCelestiaStore creates a new CelestiaStore.
func NewCelestiaStore(cfg CelestiaCfg) (*CelestiaStore, error) {
	client, err := proxy.NewClient(cfg.Rpc, cfg.AuthToken)
	if err != nil {
		return nil, err
	}

	ns, err := hex.DecodeString(cfg.Namespace)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(DefaultDataMapPath); os.IsNotExist(err) {
		os.MkdirAll(DefaultDataMapPath, 0755)
	}

	return &CelestiaStore{
		Client:     client,
		GetTimeout: time.Minute,
		Namespace:  ns,
		cfg:        cfg,
		fileStore:  FileStore{directory: DefaultDataMapPath},
	}, nil
}

func (c *CelestiaStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	// get ids from data map
	dataRead, err := c.fileStore.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var dataMap CelestiaMap
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

func (c *CelestiaStore) Put(ctx context.Context, key []byte, value []byte) error {
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

	dataMap := CelestiaMap{
		Key: ids,
	}
	dataWrite, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}
	if err := c.fileStore.Put(ctx, key, dataWrite); err != nil {
		return err
	}

	return nil
}
