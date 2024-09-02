// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PlasmaDaTranslationHubBatchSubmitInput is an auto generated low-level Go binding around an user-defined struct.
type PlasmaDaTranslationHubBatchSubmitInput struct {
	DataHash [32]byte
	Da       [32]byte
	Cid      []byte
}

// PlasmaDaTranslationHubDaCid is an auto generated low-level Go binding around an user-defined struct.
type PlasmaDaTranslationHubDaCid struct {
	Da  [32]byte
	Cid []byte
}

// PlasmaDaTranslationHubDelegatedSubmitInput is an auto generated low-level Go binding around an user-defined struct.
type PlasmaDaTranslationHubDelegatedSubmitInput struct {
	DataHash  [32]byte
	Da        [32]byte
	Cid       []byte
	Signature []byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DuplicatedSubmission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSubmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSubmitter\",\"type\":\"address\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"TranslationSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPlasmaDaTranslationHub.DelegatedSubmitInput[]\",\"name\":\"submissions\",\"type\":\"tuple[]\"}],\"name\":\"batchDelegatedSubmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"internalType\":\"structPlasmaDaTranslationHub.BatchSubmitInput[]\",\"name\":\"submissions\",\"type\":\"tuple[]\"}],\"name\":\"batchSubmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPlasmaDaTranslationHub.DelegatedSubmitInput\",\"name\":\"submission\",\"type\":\"tuple\"}],\"name\":\"delegatedSubmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"old\",\"type\":\"address\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"internalType\":\"structPlasmaDaTranslationHub.DaCid[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"getExtendedAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"translation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"da\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPlasmaDaTranslationHub.DelegatedSubmitInput\",\"name\":\"submission\",\"type\":\"tuple\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contracts.Contract.Eip712Domain(&_Contracts.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contracts.Contract.Eip712Domain(&_Contracts.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address submitter, bytes32 dataHash) view returns((bytes32,bytes)[] result)
func (_Contracts *ContractsCaller) Get(opts *bind.CallOpts, submitter common.Address, dataHash [32]byte) ([]PlasmaDaTranslationHubDaCid, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "get", submitter, dataHash)

	if err != nil {
		return *new([]PlasmaDaTranslationHubDaCid), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlasmaDaTranslationHubDaCid)).(*[]PlasmaDaTranslationHubDaCid)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address submitter, bytes32 dataHash) view returns((bytes32,bytes)[] result)
func (_Contracts *ContractsSession) Get(submitter common.Address, dataHash [32]byte) ([]PlasmaDaTranslationHubDaCid, error) {
	return _Contracts.Contract.Get(&_Contracts.CallOpts, submitter, dataHash)
}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address submitter, bytes32 dataHash) view returns((bytes32,bytes)[] result)
func (_Contracts *ContractsCallerSession) Get(submitter common.Address, dataHash [32]byte) ([]PlasmaDaTranslationHubDaCid, error) {
	return _Contracts.Contract.Get(&_Contracts.CallOpts, submitter, dataHash)
}

// GetExtendedAddresses is a free data retrieval call binding the contract method 0x1482c252.
//
// Solidity: function getExtendedAddresses(address submitter) view returns(address[])
func (_Contracts *ContractsCaller) GetExtendedAddresses(opts *bind.CallOpts, submitter common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getExtendedAddresses", submitter)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExtendedAddresses is a free data retrieval call binding the contract method 0x1482c252.
//
// Solidity: function getExtendedAddresses(address submitter) view returns(address[])
func (_Contracts *ContractsSession) GetExtendedAddresses(submitter common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetExtendedAddresses(&_Contracts.CallOpts, submitter)
}

// GetExtendedAddresses is a free data retrieval call binding the contract method 0x1482c252.
//
// Solidity: function getExtendedAddresses(address submitter) view returns(address[])
func (_Contracts *ContractsCallerSession) GetExtendedAddresses(submitter common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetExtendedAddresses(&_Contracts.CallOpts, submitter)
}

// Translation is a free data retrieval call binding the contract method 0x872bb0a2.
//
// Solidity: function translation(address , bytes32 , bytes32 ) view returns(bytes)
func (_Contracts *ContractsCaller) Translation(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "translation", arg0, arg1, arg2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Translation is a free data retrieval call binding the contract method 0x872bb0a2.
//
// Solidity: function translation(address , bytes32 , bytes32 ) view returns(bytes)
func (_Contracts *ContractsSession) Translation(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) ([]byte, error) {
	return _Contracts.Contract.Translation(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// Translation is a free data retrieval call binding the contract method 0x872bb0a2.
//
// Solidity: function translation(address , bytes32 , bytes32 ) view returns(bytes)
func (_Contracts *ContractsCallerSession) Translation(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) ([]byte, error) {
	return _Contracts.Contract.Translation(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// VerifySignature is a free data retrieval call binding the contract method 0xdc72ba2c.
//
// Solidity: function verifySignature(address submitter, (bytes32,bytes32,bytes,bytes) submission) view returns(bool)
func (_Contracts *ContractsCaller) VerifySignature(opts *bind.CallOpts, submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifySignature", submitter, submission)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0xdc72ba2c.
//
// Solidity: function verifySignature(address submitter, (bytes32,bytes32,bytes,bytes) submission) view returns(bool)
func (_Contracts *ContractsSession) VerifySignature(submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (bool, error) {
	return _Contracts.Contract.VerifySignature(&_Contracts.CallOpts, submitter, submission)
}

// VerifySignature is a free data retrieval call binding the contract method 0xdc72ba2c.
//
// Solidity: function verifySignature(address submitter, (bytes32,bytes32,bytes,bytes) submission) view returns(bool)
func (_Contracts *ContractsCallerSession) VerifySignature(submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (bool, error) {
	return _Contracts.Contract.VerifySignature(&_Contracts.CallOpts, submitter, submission)
}

// BatchDelegatedSubmit is a paid mutator transaction binding the contract method 0x59c207bd.
//
// Solidity: function batchDelegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes)[] submissions) returns()
func (_Contracts *ContractsTransactor) BatchDelegatedSubmit(opts *bind.TransactOpts, submitter common.Address, submissions []PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "batchDelegatedSubmit", submitter, submissions)
}

// BatchDelegatedSubmit is a paid mutator transaction binding the contract method 0x59c207bd.
//
// Solidity: function batchDelegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes)[] submissions) returns()
func (_Contracts *ContractsSession) BatchDelegatedSubmit(submitter common.Address, submissions []PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.BatchDelegatedSubmit(&_Contracts.TransactOpts, submitter, submissions)
}

// BatchDelegatedSubmit is a paid mutator transaction binding the contract method 0x59c207bd.
//
// Solidity: function batchDelegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes)[] submissions) returns()
func (_Contracts *ContractsTransactorSession) BatchDelegatedSubmit(submitter common.Address, submissions []PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.BatchDelegatedSubmit(&_Contracts.TransactOpts, submitter, submissions)
}

// BatchSubmit is a paid mutator transaction binding the contract method 0x137fd8b8.
//
// Solidity: function batchSubmit((bytes32,bytes32,bytes)[] submissions) returns()
func (_Contracts *ContractsTransactor) BatchSubmit(opts *bind.TransactOpts, submissions []PlasmaDaTranslationHubBatchSubmitInput) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "batchSubmit", submissions)
}

// BatchSubmit is a paid mutator transaction binding the contract method 0x137fd8b8.
//
// Solidity: function batchSubmit((bytes32,bytes32,bytes)[] submissions) returns()
func (_Contracts *ContractsSession) BatchSubmit(submissions []PlasmaDaTranslationHubBatchSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.BatchSubmit(&_Contracts.TransactOpts, submissions)
}

// BatchSubmit is a paid mutator transaction binding the contract method 0x137fd8b8.
//
// Solidity: function batchSubmit((bytes32,bytes32,bytes)[] submissions) returns()
func (_Contracts *ContractsTransactorSession) BatchSubmit(submissions []PlasmaDaTranslationHubBatchSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.BatchSubmit(&_Contracts.TransactOpts, submissions)
}

// DelegatedSubmit is a paid mutator transaction binding the contract method 0xf1d29ff3.
//
// Solidity: function delegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes) submission) returns()
func (_Contracts *ContractsTransactor) DelegatedSubmit(opts *bind.TransactOpts, submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "delegatedSubmit", submitter, submission)
}

// DelegatedSubmit is a paid mutator transaction binding the contract method 0xf1d29ff3.
//
// Solidity: function delegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes) submission) returns()
func (_Contracts *ContractsSession) DelegatedSubmit(submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.DelegatedSubmit(&_Contracts.TransactOpts, submitter, submission)
}

// DelegatedSubmit is a paid mutator transaction binding the contract method 0xf1d29ff3.
//
// Solidity: function delegatedSubmit(address submitter, (bytes32,bytes32,bytes,bytes) submission) returns()
func (_Contracts *ContractsTransactorSession) DelegatedSubmit(submitter common.Address, submission PlasmaDaTranslationHubDelegatedSubmitInput) (*types.Transaction, error) {
	return _Contracts.Contract.DelegatedSubmit(&_Contracts.TransactOpts, submitter, submission)
}

// Extend is a paid mutator transaction binding the contract method 0x82005715.
//
// Solidity: function extend(address old) returns()
func (_Contracts *ContractsTransactor) Extend(opts *bind.TransactOpts, old common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "extend", old)
}

// Extend is a paid mutator transaction binding the contract method 0x82005715.
//
// Solidity: function extend(address old) returns()
func (_Contracts *ContractsSession) Extend(old common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Extend(&_Contracts.TransactOpts, old)
}

// Extend is a paid mutator transaction binding the contract method 0x82005715.
//
// Solidity: function extend(address old) returns()
func (_Contracts *ContractsTransactorSession) Extend(old common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Extend(&_Contracts.TransactOpts, old)
}

// Submit is a paid mutator transaction binding the contract method 0x769f4b66.
//
// Solidity: function submit(bytes32 dataHash, bytes32 da, bytes cid) returns()
func (_Contracts *ContractsTransactor) Submit(opts *bind.TransactOpts, dataHash [32]byte, da [32]byte, cid []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submit", dataHash, da, cid)
}

// Submit is a paid mutator transaction binding the contract method 0x769f4b66.
//
// Solidity: function submit(bytes32 dataHash, bytes32 da, bytes cid) returns()
func (_Contracts *ContractsSession) Submit(dataHash [32]byte, da [32]byte, cid []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Submit(&_Contracts.TransactOpts, dataHash, da, cid)
}

// Submit is a paid mutator transaction binding the contract method 0x769f4b66.
//
// Solidity: function submit(bytes32 dataHash, bytes32 da, bytes cid) returns()
func (_Contracts *ContractsTransactorSession) Submit(dataHash [32]byte, da [32]byte, cid []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Submit(&_Contracts.TransactOpts, dataHash, da, cid)
}

// ContractsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Contracts contract.
type ContractsEIP712DomainChangedIterator struct {
	Event *ContractsEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEIP712DomainChanged represents a EIP712DomainChanged event raised by the Contracts contract.
type ContractsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractsEIP712DomainChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsEIP712DomainChangedIterator{contract: _Contracts.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEIP712DomainChanged)
				if err := _Contracts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) ParseEIP712DomainChanged(log types.Log) (*ContractsEIP712DomainChanged, error) {
	event := new(ContractsEIP712DomainChanged)
	if err := _Contracts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Contracts contract.
type ContractsExtendIterator struct {
	Event *ContractsExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsExtend represents a Extend event raised by the Contracts contract.
type ContractsExtend struct {
	NewSubmitter common.Address
	OldSubmitter common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0xaee0438aa8e328b164c60ff0e6aa76c79f51ef6f5af0964b9ba0b490f6f6688b.
//
// Solidity: event Extend(address indexed newSubmitter, address indexed oldSubmitter)
func (_Contracts *ContractsFilterer) FilterExtend(opts *bind.FilterOpts, newSubmitter []common.Address, oldSubmitter []common.Address) (*ContractsExtendIterator, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}
	var oldSubmitterRule []interface{}
	for _, oldSubmitterItem := range oldSubmitter {
		oldSubmitterRule = append(oldSubmitterRule, oldSubmitterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Extend", newSubmitterRule, oldSubmitterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsExtendIterator{contract: _Contracts.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0xaee0438aa8e328b164c60ff0e6aa76c79f51ef6f5af0964b9ba0b490f6f6688b.
//
// Solidity: event Extend(address indexed newSubmitter, address indexed oldSubmitter)
func (_Contracts *ContractsFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *ContractsExtend, newSubmitter []common.Address, oldSubmitter []common.Address) (event.Subscription, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}
	var oldSubmitterRule []interface{}
	for _, oldSubmitterItem := range oldSubmitter {
		oldSubmitterRule = append(oldSubmitterRule, oldSubmitterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Extend", newSubmitterRule, oldSubmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsExtend)
				if err := _Contracts.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0xaee0438aa8e328b164c60ff0e6aa76c79f51ef6f5af0964b9ba0b490f6f6688b.
//
// Solidity: event Extend(address indexed newSubmitter, address indexed oldSubmitter)
func (_Contracts *ContractsFilterer) ParseExtend(log types.Log) (*ContractsExtend, error) {
	event := new(ContractsExtend)
	if err := _Contracts.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTranslationSubmittedIterator is returned from FilterTranslationSubmitted and is used to iterate over the raw logs and unpacked data for TranslationSubmitted events raised by the Contracts contract.
type ContractsTranslationSubmittedIterator struct {
	Event *ContractsTranslationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTranslationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTranslationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTranslationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTranslationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTranslationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTranslationSubmitted represents a TranslationSubmitted event raised by the Contracts contract.
type ContractsTranslationSubmitted struct {
	Submitter common.Address
	DataHash  [32]byte
	Da        [32]byte
	Cid       []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTranslationSubmitted is a free log retrieval operation binding the contract event 0xc48e72b186ea777045daf587e36192f87708ac30d768e5f8bc7284b85e6ad30c.
//
// Solidity: event TranslationSubmitted(address indexed submitter, bytes32 indexed dataHash, bytes32 indexed da, bytes cid)
func (_Contracts *ContractsFilterer) FilterTranslationSubmitted(opts *bind.FilterOpts, submitter []common.Address, dataHash [][32]byte, da [][32]byte) (*ContractsTranslationSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}
	var daRule []interface{}
	for _, daItem := range da {
		daRule = append(daRule, daItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TranslationSubmitted", submitterRule, dataHashRule, daRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTranslationSubmittedIterator{contract: _Contracts.contract, event: "TranslationSubmitted", logs: logs, sub: sub}, nil
}

// WatchTranslationSubmitted is a free log subscription operation binding the contract event 0xc48e72b186ea777045daf587e36192f87708ac30d768e5f8bc7284b85e6ad30c.
//
// Solidity: event TranslationSubmitted(address indexed submitter, bytes32 indexed dataHash, bytes32 indexed da, bytes cid)
func (_Contracts *ContractsFilterer) WatchTranslationSubmitted(opts *bind.WatchOpts, sink chan<- *ContractsTranslationSubmitted, submitter []common.Address, dataHash [][32]byte, da [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}
	var daRule []interface{}
	for _, daItem := range da {
		daRule = append(daRule, daItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TranslationSubmitted", submitterRule, dataHashRule, daRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTranslationSubmitted)
				if err := _Contracts.contract.UnpackLog(event, "TranslationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTranslationSubmitted is a log parse operation binding the contract event 0xc48e72b186ea777045daf587e36192f87708ac30d768e5f8bc7284b85e6ad30c.
//
// Solidity: event TranslationSubmitted(address indexed submitter, bytes32 indexed dataHash, bytes32 indexed da, bytes cid)
func (_Contracts *ContractsFilterer) ParseTranslationSubmitted(log types.Log) (*ContractsTranslationSubmitted, error) {
	event := new(ContractsTranslationSubmitted)
	if err := _Contracts.contract.UnpackLog(event, "TranslationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
