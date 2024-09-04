package arweave

import (
	"alt-da/config"
	"alt-da/evm"
	"context"
	"testing"
)

func TestNewArStore(t *testing.T) {
	t.Skip("skip test")
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")
	cfg := DefaultArConfig()
	cfg.WalletPath = "path_to/data/XisVwtTblF4tZQ92Pgl4R1obWI2w76aEH4STmxhGJ9U.json"
	s, err := NewArStore(cfg, daId, submitter)
	if err != nil {
		t.Fatal(err)
	}

	if s == nil {
		t.Fatal("store is nil")
	}

	t.Log("store created")
}

func TestArStore_Put(t *testing.T) {
	t.Skip("skip test")
	cfgApp := config.DefaultConfig()
	submitter, err := evm.NewSubmitter(cfgApp)
	if err != nil {
		t.Fatal(err)
	}
	daId := [32]byte{}
	copy(daId[:], "0x00da")

	cfg := DefaultArConfig()
	cfg.WalletPath = "path_to/data/XisVwtTblF4tZQ92Pgl4R1obWI2w76aEH4STmxhGJ9U.json"
	cfg.ClientUrl = "http://localhost:8080"
	s, err := NewArStore(cfg, daId, submitter)
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

	if string(data) != "value12" {
		t.Fatal("data mismatch")
	}
}
