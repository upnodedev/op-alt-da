package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileStore(t *testing.T) {
	cfg := DefaultFileStoreCfg()
	store, err := NewFileStore(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, store)
}

func TestFileStore_Put(t *testing.T) {
	cfg := DefaultFileStoreCfg()
	store, err := NewFileStore(cfg)
	assert.NoError(t, err)
	err = store.Put(nil, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(nil, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
