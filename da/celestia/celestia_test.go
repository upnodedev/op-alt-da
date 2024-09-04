package celestia

import (
	"alt-da/config"
	"alt-da/evm"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCelestiaStore(t *testing.T) {
	t.Skip("skip test")
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")
	store, err := NewCelestiaStore(DefaultCelestiaConfig(), daId, submitter)
	assert.NoError(t, err)
	assert.NotEqual(t, store, nil)
}

func TestCelestiaStore_Put(t *testing.T) {
	t.Skip("skip test")
	ctx := context.Background()
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")
	store, err := NewCelestiaStore(DefaultCelestiaConfig(), daId, submitter)
	assert.NoError(t, err)

	err = store.Put(ctx, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(ctx, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
