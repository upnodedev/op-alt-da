package da

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCelestiaStore(t *testing.T) {
	store, err := NewCelestiaStore(DefaultCelestiaConfig())
	assert.NoError(t, err)
	assert.NotEqual(t, store, nil)
}

func TestCelestiaStore_Put(t *testing.T) {
	ctx := context.Background()
	store, err := NewCelestiaStore(DefaultCelestiaConfig())
	assert.NoError(t, err)

	err = store.Put(ctx, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(ctx, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
