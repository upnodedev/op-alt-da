package arweave

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"math/big"
	"os"
	"path"
	"plasma/common"
)

const (
	DefaultArMapPath = ".plasma-da/data/ar"
	DaAr             = "ar"
)

type MappingTx struct {
	TxID string `json:"tx_id"`
}

type ArStore struct {
	Wallet *goar.Wallet

	// temporary mapping the transaction in file
	mappingPath string
}

func NewArStore(cfg Config, homeDir string) (*ArStore, error) {
	wallet, err := goar.NewWalletFromPath(cfg.WalletPath, cfg.ClientUrl)
	if err != nil {
		return nil, err
	}

	mapPath := path.Join(homeDir, DefaultArMapPath)
	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		if err := os.MkdirAll(mapPath, 0755); err != nil {
			return nil, err
		}
	}

	return &ArStore{
		Wallet:      wallet,
		mappingPath: mapPath,
	}, nil
}

func (s *ArStore) Get(_ context.Context, key []byte) ([]byte, error) {
	if s.Wallet == nil {
		return nil, common.ErrWalletNotFound
	}
	// get tx id from data map
	dataRead, err := s.readFile(key)
	if err != nil {
		return nil, err
	}

	var dataMap MappingTx
	if err := json.Unmarshal(dataRead, &dataMap); err != nil {
		return nil, err
	}

	// get data from arweave
	return s.Wallet.Client.GetTransactionData(dataMap.TxID)
}

func (s *ArStore) Put(_ context.Context, key []byte, value []byte) error {
	if s.Wallet == nil {
		return common.ErrWalletNotFound
	}
	// check balance of wallet
	balance, err := s.Wallet.Client.GetWalletBalance(s.Wallet.Signer.Address)
	if err != nil {
		return err
	}
	if balance.Cmp(big.NewFloat(0.05)) < 0 {
		return common.ErrInsufficientBalance
	}

	// upload data to arweave
	tx, err := s.Wallet.SendData(value, []types.Tag{
		{
			Name:  "key",
			Value: string(key),
		},
	})
	if err != nil {
		return err
	}

	// save mapping
	dataMap := MappingTx{
		TxID: tx.ID,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}
	return s.writeFile(key, data)
}

func (s *ArStore) readFile(key []byte) ([]byte, error) {
	data, err := os.ReadFile(s.fileName(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (s *ArStore) writeFile(key []byte, value []byte) error {
	return os.WriteFile(s.fileName(key), value, 0600)
}

func (s *ArStore) fileName(key []byte) string {
	return path.Join(s.mappingPath, hex.EncodeToString(key))
}
