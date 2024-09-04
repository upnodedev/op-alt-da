package arweave

import (
	"alt-da/common"
	"alt-da/evm"
	"context"
	"crypto/sha256"
	"encoding/json"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"math/big"
)

const (
	DaAr = "ar"
)

type MappingTx struct {
	TxID string `json:"tx_id"`
}

type ArStore struct {
	Wallet    *goar.Wallet
	daId      [32]byte
	submitter *evm.Submitter
}

func NewArStore(cfg Config, daId [32]byte, submitter *evm.Submitter) (*ArStore, error) {
	wallet, err := goar.NewWalletFromPath(cfg.WalletPath, cfg.ClientUrl)
	if err != nil {
		return nil, err
	}

	return &ArStore{
		Wallet:    wallet,
		daId:      daId,
		submitter: submitter,
	}, nil
}

func (s *ArStore) Get(_ context.Context, key []byte) ([]byte, error) {
	if s.Wallet == nil {
		return nil, common.ErrWalletNotFound
	}
	// get tx id from plasma hub contract
	dataRead, err := s.submitter.GetSubmitter(s.submitter.Transactor.Address(), sha256.Sum256(key), s.daId)
	if err != nil {
		return nil, err
	}
	if len(dataRead) == 0 {
		return nil, common.ErrDataNotFound
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

	_, err = s.submitter.SubmitData(sha256.Sum256(key), s.daId, data)
	if err != nil {
		return err
	}
	return nil
}
