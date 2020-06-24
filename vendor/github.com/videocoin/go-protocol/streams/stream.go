// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package streams

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StreamABI is the input ABI used to generate the binding from.
const StreamABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"profiles\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"AccountFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofScrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofSubmited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"ChunkProofValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OutOfFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"ServiceFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"service\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunckId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ServiceShareCollected\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"wattage\",\"type\":\"uint256[]\"}],\"name\":\"addInputChunkId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"client\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"endStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getCandidateProof\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInChunkCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInChunks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"}],\"name\":\"getOutChunks\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProfileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getProofCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"getValidProof\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputChunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proof\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getprofiles\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"hasValidProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isChunk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"}],\"name\":\"isProfileTranscoded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTranscodingDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outStreams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"required\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatedChunks\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"refundAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"scrapProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proof\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outChunkId\",\"type\":\"uint256\"}],\"name\":\"submitProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"validateProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wattages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StreamBin is the compiled bytecode used for deploying new contracts.
var StreamBin = "0x60806040523480156200001157600080fd5b50604051620027c5380380620027c5833981810160405260608110156200003757600080fd5b810190808051906020019092919080519060200190929190805160405193929190846401000000008211156200006c57600080fd5b838201915060208201858111156200008357600080fd5b8251866020820283011164010000000082111715620000a157600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000da578082015181840152602081019050620000bd565b5050505090500160405250505081600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200012357600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508260028190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060089080519060200190620001c49291906200027a565b5060008090505b6008805490508110156200027057620001e3620002cc565b6001816000019015159081151581525050818160200181815250508060076000600885815481106200021157fe5b9060005260206000200154815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155905050508080600101915050620001cb565b5050505062000317565b828054828255906000526020600020908101928215620002b9579160200282015b82811115620002b85782518255916020019190600101906200029b565b5b509050620002c89190620002ef565b5090565b604051806060016040528060001515815260200160008152602001600081525090565b6200031491905b8082111562000310576000816000905550600101620002f6565b5090565b90565b61249e80620003276000396000f3fe6080604052600436106101b65760003560e01c806362372298116100ec578063bf032f531161008a578063d0e30db011610064578063d0e30db014610a8e578063d78e647f14610a98578063eda0ce1714610add578063fc1028bc14610b36576101b6565b8063bf032f53146108fb578063c5d0b14c146109c1578063c617519314610a22576101b6565b80637b40855d116100c65780637b40855d1461081b578063af640d0f14610874578063bb57a5ed1461089f578063bbe58b0c146108b6576101b6565b8063623722981461070757806373f93b2a14610797578063747f7589146107c2576101b6565b80632f750f20116101595780633fa911ae116101335780633fa911ae1461060d578063481c6a751461063c5780634c0b715c14610693578063590e1ae3146106f0576101b6565b80632f750f2014610480578063356939ab146105135780633697611a146105e2576101b6565b806312fa6feb1161019557806312fa6feb1461030e5780631bb62fc41461033d5780631f54e5bd1461039057806328cc413a146103e3576101b6565b8062ca5d92146101bb5780630d8e6e2c14610227578063109e94cf146102b7575b600080fd5b3480156101c757600080fd5b506101d0610b65565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102135780820151818401526020810190506101f8565b505050509050019250505060405180910390f35b34801561023357600080fd5b5061023c610bbd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561027c578082015181840152602081019050610261565b50505050905090810190601f1680156102a95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102c357600080fd5b506102cc610d25565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561031a57600080fd5b50610323610d4a565b604051808215151515815260200191505060405180910390f35b34801561034957600080fd5b506103766004803603602081101561036057600080fd5b8101908080359060200190929190505050610d5d565b604051808215151515815260200191505060405180910390f35b34801561039c57600080fd5b506103c9600480360360208110156103b357600080fd5b8101908080359060200190929190505050610d7d565b604051808215151515815260200191505060405180910390f35b3480156103ef57600080fd5b506104306004803603606081101561040657600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610dc4565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b34801561048c57600080fd5b506104c3600480360360408110156104a357600080fd5b810190808035906020019092919080359060200190929190505050610ef2565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b34801561051f57600080fd5b506105e06004803603604081101561053657600080fd5b81019080803590602001909291908035906020019064010000000081111561055d57600080fd5b82018360208201111561056f57600080fd5b8035906020019184602083028401116401000000008311171561059157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610f83565b005b3480156105ee57600080fd5b506105f76110a5565b6040518082815260200191505060405180910390f35b34801561061957600080fd5b506106226110b2565b604051808215151515815260200191505060405180910390f35b34801561064857600080fd5b5061065161119b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561069f57600080fd5b506106d6600480360360408110156106b657600080fd5b8101908080359060200190929190803590602001909291905050506111c1565b604051808215151515815260200191505060405180910390f35b3480156106fc57600080fd5b506107056112a1565b005b34801561071357600080fd5b506107406004803603602081101561072a57600080fd5b8101908080359060200190929190505050611359565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610783578082015181840152602081019050610768565b505050509050019250505060405180910390f35b3480156107a357600080fd5b506107ac611468565b6040518082815260200191505060405180910390f35b3480156107ce57600080fd5b50610819600480360360808110156107e557600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611475565b005b34801561082757600080fd5b5061085e6004803603604081101561083e57600080fd5b810190808035906020019092919080359060200190929190505050611651565b6040518082815260200191505060405180910390f35b34801561088057600080fd5b506108896116e4565b6040518082815260200191505060405180910390f35b3480156108ab57600080fd5b506108b46116ea565b005b3480156108c257600080fd5b506108f9600480360360408110156108d957600080fd5b81019080803590602001909291908035906020019092919050505061177b565b005b34801561090757600080fd5b5061093e6004803603604081101561091e57600080fd5b810190808035906020019092919080359060200190929190505050611998565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200194505050505060405180910390f35b3480156109cd57600080fd5b506109fa600480360360208110156109e457600080fd5b8101908080359060200190929190505050611aa3565b6040518084151515158152602001838152602001828152602001935050505060405180910390f35b348015610a2e57600080fd5b50610a37611ada565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610a7a578082015181840152602081019050610a5f565b505050509050019250505060405180910390f35b610a96611b32565b005b348015610aa457600080fd5b50610adb60048036036040811015610abb57600080fd5b810190808035906020019092919080359060200190929190505050611b61565b005b348015610ae957600080fd5b50610b2060048036036040811015610b0057600080fd5b810190808035906020019092919080359060200190929190505050611f3d565b6040518082815260200191505060405180910390f35b348015610b4257600080fd5b50610b4b611f6b565b604051808215151515815260200191505060405180910390f35b60606008805480602002602001604051908101604052809291908181526020018280548015610bb357602002820191906000526020600020905b815481526020019060010190808311610b9f575b5050505050905090565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d8e6e2c6040518163ffffffff1660e01b815260040160006040518083038186803b158015610c2757600080fd5b505afa158015610c3b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610c6557600080fd5b8101908080516040519392919084640100000000821115610c8557600080fd5b83820191506020820185811115610c9b57600080fd5b8251866001820283011164010000000082111715610cb857600080fd5b8083526020830192505050908051906020019080838360005b83811015610cec578082015181840152602081019050610cd1565b50505050905090810190601f168015610d195780820380516001836020036101000a031916815260200191505b50604052505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900460ff1681565b60046020528060005260406000206000915054906101000a900460ff1681565b6000806007600084815260200190815260200160002090508060000160009054906101000a900460ff16610db057600080fd5b600680549050816002015414915050919050565b60008060008060076000888152602001908152602001600020600301600087815260200190815260200160002090506007600088815260200190815260200160002060000160009054906101000a900460ff168015610e4057506004600087815260200190815260200160002060009054906101000a900460ff165b610e4957600080fd5b80600001805490508510610e5c57600080fd5b806000018581548110610e6b57fe5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816000018681548110610eac57fe5b906000526020600020906003020160010154826000018781548110610ecd57fe5b9060005260206000209060030201600201548292509350935093505093509350939050565b600080600080600760008781526020019081526020016000206003016000868152602001908152602001600020905060008160000180549050905060008111610f3a57600080fd5b60008260010154905081811015610f6457610f56888883610dc4565b955095509550505050610f7c565b60008060008292508191508090509550955095505050505b9250925092565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fdd57600080fd5b6004600083815260200190815260200160002060009054906101000a900460ff161580156110185750600360009054906101000a900460ff16155b61102157600080fd5b60016004600084815260200190815260200160002060006101000a81548160ff021916908315150217905550806005600084815260200190815260200160002090805190602001906110749291906123d6565b5060068290806001815401808255809150509060018203906000526020600020016000909192909190915055505050565b6000600880549050905090565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630f5147176002546040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561112a57600080fd5b505afa15801561113e573d6000803e3d6000fd5b505050506040513d602081101561115457600080fd5b810190808051906020019092919050505090506000611171611f6b565b80156111895750600360009054906101000a900460ff165b905081806111945750805b9250505090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806007600085815260200190815260200160002090506004600084815260200190815260200160002060009054906101000a900460ff16801561121457508060000160009054906101000a900460ff165b61121d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff1660076000868152602001908152602001600020600301600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141591505092915050565b6112a96110b2565b6112b257600080fd5b60004790506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561131e573d6000803e3d6000fd5b507f3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c816040518082815260200191505060405180910390a150565b606061136482610d7d565b61136d57600080fd5b600060076000848152602001908152602001600020905060606006805490506040519080825280602002602001820160405280156113ba5781602001602082028038833980820191505090505b50905060008090505b60068054905081101561145d576000600682815481106113df57fe5b9060005260206000200154905060008460030160008381526020019081526020016000209050600081600101549050600082600001828154811061141f57fe5b90600052602060002090600302016001015490508086868151811061144057fe5b6020026020010181815250505050505080806001019150506113c3565b508092505050919050565b6000600680549050905090565b600060076000868152602001908152602001600020600301600085815260200190815260200160002090506004600085815260200190815260200160002060009054906101000a900460ff16801561151d5750600073ffffffffffffffffffffffffffffffffffffffff168160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61152657600080fd5b6007600086815260200190815260200160002060000160009054906101000a900460ff1661155357600080fd5b8060000160405180606001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001848152602001858152509080600181540180825580915050906001820390600052602060002090600302016000909192909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155505050600181600001805490500385857fe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa7160405160405180910390a45050505050565b60008060076000858152602001908152602001600020600301600084815260200190815260200160002090506007600085815260200190815260200160002060000160009054906101000a900460ff1680156116ca57506004600084815260200190815260200160002060009054906101000a900460ff165b6116d357600080fd5b806000018054905091505092915050565b60025481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461174457600080fd5b600360009054906101000a900460ff161561175e57600080fd5b6001600360006101000a81548160ff021916908315150217905550565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561181a57600080fd5b505afa15801561182e573d6000803e3d6000fd5b505050506040513d602081101561184457600080fd5b810190808051906020019092919050505061185e57600080fd5b600060076000848152602001908152602001600020600301600083815260200190815260200160002090506004600083815260200190815260200160002060009054906101000a900460ff1680156119065750600073ffffffffffffffffffffffffffffffffffffffff168160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61190f57600080fd5b806000018054905061192f60018360010154611fc490919063ffffffff16565b111561193a57600080fd5b60008160010154905061195b60018360010154611fc490919063ffffffff16565b82600101819055508083857fb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a60405160405180910390a450505050565b6000806000806119a886866111c1565b6119b157600080fd5b600060076000888152602001908152602001600020600301600087815260200190815260200160002090506000816001015490508160000181815481106119f457fe5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836000018381548110611a5a57fe5b906000526020600020906003020160010154846000018481548110611a7b57fe5b9060005260206000209060030201600201548393509550955095509550505092959194509250565b60076020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060020154905083565b60606006805480602002602001604051908101604052809291908181526020018280548015611b2857602002820191906000526020600020905b815481526020019060010190808311611b14575b5050505050905090565b347f2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea60405160405180910390a2565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611c0057600080fd5b505afa158015611c14573d6000803e3d6000fd5b505050506040513d6020811015611c2a57600080fd5b8101908080519060200190929190505050611c4457600080fd5b60006007600084815260200190815260200160002090508060000160009054906101000a900460ff16611c7657600080fd5b600081600301600084815260200190815260200160002090506004600084815260200190815260200160002060009054906101000a900460ff168015611d0c5750600073ffffffffffffffffffffffffffffffffffffffff168160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b611d1557600080fd5b8060000180549050816001015410611d2c57600080fd5b600081600001826001015481548110611d4157fe5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323be9d276040518163ffffffff1660e01b815260040160206040518083038186803b158015611ddf57600080fd5b505afa158015611df3573d6000803e3d6000fd5b505050506040513d6020811015611e0957600080fd5b81019080805190602001909291905050509050600060056000878152602001908152602001600020856001015481548110611e4057fe5b906000526020600020015490506000611e756064611e67858561204c90919063ffffffff16565b6120d290919063ffffffff16565b9050611e8a818361211c90919063ffffffff16565b91506000611e99858484612166565b905080611eac5750505050505050611f39565b338660020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866002016000815480929190600101919050555087897f6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e3760405160405180910390a3505050505050505b5050565b60056020528160005260406000208181548110611f5657fe5b90600052602060002001600091509150505481565b600080600090505b600880549050811015611fbb57611fa060088281548110611f9057fe5b9060005260206000200154610d7d565b611fae576000915050611fc1565b8080600101915050611f73565b50600190505b90565b600080828401905083811015612042576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60008083141561205f57600090506120cc565b600082840290508284828161207057fe5b04146120c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806124496021913960400191505060405180910390fd5b809150505b92915050565b600061211483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250612250565b905092915050565b600061215e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612316565b905092915050565b60008047905060008385019050600082148061218157508082105b156121bd577f5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d60405160405180910390a1600092505050612249565b8573ffffffffffffffffffffffffffffffffffffffff167fbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7866040518082815260200191505060405180910390a27f65d7a830175a146c5afc99c09332114c9fe32fc76e6646d79b1409350cade344846040518082815260200191505060405180910390a16001925050505b9392505050565b600080831182906122fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156122c15780820151818401526020810190506122a6565b50505050905090810190601f1680156122ee5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161230857fe5b049050809150509392505050565b60008383111582906123c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561238857808201518184015260208101905061236d565b50505050905090810190601f1680156123b55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b828054828255906000526020600020908101928215612412579160200282015b828111156124115782518255916020019190600101906123f6565b5b50905061241f9190612423565b5090565b61244591905b80821115612441576000816000905550600101612429565b5090565b9056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a723158204c67c794963d036c7819b4324f84698ecf61b3e61db5d5791c8a87e54bee6f6364736f6c63430005100032"

// DeployStream deploys a new Ethereum contract, binding an instance of Stream to it.
func DeployStream(auth *bind.TransactOpts, backend bind.ContractBackend, _id *big.Int, client common.Address, profiles []*big.Int) (common.Address, *types.Transaction, *Stream, error) {
	parsed, err := abi.JSON(strings.NewReader(StreamABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StreamBin), backend, _id, client, profiles)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stream{StreamCaller: StreamCaller{contract: contract}, StreamTransactor: StreamTransactor{contract: contract}, StreamFilterer: StreamFilterer{contract: contract}}, nil
}

// Stream is an auto generated Go binding around an Ethereum contract.
type Stream struct {
	StreamCaller     // Read-only binding to the contract
	StreamTransactor // Write-only binding to the contract
	StreamFilterer   // Log filterer for contract events
}

// StreamCaller is an auto generated read-only Go binding around an Ethereum contract.
type StreamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StreamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StreamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StreamSession struct {
	Contract     *Stream           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StreamCallerSession struct {
	Contract *StreamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StreamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StreamTransactorSession struct {
	Contract     *StreamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamRaw is an auto generated low-level Go binding around an Ethereum contract.
type StreamRaw struct {
	Contract *Stream // Generic contract binding to access the raw methods on
}

// StreamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StreamCallerRaw struct {
	Contract *StreamCaller // Generic read-only contract binding to access the raw methods on
}

// StreamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StreamTransactorRaw struct {
	Contract *StreamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStream creates a new instance of Stream, bound to a specific deployed contract.
func NewStream(address common.Address, backend bind.ContractBackend) (*Stream, error) {
	contract, err := bindStream(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stream{StreamCaller: StreamCaller{contract: contract}, StreamTransactor: StreamTransactor{contract: contract}, StreamFilterer: StreamFilterer{contract: contract}}, nil
}

// NewStreamCaller creates a new read-only instance of Stream, bound to a specific deployed contract.
func NewStreamCaller(address common.Address, caller bind.ContractCaller) (*StreamCaller, error) {
	contract, err := bindStream(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StreamCaller{contract: contract}, nil
}

// NewStreamTransactor creates a new write-only instance of Stream, bound to a specific deployed contract.
func NewStreamTransactor(address common.Address, transactor bind.ContractTransactor) (*StreamTransactor, error) {
	contract, err := bindStream(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StreamTransactor{contract: contract}, nil
}

// NewStreamFilterer creates a new log filterer instance of Stream, bound to a specific deployed contract.
func NewStreamFilterer(address common.Address, filterer bind.ContractFilterer) (*StreamFilterer, error) {
	contract, err := bindStream(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StreamFilterer{contract: contract}, nil
}

// bindStream binds a generic wrapper to an already deployed contract.
func bindStream(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StreamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.StreamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transact(opts, method, params...)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "client")
	return *ret0, err
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() constant returns(address)
func (_Stream *StreamCallerSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "ended")
	return *ret0, err
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() constant returns(bool)
func (_Stream *StreamCallerSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetCandidateProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getCandidateProof", profile, chunkId)
	return *ret, err
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetCandidateProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetCandidateProof(&_Stream.CallOpts, profile, chunkId)
}

// GetCandidateProof is a free data retrieval call binding the contract method 0x2f750f20.
//
// Solidity: function getCandidateProof(uint256 profile, uint256 chunkId) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetCandidateProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetCandidateProof(&_Stream.CallOpts, profile, chunkId)
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamCaller) GetInChunkCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getInChunkCount")
	return *ret0, err
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamSession) GetInChunkCount() (*big.Int, error) {
	return _Stream.Contract.GetInChunkCount(&_Stream.CallOpts)
}

// GetInChunkCount is a free data retrieval call binding the contract method 0x73f93b2a.
//
// Solidity: function getInChunkCount() constant returns(uint256)
func (_Stream *StreamCallerSession) GetInChunkCount() (*big.Int, error) {
	return _Stream.Contract.GetInChunkCount(&_Stream.CallOpts)
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamCaller) GetInChunks(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getInChunks")
	return *ret0, err
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamSession) GetInChunks() ([]*big.Int, error) {
	return _Stream.Contract.GetInChunks(&_Stream.CallOpts)
}

// GetInChunks is a free data retrieval call binding the contract method 0xc6175193.
//
// Solidity: function getInChunks() constant returns(uint256[])
func (_Stream *StreamCallerSession) GetInChunks() ([]*big.Int, error) {
	return _Stream.Contract.GetInChunks(&_Stream.CallOpts)
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamCaller) GetOutChunks(opts *bind.CallOpts, profile *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getOutChunks", profile)
	return *ret0, err
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamSession) GetOutChunks(profile *big.Int) ([]*big.Int, error) {
	return _Stream.Contract.GetOutChunks(&_Stream.CallOpts, profile)
}

// GetOutChunks is a free data retrieval call binding the contract method 0x62372298.
//
// Solidity: function getOutChunks(uint256 profile) constant returns(uint256[])
func (_Stream *StreamCallerSession) GetOutChunks(profile *big.Int) ([]*big.Int, error) {
	return _Stream.Contract.GetOutChunks(&_Stream.CallOpts, profile)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamCaller) GetProfileCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getProfileCount")
	return *ret0, err
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamSession) GetProfileCount() (*big.Int, error) {
	return _Stream.Contract.GetProfileCount(&_Stream.CallOpts)
}

// GetProfileCount is a free data retrieval call binding the contract method 0x3697611a.
//
// Solidity: function getProfileCount() constant returns(uint256)
func (_Stream *StreamCallerSession) GetProfileCount() (*big.Int, error) {
	return _Stream.Contract.GetProfileCount(&_Stream.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getProof", profile, chunkId, idx)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetProof(profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetProof(&_Stream.CallOpts, profile, chunkId, idx)
}

// GetProof is a free data retrieval call binding the contract method 0x28cc413a.
//
// Solidity: function getProof(uint256 profile, uint256 chunkId, uint256 idx) constant returns(address miner, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetProof(profile *big.Int, chunkId *big.Int, idx *big.Int) (struct {
	Miner         common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetProof(&_Stream.CallOpts, profile, chunkId, idx)
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamCaller) GetProofCount(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getProofCount", profile, chunkId)
	return *ret0, err
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamSession) GetProofCount(profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	return _Stream.Contract.GetProofCount(&_Stream.CallOpts, profile, chunkId)
}

// GetProofCount is a free data retrieval call binding the contract method 0x7b40855d.
//
// Solidity: function getProofCount(uint256 profile, uint256 chunkId) constant returns(uint256)
func (_Stream *StreamCallerSession) GetProofCount(profile *big.Int, chunkId *big.Int) (*big.Int, error) {
	return _Stream.Contract.GetProofCount(&_Stream.CallOpts, profile, chunkId)
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCaller) GetValidProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	ret := new(struct {
		Miner         common.Address
		Validator     common.Address
		OutputChunkId *big.Int
		Proof         *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "getValidProof", profile, chunkId)
	return *ret, err
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamSession) GetValidProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetValidProof(&_Stream.CallOpts, profile, chunkId)
}

// GetValidProof is a free data retrieval call binding the contract method 0xbf032f53.
//
// Solidity: function getValidProof(uint256 profile, uint256 chunkId) constant returns(address miner, address validator, uint256 outputChunkId, uint256 proof)
func (_Stream *StreamCallerSession) GetValidProof(profile *big.Int, chunkId *big.Int) (struct {
	Miner         common.Address
	Validator     common.Address
	OutputChunkId *big.Int
	Proof         *big.Int
}, error) {
	return _Stream.Contract.GetValidProof(&_Stream.CallOpts, profile, chunkId)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Stream *StreamCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getVersion")
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Stream *StreamSession) GetVersion() (string, error) {
	return _Stream.Contract.GetVersion(&_Stream.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Stream *StreamCallerSession) GetVersion() (string, error) {
	return _Stream.Contract.GetVersion(&_Stream.CallOpts)
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamCaller) Getprofiles(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "getprofiles")
	return *ret0, err
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamSession) Getprofiles() ([]*big.Int, error) {
	return _Stream.Contract.Getprofiles(&_Stream.CallOpts)
}

// Getprofiles is a free data retrieval call binding the contract method 0x00ca5d92.
//
// Solidity: function getprofiles() constant returns(uint256[])
func (_Stream *StreamCallerSession) Getprofiles() ([]*big.Int, error) {
	return _Stream.Contract.Getprofiles(&_Stream.CallOpts)
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamCaller) HasValidProof(opts *bind.CallOpts, profile *big.Int, chunkId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "hasValidProof", profile, chunkId)
	return *ret0, err
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamSession) HasValidProof(profile *big.Int, chunkId *big.Int) (bool, error) {
	return _Stream.Contract.HasValidProof(&_Stream.CallOpts, profile, chunkId)
}

// HasValidProof is a free data retrieval call binding the contract method 0x4c0b715c.
//
// Solidity: function hasValidProof(uint256 profile, uint256 chunkId) constant returns(bool)
func (_Stream *StreamCallerSession) HasValidProof(profile *big.Int, chunkId *big.Int) (bool, error) {
	return _Stream.Contract.HasValidProof(&_Stream.CallOpts, profile, chunkId)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(uint256)
func (_Stream *StreamCallerSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamCaller) IsChunk(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isChunk", arg0)
	return *ret0, err
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) constant returns(bool)
func (_Stream *StreamCallerSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamCaller) IsProfileTranscoded(opts *bind.CallOpts, profile *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isProfileTranscoded", profile)
	return *ret0, err
}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamSession) IsProfileTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsProfileTranscoded(&_Stream.CallOpts, profile)
}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) constant returns(bool)
func (_Stream *StreamCallerSession) IsProfileTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsProfileTranscoded(&_Stream.CallOpts, profile)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamCaller) IsTranscodingDone(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "isTranscodingDone")
	return *ret0, err
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() constant returns(bool)
func (_Stream *StreamCallerSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Stream *StreamCallerSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCaller) OutStreams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	ret := new(struct {
		Required        bool
		Index           *big.Int
		ValidatedChunks *big.Int
	})
	out := ret
	err := _Stream.contract.Call(opts, out, "outStreams", arg0)
	return *ret, err
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) constant returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCallerSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamCaller) RefundAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "refundAllowed")
	return *ret0, err
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamSession) RefundAllowed() (bool, error) {
	return _Stream.Contract.RefundAllowed(&_Stream.CallOpts)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x3fa911ae.
//
// Solidity: function refundAllowed() constant returns(bool)
func (_Stream *StreamCallerSession) RefundAllowed() (bool, error) {
	return _Stream.Contract.RefundAllowed(&_Stream.CallOpts)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamCaller) Wattages(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stream.contract.Call(opts, out, "wattages", arg0, arg1)
	return *ret0, err
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) constant returns(uint256)
func (_Stream *StreamCallerSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactor) AddInputChunkId(opts *bind.TransactOpts, chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "addInputChunkId", chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactorSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamSession) Deposit() (*types.Transaction, error) {
	return _Stream.Contract.Deposit(&_Stream.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Stream *StreamTransactorSession) Deposit() (*types.Transaction, error) {
	return _Stream.Contract.Deposit(&_Stream.TransactOpts)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactor) EndStream(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "endStream")
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactorSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamSession) Refund() (*types.Transaction, error) {
	return _Stream.Contract.Refund(&_Stream.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Stream *StreamTransactorSession) Refund() (*types.Transaction, error) {
	return _Stream.Contract.Refund(&_Stream.TransactOpts)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ScrapProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "scrapProof", profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactor) SubmitProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "submitProof", profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactorSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ValidateProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "validateProof", profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// StreamAccountFundedIterator is returned from FilterAccountFunded and is used to iterate over the raw logs and unpacked data for AccountFunded events raised by the Stream contract.
type StreamAccountFundedIterator struct {
	Event *StreamAccountFunded // Event containing the contract specifics and raw log

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
func (it *StreamAccountFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamAccountFunded)
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
		it.Event = new(StreamAccountFunded)
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
func (it *StreamAccountFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamAccountFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamAccountFunded represents a AccountFunded event raised by the Stream contract.
type StreamAccountFunded struct {
	Account   common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountFunded is a free log retrieval operation binding the contract event 0xbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7.
//
// Solidity: event AccountFunded(address indexed account, uint256 weiAmount)
func (_Stream *StreamFilterer) FilterAccountFunded(opts *bind.FilterOpts, account []common.Address) (*StreamAccountFundedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "AccountFunded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StreamAccountFundedIterator{contract: _Stream.contract, event: "AccountFunded", logs: logs, sub: sub}, nil
}

// WatchAccountFunded is a free log subscription operation binding the contract event 0xbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7.
//
// Solidity: event AccountFunded(address indexed account, uint256 weiAmount)
func (_Stream *StreamFilterer) WatchAccountFunded(opts *bind.WatchOpts, sink chan<- *StreamAccountFunded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "AccountFunded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamAccountFunded)
				if err := _Stream.contract.UnpackLog(event, "AccountFunded", log); err != nil {
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

// ParseAccountFunded is a log parse operation binding the contract event 0xbccbe05a3719eacef984f404dd2adff555adcc05fb72fb8b309bafbd462cd6f7.
//
// Solidity: event AccountFunded(address indexed account, uint256 weiAmount)
func (_Stream *StreamFilterer) ParseAccountFunded(log types.Log) (*StreamAccountFunded, error) {
	event := new(StreamAccountFunded)
	if err := _Stream.contract.UnpackLog(event, "AccountFunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamChunkProofScrappedIterator is returned from FilterChunkProofScrapped and is used to iterate over the raw logs and unpacked data for ChunkProofScrapped events raised by the Stream contract.
type StreamChunkProofScrappedIterator struct {
	Event *StreamChunkProofScrapped // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofScrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofScrapped)
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
		it.Event = new(StreamChunkProofScrapped)
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
func (it *StreamChunkProofScrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofScrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofScrapped represents a ChunkProofScrapped event raised by the Stream contract.
type StreamChunkProofScrapped struct {
	Profile *big.Int
	ChunkId *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofScrapped is a free log retrieval operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofScrapped(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (*StreamChunkProofScrappedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofScrappedIterator{contract: _Stream.contract, event: "ChunkProofScrapped", logs: logs, sub: sub}, nil
}

// WatchChunkProofScrapped is a free log subscription operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofScrapped(opts *bind.WatchOpts, sink chan<- *StreamChunkProofScrapped, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofScrapped)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofScrapped", log); err != nil {
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

// ParseChunkProofScrapped is a log parse operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) ParseChunkProofScrapped(log types.Log) (*StreamChunkProofScrapped, error) {
	event := new(StreamChunkProofScrapped)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofScrapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamChunkProofSubmitedIterator is returned from FilterChunkProofSubmited and is used to iterate over the raw logs and unpacked data for ChunkProofSubmited events raised by the Stream contract.
type StreamChunkProofSubmitedIterator struct {
	Event *StreamChunkProofSubmited // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofSubmitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofSubmited)
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
		it.Event = new(StreamChunkProofSubmited)
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
func (it *StreamChunkProofSubmitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofSubmitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofSubmited represents a ChunkProofSubmited event raised by the Stream contract.
type StreamChunkProofSubmited struct {
	ChunkId *big.Int
	Profile *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofSubmited is a free log retrieval operation binding the contract event 0xe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa71.
//
// Solidity: event ChunkProofSubmited(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofSubmited(opts *bind.FilterOpts, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (*StreamChunkProofSubmitedIterator, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofSubmited", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofSubmitedIterator{contract: _Stream.contract, event: "ChunkProofSubmited", logs: logs, sub: sub}, nil
}

// WatchChunkProofSubmited is a free log subscription operation binding the contract event 0xe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa71.
//
// Solidity: event ChunkProofSubmited(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofSubmited(opts *bind.WatchOpts, sink chan<- *StreamChunkProofSubmited, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofSubmited", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofSubmited)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofSubmited", log); err != nil {
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

// ParseChunkProofSubmited is a log parse operation binding the contract event 0xe176ecde2cccccb4849f78f5d9f8378a179f72033c2fc649f90820a324cfaa71.
//
// Solidity: event ChunkProofSubmited(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) ParseChunkProofSubmited(log types.Log) (*StreamChunkProofSubmited, error) {
	event := new(StreamChunkProofSubmited)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofSubmited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamChunkProofValidatedIterator is returned from FilterChunkProofValidated and is used to iterate over the raw logs and unpacked data for ChunkProofValidated events raised by the Stream contract.
type StreamChunkProofValidatedIterator struct {
	Event *StreamChunkProofValidated // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofValidated)
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
		it.Event = new(StreamChunkProofValidated)
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
func (it *StreamChunkProofValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofValidated represents a ChunkProofValidated event raised by the Stream contract.
type StreamChunkProofValidated struct {
	Profile *big.Int
	ChunkId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofValidated is a free log retrieval operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) FilterChunkProofValidated(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int) (*StreamChunkProofValidatedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofValidatedIterator{contract: _Stream.contract, event: "ChunkProofValidated", logs: logs, sub: sub}, nil
}

// WatchChunkProofValidated is a free log subscription operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) WatchChunkProofValidated(opts *bind.WatchOpts, sink chan<- *StreamChunkProofValidated, profile []*big.Int, chunkId []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofValidated)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofValidated", log); err != nil {
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

// ParseChunkProofValidated is a log parse operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) ParseChunkProofValidated(log types.Log) (*StreamChunkProofValidated, error) {
	event := new(StreamChunkProofValidated)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofValidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Stream contract.
type StreamDepositedIterator struct {
	Event *StreamDeposited // Event containing the contract specifics and raw log

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
func (it *StreamDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamDeposited)
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
		it.Event = new(StreamDeposited)
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
func (it *StreamDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamDeposited represents a Deposited event raised by the Stream contract.
type StreamDeposited struct {
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea.
//
// Solidity: event Deposited(uint256 indexed weiAmount)
func (_Stream *StreamFilterer) FilterDeposited(opts *bind.FilterOpts, weiAmount []*big.Int) (*StreamDepositedIterator, error) {

	var weiAmountRule []interface{}
	for _, weiAmountItem := range weiAmount {
		weiAmountRule = append(weiAmountRule, weiAmountItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "Deposited", weiAmountRule)
	if err != nil {
		return nil, err
	}
	return &StreamDepositedIterator{contract: _Stream.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea.
//
// Solidity: event Deposited(uint256 indexed weiAmount)
func (_Stream *StreamFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StreamDeposited, weiAmount []*big.Int) (event.Subscription, error) {

	var weiAmountRule []interface{}
	for _, weiAmountItem := range weiAmount {
		weiAmountRule = append(weiAmountRule, weiAmountItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "Deposited", weiAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamDeposited)
				if err := _Stream.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2a89b2e3d580398d6dc2db5e0f336b52602bbaa51afa9bb5cdf59239cf0d2bea.
//
// Solidity: event Deposited(uint256 indexed weiAmount)
func (_Stream *StreamFilterer) ParseDeposited(log types.Log) (*StreamDeposited, error) {
	event := new(StreamDeposited)
	if err := _Stream.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamOutOfFundsIterator is returned from FilterOutOfFunds and is used to iterate over the raw logs and unpacked data for OutOfFunds events raised by the Stream contract.
type StreamOutOfFundsIterator struct {
	Event *StreamOutOfFunds // Event containing the contract specifics and raw log

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
func (it *StreamOutOfFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamOutOfFunds)
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
		it.Event = new(StreamOutOfFunds)
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
func (it *StreamOutOfFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamOutOfFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamOutOfFunds represents a OutOfFunds event raised by the Stream contract.
type StreamOutOfFunds struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOutOfFunds is a free log retrieval operation binding the contract event 0x5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d.
//
// Solidity: event OutOfFunds()
func (_Stream *StreamFilterer) FilterOutOfFunds(opts *bind.FilterOpts) (*StreamOutOfFundsIterator, error) {

	logs, sub, err := _Stream.contract.FilterLogs(opts, "OutOfFunds")
	if err != nil {
		return nil, err
	}
	return &StreamOutOfFundsIterator{contract: _Stream.contract, event: "OutOfFunds", logs: logs, sub: sub}, nil
}

// WatchOutOfFunds is a free log subscription operation binding the contract event 0x5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d.
//
// Solidity: event OutOfFunds()
func (_Stream *StreamFilterer) WatchOutOfFunds(opts *bind.WatchOpts, sink chan<- *StreamOutOfFunds) (event.Subscription, error) {

	logs, sub, err := _Stream.contract.WatchLogs(opts, "OutOfFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamOutOfFunds)
				if err := _Stream.contract.UnpackLog(event, "OutOfFunds", log); err != nil {
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

// ParseOutOfFunds is a log parse operation binding the contract event 0x5cc4d4e2397f909efdc05489f71c523c59913a6ffde292799faaca096635e16d.
//
// Solidity: event OutOfFunds()
func (_Stream *StreamFilterer) ParseOutOfFunds(log types.Log) (*StreamOutOfFunds, error) {
	event := new(StreamOutOfFunds)
	if err := _Stream.contract.UnpackLog(event, "OutOfFunds", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Stream contract.
type StreamRefundedIterator struct {
	Event *StreamRefunded // Event containing the contract specifics and raw log

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
func (it *StreamRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamRefunded)
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
		it.Event = new(StreamRefunded)
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
func (it *StreamRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamRefunded represents a Refunded event raised by the Stream contract.
type StreamRefunded struct {
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 weiAmount)
func (_Stream *StreamFilterer) FilterRefunded(opts *bind.FilterOpts) (*StreamRefundedIterator, error) {

	logs, sub, err := _Stream.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &StreamRefundedIterator{contract: _Stream.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 weiAmount)
func (_Stream *StreamFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *StreamRefunded) (event.Subscription, error) {

	logs, sub, err := _Stream.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamRefunded)
				if err := _Stream.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 weiAmount)
func (_Stream *StreamFilterer) ParseRefunded(log types.Log) (*StreamRefunded, error) {
	event := new(StreamRefunded)
	if err := _Stream.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamServiceFundedIterator is returned from FilterServiceFunded and is used to iterate over the raw logs and unpacked data for ServiceFunded events raised by the Stream contract.
type StreamServiceFundedIterator struct {
	Event *StreamServiceFunded // Event containing the contract specifics and raw log

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
func (it *StreamServiceFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamServiceFunded)
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
		it.Event = new(StreamServiceFunded)
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
func (it *StreamServiceFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamServiceFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamServiceFunded represents a ServiceFunded event raised by the Stream contract.
type StreamServiceFunded struct {
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceFunded is a free log retrieval operation binding the contract event 0x65d7a830175a146c5afc99c09332114c9fe32fc76e6646d79b1409350cade344.
//
// Solidity: event ServiceFunded(uint256 weiAmount)
func (_Stream *StreamFilterer) FilterServiceFunded(opts *bind.FilterOpts) (*StreamServiceFundedIterator, error) {

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ServiceFunded")
	if err != nil {
		return nil, err
	}
	return &StreamServiceFundedIterator{contract: _Stream.contract, event: "ServiceFunded", logs: logs, sub: sub}, nil
}

// WatchServiceFunded is a free log subscription operation binding the contract event 0x65d7a830175a146c5afc99c09332114c9fe32fc76e6646d79b1409350cade344.
//
// Solidity: event ServiceFunded(uint256 weiAmount)
func (_Stream *StreamFilterer) WatchServiceFunded(opts *bind.WatchOpts, sink chan<- *StreamServiceFunded) (event.Subscription, error) {

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ServiceFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamServiceFunded)
				if err := _Stream.contract.UnpackLog(event, "ServiceFunded", log); err != nil {
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

// ParseServiceFunded is a log parse operation binding the contract event 0x65d7a830175a146c5afc99c09332114c9fe32fc76e6646d79b1409350cade344.
//
// Solidity: event ServiceFunded(uint256 weiAmount)
func (_Stream *StreamFilterer) ParseServiceFunded(log types.Log) (*StreamServiceFunded, error) {
	event := new(StreamServiceFunded)
	if err := _Stream.contract.UnpackLog(event, "ServiceFunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StreamServiceShareCollectedIterator is returned from FilterServiceShareCollected and is used to iterate over the raw logs and unpacked data for ServiceShareCollected events raised by the Stream contract.
type StreamServiceShareCollectedIterator struct {
	Event *StreamServiceShareCollected // Event containing the contract specifics and raw log

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
func (it *StreamServiceShareCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamServiceShareCollected)
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
		it.Event = new(StreamServiceShareCollected)
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
func (it *StreamServiceShareCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamServiceShareCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamServiceShareCollected represents a ServiceShareCollected event raised by the Stream contract.
type StreamServiceShareCollected struct {
	Service  common.Address
	ChunckId *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterServiceShareCollected is a free log retrieval operation binding the contract event 0x3da5acdfe0e6a52b58d620773100ec8b0f21189f43733017f48bc6d57d950a33.
//
// Solidity: event ServiceShareCollected(address indexed service, uint256 indexed chunckId, uint256 amount)
func (_Stream *StreamFilterer) FilterServiceShareCollected(opts *bind.FilterOpts, service []common.Address, chunckId []*big.Int) (*StreamServiceShareCollectedIterator, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var chunckIdRule []interface{}
	for _, chunckIdItem := range chunckId {
		chunckIdRule = append(chunckIdRule, chunckIdItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ServiceShareCollected", serviceRule, chunckIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamServiceShareCollectedIterator{contract: _Stream.contract, event: "ServiceShareCollected", logs: logs, sub: sub}, nil
}

// WatchServiceShareCollected is a free log subscription operation binding the contract event 0x3da5acdfe0e6a52b58d620773100ec8b0f21189f43733017f48bc6d57d950a33.
//
// Solidity: event ServiceShareCollected(address indexed service, uint256 indexed chunckId, uint256 amount)
func (_Stream *StreamFilterer) WatchServiceShareCollected(opts *bind.WatchOpts, sink chan<- *StreamServiceShareCollected, service []common.Address, chunckId []*big.Int) (event.Subscription, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var chunckIdRule []interface{}
	for _, chunckIdItem := range chunckId {
		chunckIdRule = append(chunckIdRule, chunckIdItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ServiceShareCollected", serviceRule, chunckIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamServiceShareCollected)
				if err := _Stream.contract.UnpackLog(event, "ServiceShareCollected", log); err != nil {
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

// ParseServiceShareCollected is a log parse operation binding the contract event 0x3da5acdfe0e6a52b58d620773100ec8b0f21189f43733017f48bc6d57d950a33.
//
// Solidity: event ServiceShareCollected(address indexed service, uint256 indexed chunckId, uint256 amount)
func (_Stream *StreamFilterer) ParseServiceShareCollected(log types.Log) (*StreamServiceShareCollected, error) {
	event := new(StreamServiceShareCollected)
	if err := _Stream.contract.UnpackLog(event, "ServiceShareCollected", log); err != nil {
		return nil, err
	}
	return event, nil
}
