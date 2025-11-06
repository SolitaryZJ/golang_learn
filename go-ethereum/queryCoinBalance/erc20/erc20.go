// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// QueryCoinBalanceMetaData contains all meta data concerning the QueryCoinBalance contract.
var QueryCoinBalanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QueryCoinBalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use QueryCoinBalanceMetaData.ABI instead.
var QueryCoinBalanceABI = QueryCoinBalanceMetaData.ABI

// QueryCoinBalance is an auto generated Go binding around an Ethereum contract.
type QueryCoinBalance struct {
	QueryCoinBalanceCaller     // Read-only binding to the contract
	QueryCoinBalanceTransactor // Write-only binding to the contract
	QueryCoinBalanceFilterer   // Log filterer for contract events
}

// QueryCoinBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueryCoinBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryCoinBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueryCoinBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryCoinBalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueryCoinBalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryCoinBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueryCoinBalanceSession struct {
	Contract     *QueryCoinBalance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueryCoinBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueryCoinBalanceCallerSession struct {
	Contract *QueryCoinBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// QueryCoinBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueryCoinBalanceTransactorSession struct {
	Contract     *QueryCoinBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QueryCoinBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueryCoinBalanceRaw struct {
	Contract *QueryCoinBalance // Generic contract binding to access the raw methods on
}

// QueryCoinBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueryCoinBalanceCallerRaw struct {
	Contract *QueryCoinBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// QueryCoinBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueryCoinBalanceTransactorRaw struct {
	Contract *QueryCoinBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueryCoinBalance creates a new instance of QueryCoinBalance, bound to a specific deployed contract.
func NewQueryCoinBalance(address common.Address, backend bind.ContractBackend) (*QueryCoinBalance, error) {
	contract, err := bindQueryCoinBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalance{QueryCoinBalanceCaller: QueryCoinBalanceCaller{contract: contract}, QueryCoinBalanceTransactor: QueryCoinBalanceTransactor{contract: contract}, QueryCoinBalanceFilterer: QueryCoinBalanceFilterer{contract: contract}}, nil
}

// NewQueryCoinBalanceCaller creates a new read-only instance of QueryCoinBalance, bound to a specific deployed contract.
func NewQueryCoinBalanceCaller(address common.Address, caller bind.ContractCaller) (*QueryCoinBalanceCaller, error) {
	contract, err := bindQueryCoinBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalanceCaller{contract: contract}, nil
}

// NewQueryCoinBalanceTransactor creates a new write-only instance of QueryCoinBalance, bound to a specific deployed contract.
func NewQueryCoinBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*QueryCoinBalanceTransactor, error) {
	contract, err := bindQueryCoinBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalanceTransactor{contract: contract}, nil
}

// NewQueryCoinBalanceFilterer creates a new log filterer instance of QueryCoinBalance, bound to a specific deployed contract.
func NewQueryCoinBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*QueryCoinBalanceFilterer, error) {
	contract, err := bindQueryCoinBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalanceFilterer{contract: contract}, nil
}

// bindQueryCoinBalance binds a generic wrapper to an already deployed contract.
func bindQueryCoinBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QueryCoinBalanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueryCoinBalance *QueryCoinBalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueryCoinBalance.Contract.QueryCoinBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueryCoinBalance *QueryCoinBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.QueryCoinBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueryCoinBalance *QueryCoinBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.QueryCoinBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueryCoinBalance *QueryCoinBalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueryCoinBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueryCoinBalance *QueryCoinBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueryCoinBalance *QueryCoinBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _QueryCoinBalance.Contract.Allowance(&_QueryCoinBalance.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _QueryCoinBalance.Contract.Allowance(&_QueryCoinBalance.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _QueryCoinBalance.Contract.BalanceOf(&_QueryCoinBalance.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _QueryCoinBalance.Contract.BalanceOf(&_QueryCoinBalance.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QueryCoinBalance *QueryCoinBalanceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QueryCoinBalance *QueryCoinBalanceSession) Decimals() (uint8, error) {
	return _QueryCoinBalance.Contract.Decimals(&_QueryCoinBalance.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) Decimals() (uint8, error) {
	return _QueryCoinBalance.Contract.Decimals(&_QueryCoinBalance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceSession) Name() (string, error) {
	return _QueryCoinBalance.Contract.Name(&_QueryCoinBalance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) Name() (string, error) {
	return _QueryCoinBalance.Contract.Name(&_QueryCoinBalance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceSession) Symbol() (string, error) {
	return _QueryCoinBalance.Contract.Symbol(&_QueryCoinBalance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) Symbol() (string, error) {
	return _QueryCoinBalance.Contract.Symbol(&_QueryCoinBalance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QueryCoinBalance.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceSession) TotalSupply() (*big.Int, error) {
	return _QueryCoinBalance.Contract.TotalSupply(&_QueryCoinBalance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QueryCoinBalance *QueryCoinBalanceCallerSession) TotalSupply() (*big.Int, error) {
	return _QueryCoinBalance.Contract.TotalSupply(&_QueryCoinBalance.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.Approve(&_QueryCoinBalance.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.Approve(&_QueryCoinBalance.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.Transfer(&_QueryCoinBalance.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.Transfer(&_QueryCoinBalance.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.TransferFrom(&_QueryCoinBalance.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_QueryCoinBalance *QueryCoinBalanceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _QueryCoinBalance.Contract.TransferFrom(&_QueryCoinBalance.TransactOpts, from, to, value)
}

// QueryCoinBalanceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the QueryCoinBalance contract.
type QueryCoinBalanceApprovalIterator struct {
	Event *QueryCoinBalanceApproval // Event containing the contract specifics and raw log

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
func (it *QueryCoinBalanceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QueryCoinBalanceApproval)
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
		it.Event = new(QueryCoinBalanceApproval)
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
func (it *QueryCoinBalanceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QueryCoinBalanceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QueryCoinBalanceApproval represents a Approval event raised by the QueryCoinBalance contract.
type QueryCoinBalanceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*QueryCoinBalanceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _QueryCoinBalance.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalanceApprovalIterator{contract: _QueryCoinBalance.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QueryCoinBalanceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _QueryCoinBalance.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QueryCoinBalanceApproval)
				if err := _QueryCoinBalance.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) ParseApproval(log types.Log) (*QueryCoinBalanceApproval, error) {
	event := new(QueryCoinBalanceApproval)
	if err := _QueryCoinBalance.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QueryCoinBalanceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the QueryCoinBalance contract.
type QueryCoinBalanceTransferIterator struct {
	Event *QueryCoinBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *QueryCoinBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QueryCoinBalanceTransfer)
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
		it.Event = new(QueryCoinBalanceTransfer)
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
func (it *QueryCoinBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QueryCoinBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QueryCoinBalanceTransfer represents a Transfer event raised by the QueryCoinBalance contract.
type QueryCoinBalanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*QueryCoinBalanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QueryCoinBalance.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QueryCoinBalanceTransferIterator{contract: _QueryCoinBalance.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QueryCoinBalanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QueryCoinBalance.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QueryCoinBalanceTransfer)
				if err := _QueryCoinBalance.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_QueryCoinBalance *QueryCoinBalanceFilterer) ParseTransfer(log types.Log) (*QueryCoinBalanceTransfer, error) {
	event := new(QueryCoinBalanceTransfer)
	if err := _QueryCoinBalance.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
