package ipfs

import (
	"alt-da/config"
	"alt-da/evm"
	"bytes"
	"context"
	"testing"
)

func TestNewIpfsStore(t *testing.T) {
	t.Skip("skip test")
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")
	s, err := NewIpfsStore(DefaultIpfsConfig(), daId, submitter)
	if err != nil {
		t.Fatal(err)
	}

	if s == nil {
		t.Fatal("store is nil")
	}

	t.Log("store created")
}

func TestIpfsStore_Put(t *testing.T) {
	t.Skip("skip test")
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")
	s, err := NewIpfsStore(DefaultIpfsConfig(), daId, submitter)
	if err != nil {
		t.Fatal(err)
	}

	if s == nil {
		t.Fatal("store is nil")
	}

	err = s.Put(context.Background(), []byte("key"), []byte("value12"))
	if err != nil {
		t.Fatal(err)
	}

	data, err := s.Get(context.Background(), []byte("key"))
	if err != nil {
		t.Fatal(err)
	}
	println(string(data))

	if !bytes.Equal(data, []byte("value12")) {
		t.Fatal("data mismatch")
	}
}
