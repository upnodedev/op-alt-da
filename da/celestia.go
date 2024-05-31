package da

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"time"
)

// ErrBlobTooLarge is returned when the blob is too large.
var ErrBlobTooLarge = errors.New("blob too large")

// ErrFailedToSubmit is returned when the blob submission fails.
var ErrFailedToSubmit = errors.New("failed to submit blob")

var ErrBlobNotFound = errors.New("blob: not found")

type CelestiaCfg struct {
	Rpc                 string
	AuthToken           string
	Namespace           string
	EthFallbackDisabled bool
	MaxBlobSize         uint64
	GasPrice            float64
}

type CelestiaStore struct {
	Client     da.DA
	GetTimeout time.Duration
	Namespace  da.Namespace
	cfg        CelestiaCfg
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

	return &CelestiaStore{
		Client:     client,
		GetTimeout: time.Minute,
		Namespace:  ns,
		cfg:        cfg,
	}, nil
}

func (c *CelestiaStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	// TODO: implement me
	panic("not implemented")
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

	// TODO: mapping key with ids

	return nil
}
