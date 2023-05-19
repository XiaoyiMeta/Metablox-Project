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

// CappedNFTStakingMetaData contains all meta data concerning the CappedNFTStaking contract.
var CappedNFTStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingNFT_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerSecond_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakePerAddress_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCap\",\"type\":\"uint256\"}],\"name\":\"CapacityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MaxStakePerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardLoaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingNFTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeElapsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_\",\"type\":\"address\"}],\"name\":\"isWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"loadReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLoaded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCap\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxStakePerAddress\",\"type\":\"uint256\"}],\"name\":\"setMaxStakePerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenIdEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenIdReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenIdStakingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"userEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"userReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"userStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"userTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CappedNFTStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use CappedNFTStakingMetaData.ABI instead.
var CappedNFTStakingABI = CappedNFTStakingMetaData.ABI

// CappedNFTStaking is an auto generated Go binding around an Ethereum contract.
type CappedNFTStaking struct {
	CappedNFTStakingCaller     // Read-only binding to the contract
	CappedNFTStakingTransactor // Write-only binding to the contract
	CappedNFTStakingFilterer   // Log filterer for contract events
}

// CappedNFTStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedNFTStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedNFTStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedNFTStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedNFTStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedNFTStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedNFTStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedNFTStakingSession struct {
	Contract     *CappedNFTStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedNFTStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedNFTStakingCallerSession struct {
	Contract *CappedNFTStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CappedNFTStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedNFTStakingTransactorSession struct {
	Contract     *CappedNFTStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CappedNFTStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedNFTStakingRaw struct {
	Contract *CappedNFTStaking // Generic contract binding to access the raw methods on
}

// CappedNFTStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedNFTStakingCallerRaw struct {
	Contract *CappedNFTStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CappedNFTStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedNFTStakingTransactorRaw struct {
	Contract *CappedNFTStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedNFTStaking creates a new instance of CappedNFTStaking, bound to a specific deployed contract.
func NewCappedNFTStaking(address common.Address, backend bind.ContractBackend) (*CappedNFTStaking, error) {
	contract, err := bindCappedNFTStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStaking{CappedNFTStakingCaller: CappedNFTStakingCaller{contract: contract}, CappedNFTStakingTransactor: CappedNFTStakingTransactor{contract: contract}, CappedNFTStakingFilterer: CappedNFTStakingFilterer{contract: contract}}, nil
}

// NewCappedNFTStakingCaller creates a new read-only instance of CappedNFTStaking, bound to a specific deployed contract.
func NewCappedNFTStakingCaller(address common.Address, caller bind.ContractCaller) (*CappedNFTStakingCaller, error) {
	contract, err := bindCappedNFTStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingCaller{contract: contract}, nil
}

// NewCappedNFTStakingTransactor creates a new write-only instance of CappedNFTStaking, bound to a specific deployed contract.
func NewCappedNFTStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedNFTStakingTransactor, error) {
	contract, err := bindCappedNFTStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingTransactor{contract: contract}, nil
}

// NewCappedNFTStakingFilterer creates a new log filterer instance of CappedNFTStaking, bound to a specific deployed contract.
func NewCappedNFTStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedNFTStakingFilterer, error) {
	contract, err := bindCappedNFTStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingFilterer{contract: contract}, nil
}

// bindCappedNFTStaking binds a generic wrapper to an already deployed contract.
func bindCappedNFTStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CappedNFTStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedNFTStaking *CappedNFTStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedNFTStaking.Contract.CappedNFTStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedNFTStaking *CappedNFTStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.CappedNFTStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedNFTStaking *CappedNFTStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.CappedNFTStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedNFTStaking *CappedNFTStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedNFTStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedNFTStaking *CappedNFTStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedNFTStaking *CappedNFTStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.contract.Transact(opts, method, params...)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "capacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) Capacity() (*big.Int, error) {
	return _CappedNFTStaking.Contract.Capacity(&_CappedNFTStaking.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) Capacity() (*big.Int, error) {
	return _CappedNFTStaking.Contract.Capacity(&_CappedNFTStaking.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) Duration() (*big.Int, error) {
	return _CappedNFTStaking.Contract.Duration(&_CappedNFTStaking.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) Duration() (*big.Int, error) {
	return _CappedNFTStaking.Contract.Duration(&_CappedNFTStaking.CallOpts)
}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) GetRewardTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "getRewardTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) GetRewardTokenBalance() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetRewardTokenBalance(&_CappedNFTStaking.CallOpts)
}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) GetRewardTokenBalance() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetRewardTokenBalance(&_CappedNFTStaking.CallOpts)
}

// GetStakingNFTBalance is a free data retrieval call binding the contract method 0x61db99aa.
//
// Solidity: function getStakingNFTBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) GetStakingNFTBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "getStakingNFTBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingNFTBalance is a free data retrieval call binding the contract method 0x61db99aa.
//
// Solidity: function getStakingNFTBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) GetStakingNFTBalance() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetStakingNFTBalance(&_CappedNFTStaking.CallOpts)
}

// GetStakingNFTBalance is a free data retrieval call binding the contract method 0x61db99aa.
//
// Solidity: function getStakingNFTBalance() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) GetStakingNFTBalance() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetStakingNFTBalance(&_CappedNFTStaking.CallOpts)
}

// GetTimeElapsed is a free data retrieval call binding the contract method 0x4faa2d54.
//
// Solidity: function getTimeElapsed() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) GetTimeElapsed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "getTimeElapsed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeElapsed is a free data retrieval call binding the contract method 0x4faa2d54.
//
// Solidity: function getTimeElapsed() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) GetTimeElapsed() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetTimeElapsed(&_CappedNFTStaking.CallOpts)
}

// GetTimeElapsed is a free data retrieval call binding the contract method 0x4faa2d54.
//
// Solidity: function getTimeElapsed() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) GetTimeElapsed() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetTimeElapsed(&_CappedNFTStaking.CallOpts)
}

// GetTimeRemaining is a free data retrieval call binding the contract method 0xdac6270d.
//
// Solidity: function getTimeRemaining() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) GetTimeRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "getTimeRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeRemaining is a free data retrieval call binding the contract method 0xdac6270d.
//
// Solidity: function getTimeRemaining() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) GetTimeRemaining() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetTimeRemaining(&_CappedNFTStaking.CallOpts)
}

// GetTimeRemaining is a free data retrieval call binding the contract method 0xdac6270d.
//
// Solidity: function getTimeRemaining() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) GetTimeRemaining() (*big.Int, error) {
	return _CappedNFTStaking.Contract.GetTimeRemaining(&_CappedNFTStaking.CallOpts)
}

// IsWithdrawn is a free data retrieval call binding the contract method 0xa22c4ad0.
//
// Solidity: function isWithdrawn(address staker_) view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCaller) IsWithdrawn(opts *bind.CallOpts, staker_ common.Address) (bool, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "isWithdrawn", staker_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawn is a free data retrieval call binding the contract method 0xa22c4ad0.
//
// Solidity: function isWithdrawn(address staker_) view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingSession) IsWithdrawn(staker_ common.Address) (bool, error) {
	return _CappedNFTStaking.Contract.IsWithdrawn(&_CappedNFTStaking.CallOpts, staker_)
}

// IsWithdrawn is a free data retrieval call binding the contract method 0xa22c4ad0.
//
// Solidity: function isWithdrawn(address staker_) view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) IsWithdrawn(staker_ common.Address) (bool, error) {
	return _CappedNFTStaking.Contract.IsWithdrawn(&_CappedNFTStaking.CallOpts, staker_)
}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) MaxStakePerAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "maxStakePerAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) MaxStakePerAddress() (*big.Int, error) {
	return _CappedNFTStaking.Contract.MaxStakePerAddress(&_CappedNFTStaking.CallOpts)
}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) MaxStakePerAddress() (*big.Int, error) {
	return _CappedNFTStaking.Contract.MaxStakePerAddress(&_CappedNFTStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingSession) Owner() (common.Address, error) {
	return _CappedNFTStaking.Contract.Owner(&_CappedNFTStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) Owner() (common.Address, error) {
	return _CappedNFTStaking.Contract.Owner(&_CappedNFTStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingSession) Paused() (bool, error) {
	return _CappedNFTStaking.Contract.Paused(&_CappedNFTStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) Paused() (bool, error) {
	return _CappedNFTStaking.Contract.Paused(&_CappedNFTStaking.CallOpts)
}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCaller) RewardLoaded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "rewardLoaded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingSession) RewardLoaded() (bool, error) {
	return _CappedNFTStaking.Contract.RewardLoaded(&_CappedNFTStaking.CallOpts)
}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) RewardLoaded() (bool, error) {
	return _CappedNFTStaking.Contract.RewardLoaded(&_CappedNFTStaking.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) RewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "rewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) RewardPerSecond() (*big.Int, error) {
	return _CappedNFTStaking.Contract.RewardPerSecond(&_CappedNFTStaking.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) RewardPerSecond() (*big.Int, error) {
	return _CappedNFTStaking.Contract.RewardPerSecond(&_CappedNFTStaking.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingSession) RewardToken() (common.Address, error) {
	return _CappedNFTStaking.Contract.RewardToken(&_CappedNFTStaking.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) RewardToken() (common.Address, error) {
	return _CappedNFTStaking.Contract.RewardToken(&_CappedNFTStaking.CallOpts)
}

// StakingNFT is a free data retrieval call binding the contract method 0x0ce71e32.
//
// Solidity: function stakingNFT() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCaller) StakingNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "stakingNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingNFT is a free data retrieval call binding the contract method 0x0ce71e32.
//
// Solidity: function stakingNFT() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingSession) StakingNFT() (common.Address, error) {
	return _CappedNFTStaking.Contract.StakingNFT(&_CappedNFTStaking.CallOpts)
}

// StakingNFT is a free data retrieval call binding the contract method 0x0ce71e32.
//
// Solidity: function stakingNFT() view returns(address)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) StakingNFT() (common.Address, error) {
	return _CappedNFTStaking.Contract.StakingNFT(&_CappedNFTStaking.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) StartTime() (*big.Int, error) {
	return _CappedNFTStaking.Contract.StartTime(&_CappedNFTStaking.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) StartTime() (*big.Int, error) {
	return _CappedNFTStaking.Contract.StartTime(&_CappedNFTStaking.CallOpts)
}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) StopTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "stopTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) StopTime() (*big.Int, error) {
	return _CappedNFTStaking.Contract.StopTime(&_CappedNFTStaking.CallOpts)
}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) StopTime() (*big.Int, error) {
	return _CappedNFTStaking.Contract.StopTime(&_CappedNFTStaking.CallOpts)
}

// TokenIdEarned is a free data retrieval call binding the contract method 0x8500e0d8.
//
// Solidity: function tokenIdEarned(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TokenIdEarned(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "tokenIdEarned", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdEarned is a free data retrieval call binding the contract method 0x8500e0d8.
//
// Solidity: function tokenIdEarned(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TokenIdEarned(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdEarned(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TokenIdEarned is a free data retrieval call binding the contract method 0x8500e0d8.
//
// Solidity: function tokenIdEarned(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TokenIdEarned(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdEarned(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TokenIdReward is a free data retrieval call binding the contract method 0xcb76599d.
//
// Solidity: function tokenIdReward(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TokenIdReward(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "tokenIdReward", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdReward is a free data retrieval call binding the contract method 0xcb76599d.
//
// Solidity: function tokenIdReward(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TokenIdReward(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdReward(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TokenIdReward is a free data retrieval call binding the contract method 0xcb76599d.
//
// Solidity: function tokenIdReward(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TokenIdReward(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdReward(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TokenIdStakingTime is a free data retrieval call binding the contract method 0x8d260099.
//
// Solidity: function tokenIdStakingTime(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TokenIdStakingTime(opts *bind.CallOpts, tokenId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "tokenIdStakingTime", tokenId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdStakingTime is a free data retrieval call binding the contract method 0x8d260099.
//
// Solidity: function tokenIdStakingTime(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TokenIdStakingTime(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdStakingTime(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TokenIdStakingTime is a free data retrieval call binding the contract method 0x8d260099.
//
// Solidity: function tokenIdStakingTime(uint256 tokenId_) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TokenIdStakingTime(tokenId_ *big.Int) (*big.Int, error) {
	return _CappedNFTStaking.Contract.TokenIdStakingTime(&_CappedNFTStaking.CallOpts, tokenId_)
}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "totalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TotalReward() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalReward(&_CappedNFTStaking.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0x750142e6.
//
// Solidity: function totalReward() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TotalReward() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalReward(&_CappedNFTStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TotalStaked() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalStaked(&_CappedNFTStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TotalStaked() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalStaked(&_CappedNFTStaking.CallOpts)
}

// TotalStakedSeconds is a free data retrieval call binding the contract method 0x147b4007.
//
// Solidity: function totalStakedSeconds() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) TotalStakedSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "totalStakedSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedSeconds is a free data retrieval call binding the contract method 0x147b4007.
//
// Solidity: function totalStakedSeconds() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) TotalStakedSeconds() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalStakedSeconds(&_CappedNFTStaking.CallOpts)
}

// TotalStakedSeconds is a free data retrieval call binding the contract method 0x147b4007.
//
// Solidity: function totalStakedSeconds() view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) TotalStakedSeconds() (*big.Int, error) {
	return _CappedNFTStaking.Contract.TotalStakedSeconds(&_CappedNFTStaking.CallOpts)
}

// UserEarned is a free data retrieval call binding the contract method 0xcba72540.
//
// Solidity: function userEarned(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) UserEarned(opts *bind.CallOpts, _stakeholder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "userEarned", _stakeholder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserEarned is a free data retrieval call binding the contract method 0xcba72540.
//
// Solidity: function userEarned(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) UserEarned(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserEarned(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserEarned is a free data retrieval call binding the contract method 0xcba72540.
//
// Solidity: function userEarned(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) UserEarned(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserEarned(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) UserReward(opts *bind.CallOpts, _stakeholder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "userReward", _stakeholder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) UserReward(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserReward(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) UserReward(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserReward(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserStaked is a free data retrieval call binding the contract method 0xacc3a939.
//
// Solidity: function userStaked(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCaller) UserStaked(opts *bind.CallOpts, _stakeholder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "userStaked", _stakeholder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStaked is a free data retrieval call binding the contract method 0xacc3a939.
//
// Solidity: function userStaked(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingSession) UserStaked(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserStaked(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserStaked is a free data retrieval call binding the contract method 0xacc3a939.
//
// Solidity: function userStaked(address _stakeholder) view returns(uint256)
func (_CappedNFTStaking *CappedNFTStakingCallerSession) UserStaked(_stakeholder common.Address) (*big.Int, error) {
	return _CappedNFTStaking.Contract.UserStaked(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserTokenIds is a free data retrieval call binding the contract method 0xe8c40392.
//
// Solidity: function userTokenIds(address _stakeholder) view returns(uint256[])
func (_CappedNFTStaking *CappedNFTStakingCaller) UserTokenIds(opts *bind.CallOpts, _stakeholder common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CappedNFTStaking.contract.Call(opts, &out, "userTokenIds", _stakeholder)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// UserTokenIds is a free data retrieval call binding the contract method 0xe8c40392.
//
// Solidity: function userTokenIds(address _stakeholder) view returns(uint256[])
func (_CappedNFTStaking *CappedNFTStakingSession) UserTokenIds(_stakeholder common.Address) ([]*big.Int, error) {
	return _CappedNFTStaking.Contract.UserTokenIds(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// UserTokenIds is a free data retrieval call binding the contract method 0xe8c40392.
//
// Solidity: function userTokenIds(address _stakeholder) view returns(uint256[])
func (_CappedNFTStaking *CappedNFTStakingCallerSession) UserTokenIds(_stakeholder common.Address) ([]*big.Int, error) {
	return _CappedNFTStaking.Contract.UserTokenIds(&_CappedNFTStaking.CallOpts, _stakeholder)
}

// LoadReward is a paid mutator transaction binding the contract method 0x9b6c286d.
//
// Solidity: function loadReward(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) LoadReward(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "loadReward", token_)
}

// LoadReward is a paid mutator transaction binding the contract method 0x9b6c286d.
//
// Solidity: function loadReward(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) LoadReward(token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.LoadReward(&_CappedNFTStaking.TransactOpts, token_)
}

// LoadReward is a paid mutator transaction binding the contract method 0x9b6c286d.
//
// Solidity: function loadReward(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) LoadReward(token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.LoadReward(&_CappedNFTStaking.TransactOpts, token_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CappedNFTStaking *CappedNFTStakingTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CappedNFTStaking *CappedNFTStakingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.OnERC721Received(&_CappedNFTStaking.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.OnERC721Received(&_CappedNFTStaking.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedNFTStaking *CappedNFTStakingSession) Pause() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Pause(&_CappedNFTStaking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) Pause() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Pause(&_CappedNFTStaking.TransactOpts)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) RecoverTokens(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "recoverTokens", token_)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) RecoverTokens(token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.RecoverTokens(&_CappedNFTStaking.TransactOpts, token_)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address token_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) RecoverTokens(token_ common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.RecoverTokens(&_CappedNFTStaking.TransactOpts, token_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedNFTStaking *CappedNFTStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.RenounceOwnership(&_CappedNFTStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.RenounceOwnership(&_CappedNFTStaking.TransactOpts)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 _newCap) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) SetCapacity(opts *bind.TransactOpts, _newCap *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "setCapacity", _newCap)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 _newCap) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) SetCapacity(_newCap *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.SetCapacity(&_CappedNFTStaking.TransactOpts, _newCap)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 _newCap) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) SetCapacity(_newCap *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.SetCapacity(&_CappedNFTStaking.TransactOpts, _newCap)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) SetMaxStakePerAddress(opts *bind.TransactOpts, _newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "setMaxStakePerAddress", _newMaxStakePerAddress)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) SetMaxStakePerAddress(_newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.SetMaxStakePerAddress(&_CappedNFTStaking.TransactOpts, _newMaxStakePerAddress)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) SetMaxStakePerAddress(_newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.SetMaxStakePerAddress(&_CappedNFTStaking.TransactOpts, _newMaxStakePerAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) Stake(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "stake", tokenIds_)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) Stake(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Stake(&_CappedNFTStaking.TransactOpts, tokenIds_)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) Stake(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Stake(&_CappedNFTStaking.TransactOpts, tokenIds_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.TransferOwnership(&_CappedNFTStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.TransferOwnership(&_CappedNFTStaking.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedNFTStaking *CappedNFTStakingSession) Unpause() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Unpause(&_CappedNFTStaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Unpause(&_CappedNFTStaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) Unstake(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "unstake", tokenIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingSession) Unstake(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Unstake(&_CappedNFTStaking.TransactOpts, tokenIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds_) returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) Unstake(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Unstake(&_CappedNFTStaking.TransactOpts, tokenIds_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedNFTStaking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedNFTStaking *CappedNFTStakingSession) Withdraw() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Withdraw(&_CappedNFTStaking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedNFTStaking *CappedNFTStakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CappedNFTStaking.Contract.Withdraw(&_CappedNFTStaking.TransactOpts)
}

// CappedNFTStakingCapacityChangedIterator is returned from FilterCapacityChanged and is used to iterate over the raw logs and unpacked data for CapacityChanged events raised by the CappedNFTStaking contract.
type CappedNFTStakingCapacityChangedIterator struct {
	Event *CappedNFTStakingCapacityChanged // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingCapacityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingCapacityChanged)
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
		it.Event = new(CappedNFTStakingCapacityChanged)
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
func (it *CappedNFTStakingCapacityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingCapacityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingCapacityChanged represents a CapacityChanged event raised by the CappedNFTStaking contract.
type CappedNFTStakingCapacityChanged struct {
	OldCap *big.Int
	NewCap *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCapacityChanged is a free log retrieval operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterCapacityChanged(opts *bind.FilterOpts) (*CappedNFTStakingCapacityChangedIterator, error) {

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "CapacityChanged")
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingCapacityChangedIterator{contract: _CappedNFTStaking.contract, event: "CapacityChanged", logs: logs, sub: sub}, nil
}

// WatchCapacityChanged is a free log subscription operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchCapacityChanged(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingCapacityChanged) (event.Subscription, error) {

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "CapacityChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingCapacityChanged)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "CapacityChanged", log); err != nil {
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

// ParseCapacityChanged is a log parse operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseCapacityChanged(log types.Log) (*CappedNFTStakingCapacityChanged, error) {
	event := new(CappedNFTStakingCapacityChanged)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "CapacityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingMaxStakePerAddressChangedIterator is returned from FilterMaxStakePerAddressChanged and is used to iterate over the raw logs and unpacked data for MaxStakePerAddressChanged events raised by the CappedNFTStaking contract.
type CappedNFTStakingMaxStakePerAddressChangedIterator struct {
	Event *CappedNFTStakingMaxStakePerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingMaxStakePerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingMaxStakePerAddressChanged)
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
		it.Event = new(CappedNFTStakingMaxStakePerAddressChanged)
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
func (it *CappedNFTStakingMaxStakePerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingMaxStakePerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingMaxStakePerAddressChanged represents a MaxStakePerAddressChanged event raised by the CappedNFTStaking contract.
type CappedNFTStakingMaxStakePerAddressChanged struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxStakePerAddressChanged is a free log retrieval operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterMaxStakePerAddressChanged(opts *bind.FilterOpts) (*CappedNFTStakingMaxStakePerAddressChangedIterator, error) {

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "MaxStakePerAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingMaxStakePerAddressChangedIterator{contract: _CappedNFTStaking.contract, event: "MaxStakePerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchMaxStakePerAddressChanged is a free log subscription operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchMaxStakePerAddressChanged(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingMaxStakePerAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "MaxStakePerAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingMaxStakePerAddressChanged)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "MaxStakePerAddressChanged", log); err != nil {
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

// ParseMaxStakePerAddressChanged is a log parse operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseMaxStakePerAddressChanged(log types.Log) (*CappedNFTStakingMaxStakePerAddressChanged, error) {
	event := new(CappedNFTStakingMaxStakePerAddressChanged)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "MaxStakePerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CappedNFTStaking contract.
type CappedNFTStakingOwnershipTransferredIterator struct {
	Event *CappedNFTStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingOwnershipTransferred)
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
		it.Event = new(CappedNFTStakingOwnershipTransferred)
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
func (it *CappedNFTStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingOwnershipTransferred represents a OwnershipTransferred event raised by the CappedNFTStaking contract.
type CappedNFTStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CappedNFTStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingOwnershipTransferredIterator{contract: _CappedNFTStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingOwnershipTransferred)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseOwnershipTransferred(log types.Log) (*CappedNFTStakingOwnershipTransferred, error) {
	event := new(CappedNFTStakingOwnershipTransferred)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CappedNFTStaking contract.
type CappedNFTStakingPausedIterator struct {
	Event *CappedNFTStakingPaused // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingPaused)
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
		it.Event = new(CappedNFTStakingPaused)
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
func (it *CappedNFTStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingPaused represents a Paused event raised by the CappedNFTStaking contract.
type CappedNFTStakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterPaused(opts *bind.FilterOpts) (*CappedNFTStakingPausedIterator, error) {

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingPausedIterator{contract: _CappedNFTStaking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingPaused) (event.Subscription, error) {

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingPaused)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParsePaused(log types.Log) (*CappedNFTStakingPaused, error) {
	event := new(CappedNFTStakingPaused)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the CappedNFTStaking contract.
type CappedNFTStakingRecoveredIterator struct {
	Event *CappedNFTStakingRecovered // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingRecovered)
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
		it.Event = new(CappedNFTStakingRecovered)
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
func (it *CappedNFTStakingRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingRecovered represents a Recovered event raised by the CappedNFTStaking contract.
type CappedNFTStakingRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterRecovered(opts *bind.FilterOpts, token []common.Address) (*CappedNFTStakingRecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Recovered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingRecoveredIterator{contract: _CappedNFTStaking.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingRecovered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Recovered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingRecovered)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseRecovered(log types.Log) (*CappedNFTStakingRecovered, error) {
	event := new(CappedNFTStakingRecovered)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingRewardLoadedIterator is returned from FilterRewardLoaded and is used to iterate over the raw logs and unpacked data for RewardLoaded events raised by the CappedNFTStaking contract.
type CappedNFTStakingRewardLoadedIterator struct {
	Event *CappedNFTStakingRewardLoaded // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingRewardLoadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingRewardLoaded)
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
		it.Event = new(CappedNFTStakingRewardLoaded)
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
func (it *CappedNFTStakingRewardLoadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingRewardLoadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingRewardLoaded represents a RewardLoaded event raised by the CappedNFTStaking contract.
type CappedNFTStakingRewardLoaded struct {
	RewardToken  common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardLoaded is a free log retrieval operation binding the contract event 0x774329b2d3136032a02608fe7642480dcd3921778de629fb725ff8bc15e68807.
//
// Solidity: event RewardLoaded(address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterRewardLoaded(opts *bind.FilterOpts, rewardToken []common.Address) (*CappedNFTStakingRewardLoadedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "RewardLoaded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingRewardLoadedIterator{contract: _CappedNFTStaking.contract, event: "RewardLoaded", logs: logs, sub: sub}, nil
}

// WatchRewardLoaded is a free log subscription operation binding the contract event 0x774329b2d3136032a02608fe7642480dcd3921778de629fb725ff8bc15e68807.
//
// Solidity: event RewardLoaded(address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchRewardLoaded(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingRewardLoaded, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "RewardLoaded", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingRewardLoaded)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "RewardLoaded", log); err != nil {
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

// ParseRewardLoaded is a log parse operation binding the contract event 0x774329b2d3136032a02608fe7642480dcd3921778de629fb725ff8bc15e68807.
//
// Solidity: event RewardLoaded(address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseRewardLoaded(log types.Log) (*CappedNFTStakingRewardLoaded, error) {
	event := new(CappedNFTStakingRewardLoaded)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "RewardLoaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CappedNFTStaking contract.
type CappedNFTStakingStakedIterator struct {
	Event *CappedNFTStakingStaked // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingStaked)
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
		it.Event = new(CappedNFTStakingStaked)
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
func (it *CappedNFTStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingStaked represents a Staked event raised by the CappedNFTStaking contract.
type CappedNFTStakingStaked struct {
	Staker  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*CappedNFTStakingStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingStakedIterator{contract: _CappedNFTStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingStaked)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseStaked(log types.Log) (*CappedNFTStakingStaked, error) {
	event := new(CappedNFTStakingStaked)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CappedNFTStaking contract.
type CappedNFTStakingUnpausedIterator struct {
	Event *CappedNFTStakingUnpaused // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingUnpaused)
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
		it.Event = new(CappedNFTStakingUnpaused)
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
func (it *CappedNFTStakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingUnpaused represents a Unpaused event raised by the CappedNFTStaking contract.
type CappedNFTStakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CappedNFTStakingUnpausedIterator, error) {

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingUnpausedIterator{contract: _CappedNFTStaking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingUnpaused)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseUnpaused(log types.Log) (*CappedNFTStakingUnpaused, error) {
	event := new(CappedNFTStakingUnpaused)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the CappedNFTStaking contract.
type CappedNFTStakingUnstakedIterator struct {
	Event *CappedNFTStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingUnstaked)
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
		it.Event = new(CappedNFTStakingUnstaked)
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
func (it *CappedNFTStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingUnstaked represents a Unstaked event raised by the CappedNFTStaking contract.
type CappedNFTStakingUnstaked struct {
	Staker  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, staker []common.Address) (*CappedNFTStakingUnstakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Unstaked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingUnstakedIterator{contract: _CappedNFTStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingUnstaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Unstaked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingUnstaked)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed staker, uint256 tokenId)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseUnstaked(log types.Log) (*CappedNFTStakingUnstaked, error) {
	event := new(CappedNFTStakingUnstaked)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedNFTStakingWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the CappedNFTStaking contract.
type CappedNFTStakingWithdrawalIterator struct {
	Event *CappedNFTStakingWithdrawal // Event containing the contract specifics and raw log

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
func (it *CappedNFTStakingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedNFTStakingWithdrawal)
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
		it.Event = new(CappedNFTStakingWithdrawal)
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
func (it *CappedNFTStakingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedNFTStakingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedNFTStakingWithdrawal represents a Withdrawal event raised by the CappedNFTStaking contract.
type CappedNFTStakingWithdrawal struct {
	Staker       common.Address
	RewardToken  common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed staker, address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) FilterWithdrawal(opts *bind.FilterOpts, staker []common.Address, rewardToken []common.Address) (*CappedNFTStakingWithdrawalIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.FilterLogs(opts, "Withdrawal", stakerRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &CappedNFTStakingWithdrawalIterator{contract: _CappedNFTStaking.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed staker, address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *CappedNFTStakingWithdrawal, staker []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CappedNFTStaking.contract.WatchLogs(opts, "Withdrawal", stakerRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedNFTStakingWithdrawal)
				if err := _CappedNFTStaking.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed staker, address indexed rewardToken, uint256 rewardAmount)
func (_CappedNFTStaking *CappedNFTStakingFilterer) ParseWithdrawal(log types.Log) (*CappedNFTStakingWithdrawal, error) {
	event := new(CappedNFTStakingWithdrawal)
	if err := _CappedNFTStaking.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
