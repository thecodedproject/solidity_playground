// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tests

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BoxABI is the input ABI used to generate the binding from.
const BoxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BoxFuncSigs maps the 4-byte function signature to its string representation.
var BoxFuncSigs = map[string]string{
	"2e64cec1": "retrieve()",
	"6057361d": "store(uint256)",
}

// BoxBin is the compiled bytecode used for deploying new contracts.
var BoxBin = "0x608060405234801561001057600080fd5b5060e58061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80632e64cec11460375780636057361d14604c575b600080fd5b60005460405190815260200160405180910390f35b605b60573660046098565b605d565b005b60008190556040518181527f93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c599060200160405180910390a150565b60006020828403121560a8578081fd5b503591905056fea264697066735822122012a3f004c65affe6c5e13ab4131e4104474a612a6dfaf5ae277d4ef5bded32b864736f6c63430008040033"

// DeployBox deploys a new Ethereum contract, binding an instance of Box to it.
func DeployBox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Box, error) {
	parsed, err := abi.JSON(strings.NewReader(BoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BoxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Box{BoxCaller: BoxCaller{contract: contract}, BoxTransactor: BoxTransactor{contract: contract}, BoxFilterer: BoxFilterer{contract: contract}}, nil
}

// Box is an auto generated Go binding around an Ethereum contract.
type Box struct {
	BoxCaller     // Read-only binding to the contract
	BoxTransactor // Write-only binding to the contract
	BoxFilterer   // Log filterer for contract events
}

// BoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoxSession struct {
	Contract     *Box              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoxCallerSession struct {
	Contract *BoxCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoxTransactorSession struct {
	Contract     *BoxTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoxRaw struct {
	Contract *Box // Generic contract binding to access the raw methods on
}

// BoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoxCallerRaw struct {
	Contract *BoxCaller // Generic read-only contract binding to access the raw methods on
}

// BoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoxTransactorRaw struct {
	Contract *BoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBox creates a new instance of Box, bound to a specific deployed contract.
func NewBox(address common.Address, backend bind.ContractBackend) (*Box, error) {
	contract, err := bindBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Box{BoxCaller: BoxCaller{contract: contract}, BoxTransactor: BoxTransactor{contract: contract}, BoxFilterer: BoxFilterer{contract: contract}}, nil
}

// NewBoxCaller creates a new read-only instance of Box, bound to a specific deployed contract.
func NewBoxCaller(address common.Address, caller bind.ContractCaller) (*BoxCaller, error) {
	contract, err := bindBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoxCaller{contract: contract}, nil
}

// NewBoxTransactor creates a new write-only instance of Box, bound to a specific deployed contract.
func NewBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*BoxTransactor, error) {
	contract, err := bindBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoxTransactor{contract: contract}, nil
}

// NewBoxFilterer creates a new log filterer instance of Box, bound to a specific deployed contract.
func NewBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*BoxFilterer, error) {
	contract, err := bindBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoxFilterer{contract: contract}, nil
}

// bindBox binds a generic wrapper to an already deployed contract.
func bindBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Box *BoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Box.Contract.BoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Box *BoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Box.Contract.BoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Box *BoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Box.Contract.BoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Box *BoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Box.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Box *BoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Box.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Box *BoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Box.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Box *BoxCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Box.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Box *BoxSession) Retrieve() (*big.Int, error) {
	return _Box.Contract.Retrieve(&_Box.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Box *BoxCallerSession) Retrieve() (*big.Int, error) {
	return _Box.Contract.Retrieve(&_Box.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 newValue) returns()
func (_Box *BoxTransactor) Store(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Box.contract.Transact(opts, "store", newValue)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 newValue) returns()
func (_Box *BoxSession) Store(newValue *big.Int) (*types.Transaction, error) {
	return _Box.Contract.Store(&_Box.TransactOpts, newValue)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 newValue) returns()
func (_Box *BoxTransactorSession) Store(newValue *big.Int) (*types.Transaction, error) {
	return _Box.Contract.Store(&_Box.TransactOpts, newValue)
}

// BoxValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the Box contract.
type BoxValueChangedIterator struct {
	Event *BoxValueChanged // Event containing the contract specifics and raw log

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
func (it *BoxValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoxValueChanged)
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
		it.Event = new(BoxValueChanged)
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
func (it *BoxValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoxValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoxValueChanged represents a ValueChanged event raised by the Box contract.
type BoxValueChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_Box *BoxFilterer) FilterValueChanged(opts *bind.FilterOpts) (*BoxValueChangedIterator, error) {

	logs, sub, err := _Box.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &BoxValueChangedIterator{contract: _Box.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_Box *BoxFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *BoxValueChanged) (event.Subscription, error) {

	logs, sub, err := _Box.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoxValueChanged)
				if err := _Box.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 newValue)
func (_Box *BoxFilterer) ParseValueChanged(log types.Log) (*BoxValueChanged, error) {
	event := new(BoxValueChanged)
	if err := _Box.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
