// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recarga

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

// RecargaMetaData contains all meta data concerning the Recarga contract.
var RecargaMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"energia\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valor\",\"type\":\"uint256\"}],\"name\":\"RecargaFinalizada\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contador\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"energiaConsumida\",\"type\":\"uint256\"}],\"name\":\"finalizarSessao\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iniciarSessao\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precoPorKWh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessoes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"energiaConsumida\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valorPago\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalizada\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600560005534801561001557600080fd5b50610842806100256000396000f3fe60806040526004361061004a5760003560e01c8063412739051461004f5780634cd57d631461006b57806353d27cc214610096578063a718ae96146100c1578063d5e628dc14610101575b600080fd5b61006960048036038101906100649190610481565b61012c565b005b34801561007757600080fd5b506100806102ed565b60405161008d91906104d0565b60405180910390f35b3480156100a257600080fd5b506100ab6103dd565b6040516100b891906104d0565b60405180910390f35b3480156100cd57600080fd5b506100e860048036038101906100e391906104eb565b6103e3565b6040516100f89493929190610574565b60405180910390f35b34801561010d57600080fd5b50610116610440565b60405161012391906104d0565b60405180910390f35b60006001600084815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101cc90610616565b60405180910390fd5b8060030160009054906101000a900460ff1615610227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021e90610682565b60405180910390fd5b600080548361023691906106d1565b90508034101561027b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102729061075f565b60405180910390fd5b82826001018190555034826002018190555060018260030160006101000a81548160ff0219169083151502179055507f3ea1b19a3836fa0ed6579a4c083be8246ddc2b32d3930bc7bd5dd540027a41d6843385346040516102df949392919061077f565b60405180910390a150505050565b60008060026000815480929190610303906107c4565b91905055905060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600015158152506001600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055509050508091505090565b60025481565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900460ff16905084565b60005481565b600080fd5b6000819050919050565b61045e8161044b565b811461046957600080fd5b50565b60008135905061047b81610455565b92915050565b6000806040838503121561049857610497610446565b5b60006104a68582860161046c565b92505060206104b78582860161046c565b9150509250929050565b6104ca8161044b565b82525050565b60006020820190506104e560008301846104c1565b92915050565b60006020828403121561050157610500610446565b5b600061050f8482850161046c565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061054382610518565b9050919050565b61055381610538565b82525050565b60008115159050919050565b61056e81610559565b82525050565b6000608082019050610589600083018761054a565b61059660208301866104c1565b6105a360408301856104c1565b6105b06060830184610565565b95945050505050565b600082825260208201905092915050565b7f4ec3a36f206175746f72697a61646f0000000000000000000000000000000000600082015250565b6000610600600f836105b9565b915061060b826105ca565b602082019050919050565b6000602082019050818103600083015261062f816105f3565b9050919050565b7f4ac3a12066696e616c697a616461000000000000000000000000000000000000600082015250565b600061066c600e836105b9565b915061067782610636565b602082019050919050565b6000602082019050818103600083015261069b8161065f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106dc8261044b565b91506106e78361044b565b92508282026106f58161044b565b9150828204841483151761070c5761070b6106a2565b5b5092915050565b7f506167616d656e746f20696e737566696369656e746500000000000000000000600082015250565b60006107496016836105b9565b915061075482610713565b602082019050919050565b600060208201905081810360008301526107788161073c565b9050919050565b600060808201905061079460008301876104c1565b6107a1602083018661054a565b6107ae60408301856104c1565b6107bb60608301846104c1565b95945050505050565b60006107cf8261044b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610801576108006106a2565b5b60018201905091905056fea2646970667358221220cdb6a4b43555c0912c546daacf047a38f425da5249dd8164537187dff3ef69c664736f6c63430008140033",
}

// RecargaABI is the input ABI used to generate the binding from.
// Deprecated: Use RecargaMetaData.ABI instead.
var RecargaABI = RecargaMetaData.ABI

// RecargaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecargaMetaData.Bin instead.
var RecargaBin = RecargaMetaData.Bin

// DeployRecarga deploys a new Ethereum contract, binding an instance of Recarga to it.
func DeployRecarga(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recarga, error) {
	parsed, err := RecargaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecargaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recarga{RecargaCaller: RecargaCaller{contract: contract}, RecargaTransactor: RecargaTransactor{contract: contract}, RecargaFilterer: RecargaFilterer{contract: contract}}, nil
}

// Recarga is an auto generated Go binding around an Ethereum contract.
type Recarga struct {
	RecargaCaller     // Read-only binding to the contract
	RecargaTransactor // Write-only binding to the contract
	RecargaFilterer   // Log filterer for contract events
}

// RecargaCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecargaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecargaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecargaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecargaSession struct {
	Contract     *Recarga          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecargaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecargaCallerSession struct {
	Contract *RecargaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecargaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecargaTransactorSession struct {
	Contract     *RecargaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecargaRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecargaRaw struct {
	Contract *Recarga // Generic contract binding to access the raw methods on
}

// RecargaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecargaCallerRaw struct {
	Contract *RecargaCaller // Generic read-only contract binding to access the raw methods on
}

// RecargaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecargaTransactorRaw struct {
	Contract *RecargaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecarga creates a new instance of Recarga, bound to a specific deployed contract.
func NewRecarga(address common.Address, backend bind.ContractBackend) (*Recarga, error) {
	contract, err := bindRecarga(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recarga{RecargaCaller: RecargaCaller{contract: contract}, RecargaTransactor: RecargaTransactor{contract: contract}, RecargaFilterer: RecargaFilterer{contract: contract}}, nil
}

// NewRecargaCaller creates a new read-only instance of Recarga, bound to a specific deployed contract.
func NewRecargaCaller(address common.Address, caller bind.ContractCaller) (*RecargaCaller, error) {
	contract, err := bindRecarga(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaCaller{contract: contract}, nil
}

// NewRecargaTransactor creates a new write-only instance of Recarga, bound to a specific deployed contract.
func NewRecargaTransactor(address common.Address, transactor bind.ContractTransactor) (*RecargaTransactor, error) {
	contract, err := bindRecarga(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaTransactor{contract: contract}, nil
}

// NewRecargaFilterer creates a new log filterer instance of Recarga, bound to a specific deployed contract.
func NewRecargaFilterer(address common.Address, filterer bind.ContractFilterer) (*RecargaFilterer, error) {
	contract, err := bindRecarga(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecargaFilterer{contract: contract}, nil
}

// bindRecarga binds a generic wrapper to an already deployed contract.
func bindRecarga(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecargaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recarga *RecargaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recarga.Contract.RecargaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recarga *RecargaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recarga.Contract.RecargaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recarga *RecargaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recarga.Contract.RecargaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recarga *RecargaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recarga.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recarga *RecargaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recarga.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recarga *RecargaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recarga.Contract.contract.Transact(opts, method, params...)
}

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Recarga *RecargaCaller) Contador(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "contador")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Recarga *RecargaSession) Contador() (*big.Int, error) {
	return _Recarga.Contract.Contador(&_Recarga.CallOpts)
}

// Contador is a free data retrieval call binding the contract method 0x53d27cc2.
//
// Solidity: function contador() view returns(uint256)
func (_Recarga *RecargaCallerSession) Contador() (*big.Int, error) {
	return _Recarga.Contract.Contador(&_Recarga.CallOpts)
}

// PrecoPorKWh is a free data retrieval call binding the contract method 0xd5e628dc.
//
// Solidity: function precoPorKWh() view returns(uint256)
func (_Recarga *RecargaCaller) PrecoPorKWh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "precoPorKWh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecoPorKWh is a free data retrieval call binding the contract method 0xd5e628dc.
//
// Solidity: function precoPorKWh() view returns(uint256)
func (_Recarga *RecargaSession) PrecoPorKWh() (*big.Int, error) {
	return _Recarga.Contract.PrecoPorKWh(&_Recarga.CallOpts)
}

// PrecoPorKWh is a free data retrieval call binding the contract method 0xd5e628dc.
//
// Solidity: function precoPorKWh() view returns(uint256)
func (_Recarga *RecargaCallerSession) PrecoPorKWh() (*big.Int, error) {
	return _Recarga.Contract.PrecoPorKWh(&_Recarga.CallOpts)
}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(address usuario, uint256 energiaConsumida, uint256 valorPago, bool finalizada)
func (_Recarga *RecargaCaller) Sessoes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Usuario          common.Address
	EnergiaConsumida *big.Int
	ValorPago        *big.Int
	Finalizada       bool
}, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "sessoes", arg0)

	outstruct := new(struct {
		Usuario          common.Address
		EnergiaConsumida *big.Int
		ValorPago        *big.Int
		Finalizada       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Usuario = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EnergiaConsumida = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValorPago = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Finalizada = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(address usuario, uint256 energiaConsumida, uint256 valorPago, bool finalizada)
func (_Recarga *RecargaSession) Sessoes(arg0 *big.Int) (struct {
	Usuario          common.Address
	EnergiaConsumida *big.Int
	ValorPago        *big.Int
	Finalizada       bool
}, error) {
	return _Recarga.Contract.Sessoes(&_Recarga.CallOpts, arg0)
}

// Sessoes is a free data retrieval call binding the contract method 0xa718ae96.
//
// Solidity: function sessoes(uint256 ) view returns(address usuario, uint256 energiaConsumida, uint256 valorPago, bool finalizada)
func (_Recarga *RecargaCallerSession) Sessoes(arg0 *big.Int) (struct {
	Usuario          common.Address
	EnergiaConsumida *big.Int
	ValorPago        *big.Int
	Finalizada       bool
}, error) {
	return _Recarga.Contract.Sessoes(&_Recarga.CallOpts, arg0)
}

// FinalizarSessao is a paid mutator transaction binding the contract method 0x41273905.
//
// Solidity: function finalizarSessao(uint256 id, uint256 energiaConsumida) payable returns()
func (_Recarga *RecargaTransactor) FinalizarSessao(opts *bind.TransactOpts, id *big.Int, energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recarga.contract.Transact(opts, "finalizarSessao", id, energiaConsumida)
}

// FinalizarSessao is a paid mutator transaction binding the contract method 0x41273905.
//
// Solidity: function finalizarSessao(uint256 id, uint256 energiaConsumida) payable returns()
func (_Recarga *RecargaSession) FinalizarSessao(id *big.Int, energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.FinalizarSessao(&_Recarga.TransactOpts, id, energiaConsumida)
}

// FinalizarSessao is a paid mutator transaction binding the contract method 0x41273905.
//
// Solidity: function finalizarSessao(uint256 id, uint256 energiaConsumida) payable returns()
func (_Recarga *RecargaTransactorSession) FinalizarSessao(id *big.Int, energiaConsumida *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.FinalizarSessao(&_Recarga.TransactOpts, id, energiaConsumida)
}

// IniciarSessao is a paid mutator transaction binding the contract method 0x4cd57d63.
//
// Solidity: function iniciarSessao() returns(uint256)
func (_Recarga *RecargaTransactor) IniciarSessao(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recarga.contract.Transact(opts, "iniciarSessao")
}

// IniciarSessao is a paid mutator transaction binding the contract method 0x4cd57d63.
//
// Solidity: function iniciarSessao() returns(uint256)
func (_Recarga *RecargaSession) IniciarSessao() (*types.Transaction, error) {
	return _Recarga.Contract.IniciarSessao(&_Recarga.TransactOpts)
}

// IniciarSessao is a paid mutator transaction binding the contract method 0x4cd57d63.
//
// Solidity: function iniciarSessao() returns(uint256)
func (_Recarga *RecargaTransactorSession) IniciarSessao() (*types.Transaction, error) {
	return _Recarga.Contract.IniciarSessao(&_Recarga.TransactOpts)
}

// RecargaRecargaFinalizadaIterator is returned from FilterRecargaFinalizada and is used to iterate over the raw logs and unpacked data for RecargaFinalizada events raised by the Recarga contract.
type RecargaRecargaFinalizadaIterator struct {
	Event *RecargaRecargaFinalizada // Event containing the contract specifics and raw log

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
func (it *RecargaRecargaFinalizadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecargaRecargaFinalizada)
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
		it.Event = new(RecargaRecargaFinalizada)
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
func (it *RecargaRecargaFinalizadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecargaRecargaFinalizadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecargaRecargaFinalizada represents a RecargaFinalizada event raised by the Recarga contract.
type RecargaRecargaFinalizada struct {
	Id      *big.Int
	Usuario common.Address
	Energia *big.Int
	Valor   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecargaFinalizada is a free log retrieval operation binding the contract event 0x3ea1b19a3836fa0ed6579a4c083be8246ddc2b32d3930bc7bd5dd540027a41d6.
//
// Solidity: event RecargaFinalizada(uint256 id, address usuario, uint256 energia, uint256 valor)
func (_Recarga *RecargaFilterer) FilterRecargaFinalizada(opts *bind.FilterOpts) (*RecargaRecargaFinalizadaIterator, error) {

	logs, sub, err := _Recarga.contract.FilterLogs(opts, "RecargaFinalizada")
	if err != nil {
		return nil, err
	}
	return &RecargaRecargaFinalizadaIterator{contract: _Recarga.contract, event: "RecargaFinalizada", logs: logs, sub: sub}, nil
}

// WatchRecargaFinalizada is a free log subscription operation binding the contract event 0x3ea1b19a3836fa0ed6579a4c083be8246ddc2b32d3930bc7bd5dd540027a41d6.
//
// Solidity: event RecargaFinalizada(uint256 id, address usuario, uint256 energia, uint256 valor)
func (_Recarga *RecargaFilterer) WatchRecargaFinalizada(opts *bind.WatchOpts, sink chan<- *RecargaRecargaFinalizada) (event.Subscription, error) {

	logs, sub, err := _Recarga.contract.WatchLogs(opts, "RecargaFinalizada")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecargaRecargaFinalizada)
				if err := _Recarga.contract.UnpackLog(event, "RecargaFinalizada", log); err != nil {
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

// ParseRecargaFinalizada is a log parse operation binding the contract event 0x3ea1b19a3836fa0ed6579a4c083be8246ddc2b32d3930bc7bd5dd540027a41d6.
//
// Solidity: event RecargaFinalizada(uint256 id, address usuario, uint256 energia, uint256 valor)
func (_Recarga *RecargaFilterer) ParseRecargaFinalizada(log types.Log) (*RecargaRecargaFinalizada, error) {
	event := new(RecargaRecargaFinalizada)
	if err := _Recarga.contract.UnpackLog(event, "RecargaFinalizada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
