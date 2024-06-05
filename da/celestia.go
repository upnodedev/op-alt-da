package da

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"github.com/spf13/viper"
	"os"
	"time"
)

// ErrBlobTooLarge is returned when the blob is too large.
var ErrBlobTooLarge = errors.New("blob too large")

// ErrFailedToSubmit is returned when the blob submission fails.
var ErrFailedToSubmit = errors.New("failed to submit blob")

var ErrBlobNotFound = errors.New("blob: not found")

const (
	Rpc                 = "celestia.rpc_port"
	AuthToken           = "celestia.auth_token"
	Namespace           = "celestia.namespace"
	EthFallbackDisabled = "celestia.eth_fallback_disabled"
	MaxBlobSize         = "celestia.max_blob_size"
	GasPrice            = "celestia.gas_price"
)

const (
	DefaultDataMapPath = "plasma-da/data/celestia"
	DaCelestia         = "celestia"
)

type CelestiaCfg struct {
	Rpc                 string
	AuthToken           string
	Namespace           string
	EthFallbackDisabled bool
	MaxBlobSize         uint64
	GasPrice            float64
}

func DefaultCelestiaConfig() CelestiaCfg {
	return CelestiaCfg{
		Rpc:                 "http://localhost:7980",
		AuthToken:           "",
		Namespace:           "",
		EthFallbackDisabled: false,
		MaxBlobSize:         2000,
		GasPrice:            0,
	}
}

func NewCelestiaCfg() CelestiaCfg {
	cfg := DefaultCelestiaConfig()

	if rpc := viper.GetString(Rpc); rpc != "" {
		cfg.Rpc = rpc
	}
	if authToken := viper.GetString(AuthToken); authToken != "" {
		cfg.AuthToken = authToken
	}
	if namespace := viper.GetString(Namespace); namespace != "" {
		cfg.Namespace = namespace
	}
	if ethFallbackDisabled := viper.GetBool(EthFallbackDisabled); ethFallbackDisabled {
		cfg.EthFallbackDisabled = ethFallbackDisabled
	}
	if maxBlobSize := viper.GetUint64(MaxBlobSize); maxBlobSize > 0 {
		cfg.MaxBlobSize = maxBlobSize
	}
	if gasPrice := viper.GetFloat64(GasPrice); gasPrice > 0 {
		cfg.GasPrice = gasPrice
	}

	return cfg
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
		if err := os.MkdirAll(DefaultDataMapPath, 0755); err != nil {
			return nil, err
		}
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
