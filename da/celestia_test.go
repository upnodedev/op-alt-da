package da

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCelestiaStore(t *testing.T) {
	cfg := CelestiaCfg{
		Rpc:                 "http://localhost:7980",
		AuthToken:           "",
		Namespace:           "",
		EthFallbackDisabled: false,
		MaxBlobSize:         2000,
		GasPrice:            0,
	}

	store, err := NewCelestiaStore(cfg)
	assert.NoError(t, err)
	assert.NotEqual(t, store, nil)
}

func TestCelestiaStore_Put(t *testing.T) {
	cfg := CelestiaCfg{
		Rpc:                 "http://localhost:7980",
		AuthToken:           "",
		Namespace:           "",
		EthFallbackDisabled: false,
		MaxBlobSize:         2000,
		GasPrice:            0,
	}

	ctx := context.Background()
	store, err := NewCelestiaStore(cfg)
	assert.NoError(t, err)

	err = store.Put(ctx, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(ctx, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
