package da

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFileStore(t *testing.T) {
	store := NewFileStore("data")
	if store == nil {
		t.Error("expected store to be non-nil")
	}

	assert.NotNil(t, store)
	assert.Equal(t, store.directory, "data")
}

func TestFileStore_Put(t *testing.T) {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}
	store := NewFileStore("data")
	err := store.Put(nil, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(nil, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
