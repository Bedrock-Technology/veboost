// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

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
)

// AirdropDist is an auto generated low-level Go binding around an user-defined struct.
type AirdropDist struct {
	Root        [32]byte
	ActivatedAt uint32
	Duration    uint32
	Disabled    bool
}

// AirdropMetaData contains all meta data concerning the Airdrop contract.
var AirdropMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"brToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoot\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAirdrop.Dist\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasClaimed\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_users\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_votingEscrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_brToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAirdrop\",\"inputs\":[{\"name\":\"_disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"_newRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDuration\",\"inputs\":[{\"name\":\"_duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateRoot\",\"inputs\":[{\"name\":\"_newRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votingEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AirdropClaimed\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"veNFTId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionDisabledSet\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"preStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootSubmit\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardsValidTime\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootUpdate\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"preRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidDurationUpdate\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"preValidDuration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"validDuration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// AirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use AirdropMetaData.ABI instead.
var AirdropABI = AirdropMetaData.ABI

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Airdrop.Contract.DEFAULTADMINROLE(&_Airdrop.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Airdrop.Contract.DEFAULTADMINROLE(&_Airdrop.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) OPERATORROLE() ([32]byte, error) {
	return _Airdrop.Contract.OPERATORROLE(&_Airdrop.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Airdrop.Contract.OPERATORROLE(&_Airdrop.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) PAUSERROLE() ([32]byte, error) {
	return _Airdrop.Contract.PAUSERROLE(&_Airdrop.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Airdrop.Contract.PAUSERROLE(&_Airdrop.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_Airdrop *AirdropCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_Airdrop *AirdropSession) ActivationDelay() (uint32, error) {
	return _Airdrop.Contract.ActivationDelay(&_Airdrop.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_Airdrop *AirdropCallerSession) ActivationDelay() (uint32, error) {
	return _Airdrop.Contract.ActivationDelay(&_Airdrop.CallOpts)
}

// BrToken is a free data retrieval call binding the contract method 0x3c0d2f2b.
//
// Solidity: function brToken() view returns(address)
func (_Airdrop *AirdropCaller) BrToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "brToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrToken is a free data retrieval call binding the contract method 0x3c0d2f2b.
//
// Solidity: function brToken() view returns(address)
func (_Airdrop *AirdropSession) BrToken() (common.Address, error) {
	return _Airdrop.Contract.BrToken(&_Airdrop.CallOpts)
}

// BrToken is a free data retrieval call binding the contract method 0x3c0d2f2b.
//
// Solidity: function brToken() view returns(address)
func (_Airdrop *AirdropCallerSession) BrToken() (common.Address, error) {
	return _Airdrop.Contract.BrToken(&_Airdrop.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropSession) CurrentEpoch() (*big.Int, error) {
	return _Airdrop.Contract.CurrentEpoch(&_Airdrop.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Airdrop.Contract.CurrentEpoch(&_Airdrop.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Airdrop.Contract.GetRoleAdmin(&_Airdrop.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Airdrop.Contract.GetRoleAdmin(&_Airdrop.CallOpts, role)
}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint32,uint32,bool))
func (_Airdrop *AirdropCaller) GetRoot(opts *bind.CallOpts, _epoch *big.Int) (AirdropDist, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getRoot", _epoch)

	if err != nil {
		return *new(AirdropDist), err
	}

	out0 := *abi.ConvertType(out[0], new(AirdropDist)).(*AirdropDist)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint32,uint32,bool))
func (_Airdrop *AirdropSession) GetRoot(_epoch *big.Int) (AirdropDist, error) {
	return _Airdrop.Contract.GetRoot(&_Airdrop.CallOpts, _epoch)
}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint32,uint32,bool))
func (_Airdrop *AirdropCallerSession) GetRoot(_epoch *big.Int) (AirdropDist, error) {
	return _Airdrop.Contract.GetRoot(&_Airdrop.CallOpts, _epoch)
}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropCaller) HasClaimed(opts *bind.CallOpts, _epoch *big.Int, _users []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "hasClaimed", _epoch, _users)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropSession) HasClaimed(_epoch *big.Int, _users []common.Address) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed(&_Airdrop.CallOpts, _epoch, _users)
}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropCallerSession) HasClaimed(_epoch *big.Int, _users []common.Address) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed(&_Airdrop.CallOpts, _epoch, _users)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Airdrop.Contract.HasRole(&_Airdrop.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Airdrop.Contract.HasRole(&_Airdrop.CallOpts, role, account)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Airdrop *AirdropCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Airdrop *AirdropSession) IsActive() (bool, error) {
	return _Airdrop.Contract.IsActive(&_Airdrop.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Airdrop *AirdropCallerSession) IsActive() (bool, error) {
	return _Airdrop.Contract.IsActive(&_Airdrop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCallerSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_Airdrop *AirdropCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "votingEscrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_Airdrop *AirdropSession) VotingEscrow() (common.Address, error) {
	return _Airdrop.Contract.VotingEscrow(&_Airdrop.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0x4f2bfe5b.
//
// Solidity: function votingEscrow() view returns(address)
func (_Airdrop *AirdropCallerSession) VotingEscrow() (common.Address, error) {
	return _Airdrop.Contract.VotingEscrow(&_Airdrop.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropTransactor) Claim(opts *bind.TransactOpts, _amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "claim", _amount, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropSession) Claim(_amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim(&_Airdrop.TransactOpts, _amount, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropTransactorSession) Claim(_amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim(&_Airdrop.TransactOpts, _amount, _proof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.GrantRole(&_Airdrop.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.GrantRole(&_Airdrop.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfb6bf74c.
//
// Solidity: function initialize(uint32 _activationDelay, address _votingEscrow, address _brToken, address _admin) returns()
func (_Airdrop *AirdropTransactor) Initialize(opts *bind.TransactOpts, _activationDelay uint32, _votingEscrow common.Address, _brToken common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "initialize", _activationDelay, _votingEscrow, _brToken, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xfb6bf74c.
//
// Solidity: function initialize(uint32 _activationDelay, address _votingEscrow, address _brToken, address _admin) returns()
func (_Airdrop *AirdropSession) Initialize(_activationDelay uint32, _votingEscrow common.Address, _brToken common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Initialize(&_Airdrop.TransactOpts, _activationDelay, _votingEscrow, _brToken, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xfb6bf74c.
//
// Solidity: function initialize(uint32 _activationDelay, address _votingEscrow, address _brToken, address _admin) returns()
func (_Airdrop *AirdropTransactorSession) Initialize(_activationDelay uint32, _votingEscrow common.Address, _brToken common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Initialize(&_Airdrop.TransactOpts, _activationDelay, _votingEscrow, _brToken, _admin)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactorSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceRole(&_Airdrop.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceRole(&_Airdrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RevokeRole(&_Airdrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RevokeRole(&_Airdrop.TransactOpts, role, account)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x11111487.
//
// Solidity: function setAirdrop(bool _disabled) returns()
func (_Airdrop *AirdropTransactor) SetAirdrop(opts *bind.TransactOpts, _disabled bool) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setAirdrop", _disabled)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x11111487.
//
// Solidity: function setAirdrop(bool _disabled) returns()
func (_Airdrop *AirdropSession) SetAirdrop(_disabled bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAirdrop(&_Airdrop.TransactOpts, _disabled)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x11111487.
//
// Solidity: function setAirdrop(bool _disabled) returns()
func (_Airdrop *AirdropTransactorSession) SetAirdrop(_disabled bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAirdrop(&_Airdrop.TransactOpts, _disabled)
}

// SetDelay is a paid mutator transaction binding the contract method 0x10348665.
//
// Solidity: function setDelay(uint32 _activationDelay) returns()
func (_Airdrop *AirdropTransactor) SetDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setDelay", _activationDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x10348665.
//
// Solidity: function setDelay(uint32 _activationDelay) returns()
func (_Airdrop *AirdropSession) SetDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.SetDelay(&_Airdrop.TransactOpts, _activationDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x10348665.
//
// Solidity: function setDelay(uint32 _activationDelay) returns()
func (_Airdrop *AirdropTransactorSession) SetDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.SetDelay(&_Airdrop.TransactOpts, _activationDelay)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint32 _duration) returns()
func (_Airdrop *AirdropTransactor) SubmitRoot(opts *bind.TransactOpts, _newRoot [32]byte, _duration uint32) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "submitRoot", _newRoot, _duration)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint32 _duration) returns()
func (_Airdrop *AirdropSession) SubmitRoot(_newRoot [32]byte, _duration uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.SubmitRoot(&_Airdrop.TransactOpts, _newRoot, _duration)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint32 _duration) returns()
func (_Airdrop *AirdropTransactorSession) SubmitRoot(_newRoot [32]byte, _duration uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.SubmitRoot(&_Airdrop.TransactOpts, _newRoot, _duration)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactorSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xb68fb682.
//
// Solidity: function updateDuration(uint32 _duration) returns()
func (_Airdrop *AirdropTransactor) UpdateDuration(opts *bind.TransactOpts, _duration uint32) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateDuration", _duration)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xb68fb682.
//
// Solidity: function updateDuration(uint32 _duration) returns()
func (_Airdrop *AirdropSession) UpdateDuration(_duration uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateDuration(&_Airdrop.TransactOpts, _duration)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xb68fb682.
//
// Solidity: function updateDuration(uint32 _duration) returns()
func (_Airdrop *AirdropTransactorSession) UpdateDuration(_duration uint32) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateDuration(&_Airdrop.TransactOpts, _duration)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x21ff9970.
//
// Solidity: function updateRoot(bytes32 _newRoot) returns()
func (_Airdrop *AirdropTransactor) UpdateRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateRoot", _newRoot)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x21ff9970.
//
// Solidity: function updateRoot(bytes32 _newRoot) returns()
func (_Airdrop *AirdropSession) UpdateRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateRoot(&_Airdrop.TransactOpts, _newRoot)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x21ff9970.
//
// Solidity: function updateRoot(bytes32 _newRoot) returns()
func (_Airdrop *AirdropTransactorSession) UpdateRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateRoot(&_Airdrop.TransactOpts, _newRoot)
}

// AirdropActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the Airdrop contract.
type AirdropActivationDelaySetIterator struct {
	Event *AirdropActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *AirdropActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropActivationDelaySet)
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
		it.Event = new(AirdropActivationDelaySet)
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
func (it *AirdropActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropActivationDelaySet represents a ActivationDelaySet event raised by the Airdrop contract.
type AirdropActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_Airdrop *AirdropFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*AirdropActivationDelaySetIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AirdropActivationDelaySetIterator{contract: _Airdrop.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_Airdrop *AirdropFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *AirdropActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropActivationDelaySet)
				if err := _Airdrop.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_Airdrop *AirdropFilterer) ParseActivationDelaySet(log types.Log) (*AirdropActivationDelaySet, error) {
	event := new(AirdropActivationDelaySet)
	if err := _Airdrop.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropAirdropClaimedIterator is returned from FilterAirdropClaimed and is used to iterate over the raw logs and unpacked data for AirdropClaimed events raised by the Airdrop contract.
type AirdropAirdropClaimedIterator struct {
	Event *AirdropAirdropClaimed // Event containing the contract specifics and raw log

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
func (it *AirdropAirdropClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirdropClaimed)
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
		it.Event = new(AirdropAirdropClaimed)
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
func (it *AirdropAirdropClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirdropClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirdropClaimed represents a AirdropClaimed event raised by the Airdrop contract.
type AirdropAirdropClaimed struct {
	Epoch   *big.Int
	User    common.Address
	Amount  *big.Int
	VeNFTId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAirdropClaimed is a free log retrieval operation binding the contract event 0x2db7b200aa423ebe7828d265f5dfeab33f3d1d049f4097ac814b05571ac97eac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, uint256 amount, uint256 veNFTId)
func (_Airdrop *AirdropFilterer) FilterAirdropClaimed(opts *bind.FilterOpts, epoch []*big.Int, user []common.Address) (*AirdropAirdropClaimedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AirdropClaimed", epochRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAirdropClaimedIterator{contract: _Airdrop.contract, event: "AirdropClaimed", logs: logs, sub: sub}, nil
}

// WatchAirdropClaimed is a free log subscription operation binding the contract event 0x2db7b200aa423ebe7828d265f5dfeab33f3d1d049f4097ac814b05571ac97eac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, uint256 amount, uint256 veNFTId)
func (_Airdrop *AirdropFilterer) WatchAirdropClaimed(opts *bind.WatchOpts, sink chan<- *AirdropAirdropClaimed, epoch []*big.Int, user []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AirdropClaimed", epochRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirdropClaimed)
				if err := _Airdrop.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
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

// ParseAirdropClaimed is a log parse operation binding the contract event 0x2db7b200aa423ebe7828d265f5dfeab33f3d1d049f4097ac814b05571ac97eac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, uint256 amount, uint256 veNFTId)
func (_Airdrop *AirdropFilterer) ParseAirdropClaimed(log types.Log) (*AirdropAirdropClaimed, error) {
	event := new(AirdropAirdropClaimed)
	if err := _Airdrop.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropDistributionDisabledSetIterator is returned from FilterDistributionDisabledSet and is used to iterate over the raw logs and unpacked data for DistributionDisabledSet events raised by the Airdrop contract.
type AirdropDistributionDisabledSetIterator struct {
	Event *AirdropDistributionDisabledSet // Event containing the contract specifics and raw log

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
func (it *AirdropDistributionDisabledSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropDistributionDisabledSet)
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
		it.Event = new(AirdropDistributionDisabledSet)
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
func (it *AirdropDistributionDisabledSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropDistributionDisabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropDistributionDisabledSet represents a DistributionDisabledSet event raised by the Airdrop contract.
type AirdropDistributionDisabledSet struct {
	Epoch     *big.Int
	PreStatus bool
	Status    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionDisabledSet is a free log retrieval operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) FilterDistributionDisabledSet(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropDistributionDisabledSetIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "DistributionDisabledSet", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropDistributionDisabledSetIterator{contract: _Airdrop.contract, event: "DistributionDisabledSet", logs: logs, sub: sub}, nil
}

// WatchDistributionDisabledSet is a free log subscription operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) WatchDistributionDisabledSet(opts *bind.WatchOpts, sink chan<- *AirdropDistributionDisabledSet, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "DistributionDisabledSet", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropDistributionDisabledSet)
				if err := _Airdrop.contract.UnpackLog(event, "DistributionDisabledSet", log); err != nil {
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

// ParseDistributionDisabledSet is a log parse operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) ParseDistributionDisabledSet(log types.Log) (*AirdropDistributionDisabledSet, error) {
	event := new(AirdropDistributionDisabledSet)
	if err := _Airdrop.contract.UnpackLog(event, "DistributionDisabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Airdrop contract.
type AirdropInitializedIterator struct {
	Event *AirdropInitialized // Event containing the contract specifics and raw log

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
func (it *AirdropInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropInitialized)
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
		it.Event = new(AirdropInitialized)
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
func (it *AirdropInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropInitialized represents a Initialized event raised by the Airdrop contract.
type AirdropInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Airdrop *AirdropFilterer) FilterInitialized(opts *bind.FilterOpts) (*AirdropInitializedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AirdropInitializedIterator{contract: _Airdrop.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Airdrop *AirdropFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AirdropInitialized) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropInitialized)
				if err := _Airdrop.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Airdrop *AirdropFilterer) ParseInitialized(log types.Log) (*AirdropInitialized, error) {
	event := new(AirdropInitialized)
	if err := _Airdrop.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropMerkleRootSubmitIterator is returned from FilterMerkleRootSubmit and is used to iterate over the raw logs and unpacked data for MerkleRootSubmit events raised by the Airdrop contract.
type AirdropMerkleRootSubmitIterator struct {
	Event *AirdropMerkleRootSubmit // Event containing the contract specifics and raw log

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
func (it *AirdropMerkleRootSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropMerkleRootSubmit)
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
		it.Event = new(AirdropMerkleRootSubmit)
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
func (it *AirdropMerkleRootSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropMerkleRootSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropMerkleRootSubmit represents a MerkleRootSubmit event raised by the Airdrop contract.
type AirdropMerkleRootSubmit struct {
	Epoch            *big.Int
	Root             [32]byte
	RewardsValidTime uint32
	ActivatedAt      uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootSubmit is a free log retrieval operation binding the contract event 0xd73de93535b69fff8d9446ce761a44eaa99231144738880817d8bd22131460b5.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint32 rewardsValidTime, uint32 activatedAt)
func (_Airdrop *AirdropFilterer) FilterMerkleRootSubmit(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropMerkleRootSubmitIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "MerkleRootSubmit", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropMerkleRootSubmitIterator{contract: _Airdrop.contract, event: "MerkleRootSubmit", logs: logs, sub: sub}, nil
}

// WatchMerkleRootSubmit is a free log subscription operation binding the contract event 0xd73de93535b69fff8d9446ce761a44eaa99231144738880817d8bd22131460b5.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint32 rewardsValidTime, uint32 activatedAt)
func (_Airdrop *AirdropFilterer) WatchMerkleRootSubmit(opts *bind.WatchOpts, sink chan<- *AirdropMerkleRootSubmit, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "MerkleRootSubmit", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropMerkleRootSubmit)
				if err := _Airdrop.contract.UnpackLog(event, "MerkleRootSubmit", log); err != nil {
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

// ParseMerkleRootSubmit is a log parse operation binding the contract event 0xd73de93535b69fff8d9446ce761a44eaa99231144738880817d8bd22131460b5.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint32 rewardsValidTime, uint32 activatedAt)
func (_Airdrop *AirdropFilterer) ParseMerkleRootSubmit(log types.Log) (*AirdropMerkleRootSubmit, error) {
	event := new(AirdropMerkleRootSubmit)
	if err := _Airdrop.contract.UnpackLog(event, "MerkleRootSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropMerkleRootUpdateIterator is returned from FilterMerkleRootUpdate and is used to iterate over the raw logs and unpacked data for MerkleRootUpdate events raised by the Airdrop contract.
type AirdropMerkleRootUpdateIterator struct {
	Event *AirdropMerkleRootUpdate // Event containing the contract specifics and raw log

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
func (it *AirdropMerkleRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropMerkleRootUpdate)
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
		it.Event = new(AirdropMerkleRootUpdate)
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
func (it *AirdropMerkleRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropMerkleRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropMerkleRootUpdate represents a MerkleRootUpdate event raised by the Airdrop contract.
type AirdropMerkleRootUpdate struct {
	Epoch   *big.Int
	PreRoot [32]byte
	Root    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdate is a free log retrieval operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) FilterMerkleRootUpdate(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropMerkleRootUpdateIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "MerkleRootUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropMerkleRootUpdateIterator{contract: _Airdrop.contract, event: "MerkleRootUpdate", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdate is a free log subscription operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) WatchMerkleRootUpdate(opts *bind.WatchOpts, sink chan<- *AirdropMerkleRootUpdate, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "MerkleRootUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropMerkleRootUpdate)
				if err := _Airdrop.contract.UnpackLog(event, "MerkleRootUpdate", log); err != nil {
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

// ParseMerkleRootUpdate is a log parse operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) ParseMerkleRootUpdate(log types.Log) (*AirdropMerkleRootUpdate, error) {
	event := new(AirdropMerkleRootUpdate)
	if err := _Airdrop.contract.UnpackLog(event, "MerkleRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Airdrop contract.
type AirdropPausedIterator struct {
	Event *AirdropPaused // Event containing the contract specifics and raw log

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
func (it *AirdropPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropPaused)
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
		it.Event = new(AirdropPaused)
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
func (it *AirdropPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropPaused represents a Paused event raised by the Airdrop contract.
type AirdropPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) FilterPaused(opts *bind.FilterOpts) (*AirdropPausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AirdropPausedIterator{contract: _Airdrop.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AirdropPaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropPaused)
				if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParsePaused(log types.Log) (*AirdropPaused, error) {
	event := new(AirdropPaused)
	if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Airdrop contract.
type AirdropRoleAdminChangedIterator struct {
	Event *AirdropRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AirdropRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleAdminChanged)
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
		it.Event = new(AirdropRoleAdminChanged)
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
func (it *AirdropRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleAdminChanged represents a RoleAdminChanged event raised by the Airdrop contract.
type AirdropRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Airdrop *AirdropFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AirdropRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleAdminChangedIterator{contract: _Airdrop.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Airdrop *AirdropFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AirdropRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleAdminChanged)
				if err := _Airdrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Airdrop *AirdropFilterer) ParseRoleAdminChanged(log types.Log) (*AirdropRoleAdminChanged, error) {
	event := new(AirdropRoleAdminChanged)
	if err := _Airdrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Airdrop contract.
type AirdropRoleGrantedIterator struct {
	Event *AirdropRoleGranted // Event containing the contract specifics and raw log

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
func (it *AirdropRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleGranted)
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
		it.Event = new(AirdropRoleGranted)
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
func (it *AirdropRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleGranted represents a RoleGranted event raised by the Airdrop contract.
type AirdropRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AirdropRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleGrantedIterator{contract: _Airdrop.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AirdropRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleGranted)
				if err := _Airdrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) ParseRoleGranted(log types.Log) (*AirdropRoleGranted, error) {
	event := new(AirdropRoleGranted)
	if err := _Airdrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Airdrop contract.
type AirdropRoleRevokedIterator struct {
	Event *AirdropRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AirdropRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleRevoked)
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
		it.Event = new(AirdropRoleRevoked)
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
func (it *AirdropRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleRevoked represents a RoleRevoked event raised by the Airdrop contract.
type AirdropRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AirdropRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleRevokedIterator{contract: _Airdrop.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AirdropRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleRevoked)
				if err := _Airdrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) ParseRoleRevoked(log types.Log) (*AirdropRoleRevoked, error) {
	event := new(AirdropRoleRevoked)
	if err := _Airdrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Airdrop contract.
type AirdropUnpausedIterator struct {
	Event *AirdropUnpaused // Event containing the contract specifics and raw log

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
func (it *AirdropUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropUnpaused)
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
		it.Event = new(AirdropUnpaused)
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
func (it *AirdropUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropUnpaused represents a Unpaused event raised by the Airdrop contract.
type AirdropUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AirdropUnpausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AirdropUnpausedIterator{contract: _Airdrop.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AirdropUnpaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropUnpaused)
				if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseUnpaused(log types.Log) (*AirdropUnpaused, error) {
	event := new(AirdropUnpaused)
	if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropValidDurationUpdateIterator is returned from FilterValidDurationUpdate and is used to iterate over the raw logs and unpacked data for ValidDurationUpdate events raised by the Airdrop contract.
type AirdropValidDurationUpdateIterator struct {
	Event *AirdropValidDurationUpdate // Event containing the contract specifics and raw log

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
func (it *AirdropValidDurationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropValidDurationUpdate)
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
		it.Event = new(AirdropValidDurationUpdate)
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
func (it *AirdropValidDurationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropValidDurationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropValidDurationUpdate represents a ValidDurationUpdate event raised by the Airdrop contract.
type AirdropValidDurationUpdate struct {
	Epoch            *big.Int
	PreValidDuration uint32
	ValidDuration    uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidDurationUpdate is a free log retrieval operation binding the contract event 0x3ab4c684dd32df4f0e46296012a3ba54a9268793454a7bf5fec78a277683be31.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint32 preValidDuration, uint32 validDuration)
func (_Airdrop *AirdropFilterer) FilterValidDurationUpdate(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropValidDurationUpdateIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ValidDurationUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropValidDurationUpdateIterator{contract: _Airdrop.contract, event: "ValidDurationUpdate", logs: logs, sub: sub}, nil
}

// WatchValidDurationUpdate is a free log subscription operation binding the contract event 0x3ab4c684dd32df4f0e46296012a3ba54a9268793454a7bf5fec78a277683be31.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint32 preValidDuration, uint32 validDuration)
func (_Airdrop *AirdropFilterer) WatchValidDurationUpdate(opts *bind.WatchOpts, sink chan<- *AirdropValidDurationUpdate, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ValidDurationUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropValidDurationUpdate)
				if err := _Airdrop.contract.UnpackLog(event, "ValidDurationUpdate", log); err != nil {
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

// ParseValidDurationUpdate is a log parse operation binding the contract event 0x3ab4c684dd32df4f0e46296012a3ba54a9268793454a7bf5fec78a277683be31.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint32 preValidDuration, uint32 validDuration)
func (_Airdrop *AirdropFilterer) ParseValidDurationUpdate(log types.Log) (*AirdropValidDurationUpdate, error) {
	event := new(AirdropValidDurationUpdate)
	if err := _Airdrop.contract.UnpackLog(event, "ValidDurationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
