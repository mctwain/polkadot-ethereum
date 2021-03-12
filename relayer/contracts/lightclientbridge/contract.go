// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightclientbridge

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"contractValidatorRegistry\",\"name\":\"_validatorRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"FinalVerificationSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"InitialVerificationSuccessful\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_WAIT_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_NUM_SIGNERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validationData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"validatorClaimsBitfield\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"contractValidatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"validatorClaimsBitfield\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorSignatureCommitment\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"validatorPublicKey\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"validatorPublicKeyMerkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"newSignatureCommitment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"randomSignatureCommitments\",\"type\":\"bytes[]\"},{\"internalType\":\"uint8[]\",\"name\":\"randomSignatureBitfieldPositions\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"randomValidatorAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"randomPublicKeyMerkleProofs\",\"type\":\"bytes32[][]\"}],\"name\":\"completeSignatureCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BLOCKWAITPERIOD is a free data retrieval call binding the contract method 0xfb752c62.
//
// Solidity: function BLOCK_WAIT_PERIOD() view returns(uint256)
func (_Contract *ContractCaller) BLOCKWAITPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "BLOCK_WAIT_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKWAITPERIOD is a free data retrieval call binding the contract method 0xfb752c62.
//
// Solidity: function BLOCK_WAIT_PERIOD() view returns(uint256)
func (_Contract *ContractSession) BLOCKWAITPERIOD() (*big.Int, error) {
	return _Contract.Contract.BLOCKWAITPERIOD(&_Contract.CallOpts)
}

// BLOCKWAITPERIOD is a free data retrieval call binding the contract method 0xfb752c62.
//
// Solidity: function BLOCK_WAIT_PERIOD() view returns(uint256)
func (_Contract *ContractCallerSession) BLOCKWAITPERIOD() (*big.Int, error) {
	return _Contract.Contract.BLOCKWAITPERIOD(&_Contract.CallOpts)
}

// MAXIMUMNUMSIGNERS is a free data retrieval call binding the contract method 0x1cbb3307.
//
// Solidity: function MAXIMUM_NUM_SIGNERS() view returns(uint256)
func (_Contract *ContractCaller) MAXIMUMNUMSIGNERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAXIMUM_NUM_SIGNERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMNUMSIGNERS is a free data retrieval call binding the contract method 0x1cbb3307.
//
// Solidity: function MAXIMUM_NUM_SIGNERS() view returns(uint256)
func (_Contract *ContractSession) MAXIMUMNUMSIGNERS() (*big.Int, error) {
	return _Contract.Contract.MAXIMUMNUMSIGNERS(&_Contract.CallOpts)
}

// MAXIMUMNUMSIGNERS is a free data retrieval call binding the contract method 0x1cbb3307.
//
// Solidity: function MAXIMUM_NUM_SIGNERS() view returns(uint256)
func (_Contract *ContractCallerSession) MAXIMUMNUMSIGNERS() (*big.Int, error) {
	return _Contract.Contract.MAXIMUMNUMSIGNERS(&_Contract.CallOpts)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Contract *ContractCaller) CurrentId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Contract *ContractSession) CurrentId() (*big.Int, error) {
	return _Contract.Contract.CurrentId(&_Contract.CallOpts)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentId() (*big.Int, error) {
	return _Contract.Contract.CurrentId(&_Contract.CallOpts)
}

// ValidationData is a free data retrieval call binding the contract method 0x20bfa5cb.
//
// Solidity: function validationData(uint256 ) view returns(address senderAddress, bytes32 payload, uint256 validatorClaimsBitfield, uint256 blockNumber)
func (_Contract *ContractCaller) ValidationData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SenderAddress           common.Address
	Payload                 [32]byte
	ValidatorClaimsBitfield *big.Int
	BlockNumber             *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validationData", arg0)

	outstruct := new(struct {
		SenderAddress           common.Address
		Payload                 [32]byte
		ValidatorClaimsBitfield *big.Int
		BlockNumber             *big.Int
	})

	outstruct.SenderAddress = out[0].(common.Address)
	outstruct.Payload = out[1].([32]byte)
	outstruct.ValidatorClaimsBitfield = out[2].(*big.Int)
	outstruct.BlockNumber = out[3].(*big.Int)

	return *outstruct, err

}

// ValidationData is a free data retrieval call binding the contract method 0x20bfa5cb.
//
// Solidity: function validationData(uint256 ) view returns(address senderAddress, bytes32 payload, uint256 validatorClaimsBitfield, uint256 blockNumber)
func (_Contract *ContractSession) ValidationData(arg0 *big.Int) (struct {
	SenderAddress           common.Address
	Payload                 [32]byte
	ValidatorClaimsBitfield *big.Int
	BlockNumber             *big.Int
}, error) {
	return _Contract.Contract.ValidationData(&_Contract.CallOpts, arg0)
}

// ValidationData is a free data retrieval call binding the contract method 0x20bfa5cb.
//
// Solidity: function validationData(uint256 ) view returns(address senderAddress, bytes32 payload, uint256 validatorClaimsBitfield, uint256 blockNumber)
func (_Contract *ContractCallerSession) ValidationData(arg0 *big.Int) (struct {
	SenderAddress           common.Address
	Payload                 [32]byte
	ValidatorClaimsBitfield *big.Int
	BlockNumber             *big.Int
}, error) {
	return _Contract.Contract.ValidationData(&_Contract.CallOpts, arg0)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Contract *ContractCaller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Contract *ContractSession) ValidatorRegistry() (common.Address, error) {
	return _Contract.Contract.ValidatorRegistry(&_Contract.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Contract *ContractCallerSession) ValidatorRegistry() (common.Address, error) {
	return _Contract.Contract.ValidatorRegistry(&_Contract.CallOpts)
}

// CompleteSignatureCommitment is a paid mutator transaction binding the contract method 0x686bb405.
//
// Solidity: function completeSignatureCommitment(uint256 id, bytes32 payload, bytes[] randomSignatureCommitments, uint8[] randomSignatureBitfieldPositions, address[] randomValidatorAddresses, bytes32[][] randomPublicKeyMerkleProofs) returns()
func (_Contract *ContractTransactor) CompleteSignatureCommitment(opts *bind.TransactOpts, id *big.Int, payload [32]byte, randomSignatureCommitments [][]byte, randomSignatureBitfieldPositions []uint8, randomValidatorAddresses []common.Address, randomPublicKeyMerkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "completeSignatureCommitment", id, payload, randomSignatureCommitments, randomSignatureBitfieldPositions, randomValidatorAddresses, randomPublicKeyMerkleProofs)
}

// CompleteSignatureCommitment is a paid mutator transaction binding the contract method 0x686bb405.
//
// Solidity: function completeSignatureCommitment(uint256 id, bytes32 payload, bytes[] randomSignatureCommitments, uint8[] randomSignatureBitfieldPositions, address[] randomValidatorAddresses, bytes32[][] randomPublicKeyMerkleProofs) returns()
func (_Contract *ContractSession) CompleteSignatureCommitment(id *big.Int, payload [32]byte, randomSignatureCommitments [][]byte, randomSignatureBitfieldPositions []uint8, randomValidatorAddresses []common.Address, randomPublicKeyMerkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.CompleteSignatureCommitment(&_Contract.TransactOpts, id, payload, randomSignatureCommitments, randomSignatureBitfieldPositions, randomValidatorAddresses, randomPublicKeyMerkleProofs)
}

// CompleteSignatureCommitment is a paid mutator transaction binding the contract method 0x686bb405.
//
// Solidity: function completeSignatureCommitment(uint256 id, bytes32 payload, bytes[] randomSignatureCommitments, uint8[] randomSignatureBitfieldPositions, address[] randomValidatorAddresses, bytes32[][] randomPublicKeyMerkleProofs) returns()
func (_Contract *ContractTransactorSession) CompleteSignatureCommitment(id *big.Int, payload [32]byte, randomSignatureCommitments [][]byte, randomSignatureBitfieldPositions []uint8, randomValidatorAddresses []common.Address, randomPublicKeyMerkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.CompleteSignatureCommitment(&_Contract.TransactOpts, id, payload, randomSignatureCommitments, randomSignatureBitfieldPositions, randomValidatorAddresses, randomPublicKeyMerkleProofs)
}

// NewSignatureCommitment is a paid mutator transaction binding the contract method 0x8da129e7.
//
// Solidity: function newSignatureCommitment(bytes32 payload, uint256 validatorClaimsBitfield, bytes validatorSignatureCommitment, address validatorPublicKey, bytes32[] validatorPublicKeyMerkleProof) payable returns()
func (_Contract *ContractTransactor) NewSignatureCommitment(opts *bind.TransactOpts, payload [32]byte, validatorClaimsBitfield *big.Int, validatorSignatureCommitment []byte, validatorPublicKey common.Address, validatorPublicKeyMerkleProof [][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "newSignatureCommitment", payload, validatorClaimsBitfield, validatorSignatureCommitment, validatorPublicKey, validatorPublicKeyMerkleProof)
}

// NewSignatureCommitment is a paid mutator transaction binding the contract method 0x8da129e7.
//
// Solidity: function newSignatureCommitment(bytes32 payload, uint256 validatorClaimsBitfield, bytes validatorSignatureCommitment, address validatorPublicKey, bytes32[] validatorPublicKeyMerkleProof) payable returns()
func (_Contract *ContractSession) NewSignatureCommitment(payload [32]byte, validatorClaimsBitfield *big.Int, validatorSignatureCommitment []byte, validatorPublicKey common.Address, validatorPublicKeyMerkleProof [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NewSignatureCommitment(&_Contract.TransactOpts, payload, validatorClaimsBitfield, validatorSignatureCommitment, validatorPublicKey, validatorPublicKeyMerkleProof)
}

// NewSignatureCommitment is a paid mutator transaction binding the contract method 0x8da129e7.
//
// Solidity: function newSignatureCommitment(bytes32 payload, uint256 validatorClaimsBitfield, bytes validatorSignatureCommitment, address validatorPublicKey, bytes32[] validatorPublicKeyMerkleProof) payable returns()
func (_Contract *ContractTransactorSession) NewSignatureCommitment(payload [32]byte, validatorClaimsBitfield *big.Int, validatorSignatureCommitment []byte, validatorPublicKey common.Address, validatorPublicKeyMerkleProof [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NewSignatureCommitment(&_Contract.TransactOpts, payload, validatorClaimsBitfield, validatorSignatureCommitment, validatorPublicKey, validatorPublicKeyMerkleProof)
}

// ContractFinalVerificationSuccessfulIterator is returned from FilterFinalVerificationSuccessful and is used to iterate over the raw logs and unpacked data for FinalVerificationSuccessful events raised by the Contract contract.
type ContractFinalVerificationSuccessfulIterator struct {
	Event *ContractFinalVerificationSuccessful // Event containing the contract specifics and raw log

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
func (it *ContractFinalVerificationSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFinalVerificationSuccessful)
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
		it.Event = new(ContractFinalVerificationSuccessful)
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
func (it *ContractFinalVerificationSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFinalVerificationSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFinalVerificationSuccessful represents a FinalVerificationSuccessful event raised by the Contract contract.
type ContractFinalVerificationSuccessful struct {
	Prover  common.Address
	Payload [32]byte
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalVerificationSuccessful is a free log retrieval operation binding the contract event 0x064680bd49185bfe9d29035c0d5caec57f27902d4f2037faa81246a9d686a106.
//
// Solidity: event FinalVerificationSuccessful(address prover, bytes32 payload, uint256 id)
func (_Contract *ContractFilterer) FilterFinalVerificationSuccessful(opts *bind.FilterOpts) (*ContractFinalVerificationSuccessfulIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FinalVerificationSuccessful")
	if err != nil {
		return nil, err
	}
	return &ContractFinalVerificationSuccessfulIterator{contract: _Contract.contract, event: "FinalVerificationSuccessful", logs: logs, sub: sub}, nil
}

// WatchFinalVerificationSuccessful is a free log subscription operation binding the contract event 0x064680bd49185bfe9d29035c0d5caec57f27902d4f2037faa81246a9d686a106.
//
// Solidity: event FinalVerificationSuccessful(address prover, bytes32 payload, uint256 id)
func (_Contract *ContractFilterer) WatchFinalVerificationSuccessful(opts *bind.WatchOpts, sink chan<- *ContractFinalVerificationSuccessful) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FinalVerificationSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFinalVerificationSuccessful)
				if err := _Contract.contract.UnpackLog(event, "FinalVerificationSuccessful", log); err != nil {
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

// ParseFinalVerificationSuccessful is a log parse operation binding the contract event 0x064680bd49185bfe9d29035c0d5caec57f27902d4f2037faa81246a9d686a106.
//
// Solidity: event FinalVerificationSuccessful(address prover, bytes32 payload, uint256 id)
func (_Contract *ContractFilterer) ParseFinalVerificationSuccessful(log types.Log) (*ContractFinalVerificationSuccessful, error) {
	event := new(ContractFinalVerificationSuccessful)
	if err := _Contract.contract.UnpackLog(event, "FinalVerificationSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitialVerificationSuccessfulIterator is returned from FilterInitialVerificationSuccessful and is used to iterate over the raw logs and unpacked data for InitialVerificationSuccessful events raised by the Contract contract.
type ContractInitialVerificationSuccessfulIterator struct {
	Event *ContractInitialVerificationSuccessful // Event containing the contract specifics and raw log

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
func (it *ContractInitialVerificationSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialVerificationSuccessful)
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
		it.Event = new(ContractInitialVerificationSuccessful)
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
func (it *ContractInitialVerificationSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitialVerificationSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialVerificationSuccessful represents a InitialVerificationSuccessful event raised by the Contract contract.
type ContractInitialVerificationSuccessful struct {
	Prover      common.Address
	BlockNumber *big.Int
	Id          *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialVerificationSuccessful is a free log retrieval operation binding the contract event 0xf93e67b44fe47465ca1478dbc20efe59702e2fff4b8beecf053817d7ee29fd55.
//
// Solidity: event InitialVerificationSuccessful(address prover, uint256 blockNumber, uint256 id)
func (_Contract *ContractFilterer) FilterInitialVerificationSuccessful(opts *bind.FilterOpts) (*ContractInitialVerificationSuccessfulIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InitialVerificationSuccessful")
	if err != nil {
		return nil, err
	}
	return &ContractInitialVerificationSuccessfulIterator{contract: _Contract.contract, event: "InitialVerificationSuccessful", logs: logs, sub: sub}, nil
}

// WatchInitialVerificationSuccessful is a free log subscription operation binding the contract event 0xf93e67b44fe47465ca1478dbc20efe59702e2fff4b8beecf053817d7ee29fd55.
//
// Solidity: event InitialVerificationSuccessful(address prover, uint256 blockNumber, uint256 id)
func (_Contract *ContractFilterer) WatchInitialVerificationSuccessful(opts *bind.WatchOpts, sink chan<- *ContractInitialVerificationSuccessful) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InitialVerificationSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialVerificationSuccessful)
				if err := _Contract.contract.UnpackLog(event, "InitialVerificationSuccessful", log); err != nil {
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

// ParseInitialVerificationSuccessful is a log parse operation binding the contract event 0xf93e67b44fe47465ca1478dbc20efe59702e2fff4b8beecf053817d7ee29fd55.
//
// Solidity: event InitialVerificationSuccessful(address prover, uint256 blockNumber, uint256 id)
func (_Contract *ContractFilterer) ParseInitialVerificationSuccessful(log types.Log) (*ContractInitialVerificationSuccessful, error) {
	event := new(ContractInitialVerificationSuccessful)
	if err := _Contract.contract.UnpackLog(event, "InitialVerificationSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
