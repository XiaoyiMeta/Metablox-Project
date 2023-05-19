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

// IHeartbeatStatus is an auto generated low-level Go binding around an user-defined struct.
type IHeartbeatStatus struct {
	LatestBlockNumber *big.Int
	LatestTimestamp   *big.Int
}

// HeartbeatMetaData contains all meta data concerning the Heartbeat contract.
var HeartbeatMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeout_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Heartbeated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"TimeoutSetted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Validated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"heartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"isRegisterd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"latestMinerStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"latestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIHeartbeat.Status\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"latestValidatorStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"latestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIHeartbeat.Status\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeBatchTransferETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// HeartbeatABI is the input ABI used to generate the binding from.
// Deprecated: Use HeartbeatMetaData.ABI instead.
var HeartbeatABI = HeartbeatMetaData.ABI

// Heartbeat is an auto generated Go binding around an Ethereum contract.
type Heartbeat struct {
	HeartbeatCaller     // Read-only binding to the contract
	HeartbeatTransactor // Write-only binding to the contract
	HeartbeatFilterer   // Log filterer for contract events
}

// HeartbeatCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeartbeatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartbeatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeartbeatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartbeatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeartbeatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeartbeatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeartbeatSession struct {
	Contract     *Heartbeat        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeartbeatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeartbeatCallerSession struct {
	Contract *HeartbeatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HeartbeatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeartbeatTransactorSession struct {
	Contract     *HeartbeatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HeartbeatRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeartbeatRaw struct {
	Contract *Heartbeat // Generic contract binding to access the raw methods on
}

// HeartbeatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeartbeatCallerRaw struct {
	Contract *HeartbeatCaller // Generic read-only contract binding to access the raw methods on
}

// HeartbeatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeartbeatTransactorRaw struct {
	Contract *HeartbeatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeartbeat creates a new instance of Heartbeat, bound to a specific deployed contract.
func NewHeartbeat(address common.Address, backend bind.ContractBackend) (*Heartbeat, error) {
	contract, err := bindHeartbeat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Heartbeat{HeartbeatCaller: HeartbeatCaller{contract: contract}, HeartbeatTransactor: HeartbeatTransactor{contract: contract}, HeartbeatFilterer: HeartbeatFilterer{contract: contract}}, nil
}

// NewHeartbeatCaller creates a new read-only instance of Heartbeat, bound to a specific deployed contract.
func NewHeartbeatCaller(address common.Address, caller bind.ContractCaller) (*HeartbeatCaller, error) {
	contract, err := bindHeartbeat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeartbeatCaller{contract: contract}, nil
}

// NewHeartbeatTransactor creates a new write-only instance of Heartbeat, bound to a specific deployed contract.
func NewHeartbeatTransactor(address common.Address, transactor bind.ContractTransactor) (*HeartbeatTransactor, error) {
	contract, err := bindHeartbeat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeartbeatTransactor{contract: contract}, nil
}

// NewHeartbeatFilterer creates a new log filterer instance of Heartbeat, bound to a specific deployed contract.
func NewHeartbeatFilterer(address common.Address, filterer bind.ContractFilterer) (*HeartbeatFilterer, error) {
	contract, err := bindHeartbeat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeartbeatFilterer{contract: contract}, nil
}

// bindHeartbeat binds a generic wrapper to an already deployed contract.
func bindHeartbeat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HeartbeatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heartbeat *HeartbeatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Heartbeat.Contract.HeartbeatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heartbeat *HeartbeatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.Contract.HeartbeatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heartbeat *HeartbeatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heartbeat.Contract.HeartbeatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Heartbeat *HeartbeatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Heartbeat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Heartbeat *HeartbeatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Heartbeat *HeartbeatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Heartbeat.Contract.contract.Transact(opts, method, params...)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address miner) view returns(bool)
func (_Heartbeat *HeartbeatCaller) IsActive(opts *bind.CallOpts, miner common.Address) (bool, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "isActive", miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address miner) view returns(bool)
func (_Heartbeat *HeartbeatSession) IsActive(miner common.Address) (bool, error) {
	return _Heartbeat.Contract.IsActive(&_Heartbeat.CallOpts, miner)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address miner) view returns(bool)
func (_Heartbeat *HeartbeatCallerSession) IsActive(miner common.Address) (bool, error) {
	return _Heartbeat.Contract.IsActive(&_Heartbeat.CallOpts, miner)
}

// IsRegisterd is a free data retrieval call binding the contract method 0x22810053.
//
// Solidity: function isRegisterd(address miner) view returns(bool)
func (_Heartbeat *HeartbeatCaller) IsRegisterd(opts *bind.CallOpts, miner common.Address) (bool, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "isRegisterd", miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisterd is a free data retrieval call binding the contract method 0x22810053.
//
// Solidity: function isRegisterd(address miner) view returns(bool)
func (_Heartbeat *HeartbeatSession) IsRegisterd(miner common.Address) (bool, error) {
	return _Heartbeat.Contract.IsRegisterd(&_Heartbeat.CallOpts, miner)
}

// IsRegisterd is a free data retrieval call binding the contract method 0x22810053.
//
// Solidity: function isRegisterd(address miner) view returns(bool)
func (_Heartbeat *HeartbeatCallerSession) IsRegisterd(miner common.Address) (bool, error) {
	return _Heartbeat.Contract.IsRegisterd(&_Heartbeat.CallOpts, miner)
}

// LatestMinerStatus is a free data retrieval call binding the contract method 0x8f830164.
//
// Solidity: function latestMinerStatus(address miner) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatCaller) LatestMinerStatus(opts *bind.CallOpts, miner common.Address) (IHeartbeatStatus, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "latestMinerStatus", miner)

	if err != nil {
		return *new(IHeartbeatStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IHeartbeatStatus)).(*IHeartbeatStatus)

	return out0, err

}

// LatestMinerStatus is a free data retrieval call binding the contract method 0x8f830164.
//
// Solidity: function latestMinerStatus(address miner) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatSession) LatestMinerStatus(miner common.Address) (IHeartbeatStatus, error) {
	return _Heartbeat.Contract.LatestMinerStatus(&_Heartbeat.CallOpts, miner)
}

// LatestMinerStatus is a free data retrieval call binding the contract method 0x8f830164.
//
// Solidity: function latestMinerStatus(address miner) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatCallerSession) LatestMinerStatus(miner common.Address) (IHeartbeatStatus, error) {
	return _Heartbeat.Contract.LatestMinerStatus(&_Heartbeat.CallOpts, miner)
}

// LatestValidatorStatus is a free data retrieval call binding the contract method 0x0076005d.
//
// Solidity: function latestValidatorStatus(address validator) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatCaller) LatestValidatorStatus(opts *bind.CallOpts, validator common.Address) (IHeartbeatStatus, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "latestValidatorStatus", validator)

	if err != nil {
		return *new(IHeartbeatStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IHeartbeatStatus)).(*IHeartbeatStatus)

	return out0, err

}

// LatestValidatorStatus is a free data retrieval call binding the contract method 0x0076005d.
//
// Solidity: function latestValidatorStatus(address validator) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatSession) LatestValidatorStatus(validator common.Address) (IHeartbeatStatus, error) {
	return _Heartbeat.Contract.LatestValidatorStatus(&_Heartbeat.CallOpts, validator)
}

// LatestValidatorStatus is a free data retrieval call binding the contract method 0x0076005d.
//
// Solidity: function latestValidatorStatus(address validator) view returns((uint256,uint256))
func (_Heartbeat *HeartbeatCallerSession) LatestValidatorStatus(validator common.Address) (IHeartbeatStatus, error) {
	return _Heartbeat.Contract.LatestValidatorStatus(&_Heartbeat.CallOpts, validator)
}

// MinerNum is a free data retrieval call binding the contract method 0x6d425249.
//
// Solidity: function minerNum() view returns(uint256)
func (_Heartbeat *HeartbeatCaller) MinerNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "minerNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerNum is a free data retrieval call binding the contract method 0x6d425249.
//
// Solidity: function minerNum() view returns(uint256)
func (_Heartbeat *HeartbeatSession) MinerNum() (*big.Int, error) {
	return _Heartbeat.Contract.MinerNum(&_Heartbeat.CallOpts)
}

// MinerNum is a free data retrieval call binding the contract method 0x6d425249.
//
// Solidity: function minerNum() view returns(uint256)
func (_Heartbeat *HeartbeatCallerSession) MinerNum() (*big.Int, error) {
	return _Heartbeat.Contract.MinerNum(&_Heartbeat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heartbeat *HeartbeatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heartbeat *HeartbeatSession) Owner() (common.Address, error) {
	return _Heartbeat.Contract.Owner(&_Heartbeat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Heartbeat *HeartbeatCallerSession) Owner() (common.Address, error) {
	return _Heartbeat.Contract.Owner(&_Heartbeat.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Heartbeat *HeartbeatCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Heartbeat *HeartbeatSession) Paused() (bool, error) {
	return _Heartbeat.Contract.Paused(&_Heartbeat.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Heartbeat *HeartbeatCallerSession) Paused() (bool, error) {
	return _Heartbeat.Contract.Paused(&_Heartbeat.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Heartbeat *HeartbeatCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "timeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Heartbeat *HeartbeatSession) Timeout() (*big.Int, error) {
	return _Heartbeat.Contract.Timeout(&_Heartbeat.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Heartbeat *HeartbeatCallerSession) Timeout() (*big.Int, error) {
	return _Heartbeat.Contract.Timeout(&_Heartbeat.CallOpts)
}

// ValidatorNum is a free data retrieval call binding the contract method 0xd5a6151a.
//
// Solidity: function validatorNum() view returns(uint256)
func (_Heartbeat *HeartbeatCaller) ValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Heartbeat.contract.Call(opts, &out, "validatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorNum is a free data retrieval call binding the contract method 0xd5a6151a.
//
// Solidity: function validatorNum() view returns(uint256)
func (_Heartbeat *HeartbeatSession) ValidatorNum() (*big.Int, error) {
	return _Heartbeat.Contract.ValidatorNum(&_Heartbeat.CallOpts)
}

// ValidatorNum is a free data retrieval call binding the contract method 0xd5a6151a.
//
// Solidity: function validatorNum() view returns(uint256)
func (_Heartbeat *HeartbeatCallerSession) ValidatorNum() (*big.Int, error) {
	return _Heartbeat.Contract.ValidatorNum(&_Heartbeat.CallOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Heartbeat *HeartbeatTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Heartbeat *HeartbeatSession) DepositETH() (*types.Transaction, error) {
	return _Heartbeat.Contract.DepositETH(&_Heartbeat.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Heartbeat *HeartbeatTransactorSession) DepositETH() (*types.Transaction, error) {
	return _Heartbeat.Contract.DepositETH(&_Heartbeat.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x2a3cab6d.
//
// Solidity: function heartbeat(bytes data) returns()
func (_Heartbeat *HeartbeatTransactor) Heartbeat(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "heartbeat", data)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x2a3cab6d.
//
// Solidity: function heartbeat(bytes data) returns()
func (_Heartbeat *HeartbeatSession) Heartbeat(data []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Heartbeat(&_Heartbeat.TransactOpts, data)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x2a3cab6d.
//
// Solidity: function heartbeat(bytes data) returns()
func (_Heartbeat *HeartbeatTransactorSession) Heartbeat(data []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Heartbeat(&_Heartbeat.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Heartbeat *HeartbeatTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Heartbeat *HeartbeatSession) Pause() (*types.Transaction, error) {
	return _Heartbeat.Contract.Pause(&_Heartbeat.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Heartbeat *HeartbeatTransactorSession) Pause() (*types.Transaction, error) {
	return _Heartbeat.Contract.Pause(&_Heartbeat.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x43dc3ab2.
//
// Solidity: function register(address[] miners) returns()
func (_Heartbeat *HeartbeatTransactor) Register(opts *bind.TransactOpts, miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "register", miners)
}

// Register is a paid mutator transaction binding the contract method 0x43dc3ab2.
//
// Solidity: function register(address[] miners) returns()
func (_Heartbeat *HeartbeatSession) Register(miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.Register(&_Heartbeat.TransactOpts, miners)
}

// Register is a paid mutator transaction binding the contract method 0x43dc3ab2.
//
// Solidity: function register(address[] miners) returns()
func (_Heartbeat *HeartbeatTransactorSession) Register(miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.Register(&_Heartbeat.TransactOpts, miners)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heartbeat *HeartbeatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heartbeat *HeartbeatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Heartbeat.Contract.RenounceOwnership(&_Heartbeat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Heartbeat *HeartbeatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Heartbeat.Contract.RenounceOwnership(&_Heartbeat.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x05f203d9.
//
// Solidity: function revoke(address[] miners) returns()
func (_Heartbeat *HeartbeatTransactor) Revoke(opts *bind.TransactOpts, miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "revoke", miners)
}

// Revoke is a paid mutator transaction binding the contract method 0x05f203d9.
//
// Solidity: function revoke(address[] miners) returns()
func (_Heartbeat *HeartbeatSession) Revoke(miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.Revoke(&_Heartbeat.TransactOpts, miners)
}

// Revoke is a paid mutator transaction binding the contract method 0x05f203d9.
//
// Solidity: function revoke(address[] miners) returns()
func (_Heartbeat *HeartbeatTransactorSession) Revoke(miners []common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.Revoke(&_Heartbeat.TransactOpts, miners)
}

// Revoke0 is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_Heartbeat *HeartbeatTransactor) Revoke0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "revoke0")
}

// Revoke0 is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_Heartbeat *HeartbeatSession) Revoke0() (*types.Transaction, error) {
	return _Heartbeat.Contract.Revoke0(&_Heartbeat.TransactOpts)
}

// Revoke0 is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_Heartbeat *HeartbeatTransactorSession) Revoke0() (*types.Transaction, error) {
	return _Heartbeat.Contract.Revoke0(&_Heartbeat.TransactOpts)
}

// SafeBatchTransferETH is a paid mutator transaction binding the contract method 0x4e53b6ea.
//
// Solidity: function safeBatchTransferETH(address[] users, uint256 amount) returns()
func (_Heartbeat *HeartbeatTransactor) SafeBatchTransferETH(opts *bind.TransactOpts, users []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "safeBatchTransferETH", users, amount)
}

// SafeBatchTransferETH is a paid mutator transaction binding the contract method 0x4e53b6ea.
//
// Solidity: function safeBatchTransferETH(address[] users, uint256 amount) returns()
func (_Heartbeat *HeartbeatSession) SafeBatchTransferETH(users []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.SafeBatchTransferETH(&_Heartbeat.TransactOpts, users, amount)
}

// SafeBatchTransferETH is a paid mutator transaction binding the contract method 0x4e53b6ea.
//
// Solidity: function safeBatchTransferETH(address[] users, uint256 amount) returns()
func (_Heartbeat *HeartbeatTransactorSession) SafeBatchTransferETH(users []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.SafeBatchTransferETH(&_Heartbeat.TransactOpts, users, amount)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(uint256 interval) returns()
func (_Heartbeat *HeartbeatTransactor) SetTimeout(opts *bind.TransactOpts, interval *big.Int) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "setTimeout", interval)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(uint256 interval) returns()
func (_Heartbeat *HeartbeatSession) SetTimeout(interval *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.SetTimeout(&_Heartbeat.TransactOpts, interval)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(uint256 interval) returns()
func (_Heartbeat *HeartbeatTransactorSession) SetTimeout(interval *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.SetTimeout(&_Heartbeat.TransactOpts, interval)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heartbeat *HeartbeatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heartbeat *HeartbeatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.TransferOwnership(&_Heartbeat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Heartbeat *HeartbeatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Heartbeat.Contract.TransferOwnership(&_Heartbeat.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Heartbeat *HeartbeatTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Heartbeat *HeartbeatSession) Unpause() (*types.Transaction, error) {
	return _Heartbeat.Contract.Unpause(&_Heartbeat.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Heartbeat *HeartbeatTransactorSession) Unpause() (*types.Transaction, error) {
	return _Heartbeat.Contract.Unpause(&_Heartbeat.TransactOpts)
}

// Validate is a paid mutator transaction binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) returns()
func (_Heartbeat *HeartbeatTransactor) Validate(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "validate", data)
}

// Validate is a paid mutator transaction binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) returns()
func (_Heartbeat *HeartbeatSession) Validate(data []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Validate(&_Heartbeat.TransactOpts, data)
}

// Validate is a paid mutator transaction binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) returns()
func (_Heartbeat *HeartbeatTransactorSession) Validate(data []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Validate(&_Heartbeat.TransactOpts, data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 ethAmount) returns()
func (_Heartbeat *HeartbeatTransactor) WithdrawETH(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "withdrawETH", ethAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 ethAmount) returns()
func (_Heartbeat *HeartbeatSession) WithdrawETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.WithdrawETH(&_Heartbeat.TransactOpts, ethAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 ethAmount) returns()
func (_Heartbeat *HeartbeatTransactorSession) WithdrawETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.WithdrawETH(&_Heartbeat.TransactOpts, ethAmount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Heartbeat *HeartbeatTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Heartbeat *HeartbeatSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.WithdrawToken(&_Heartbeat.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Heartbeat *HeartbeatTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Heartbeat.Contract.WithdrawToken(&_Heartbeat.TransactOpts, token, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Heartbeat *HeartbeatTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Heartbeat.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Heartbeat *HeartbeatSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Fallback(&_Heartbeat.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Heartbeat *HeartbeatTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Heartbeat.Contract.Fallback(&_Heartbeat.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heartbeat *HeartbeatTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Heartbeat.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heartbeat *HeartbeatSession) Receive() (*types.Transaction, error) {
	return _Heartbeat.Contract.Receive(&_Heartbeat.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Heartbeat *HeartbeatTransactorSession) Receive() (*types.Transaction, error) {
	return _Heartbeat.Contract.Receive(&_Heartbeat.TransactOpts)
}

// HeartbeatHeartbeatedIterator is returned from FilterHeartbeated and is used to iterate over the raw logs and unpacked data for Heartbeated events raised by the Heartbeat contract.
type HeartbeatHeartbeatedIterator struct {
	Event *HeartbeatHeartbeated // Event containing the contract specifics and raw log

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
func (it *HeartbeatHeartbeatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatHeartbeated)
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
		it.Event = new(HeartbeatHeartbeated)
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
func (it *HeartbeatHeartbeatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatHeartbeatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatHeartbeated represents a Heartbeated event raised by the Heartbeat contract.
type HeartbeatHeartbeated struct {
	Miner common.Address
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHeartbeated is a free log retrieval operation binding the contract event 0xa0d72b080a8e22c25e687e386003e1acbc452f469c36713c06ad32a3d5046747.
//
// Solidity: event Heartbeated(address indexed miner, bytes data)
func (_Heartbeat *HeartbeatFilterer) FilterHeartbeated(opts *bind.FilterOpts, miner []common.Address) (*HeartbeatHeartbeatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Heartbeated", minerRule)
	if err != nil {
		return nil, err
	}
	return &HeartbeatHeartbeatedIterator{contract: _Heartbeat.contract, event: "Heartbeated", logs: logs, sub: sub}, nil
}

// WatchHeartbeated is a free log subscription operation binding the contract event 0xa0d72b080a8e22c25e687e386003e1acbc452f469c36713c06ad32a3d5046747.
//
// Solidity: event Heartbeated(address indexed miner, bytes data)
func (_Heartbeat *HeartbeatFilterer) WatchHeartbeated(opts *bind.WatchOpts, sink chan<- *HeartbeatHeartbeated, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Heartbeated", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatHeartbeated)
				if err := _Heartbeat.contract.UnpackLog(event, "Heartbeated", log); err != nil {
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

// ParseHeartbeated is a log parse operation binding the contract event 0xa0d72b080a8e22c25e687e386003e1acbc452f469c36713c06ad32a3d5046747.
//
// Solidity: event Heartbeated(address indexed miner, bytes data)
func (_Heartbeat *HeartbeatFilterer) ParseHeartbeated(log types.Log) (*HeartbeatHeartbeated, error) {
	event := new(HeartbeatHeartbeated)
	if err := _Heartbeat.contract.UnpackLog(event, "Heartbeated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Heartbeat contract.
type HeartbeatOwnershipTransferredIterator struct {
	Event *HeartbeatOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HeartbeatOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatOwnershipTransferred)
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
		it.Event = new(HeartbeatOwnershipTransferred)
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
func (it *HeartbeatOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatOwnershipTransferred represents a OwnershipTransferred event raised by the Heartbeat contract.
type HeartbeatOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Heartbeat *HeartbeatFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HeartbeatOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HeartbeatOwnershipTransferredIterator{contract: _Heartbeat.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Heartbeat *HeartbeatFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HeartbeatOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatOwnershipTransferred)
				if err := _Heartbeat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Heartbeat *HeartbeatFilterer) ParseOwnershipTransferred(log types.Log) (*HeartbeatOwnershipTransferred, error) {
	event := new(HeartbeatOwnershipTransferred)
	if err := _Heartbeat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Heartbeat contract.
type HeartbeatPausedIterator struct {
	Event *HeartbeatPaused // Event containing the contract specifics and raw log

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
func (it *HeartbeatPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatPaused)
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
		it.Event = new(HeartbeatPaused)
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
func (it *HeartbeatPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatPaused represents a Paused event raised by the Heartbeat contract.
type HeartbeatPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Heartbeat *HeartbeatFilterer) FilterPaused(opts *bind.FilterOpts) (*HeartbeatPausedIterator, error) {

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HeartbeatPausedIterator{contract: _Heartbeat.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Heartbeat *HeartbeatFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HeartbeatPaused) (event.Subscription, error) {

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatPaused)
				if err := _Heartbeat.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Heartbeat *HeartbeatFilterer) ParsePaused(log types.Log) (*HeartbeatPaused, error) {
	event := new(HeartbeatPaused)
	if err := _Heartbeat.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Heartbeat contract.
type HeartbeatRegisteredIterator struct {
	Event *HeartbeatRegistered // Event containing the contract specifics and raw log

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
func (it *HeartbeatRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatRegistered)
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
		it.Event = new(HeartbeatRegistered)
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
func (it *HeartbeatRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatRegistered represents a Registered event raised by the Heartbeat contract.
type HeartbeatRegistered struct {
	Miner     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) FilterRegistered(opts *bind.FilterOpts, miner []common.Address) (*HeartbeatRegisteredIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Registered", minerRule)
	if err != nil {
		return nil, err
	}
	return &HeartbeatRegisteredIterator{contract: _Heartbeat.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *HeartbeatRegistered, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Registered", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatRegistered)
				if err := _Heartbeat.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) ParseRegistered(log types.Log) (*HeartbeatRegistered, error) {
	event := new(HeartbeatRegistered)
	if err := _Heartbeat.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Heartbeat contract.
type HeartbeatRevokedIterator struct {
	Event *HeartbeatRevoked // Event containing the contract specifics and raw log

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
func (it *HeartbeatRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatRevoked)
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
		it.Event = new(HeartbeatRevoked)
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
func (it *HeartbeatRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatRevoked represents a Revoked event raised by the Heartbeat contract.
type HeartbeatRevoked struct {
	Miner     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) FilterRevoked(opts *bind.FilterOpts, miner []common.Address) (*HeartbeatRevokedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Revoked", minerRule)
	if err != nil {
		return nil, err
	}
	return &HeartbeatRevokedIterator{contract: _Heartbeat.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *HeartbeatRevoked, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Revoked", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatRevoked)
				if err := _Heartbeat.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed miner, uint256 timestamp)
func (_Heartbeat *HeartbeatFilterer) ParseRevoked(log types.Log) (*HeartbeatRevoked, error) {
	event := new(HeartbeatRevoked)
	if err := _Heartbeat.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatTimeoutSettedIterator is returned from FilterTimeoutSetted and is used to iterate over the raw logs and unpacked data for TimeoutSetted events raised by the Heartbeat contract.
type HeartbeatTimeoutSettedIterator struct {
	Event *HeartbeatTimeoutSetted // Event containing the contract specifics and raw log

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
func (it *HeartbeatTimeoutSettedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatTimeoutSetted)
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
		it.Event = new(HeartbeatTimeoutSetted)
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
func (it *HeartbeatTimeoutSettedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatTimeoutSettedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatTimeoutSetted represents a TimeoutSetted event raised by the Heartbeat contract.
type HeartbeatTimeoutSetted struct {
	OldTimeout *big.Int
	NewTimeout *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTimeoutSetted is a free log retrieval operation binding the contract event 0x591db2e6588f8eea133810d911600e5d7a1659f6d2122e2e57a72539b0ee326f.
//
// Solidity: event TimeoutSetted(uint256 oldTimeout, uint256 newTimeout)
func (_Heartbeat *HeartbeatFilterer) FilterTimeoutSetted(opts *bind.FilterOpts) (*HeartbeatTimeoutSettedIterator, error) {

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "TimeoutSetted")
	if err != nil {
		return nil, err
	}
	return &HeartbeatTimeoutSettedIterator{contract: _Heartbeat.contract, event: "TimeoutSetted", logs: logs, sub: sub}, nil
}

// WatchTimeoutSetted is a free log subscription operation binding the contract event 0x591db2e6588f8eea133810d911600e5d7a1659f6d2122e2e57a72539b0ee326f.
//
// Solidity: event TimeoutSetted(uint256 oldTimeout, uint256 newTimeout)
func (_Heartbeat *HeartbeatFilterer) WatchTimeoutSetted(opts *bind.WatchOpts, sink chan<- *HeartbeatTimeoutSetted) (event.Subscription, error) {

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "TimeoutSetted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatTimeoutSetted)
				if err := _Heartbeat.contract.UnpackLog(event, "TimeoutSetted", log); err != nil {
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

// ParseTimeoutSetted is a log parse operation binding the contract event 0x591db2e6588f8eea133810d911600e5d7a1659f6d2122e2e57a72539b0ee326f.
//
// Solidity: event TimeoutSetted(uint256 oldTimeout, uint256 newTimeout)
func (_Heartbeat *HeartbeatFilterer) ParseTimeoutSetted(log types.Log) (*HeartbeatTimeoutSetted, error) {
	event := new(HeartbeatTimeoutSetted)
	if err := _Heartbeat.contract.UnpackLog(event, "TimeoutSetted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Heartbeat contract.
type HeartbeatUnpausedIterator struct {
	Event *HeartbeatUnpaused // Event containing the contract specifics and raw log

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
func (it *HeartbeatUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatUnpaused)
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
		it.Event = new(HeartbeatUnpaused)
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
func (it *HeartbeatUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatUnpaused represents a Unpaused event raised by the Heartbeat contract.
type HeartbeatUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Heartbeat *HeartbeatFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HeartbeatUnpausedIterator, error) {

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HeartbeatUnpausedIterator{contract: _Heartbeat.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Heartbeat *HeartbeatFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HeartbeatUnpaused) (event.Subscription, error) {

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatUnpaused)
				if err := _Heartbeat.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Heartbeat *HeartbeatFilterer) ParseUnpaused(log types.Log) (*HeartbeatUnpaused, error) {
	event := new(HeartbeatUnpaused)
	if err := _Heartbeat.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeartbeatValidatedIterator is returned from FilterValidated and is used to iterate over the raw logs and unpacked data for Validated events raised by the Heartbeat contract.
type HeartbeatValidatedIterator struct {
	Event *HeartbeatValidated // Event containing the contract specifics and raw log

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
func (it *HeartbeatValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeartbeatValidated)
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
		it.Event = new(HeartbeatValidated)
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
func (it *HeartbeatValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeartbeatValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeartbeatValidated represents a Validated event raised by the Heartbeat contract.
type HeartbeatValidated struct {
	Validator common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidated is a free log retrieval operation binding the contract event 0x000b059bbd70fdb1e38071f518a6b73b9414364a6da2930f4cfa530355b6f48a.
//
// Solidity: event Validated(address indexed validator, bytes data)
func (_Heartbeat *HeartbeatFilterer) FilterValidated(opts *bind.FilterOpts, validator []common.Address) (*HeartbeatValidatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Heartbeat.contract.FilterLogs(opts, "Validated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &HeartbeatValidatedIterator{contract: _Heartbeat.contract, event: "Validated", logs: logs, sub: sub}, nil
}

// WatchValidated is a free log subscription operation binding the contract event 0x000b059bbd70fdb1e38071f518a6b73b9414364a6da2930f4cfa530355b6f48a.
//
// Solidity: event Validated(address indexed validator, bytes data)
func (_Heartbeat *HeartbeatFilterer) WatchValidated(opts *bind.WatchOpts, sink chan<- *HeartbeatValidated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Heartbeat.contract.WatchLogs(opts, "Validated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeartbeatValidated)
				if err := _Heartbeat.contract.UnpackLog(event, "Validated", log); err != nil {
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

// ParseValidated is a log parse operation binding the contract event 0x000b059bbd70fdb1e38071f518a6b73b9414364a6da2930f4cfa530355b6f48a.
//
// Solidity: event Validated(address indexed validator, bytes data)
func (_Heartbeat *HeartbeatFilterer) ParseValidated(log types.Log) (*HeartbeatValidated, error) {
	event := new(HeartbeatValidated)
	if err := _Heartbeat.contract.UnpackLog(event, "Validated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
