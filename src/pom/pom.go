// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pom

import (
	"errors"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum"
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
)

// POMMeshBoxInfo is an auto generated low-level Go binding around an user-defined struct.
type POMMeshBoxInfo struct {
	Id   string
	Addr common.Address
}

// POMPOMChallenge is an auto generated low-level Go binding around an user-defined struct.
type POMPOMChallenge struct {
	Challenger common.Address
	Target     common.Address
	Witness    []common.Address
}

// PomMetaData contains all meta data concerning the Pom contract.
var PomMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddMeshBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"GetTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RemoveMeshBox\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addMeshBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChallengeList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"witness\",\"type\":\"address[]\"}],\"internalType\":\"structPOM.POMChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"getLastChallengeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMeshBoxList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structPOM.MeshBoxInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeMeshBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"sendPOMReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendPomEpochReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PomABI is the input ABI used to generate the binding from.
// Deprecated: Use PomMetaData.ABI instead.
var PomABI = PomMetaData.ABI

// Pom is an auto generated Go binding around an Ethereum contract.
type Pom struct {
	PomCaller     // Read-only binding to the contract
	PomTransactor // Write-only binding to the contract
	PomFilterer   // Log filterer for contract events
}

// PomCaller is an auto generated read-only Go binding around an Ethereum contract.
type PomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PomSession struct {
	Contract     *Pom              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PomCallerSession struct {
	Contract *PomCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PomTransactorSession struct {
	Contract     *PomTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PomRaw is an auto generated low-level Go binding around an Ethereum contract.
type PomRaw struct {
	Contract *Pom // Generic contract binding to access the raw methods on
}

// PomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PomCallerRaw struct {
	Contract *PomCaller // Generic read-only contract binding to access the raw methods on
}

// PomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PomTransactorRaw struct {
	Contract *PomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPom creates a new instance of Pom, bound to a specific deployed contract.
func NewPom(address common.Address, backend bind.ContractBackend) (*Pom, error) {
	contract, err := bindPom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pom{PomCaller: PomCaller{contract: contract}, PomTransactor: PomTransactor{contract: contract}, PomFilterer: PomFilterer{contract: contract}}, nil
}

// NewPomCaller creates a new read-only instance of Pom, bound to a specific deployed contract.
func NewPomCaller(address common.Address, caller bind.ContractCaller) (*PomCaller, error) {
	contract, err := bindPom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PomCaller{contract: contract}, nil
}

// NewPomTransactor creates a new write-only instance of Pom, bound to a specific deployed contract.
func NewPomTransactor(address common.Address, transactor bind.ContractTransactor) (*PomTransactor, error) {
	contract, err := bindPom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PomTransactor{contract: contract}, nil
}

// NewPomFilterer creates a new log filterer instance of Pom, bound to a specific deployed contract.
func NewPomFilterer(address common.Address, filterer bind.ContractFilterer) (*PomFilterer, error) {
	contract, err := bindPom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PomFilterer{contract: contract}, nil
}

// bindPom binds a generic wrapper to an already deployed contract.
func bindPom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pom *PomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pom.Contract.PomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pom *PomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pom.Contract.PomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pom *PomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pom.Contract.PomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pom *PomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pom *PomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pom *PomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pom.Contract.contract.Transact(opts, method, params...)
}

// ChallengeDelay is a free data retrieval call binding the contract method 0x6d4c9dea.
//
// Solidity: function challengeDelay() view returns(uint256)
func (_Pom *PomCaller) ChallengeDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "challengeDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeDelay is a free data retrieval call binding the contract method 0x6d4c9dea.
//
// Solidity: function challengeDelay() view returns(uint256)
func (_Pom *PomSession) ChallengeDelay() (*big.Int, error) {
	return _Pom.Contract.ChallengeDelay(&_Pom.CallOpts)
}

// ChallengeDelay is a free data retrieval call binding the contract method 0x6d4c9dea.
//
// Solidity: function challengeDelay() view returns(uint256)
func (_Pom *PomCallerSession) ChallengeDelay() (*big.Int, error) {
	return _Pom.Contract.ChallengeDelay(&_Pom.CallOpts)
}

// ChallengeInterval is a free data retrieval call binding the contract method 0x5171f8b6.
//
// Solidity: function challengeInterval() view returns(uint256)
func (_Pom *PomCaller) ChallengeInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "challengeInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeInterval is a free data retrieval call binding the contract method 0x5171f8b6.
//
// Solidity: function challengeInterval() view returns(uint256)
func (_Pom *PomSession) ChallengeInterval() (*big.Int, error) {
	return _Pom.Contract.ChallengeInterval(&_Pom.CallOpts)
}

// ChallengeInterval is a free data retrieval call binding the contract method 0x5171f8b6.
//
// Solidity: function challengeInterval() view returns(uint256)
func (_Pom *PomCallerSession) ChallengeInterval() (*big.Int, error) {
	return _Pom.Contract.ChallengeInterval(&_Pom.CallOpts)
}

// GetChallengeList is a free data retrieval call binding the contract method 0xf8c57e56.
//
// Solidity: function getChallengeList() view returns((address,address,address[])[])
func (_Pom *PomCaller) GetChallengeList(opts *bind.CallOpts) ([]POMPOMChallenge, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "getChallengeList")

	if err != nil {
		return *new([]POMPOMChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new([]POMPOMChallenge)).(*[]POMPOMChallenge)

	return out0, err

}

// GetChallengeList is a free data retrieval call binding the contract method 0xf8c57e56.
//
// Solidity: function getChallengeList() view returns((address,address,address[])[])
func (_Pom *PomSession) GetChallengeList() ([]POMPOMChallenge, error) {
	return _Pom.Contract.GetChallengeList(&_Pom.CallOpts)
}

// GetChallengeList is a free data retrieval call binding the contract method 0xf8c57e56.
//
// Solidity: function getChallengeList() view returns((address,address,address[])[])
func (_Pom *PomCallerSession) GetChallengeList() ([]POMPOMChallenge, error) {
	return _Pom.Contract.GetChallengeList(&_Pom.CallOpts)
}

// GetLastChallengeBlock is a free data retrieval call binding the contract method 0xaffea589.
//
// Solidity: function getLastChallengeBlock(address challenger) view returns(uint256)
func (_Pom *PomCaller) GetLastChallengeBlock(opts *bind.CallOpts, challenger common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "getLastChallengeBlock", challenger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastChallengeBlock is a free data retrieval call binding the contract method 0xaffea589.
//
// Solidity: function getLastChallengeBlock(address challenger) view returns(uint256)
func (_Pom *PomSession) GetLastChallengeBlock(challenger common.Address) (*big.Int, error) {
	return _Pom.Contract.GetLastChallengeBlock(&_Pom.CallOpts, challenger)
}

// GetLastChallengeBlock is a free data retrieval call binding the contract method 0xaffea589.
//
// Solidity: function getLastChallengeBlock(address challenger) view returns(uint256)
func (_Pom *PomCallerSession) GetLastChallengeBlock(challenger common.Address) (*big.Int, error) {
	return _Pom.Contract.GetLastChallengeBlock(&_Pom.CallOpts, challenger)
}

// GetMeshBoxList is a free data retrieval call binding the contract method 0x66ce5666.
//
// Solidity: function getMeshBoxList() view returns((string,address)[])
func (_Pom *PomCaller) GetMeshBoxList(opts *bind.CallOpts) ([]POMMeshBoxInfo, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "getMeshBoxList")
	if err != nil {
		return *new([]POMMeshBoxInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]POMMeshBoxInfo)).(*[]POMMeshBoxInfo)
	return out0, err

}

// GetMeshBoxList is a free data retrieval call binding the contract method 0x66ce5666.
//
// Solidity: function getMeshBoxList() view returns((string,address)[])
func (_Pom *PomSession) GetMeshBoxList() ([]POMMeshBoxInfo, error) {
	return _Pom.Contract.GetMeshBoxList(&_Pom.CallOpts)
}

// GetMeshBoxList is a free data retrieval call binding the contract method 0x66ce5666.
//
// Solidity: function getMeshBoxList() view returns((string,address)[])
func (_Pom *PomCallerSession) GetMeshBoxList() ([]POMMeshBoxInfo, error) {
	return _Pom.Contract.GetMeshBoxList(&_Pom.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pom *PomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pom.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pom *PomSession) Owner() (common.Address, error) {
	return _Pom.Contract.Owner(&_Pom.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pom *PomCallerSession) Owner() (common.Address, error) {
	return _Pom.Contract.Owner(&_Pom.CallOpts)
}

// AddMeshBox is a paid mutator transaction binding the contract method 0xa1244e6f.
//
// Solidity: function addMeshBox(string id, address addr) returns()
func (_Pom *PomTransactor) AddMeshBox(opts *bind.TransactOpts, id string, addr common.Address) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "addMeshBox", id, addr)
}

// AddMeshBox is a paid mutator transaction binding the contract method 0xa1244e6f.
//
// Solidity: function addMeshBox(string id, address addr) returns()
func (_Pom *PomSession) AddMeshBox(id string, addr common.Address) (*types.Transaction, error) {
	return _Pom.Contract.AddMeshBox(&_Pom.TransactOpts, id, addr)
}

// AddMeshBox is a paid mutator transaction binding the contract method 0xa1244e6f.
//
// Solidity: function addMeshBox(string id, address addr) returns()
func (_Pom *PomTransactorSession) AddMeshBox(id string, addr common.Address) (*types.Transaction, error) {
	return _Pom.Contract.AddMeshBox(&_Pom.TransactOpts, id, addr)
}

// GetTarget is a paid mutator transaction binding the contract method 0xf00e6a2a.
//
// Solidity: function getTarget() returns(address)
func (_Pom *PomTransactor) GetTarget(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "getTarget")
}

// GetTarget is a paid mutator transaction binding the contract method 0xf00e6a2a.
//
// Solidity: function getTarget() returns(address)
func (_Pom *PomSession) GetTarget() (*types.Transaction, error) {
	return _Pom.Contract.GetTarget(&_Pom.TransactOpts)
}

// GetTarget is a paid mutator transaction binding the contract method 0xf00e6a2a.
//
// Solidity: function getTarget() returns(address)
func (_Pom *PomTransactorSession) GetTarget() (*types.Transaction, error) {
	return _Pom.Contract.GetTarget(&_Pom.TransactOpts)
}

// RemoveMeshBox is a paid mutator transaction binding the contract method 0x5231432c.
//
// Solidity: function removeMeshBox(address addr) returns()
func (_Pom *PomTransactor) RemoveMeshBox(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "removeMeshBox", addr)
}

// RemoveMeshBox is a paid mutator transaction binding the contract method 0x5231432c.
//
// Solidity: function removeMeshBox(address addr) returns()
func (_Pom *PomSession) RemoveMeshBox(addr common.Address) (*types.Transaction, error) {
	return _Pom.Contract.RemoveMeshBox(&_Pom.TransactOpts, addr)
}

// RemoveMeshBox is a paid mutator transaction binding the contract method 0x5231432c.
//
// Solidity: function removeMeshBox(address addr) returns()
func (_Pom *PomTransactorSession) RemoveMeshBox(addr common.Address) (*types.Transaction, error) {
	return _Pom.Contract.RemoveMeshBox(&_Pom.TransactOpts, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pom *PomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pom *PomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pom.Contract.RenounceOwnership(&_Pom.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pom *PomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pom.Contract.RenounceOwnership(&_Pom.TransactOpts)
}

// SendPOMReceipt is a paid mutator transaction binding the contract method 0xee1a93e9.
//
// Solidity: function sendPOMReceipt(address target, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Pom *PomTransactor) SendPOMReceipt(opts *bind.TransactOpts, target common.Address, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "sendPOMReceipt", target, v, r, s)
}

// SendPOMReceipt is a paid mutator transaction binding the contract method 0xee1a93e9.
//
// Solidity: function sendPOMReceipt(address target, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Pom *PomSession) SendPOMReceipt(target common.Address, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Pom.Contract.SendPOMReceipt(&_Pom.TransactOpts, target, v, r, s)
}

// SendPOMReceipt is a paid mutator transaction binding the contract method 0xee1a93e9.
//
// Solidity: function sendPOMReceipt(address target, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Pom *PomTransactorSession) SendPOMReceipt(target common.Address, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Pom.Contract.SendPOMReceipt(&_Pom.TransactOpts, target, v, r, s)
}

// SendPomEpochReward is a paid mutator transaction binding the contract method 0xa18c7947.
//
// Solidity: function sendPomEpochReward(uint256 amount) returns()
func (_Pom *PomTransactor) SendPomEpochReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "sendPomEpochReward", amount)
}

// SendPomEpochReward is a paid mutator transaction binding the contract method 0xa18c7947.
//
// Solidity: function sendPomEpochReward(uint256 amount) returns()
func (_Pom *PomSession) SendPomEpochReward(amount *big.Int) (*types.Transaction, error) {
	return _Pom.Contract.SendPomEpochReward(&_Pom.TransactOpts, amount)
}

// SendPomEpochReward is a paid mutator transaction binding the contract method 0xa18c7947.
//
// Solidity: function sendPomEpochReward(uint256 amount) returns()
func (_Pom *PomTransactorSession) SendPomEpochReward(amount *big.Int) (*types.Transaction, error) {
	return _Pom.Contract.SendPomEpochReward(&_Pom.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pom *PomTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pom.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pom *PomSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pom.Contract.TransferOwnership(&_Pom.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pom *PomTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pom.Contract.TransferOwnership(&_Pom.TransactOpts, newOwner)
}

// PomAddMeshBoxIterator is returned from FilterAddMeshBox and is used to iterate over the raw logs and unpacked data for AddMeshBox events raised by the Pom contract.
type PomAddMeshBoxIterator struct {
	Event *PomAddMeshBox // Event containing the contract specifics and raw log

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
func (it *PomAddMeshBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PomAddMeshBox)
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
		it.Event = new(PomAddMeshBox)
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
func (it *PomAddMeshBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PomAddMeshBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PomAddMeshBox represents a AddMeshBox event raised by the Pom contract.
type PomAddMeshBox struct {
	Id   string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddMeshBox is a free log retrieval operation binding the contract event 0x13044b12afcaf31cb56bcec25471270d6a58524c08750c78748170bccb5776d3.
//
// Solidity: event AddMeshBox(string id, address addr)
func (_Pom *PomFilterer) FilterAddMeshBox(opts *bind.FilterOpts) (*PomAddMeshBoxIterator, error) {

	logs, sub, err := _Pom.contract.FilterLogs(opts, "AddMeshBox")
	if err != nil {
		return nil, err
	}
	return &PomAddMeshBoxIterator{contract: _Pom.contract, event: "AddMeshBox", logs: logs, sub: sub}, nil
}

// WatchAddMeshBox is a free log subscription operation binding the contract event 0x13044b12afcaf31cb56bcec25471270d6a58524c08750c78748170bccb5776d3.
//
// Solidity: event AddMeshBox(string id, address addr)
func (_Pom *PomFilterer) WatchAddMeshBox(opts *bind.WatchOpts, sink chan<- *PomAddMeshBox) (event.Subscription, error) {

	logs, sub, err := _Pom.contract.WatchLogs(opts, "AddMeshBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PomAddMeshBox)
				if err := _Pom.contract.UnpackLog(event, "AddMeshBox", log); err != nil {
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

// ParseAddMeshBox is a log parse operation binding the contract event 0x13044b12afcaf31cb56bcec25471270d6a58524c08750c78748170bccb5776d3.
//
// Solidity: event AddMeshBox(string id, address addr)
func (_Pom *PomFilterer) ParseAddMeshBox(log types.Log) (*PomAddMeshBox, error) {
	event := new(PomAddMeshBox)
	if err := _Pom.contract.UnpackLog(event, "AddMeshBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PomGetTargetIterator is returned from FilterGetTarget and is used to iterate over the raw logs and unpacked data for GetTarget events raised by the Pom contract.
type PomGetTargetIterator struct {
	Event *PomGetTarget // Event containing the contract specifics and raw log

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
func (it *PomGetTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PomGetTarget)
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
		it.Event = new(PomGetTarget)
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
func (it *PomGetTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PomGetTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PomGetTarget represents a GetTarget event raised by the Pom contract.
type PomGetTarget struct {
	Challenger common.Address
	Target     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGetTarget is a free log retrieval operation binding the contract event 0x53b0b52a85961621cf1f25fa9513d34315a2815b92e27c5d7c5603bd8ecb3385.
//
// Solidity: event GetTarget(address indexed challenger, address target)
func (_Pom *PomFilterer) FilterGetTarget(opts *bind.FilterOpts, challenger []common.Address) (*PomGetTargetIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Pom.contract.FilterLogs(opts, "GetTarget", challengerRule)
	if err != nil {
		return nil, err
	}
	return &PomGetTargetIterator{contract: _Pom.contract, event: "GetTarget", logs: logs, sub: sub}, nil
}

// WatchGetTarget is a free log subscription operation binding the contract event 0x53b0b52a85961621cf1f25fa9513d34315a2815b92e27c5d7c5603bd8ecb3385.
//
// Solidity: event GetTarget(address indexed challenger, address target)
func (_Pom *PomFilterer) WatchGetTarget(opts *bind.WatchOpts, sink chan<- *PomGetTarget, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Pom.contract.WatchLogs(opts, "GetTarget", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PomGetTarget)
				if err := _Pom.contract.UnpackLog(event, "GetTarget", log); err != nil {
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

// ParseGetTarget is a log parse operation binding the contract event 0x53b0b52a85961621cf1f25fa9513d34315a2815b92e27c5d7c5603bd8ecb3385.
//
// Solidity: event GetTarget(address indexed challenger, address target)
func (_Pom *PomFilterer) ParseGetTarget(log types.Log) (*PomGetTarget, error) {
	event := new(PomGetTarget)
	if err := _Pom.contract.UnpackLog(event, "GetTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pom contract.
type PomOwnershipTransferredIterator struct {
	Event *PomOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PomOwnershipTransferred)
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
		it.Event = new(PomOwnershipTransferred)
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
func (it *PomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PomOwnershipTransferred represents a OwnershipTransferred event raised by the Pom contract.
type PomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pom *PomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pom.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PomOwnershipTransferredIterator{contract: _Pom.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pom *PomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pom.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PomOwnershipTransferred)
				if err := _Pom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pom *PomFilterer) ParseOwnershipTransferred(log types.Log) (*PomOwnershipTransferred, error) {
	event := new(PomOwnershipTransferred)
	if err := _Pom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PomRemoveMeshBoxIterator is returned from FilterRemoveMeshBox and is used to iterate over the raw logs and unpacked data for RemoveMeshBox events raised by the Pom contract.
type PomRemoveMeshBoxIterator struct {
	Event *PomRemoveMeshBox // Event containing the contract specifics and raw log

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
func (it *PomRemoveMeshBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PomRemoveMeshBox)
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
		it.Event = new(PomRemoveMeshBox)
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
func (it *PomRemoveMeshBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PomRemoveMeshBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PomRemoveMeshBox represents a RemoveMeshBox event raised by the Pom contract.
type PomRemoveMeshBox struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemoveMeshBox is a free log retrieval operation binding the contract event 0x4fffe0b741493f7500fdc97a40de9f74946cf63385c6689f84581f6bfed7022a.
//
// Solidity: event RemoveMeshBox(address addr)
func (_Pom *PomFilterer) FilterRemoveMeshBox(opts *bind.FilterOpts) (*PomRemoveMeshBoxIterator, error) {

	logs, sub, err := _Pom.contract.FilterLogs(opts, "RemoveMeshBox")
	if err != nil {
		return nil, err
	}
	return &PomRemoveMeshBoxIterator{contract: _Pom.contract, event: "RemoveMeshBox", logs: logs, sub: sub}, nil
}

// WatchRemoveMeshBox is a free log subscription operation binding the contract event 0x4fffe0b741493f7500fdc97a40de9f74946cf63385c6689f84581f6bfed7022a.
//
// Solidity: event RemoveMeshBox(address addr)
func (_Pom *PomFilterer) WatchRemoveMeshBox(opts *bind.WatchOpts, sink chan<- *PomRemoveMeshBox) (event.Subscription, error) {

	logs, sub, err := _Pom.contract.WatchLogs(opts, "RemoveMeshBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PomRemoveMeshBox)
				if err := _Pom.contract.UnpackLog(event, "RemoveMeshBox", log); err != nil {
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

// ParseRemoveMeshBox is a log parse operation binding the contract event 0x4fffe0b741493f7500fdc97a40de9f74946cf63385c6689f84581f6bfed7022a.
//
// Solidity: event RemoveMeshBox(address addr)
func (_Pom *PomFilterer) ParseRemoveMeshBox(log types.Log) (*PomRemoveMeshBox, error) {
	event := new(PomRemoveMeshBox)
	if err := _Pom.contract.UnpackLog(event, "RemoveMeshBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
