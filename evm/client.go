package evm

import (
	"alt-da/config"
	"alt-da/evm/contracts"
	"fmt"
	"github.com/celer-network/goutils/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"os"
	"time"
)

type Submitter struct {
	Transactor      *eth.Transactor
	Client          *ethclient.Client
	AltDaHubAddress common.Address
}

func NewSubmitter(cfg config.App) (*Submitter, error) {
	rpcClient, err := ethrpc.Dial(cfg.EvmRpcUrl)
	if err != nil {
		return nil, err
	}
	ec := ethclient.NewClient(rpcClient)

	ksBytes, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	transactor, err := eth.NewTransactor(string(ksBytes), cfg.Passphrase, ec, big.NewInt(cfg.ChainId))
	if err != nil {
		return nil, err
	}

	altDaHubAddress := common.HexToAddress(cfg.AltDaHubAddr)
	return &Submitter{Transactor: transactor, Client: ec, AltDaHubAddress: altDaHubAddress}, nil
}

func (s *Submitter) SubmitData(dataHash [32]byte, da [32]byte, cid []byte) (string, error) {
	receipt, err := s.Transactor.TransactWaitMined(
		fmt.Sprintf("SubmitData dataHash=%x da=%x cid=%x", dataHash, da, cid),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*types.Transaction, error) {
			altDaHub, err := contracts.NewContractsTransactor(s.AltDaHubAddress, transactor)
			if err != nil {
				return nil, err
			}
			return altDaHub.Submit(opts, dataHash, da, cid)
		},
		eth.WithPollingInterval(1*time.Second),
	)
	if err != nil {
		return "", err
	}

	return receipt.TxHash.Hex(), nil
}

func (s *Submitter) GetSubmitter(submitter common.Address, dataHash [32]byte, daId [32]byte) ([]byte, error) {
	plasmaHub, err := contracts.NewContractsCaller(s.AltDaHubAddress, s.Client)
	if err != nil {
		return nil, err
	}

	return plasmaHub.Get(&bind.CallOpts{}, submitter, dataHash, daId)
}
