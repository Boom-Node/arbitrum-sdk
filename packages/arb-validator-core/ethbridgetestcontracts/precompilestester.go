// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// PrecompilesTesterABI is the input ABI used to generate the binding from.
const PrecompilesTesterABI = "[{\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"input\",\"type\":\"uint256[25]\"}],\"name\":\"keccakF\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"inputChunk\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashState\",\"type\":\"bytes32\"}],\"name\":\"sha256Block\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PrecompilesTesterFuncSigs maps the 4-byte function signature to its string representation.
var PrecompilesTesterFuncSigs = map[string]string{
	"ac90ed46": "keccakF(uint256[25])",
	"e479f532": "sha256Block(bytes32[2],bytes32)",
}

// PrecompilesTesterBin is the compiled bytecode used for deploying new contracts.
var PrecompilesTesterBin = "0x611587610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063ac90ed4614610045578063e479f532146100d2575b600080fd5b610099600480360361032081101561005c57600080fd5b8101908080610320019060198060200260405190810160405280929190826019602002808284376000920191909152509194506101319350505050565b604051808261032080838360005b838110156100bf5781810151838201526020016100a7565b5050505090500191505060405180910390f35b61011f600480360360608110156100e857600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201919091525091945050903591506101489050565b60408051918252519081900360200190f35b6101396114b7565b61014282610185565b92915050565b600061017e60405180604001604052808560006002811061016557fe5b6020908102919091015182528681015191015283610d73565b9392505050565b61018d6114b7565b6101956114d6565b61019d6114d6565b6101a56114b7565b6101ad6114f4565b60405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015610d68576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b8560036020020151816103f657fe5b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b85600460200201518161043157fe5b0485600460200201516002026001600160401b0316178560026005811061045457fe5b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b031617836002601981106106be57fe5b602002015260c0870151621000008104651000000000009091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b90910216919004176080840152610160870151670400000000000000906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036508000000000090910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b036302000000909102169190041761018084015261024087015165080000000000906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a087015165100000000000906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f6020020152826001602002015183601560200201511916836010602002015118876010602002015282600260200201518360166020020151191683601160200201511887601160200201528260036020020151836017602002015119168360126020020151188760126020020152826004602002015183601860200201511916836013602002015118876013602002015282600560200201518360006020020151191683601460200201511887601460200201528260066020020151836001602002015119168360156020020151188760156020020152826007602002015183600260200201511916836016602002015118876016602002015282600860200201518360036020020151191683601760200201511887601760200201528260096020020151836004602002015119168360186020020151188760186020020152818160188110610d5657fe5b602002015187511887526001016102d1565b509495945050505050565b6000610d7d611513565b50604080516108008101825263428a2f9881526371374491602082015263b5c0fbcf9181019190915263e9b5dba56060820152633956c25b60808201526359f111f160a082015263923f82a460c082015263ab1c5ed560e082015263d807aa986101008201526312835b0161012082015263243185be61014082015263550c7dc36101608201526372be5d746101808201526380deb1fe6101a0820152639bdc06a76101c082015263c19bf1746101e082015263e49b69c161020082015263efbe4786610220820152630fc19dc661024082015263240ca1cc610260820152632de92c6f610280820152634a7484aa6102a0820152635cb0a9dc6102c08201526376f988da6102e082015263983e515261030082015263a831c66d61032082015263b00327c861034082015263bf597fc761036082015263c6e00bf361038082015263d5a791476103a08201526306ca63516103c082015263142929676103e08201526327b70a85610400820152632e1b2138610420820152634d2c6dfc6104408201526353380d1361046082015263650a735461048082015263766a0abb6104a08201526381c2c92e6104c08201526392722c856104e082015263a2bfe8a161050082015263a81a664b61052082015263c24b8b7061054082015263c76c51a361056082015263d192e81961058082015263d69906246105a082015263f40e35856105c082015263106aa0706105e08201526319a4c116610600820152631e376c08610620820152632748774c6106408201526334b0bcb561066082015263391c0cb3610680820152634ed8aa4a6106a0820152635b9cca4f6106c082015263682e6ff36106e082015263748f82ee6107008201526378a5636f6107208201526384c87814610740820152638cc702086107608201526390befffa61078082015263a4506ceb6107a082015263bef9a3f76107c082015263c67178f26107e0820152611048611513565b60005b60088163ffffffff1610156110d55763ffffffff6020820260e003168660006020020151901c828263ffffffff166040811061108357fe5b63ffffffff92831660209182029290920191909152820260e003168660016020020151901c828260080163ffffffff16604081106110bd57fe5b63ffffffff909216602092909202015260010161104b565b5060106000805b60408363ffffffff16101561123157600384600f850363ffffffff166040811061110257fe5b602002015163ffffffff16901c61113385600f860363ffffffff166040811061112757fe5b60200201516012611490565b61115786600f870363ffffffff166040811061114b57fe5b60200201516007611490565b18189150600a846002850363ffffffff166040811061117257fe5b602002015163ffffffff16901c6111a3856002860363ffffffff166040811061119757fe5b60200201516013611490565b6111c7866002870363ffffffff16604081106111bb57fe5b60200201516011611490565b1818905080846007850363ffffffff16604081106111e157fe5b602002015183866010870363ffffffff16604081106111fc57fe5b6020020151010101848463ffffffff166040811061121657fe5b63ffffffff90921660209290920201526001909201916110dc565b611239611532565b600093505b60088463ffffffff16101561128a578360200260e00363ffffffff1688901c818563ffffffff166008811061126f57fe5b63ffffffff909216602092909202015260019093019261123e565b60008060008096505b60408763ffffffff1610156113df5760808401516112b2906019611490565b60808501516112c290600b611490565b60808601516112d2906006611490565b18189450878763ffffffff16604081106112e857fe5b6020020151898863ffffffff16604081106112ff57fe5b6020020151608086015160a087015160c088015161131e9291906114ae565b87876007602002015101010101925061133f84600060200201516016611490565b845161134c90600d611490565b8551611359906002611490565b6040870180516020890180518a5160c08c01805163ffffffff90811660e08f015260a08e018051821690925260808e018051821690925260608e0180518e01821690925280861690915280831690955284811690925280831891909116911618929091189290921881810186810190931687526001999099019897509092509050611293565b600096505b60088763ffffffff161015611433578660200260e00363ffffffff168b901c848863ffffffff166008811061141557fe5b60200201805163ffffffff92019190911690526001909601956113e4565b60008097505b60088863ffffffff161015611480578760200260e00363ffffffff16858963ffffffff166008811061146757fe5b602002015160019099019863ffffffff16901b17611439565b9c9b505050505050505050505050565b63ffffffff9182166020829003831681901b919092169190911c1790565b82191691161890565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061080001604052806040906020820280368337509192915050565b604051806101000160405280600890602082028036833750919291505056fea264697066735822122099e7841f86765416cac377f9dbf8c09763a31a47a620e618b8e06f786bcf0f1f64736f6c634300060c0033"

// DeployPrecompilesTester deploys a new Ethereum contract, binding an instance of PrecompilesTester to it.
func DeployPrecompilesTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrecompilesTester, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrecompilesTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrecompilesTester{PrecompilesTesterCaller: PrecompilesTesterCaller{contract: contract}, PrecompilesTesterTransactor: PrecompilesTesterTransactor{contract: contract}, PrecompilesTesterFilterer: PrecompilesTesterFilterer{contract: contract}}, nil
}

// PrecompilesTester is an auto generated Go binding around an Ethereum contract.
type PrecompilesTester struct {
	PrecompilesTesterCaller     // Read-only binding to the contract
	PrecompilesTesterTransactor // Write-only binding to the contract
	PrecompilesTesterFilterer   // Log filterer for contract events
}

// PrecompilesTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompilesTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompilesTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompilesTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompilesTesterSession struct {
	Contract     *PrecompilesTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PrecompilesTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompilesTesterCallerSession struct {
	Contract *PrecompilesTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PrecompilesTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompilesTesterTransactorSession struct {
	Contract     *PrecompilesTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PrecompilesTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompilesTesterRaw struct {
	Contract *PrecompilesTester // Generic contract binding to access the raw methods on
}

// PrecompilesTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompilesTesterCallerRaw struct {
	Contract *PrecompilesTesterCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompilesTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompilesTesterTransactorRaw struct {
	Contract *PrecompilesTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompilesTester creates a new instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTester(address common.Address, backend bind.ContractBackend) (*PrecompilesTester, error) {
	contract, err := bindPrecompilesTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTester{PrecompilesTesterCaller: PrecompilesTesterCaller{contract: contract}, PrecompilesTesterTransactor: PrecompilesTesterTransactor{contract: contract}, PrecompilesTesterFilterer: PrecompilesTesterFilterer{contract: contract}}, nil
}

// NewPrecompilesTesterCaller creates a new read-only instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterCaller(address common.Address, caller bind.ContractCaller) (*PrecompilesTesterCaller, error) {
	contract, err := bindPrecompilesTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterCaller{contract: contract}, nil
}

// NewPrecompilesTesterTransactor creates a new write-only instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompilesTesterTransactor, error) {
	contract, err := bindPrecompilesTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterTransactor{contract: contract}, nil
}

// NewPrecompilesTesterFilterer creates a new log filterer instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompilesTesterFilterer, error) {
	contract, err := bindPrecompilesTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterFilterer{contract: contract}, nil
}

// bindPrecompilesTester binds a generic wrapper to an already deployed contract.
func bindPrecompilesTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompilesTester *PrecompilesTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompilesTester.Contract.PrecompilesTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompilesTester *PrecompilesTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.PrecompilesTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompilesTester *PrecompilesTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.PrecompilesTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompilesTester *PrecompilesTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompilesTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompilesTester *PrecompilesTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompilesTester *PrecompilesTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterCaller) KeccakF(opts *bind.CallOpts, input [25]*big.Int) ([25]*big.Int, error) {
	var out []interface{}
	err := _PrecompilesTester.contract.Call(opts, &out, "keccakF", input)

	if err != nil {
		return *new([25]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([25]*big.Int)).(*[25]*big.Int)

	return out0, err

}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _PrecompilesTester.Contract.KeccakF(&_PrecompilesTester.CallOpts, input)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterCallerSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _PrecompilesTester.Contract.KeccakF(&_PrecompilesTester.CallOpts, input)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_PrecompilesTester *PrecompilesTesterCaller) Sha256Block(opts *bind.CallOpts, inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PrecompilesTester.contract.Call(opts, &out, "sha256Block", inputChunk, hashState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_PrecompilesTester *PrecompilesTesterSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _PrecompilesTester.Contract.Sha256Block(&_PrecompilesTester.CallOpts, inputChunk, hashState)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_PrecompilesTester *PrecompilesTesterCallerSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _PrecompilesTester.Contract.Sha256Block(&_PrecompilesTester.CallOpts, inputChunk, hashState)
}
