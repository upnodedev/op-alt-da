package evm

import (
	"fmt"
	"github.com/celer-network/goutils/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"os"
	"plasma/config"
	"plasma/evm/contracts"
)

type Submitter struct {
	Transactor       *eth.Transactor
	Client           *ethclient.Client
	PlasmaHubAddress common.Address
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

	plasmaHubAddress := common.HexToAddress(cfg.PlasmaHubAddr)
	return &Submitter{Transactor: transactor, Client: ec, PlasmaHubAddress: plasmaHubAddress}, nil
}

func (s *Submitter) SubmitData(dataHash [32]byte, da [32]byte, cid []byte) (string, error) {
	receipt, err := s.Transactor.TransactWaitMined(
		fmt.Sprintf("SubmitData dataHash=%x da=%x cid=%x", dataHash, da, cid),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*types.Transaction, error) {
			plasmaHub, err := contracts.NewContractsTransactor(s.PlasmaHubAddress, transactor)
			if err != nil {
				return nil, err
			}
			return plasmaHub.Submit(opts, dataHash, da, cid)
		},
	)
	if err != nil {
		return "", err
	}

	return receipt.TxHash.Hex(), nil
}

func (s *Submitter) GetSubmitter(submitter common.Address, dataHash [32]byte) ([]contracts.PlasmaDaTranslationHubDaCid, error) {
	plasmaHub, err := contracts.NewContractsCaller(s.PlasmaHubAddress, s.Client)
	if err != nil {
		return nil, err
	}

	return plasmaHub.Get(&bind.CallOpts{}, submitter, dataHash)
}
