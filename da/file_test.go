package da

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const FileData = "plasma-da/data/filestore"

func TestNewFileStore(t *testing.T) {
	store, err := NewFileStore(FileData)
	assert.NoError(t, err)
	assert.NotNil(t, store)
	assert.Equal(t, store.directory, "data")
}

func TestFileStore_Put(t *testing.T) {
	if _, err := os.Stat(FileData); os.IsNotExist(err) {
		os.MkdirAll(FileData, 0755)
	}
	store, err := NewFileStore(FileData)
	assert.NoError(t, err)
	err = store.Put(nil, []byte("key"), []byte("value"))
	assert.NoError(t, err)

	data, err := store.Get(nil, []byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("value"))
}
