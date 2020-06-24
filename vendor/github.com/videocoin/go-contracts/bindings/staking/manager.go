// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingManagerUnbondingRequest is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerUnbondingRequest struct {
	Transcoder common.Address
	Timestamp  *big.Int
	Amount     *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transcoderApprovalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_slashPoolAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"TranscoderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"readiness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnbondingRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Unjailed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"delegateManaged\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"managed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"getSlashableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTrancoderSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTranscoderState\",\"outputs\":[{\"internalType\":\"enumStakingManager.TranscoderState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"getUnbondingRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.UnbondingRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"isManaged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWithdrawalsExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"registerTranscoder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbondingManaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setApprovalPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setSelfMinStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSlashFundAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setSlashRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"}],\"name\":\"setZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transcoderApprovalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transcoders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"effectiveMinSelfStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcodersArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transcodersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"unjail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAllPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
var StakingManagerBin = "0x60806040523480156200001157600080fd5b506040516200525e3803806200525e833981810160405262000037919081019062000435565b6000620000496200016860201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350856001819055508460028190555083600381905550826004819055508160058190555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200015c336200017060201b60201c565b505050505050620006a0565b600033905090565b620001806200022360201b60201c565b620001c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001b990620005d9565b60405180910390fd5b620001dd81600a6200028960201b620030851790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a60405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166200026d6200016860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6200029b82826200033c60201b60201c565b15620002de576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002d590620005b7565b60405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620003b0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003a790620005fb565b60405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60008151905062000418816200066c565b92915050565b6000815190506200042f8162000686565b92915050565b60008060008060008060c087890312156200044f57600080fd5b60006200045f89828a016200041e565b96505060206200047289828a016200041e565b95505060406200048589828a016200041e565b94505060606200049889828a016200041e565b9350506080620004ab89828a016200041e565b92505060a0620004be89828a0162000407565b9150509295509295509295565b6000620004da601f836200061d565b91507f526f6c65733a206163636f756e7420616c72656164792068617320726f6c65006000830152602082019050919050565b60006200051c6020836200061d565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006200055e6022836200061d565b91507f526f6c65733a206163636f756e7420697320746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006020820190508181036000830152620005d281620004cb565b9050919050565b60006020820190508181036000830152620005f4816200050d565b9050919050565b6000602082019050818103600083015262000616816200054f565b9050919050565b600082825260208201905092915050565b60006200063b8262000642565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b62000677816200062e565b81146200068357600080fd5b50565b620006918162000662565b81146200069d57600080fd5b50565b614bae80620006b06000396000f3fe6080604052600436106102465760003560e01c80637edbceb111610139578063c5f530af116100b6578063ec69a1bb1161007a578063ec69a1bb1461089f578063ec810a0a146108dc578063f2fde38b14610919578063f3ae241514610942578063fa7601541461097f578063fece707d146109aa57610246565b8063c5f530af1461078b578063c96be4cb146107b6578063e2dc17f6146107f3578063e341181d14610837578063e71824bc1461086257610246565b806396fb4a9c116100fd57806396fb4a9c146106a85780639e79b122146106d1578063a79e72631461070e578063ac18de4314610737578063c57cc2001461076057610246565b80637edbceb1146105bf5780638d23fc61146105d65780638da5cb5b146106155780638efc97a1146106405780638f32d59b1461067d57610246565b8063399f57c0116101c75780635afd2faa1161018b5780635afd2faa1461050d5780635c19a95c146105245780636cf6d675146105405780636db289091461056b578063715018a6146105a857610246565b8063399f57c01461042c5780634151766b14610455578063449ecfe61461047e5780635028e2e1146104a7578063503074ef146104e457610246565b8063254124fa1161020e578063254124fa1461034957806326e348ba146103865780632c9f0f2e146103af5780632d06177a146103d85780633939e6081461040157610246565b8063029859921461024b57806314bfb5271461027657806315620cce146102b35780631e7ff8f6146102cf578063220bb14e1461030c575b600080fd5b34801561025757600080fd5b506102606109e7565b60405161026d9190614941565b60405180910390f35b34801561028257600080fd5b5061029d60048036036102989190810190613d40565b6109ed565b6040516102aa919061462e565b60405180910390f35b6102cd60048036036102c89190810190613d92565b610b02565b005b3480156102db57600080fd5b506102f660048036036102f19190810190613d40565b610bb9565b6040516103039190614941565b60405180910390f35b34801561031857600080fd5b50610333600480360361032e9190810190613d40565b610c74565b604051610340919061462e565b60405180910390f35b34801561035557600080fd5b50610370600480360361036b9190810190613dce565b610cd2565b60405161037d9190614941565b60405180910390f35b34801561039257600080fd5b506103ad60048036036103a89190810190613e59565b610dc5565b005b3480156103bb57600080fd5b506103d660048036036103d19190810190613e59565b610e23565b005b3480156103e457600080fd5b506103ff60048036036103fa9190810190613d40565b610e74565b005b34801561040d57600080fd5b50610416610f15565b6040516104239190614941565b60405180910390f35b34801561043857600080fd5b50610453600480360361044e9190810190613e59565b610f1b565b005b34801561046157600080fd5b5061047c60048036036104779190810190613e59565b6110b8565b005b34801561048a57600080fd5b506104a560048036036104a09190810190613d40565b611109565b005b3480156104b357600080fd5b506104ce60048036036104c99190810190613d92565b6112ae565b6040516104db9190614941565b60405180910390f35b3480156104f057600080fd5b5061050b60048036036105069190810190613e1d565b611438565b005b34801561051957600080fd5b50610522611517565b005b61053e60048036036105399190810190613d40565b611608565b005b34801561054c57600080fd5b506105556116ab565b6040516105629190614941565b60405180910390f35b34801561057757600080fd5b50610592600480360361058d9190810190613e1d565b6116b1565b60405161059f9190614941565b60405180910390f35b3480156105b457600080fd5b506105bd61175b565b005b3480156105cb57600080fd5b506105d4611861565b005b3480156105e257600080fd5b506105fd60048036036105f89190810190613d40565b61195e565b60405161060c93929190614985565b60405180910390f35b34801561062157600080fd5b5061062a611995565b6040516106379190614613565b60405180910390f35b34801561064c57600080fd5b5061066760048036036106629190810190613d92565b6119be565b6040516106749190614941565b60405180910390f35b34801561068957600080fd5b50610692611b83565b60405161069f919061462e565b60405180910390f35b3480156106b457600080fd5b506106cf60048036036106ca9190810190613e1d565b611be1565b005b3480156106dd57600080fd5b506106f860048036036106f39190810190613e1d565b611cc0565b6040516107059190614926565b60405180910390f35b34801561071a57600080fd5b5061073560048036036107309190810190613d69565b611d9e565b005b34801561074357600080fd5b5061075e60048036036107599190810190613d40565b611eba565b005b34801561076c57600080fd5b50610775611f5b565b604051610782919061462e565b60405180910390f35b34801561079757600080fd5b506107a0612001565b6040516107ad9190614941565b60405180910390f35b3480156107c257600080fd5b506107dd60048036036107d89190810190613d40565b612007565b6040516107ea919061462e565b60405180910390f35b3480156107ff57600080fd5b5061081a60048036036108159190810190613d40565b6122b0565b60405161082e9897969594939291906149bc565b60405180910390f35b34801561084357600080fd5b5061084c612305565b6040516108599190614941565b60405180910390f35b34801561086e57600080fd5b5061088960048036036108849190810190613d40565b61230b565b6040516108969190614649565b60405180910390f35b3480156108ab57600080fd5b506108c660048036036108c19190810190613d40565b6125e3565b6040516108d39190614941565b60405180910390f35b3480156108e857600080fd5b5061090360048036036108fe9190810190613e59565b6126ee565b6040516109109190614613565b60405180910390f35b34801561092557600080fd5b50610940600480360361093b9190810190613d40565b61272a565b005b34801561094e57600080fd5b5061096960048036036109649190810190613d40565b61277d565b604051610976919061462e565b60405180910390f35b34801561098b57600080fd5b5061099461279a565b6040516109a19190614941565b60405180910390f35b3480156109b657600080fd5b506109d160048036036109cc9190810190613d40565b6127a7565b6040516109de9190614941565b60405180910390f35b60015481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5590614866565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411610ae8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adf90614746565b60405180910390fd5b8060080160009054906101000a900460ff16915050919050565b610b0b3361277d565b610b4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4190614906565b60405180910390fd5b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160050160006101000a81548160ff021916908315150217905550610bb483836127ba565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2190614866565b60405180910390fd5b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff16915050919050565b6000610cdd3361277d565b610d1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1390614906565b60405180910390fd5b6000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff16610db0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da790614826565b60405180910390fd5b610dbb858585612b1e565b9150509392505050565b610dcd611b83565b610e0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e03906147c6565b60405180910390fd5b60008111610e1957600080fd5b8060028190555050565b610e2b611b83565b610e6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e61906147c6565b60405180910390fd5b8060038190555050565b610e7c611b83565b610ebb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb2906147c6565b60405180910390fd5b610ecf81600a61308590919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a60405160405180910390a250565b60035481565b60648110610f5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5590614846565b60405180910390fd5b60003390506000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015414610fed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe490614726565b60405180910390fd5b428160010181905550828160020181905550600254816009018190555060098290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508173ffffffffffffffffffffffffffffffffffffffff167f6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b60405160405180910390a2505050565b6110c0611b83565b6110ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f6906147c6565b60405180910390fd5b8060058190555050565b611111611b83565b611150576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611147906147c6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156111c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b790614866565b60405180910390fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161124a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124190614886565b60405180910390fd5b60008160080160006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167ffa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff60405160405180910390a25050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561131f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131690614866565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561138f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138690614866565b60405180910390fd5b600061139b84846119be565b905061142f81600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461312d90919063ffffffff16565b91505092915050565b611440611b83565b61147f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611476906147c6565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411611509576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150090614746565b60405180910390fd5b818160050181905550505050565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600401548160030154106115a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161159b906148a6565b60405180910390fd5b6000816003015490505b81600401548110156116035760006115c68233613177565b9050806115d557505050611606565b6115ed600184600301546134c590919063ffffffff16565b83600301819055505080806001019150506115ae565b50505b565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff161561169d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169490614786565b60405180910390fd5b6116a782336127ba565b5050565b60045481565b600080600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff1615611747576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173e90614786565b60405180910390fd5b611752843385612b1e565b91505092915050565b611763611b83565b6117a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611799906147c6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600401548160030154106118ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e5906148a6565b60405180910390fd5b6118fc816003015433613177565b61193b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611932906146a6565b60405180910390fd5b611953600182600301546134c590919063ffffffff16565b816003018190555050565b60086020528060005260406000206000915090508060030154908060040154908060050160009054906101000a900460ff16905083565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000836006018054905090506000809050818310611ab0578095505050505050611b7d565b60008390505b82811015611b73576000866006018281548110611acf57fe5b90600052602060002090600202016001015490506000611b39828860000160008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461351a90919063ffffffff16565b9050611b4f60648261358a90919063ffffffff16565b9050611b6481856134c590919063ffffffff16565b93505050806001019050611ab6565b5080955050505050505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611bc56135d4565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b611be9611b83565b611c28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c1f906147c6565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411611cb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ca990614746565b60405180910390fd5b818160040181905550505050565b611cc8613cca565b6000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152505091505092915050565b611da6611b83565b611de5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ddc906147c6565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611e76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e6d906148e6565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611ec2611b83565b611f01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ef8906147c6565b60405180910390fd5b611f1581600a6135dc90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd3160405160405180910390a250565b600080600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816003015490505b8160040154811015611ff757600082600201600083815260200190815260200160002090506004548160010154420310611fe95760019350505050611ffe565b508080600101915050611fa9565b5060009150505b90565b60025481565b6000612011611b83565b612050576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612047906147c6565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612103576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120fa90614866565b60405180910390fd5b600081600101541161214a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161214190614886565b60405180910390fd5b60006121558461230b565b90506001600481111561216457fe5b81600481111561217057fe5b1415801561219557506003600481111561218657fe5b81600481111561219257fe5b14155b156121a5576000925050506122ab565b60006121c0600554846000015461351a90919063ffffffff16565b90506121d660648261358a90919063ffffffff16565b90506121ef81846000015461312d90919063ffffffff16565b836000018190555082600601604051806040016040528042815260200160055481525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505061225d85613683565b6005548573ffffffffffffffffffffffffffffffffffffffff167f4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd60405160405180910390a3600193505050505b919050565b60076020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154908060050154908060080160009054906101000a900460ff16908060090154905088565b60055481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561237c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161237390614866565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008260010154141561241a576004925050506125de565b8160080160009054906101000a900460ff161561243c576002925050506125de565b60035482600101544203106125d75781600901548160000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054106124a2576001925050506125de565b60008090506000826003015490505b826004015481101561256c57600083600201600083815260200190815260200160002090508673ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612533575061255f565b60045481600101544203101561255d5761255a8160020154846134c590919063ffffffff16565b92505b505b80806001019150506124b1565b506002546125c48360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836134c590919063ffffffff16565b106125d557600393505050506125de565b505b6000925050505b919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612654576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161264b90614866565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116126de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126d590614746565b60405180910390fd5b8060060180549050915050919050565b600981815481106126fb57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612732611b83565b612771576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612768906147c6565b60405180910390fd5b61277a81613828565b50565b600061279382600a61395690919063ffffffff16565b9050919050565b6000600980549050905090565b60006127b382836112ae565b9050919050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000349050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614156128b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128ac906148c6565b60405180910390fd5b6001548110156128fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128f190614806565b60405180910390fd5b60008260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156129fa57826007018490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600601805490508260010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b612a048585613a1e565b612a1b8184600001546134c590919063ffffffff16565b8360000181905550612a77818360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546134c590919063ffffffff16565b8260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b60405160405180910390a45050505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415612b8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b86906148c6565b60405180910390fd5b6000600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050612c1f8686613a1e565b8060000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054841115612ca3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9a90614706565b60405180910390fd5b6000612cae8761230b565b9050612d04858360000160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461312d90919063ffffffff16565b8260000160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612d6085846000015461312d90919063ffffffff16565b8360000181905550600082600401549050612d89600184600401546134c590919063ffffffff16565b836004018190555060006004811115612d9e57fe5b826004811115612daa57fe5b1480612dcc575060026004811115612dbe57fe5b826004811115612dca57fe5b145b80612e0357508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b15612f66578773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5428a604051612e6892919061495c565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff168152602001600454420381526020018781525083600201600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050612f228188613177565b612f61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f58906146a6565b60405180910390fd5b613077565b8773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da560045442018a604051612fca92919061495c565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020014281526020018781525083600201600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050505b809450505050509392505050565b61308f8282613956565b156130cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130c690614686565b60405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600061316f83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613c0e565b905092915050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600201600086815260200190815260200160002090506000816002015414156131ec576000925050506134bf565b600454816001015442031015613207576000925050506134bf565b6000600760008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600082600101549050600061328e600454836134c590919063ffffffff16565b9050600080905060008090505b84600601805490508110156133655760008560060182815481106132bb57fe5b90600052602060002090600202019050600081600001549050858110806132e157508481115b156132ed57505061335a565b600061330a83600101548a6002015461351a90919063ffffffff16565b905061332060648261358a90919063ffffffff16565b9050613339818a6002015461312d90919063ffffffff16565b896002018190555061335481866134c590919063ffffffff16565b94505050505b80600101905061329b565b50600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156133ce573d6000803e3d6000fd5b50600085600201549050600086600201819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015613428573d6000803e3d6000fd5b508560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b7f544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84846040516134ab9190614941565b60405180910390a460019750505050505050505b92915050565b600080828401905083811015613510576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613507906146e6565b60405180910390fd5b8091505092915050565b60008083141561352d5760009050613584565b600082840290508284828161353e57fe5b041461357f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613576906147a6565b60405180910390fd5b809150505b92915050565b60006135cc83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613c69565b905092915050565b600033905090565b6135e68282613956565b613625576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161361c90614766565b60405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61368b611b83565b6136ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136c1906147c6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561373a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161373190614866565b60405180910390fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116137c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137bb90614886565b60405180910390fd5b60018160080160006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e760405160405180910390a25050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415613898576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161388f906146c6565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156139c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139be906147e6565b60405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000826006018054905090506000613abc86866119be565b9050613b12818460000160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461312d90919063ffffffff16565b8360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550818360010160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015613c05573d6000803e3d6000fd5b50505050505050565b6000838311158290613c56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c4d9190614664565b60405180910390fd5b5060008385039050809150509392505050565b60008083118290613cb0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ca79190614664565b60405180910390fd5b506000838581613cbc57fe5b049050809150509392505050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b600081359050613d1081614b26565b92915050565b600081359050613d2581614b3d565b92915050565b600081359050613d3a81614b54565b92915050565b600060208284031215613d5257600080fd5b6000613d6084828501613d01565b91505092915050565b600060208284031215613d7b57600080fd5b6000613d8984828501613d16565b91505092915050565b60008060408385031215613da557600080fd5b6000613db385828601613d01565b9250506020613dc485828601613d01565b9150509250929050565b600080600060608486031215613de357600080fd5b6000613df186828701613d01565b9350506020613e0286828701613d01565b9250506040613e1386828701613d2b565b9150509250925092565b60008060408385031215613e3057600080fd5b6000613e3e85828601613d01565b9250506020613e4f85828601613d2b565b9150509250929050565b600060208284031215613e6b57600080fd5b6000613e7984828501613d2b565b91505092915050565b613e8b81614a56565b82525050565b613e9a81614a56565b82525050565b613ea981614a7a565b82525050565b613eb881614ac3565b82525050565b6000613ec982614a3a565b613ed38185614a45565b9350613ee3818560208601614ad5565b613eec81614b08565b840191505092915050565b6000613f04601f83614a45565b91507f526f6c65733a206163636f756e7420616c72656164792068617320726f6c65006000830152602082019050919050565b6000613f44601883614a45565b91507f6661696c656420746f207769746864726177207374616b6500000000000000006000830152602082019050919050565b6000613f84602683614a45565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000613fea601b83614a45565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b600061402a601083614a45565b91507f4e6f7420656e6f7567682066756e6473000000000000000000000000000000006000830152602082019050919050565b600061406a601d83614a45565b91507f5472616e73636f64657220616c726561647920726567697374657265640000006000830152602082019050919050565b60006140aa601983614a45565b91507f5472616e73636f646572206e6f742072656769737465726564000000000000006000830152602082019050919050565b60006140ea602183614a45565b91507f526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c60008301527f65000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614150604283614a45565b91507f74686973206d6574686f642063616e277420626520757365642062792064656c60008301527f656761746f722074686174206465706f736974656420455243323020746f6b6560208301527f6e730000000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b60006141dc602183614a45565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614242602083614a45565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000614282602283614a45565b91507f526f6c65733a206163636f756e7420697320746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006142e8602383614a45565b91507f4d757374206465706f736974206174206c65617374206d696e696d756d20766160008301527f6c756500000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061434e604b83614a45565b91507f74686973206d6574686f642063616e206f6e6c792062652075736564206f6e6c60008301527f7920666f722064656c656761746f722074686174206465706f7369746564204560208301527f5243323020746f6b656e730000000000000000000000000000000000000000006040830152606082019050919050565b60006143da602b83614a45565b91507f52617465206d75737420626520612070657263656e746167652062657477656560008301527f6e203020616e64203130300000000000000000000000000000000000000000006020830152604082019050919050565b6000614440601683614a45565b91507f63616e277420757365207a65726f2061646472657373000000000000000000006000830152602082019050919050565b6000614480601a83614a45565b91507f52656769737465726564207472616e73636f646572206f6e6c790000000000006000830152602082019050919050565b60006144c0601383614a45565b91507f6e6f2070656e64696e67207265717565737473000000000000000000000000006000830152602082019050919050565b6000614500601583614a45565b91507f43616e60742075736520616464726573732030783000000000000000000000006000830152602082019050919050565b6000614540601b83614a45565b91507f416c72656164792073657420746f2074686973206164647265737300000000006000830152602082019050919050565b6000614580600d83614a45565b91507f6e6f742061206d616e61676572000000000000000000000000000000000000006000830152602082019050919050565b6060820160008201516145c96000850182613e82565b5060208201516145dc60208501826145f5565b5060408201516145ef60408501826145f5565b50505050565b6145fe81614ab9565b82525050565b61460d81614ab9565b82525050565b60006020820190506146286000830184613e91565b92915050565b60006020820190506146436000830184613ea0565b92915050565b600060208201905061465e6000830184613eaf565b92915050565b6000602082019050818103600083015261467e8184613ebe565b905092915050565b6000602082019050818103600083015261469f81613ef7565b9050919050565b600060208201905081810360008301526146bf81613f37565b9050919050565b600060208201905081810360008301526146df81613f77565b9050919050565b600060208201905081810360008301526146ff81613fdd565b9050919050565b6000602082019050818103600083015261471f8161401d565b9050919050565b6000602082019050818103600083015261473f8161405d565b9050919050565b6000602082019050818103600083015261475f8161409d565b9050919050565b6000602082019050818103600083015261477f816140dd565b9050919050565b6000602082019050818103600083015261479f81614143565b9050919050565b600060208201905081810360008301526147bf816141cf565b9050919050565b600060208201905081810360008301526147df81614235565b9050919050565b600060208201905081810360008301526147ff81614275565b9050919050565b6000602082019050818103600083015261481f816142db565b9050919050565b6000602082019050818103600083015261483f81614341565b9050919050565b6000602082019050818103600083015261485f816143cd565b9050919050565b6000602082019050818103600083015261487f81614433565b9050919050565b6000602082019050818103600083015261489f81614473565b9050919050565b600060208201905081810360008301526148bf816144b3565b9050919050565b600060208201905081810360008301526148df816144f3565b9050919050565b600060208201905081810360008301526148ff81614533565b9050919050565b6000602082019050818103600083015261491f81614573565b9050919050565b600060608201905061493b60008301846145b3565b92915050565b60006020820190506149566000830184614604565b92915050565b60006040820190506149716000830185614604565b61497e6020830184614604565b9392505050565b600060608201905061499a6000830186614604565b6149a76020830185614604565b6149b46040830184613ea0565b949350505050565b6000610100820190506149d2600083018b614604565b6149df602083018a614604565b6149ec6040830189614604565b6149f96060830188614604565b614a066080830187614604565b614a1360a0830186614604565b614a2060c0830185613ea0565b614a2d60e0830184614604565b9998505050505050505050565b600081519050919050565b600082825260208201905092915050565b6000614a6182614a99565b9050919050565b6000614a7382614a99565b9050919050565b60008115159050919050565b6000819050614a9482614b19565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000614ace82614a86565b9050919050565b60005b83811015614af3578082015181840152602081019050614ad8565b83811115614b02576000848401525b50505050565b6000601f19601f8301169050919050565b60058110614b2357fe5b50565b614b2f81614a56565b8114614b3a57600080fd5b50565b614b4681614a68565b8114614b5157600080fd5b50565b614b5d81614ab9565b8114614b6857600080fd5b5056fea365627a7a72315820b791bb30559d501d71ad38fab89be89f4169927a56a97c531476081e83e9dc516c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, _minDelegation *big.Int, _minSelfStake *big.Int, _transcoderApprovalPeriod *big.Int, _unbondingPeriod *big.Int, _slashRate *big.Int, _slashPoolAddress common.Address) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingManagerBin), backend, _minDelegation, _minSelfStake, _transcoderApprovalPeriod, _unbondingPeriod, _slashRate, _slashPoolAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCaller) Delegators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	ret := new(struct {
		Pending *big.Int
		Next    *big.Int
		Managed bool
	})
	out := ret
	err := _StakingManager.contract.Call(opts, out, "delegators", arg0)
	return *ret, err
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCallerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetDelegatorStake(opts *bind.CallOpts, transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getDelegatorStake", transcoderAddr, delegAddr)
	return *ret0, err
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSelfStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSelfStake", _addr)
	return *ret0, err
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSlashableAmount(opts *bind.CallOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSlashableAmount", transcoderAddr, delegatorAddr)
	return *ret0, err
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTotalStake", _addr)
	return *ret0, err
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTrancoderSlashes(opts *bind.CallOpts, transcoderAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTrancoderSlashes", transcoderAddr)
	return *ret0, err
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCaller) GetTranscoderState(opts *bind.CallOpts, transcoderAddr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTranscoderState", transcoderAddr)
	return *ret0, err
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCallerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCaller) GetUnbondingRequest(opts *bind.CallOpts, delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	var (
		ret0 = new(StakingManagerUnbondingRequest)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getUnbondingRequest", delegatorAddr, unbondingID)
	return *ret0, err
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCallerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsJailed(opts *bind.CallOpts, transcoderAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isJailed", transcoderAddr)
	return *ret0, err
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManaged(opts *bind.CallOpts, delegatorAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isManaged", delegatorAddr)
	return *ret0, err
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManager(opts *bind.CallOpts, v common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isManager", v)
	return *ret0, err
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minSelfStake")
	return *ret0, err
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCaller) PendingWithdrawalsExist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "pendingWithdrawalsExist")
	return *ret0, err
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCaller) SlashRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "slashRate")
	return *ret0, err
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscoderApprovalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcoderApprovalPeriod")
	return *ret0, err
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCaller) Transcoders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	ret := new(struct {
		Total                 *big.Int
		Timestamp             *big.Int
		RewardRate            *big.Int
		Rewards               *big.Int
		Zone                  *big.Int
		Capacity              *big.Int
		Jailed                bool
		EffectiveMinSelfStake *big.Int
	})
	out := ret
	err := _StakingManager.contract.Call(opts, out, "transcoders", arg0)
	return *ret, err
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCallerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCaller) TranscodersArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcodersArray", arg0)
	return *ret0, err
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCallerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscodersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcodersCount")
	return *ret0, err
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) AddManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "addManager", v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerTransactor) Delegate(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegate", transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerTransactorSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerTransactor) DelegateManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegateManaged", transcoderAddr, delegatorAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerTransactorSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) RegisterTranscoder(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerTranscoder", rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) RemoveManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "removeManager", v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbonding(opts *bind.TransactOpts, transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbonding", transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbondingManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbondingManaged", transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactor) SetApprovalPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setApprovalPeriod", period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactorSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactor) SetCapacity(opts *bind.TransactOpts, addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setCapacity", addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactorSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) SetSelfMinStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSelfMinStake", amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashFundAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashFundAddress", addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashRate", rate)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) SetSlashRate(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashRate(&_StakingManager.TransactOpts, rate)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashRate(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashRate(&_StakingManager.TransactOpts, rate)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactor) SetZone(opts *bind.TransactOpts, addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setZone", addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactorSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactor) Slash(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "slash", addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactorSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) Unjail(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unjail", transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawAllPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawAllPending")
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawPending")
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// StakingManagerDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the StakingManager contract.
type StakingManagerDelegatedIterator struct {
	Event *StakingManagerDelegated // Event containing the contract specifics and raw log

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
func (it *StakingManagerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerDelegated)
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
		it.Event = new(StakingManagerDelegated)
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
func (it *StakingManagerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerDelegated represents a Delegated event raised by the StakingManager contract.
type StakingManagerDelegated struct {
	Transcoder common.Address
	Delegator  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) FilterDelegated(opts *bind.FilterOpts, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (*StakingManagerDelegatedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerDelegatedIterator{contract: _StakingManager.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakingManagerDelegated, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerDelegated)
				if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) ParseDelegated(log types.Log) (*StakingManagerDelegated, error) {
	event := new(StakingManagerDelegated)
	if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the StakingManager contract.
type StakingManagerJailedIterator struct {
	Event *StakingManagerJailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerJailed)
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
		it.Event = new(StakingManagerJailed)
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
func (it *StakingManagerJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerJailed represents a Jailed event raised by the StakingManager contract.
type StakingManagerJailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterJailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerJailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerJailedIterator{contract: _StakingManager.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakingManagerJailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerJailed)
				if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
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

// ParseJailed is a log parse operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseJailed(log types.Log) (*StakingManagerJailed, error) {
	event := new(StakingManagerJailed)
	if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the StakingManager contract.
type StakingManagerManagerAddedIterator struct {
	Event *StakingManagerManagerAdded // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerAdded)
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
		it.Event = new(StakingManagerManagerAdded)
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
func (it *StakingManagerManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerAdded represents a ManagerAdded event raised by the StakingManager contract.
type StakingManagerManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerAddedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerAddedIterator{contract: _StakingManager.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerAdded, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerAdded)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerAdded(log types.Log) (*StakingManagerManagerAdded, error) {
	event := new(StakingManagerManagerAdded)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the StakingManager contract.
type StakingManagerManagerRemovedIterator struct {
	Event *StakingManagerManagerRemoved // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerRemoved)
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
		it.Event = new(StakingManagerManagerRemoved)
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
func (it *StakingManagerManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerRemoved represents a ManagerRemoved event raised by the StakingManager contract.
type StakingManagerManagerRemoved struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerRemovedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerRemovedIterator{contract: _StakingManager.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerRemoved, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerRemoved)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerRemoved(log types.Log) (*StakingManagerManagerRemoved, error) {
	event := new(StakingManagerManagerRemoved)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the StakingManager contract.
type StakingManagerSlashedIterator struct {
	Event *StakingManagerSlashed // Event containing the contract specifics and raw log

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
func (it *StakingManagerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerSlashed)
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
		it.Event = new(StakingManagerSlashed)
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
func (it *StakingManagerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerSlashed represents a Slashed event raised by the StakingManager contract.
type StakingManagerSlashed struct {
	Transcoder common.Address
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) FilterSlashed(opts *bind.FilterOpts, transcoder []common.Address, rate []*big.Int) (*StakingManagerSlashedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerSlashedIterator{contract: _StakingManager.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingManagerSlashed, transcoder []common.Address, rate []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerSlashed)
				if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) ParseSlashed(log types.Log) (*StakingManagerSlashed, error) {
	event := new(StakingManagerSlashed)
	if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerStakeWithdrawalIterator is returned from FilterStakeWithdrawal and is used to iterate over the raw logs and unpacked data for StakeWithdrawal events raised by the StakingManager contract.
type StakingManagerStakeWithdrawalIterator struct {
	Event *StakingManagerStakeWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeWithdrawal)
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
		it.Event = new(StakingManagerStakeWithdrawal)
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
func (it *StakingManagerStakeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeWithdrawal represents a StakeWithdrawal event raised by the StakingManager contract.
type StakingManagerStakeWithdrawal struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawal is a free log retrieval operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeWithdrawal(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerStakeWithdrawalIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeWithdrawalIterator{contract: _StakingManager.contract, event: "StakeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawal is a free log subscription operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeWithdrawal, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeWithdrawal)
				if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
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

// ParseStakeWithdrawal is a log parse operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeWithdrawal(log types.Log) (*StakingManagerStakeWithdrawal, error) {
	event := new(StakingManagerStakeWithdrawal)
	if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerTranscoderRegisteredIterator is returned from FilterTranscoderRegistered and is used to iterate over the raw logs and unpacked data for TranscoderRegistered events raised by the StakingManager contract.
type StakingManagerTranscoderRegisteredIterator struct {
	Event *StakingManagerTranscoderRegistered // Event containing the contract specifics and raw log

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
func (it *StakingManagerTranscoderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerTranscoderRegistered)
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
		it.Event = new(StakingManagerTranscoderRegistered)
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
func (it *StakingManagerTranscoderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerTranscoderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerTranscoderRegistered represents a TranscoderRegistered event raised by the StakingManager contract.
type StakingManagerTranscoderRegistered struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTranscoderRegistered is a free log retrieval operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterTranscoderRegistered(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerTranscoderRegisteredIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTranscoderRegisteredIterator{contract: _StakingManager.contract, event: "TranscoderRegistered", logs: logs, sub: sub}, nil
}

// WatchTranscoderRegistered is a free log subscription operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchTranscoderRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerTranscoderRegistered, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerTranscoderRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
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

// ParseTranscoderRegistered is a log parse operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseTranscoderRegistered(log types.Log) (*StakingManagerTranscoderRegistered, error) {
	event := new(StakingManagerTranscoderRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnbondingRequestedIterator is returned from FilterUnbondingRequested and is used to iterate over the raw logs and unpacked data for UnbondingRequested events raised by the StakingManager contract.
type StakingManagerUnbondingRequestedIterator struct {
	Event *StakingManagerUnbondingRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnbondingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnbondingRequested)
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
		it.Event = new(StakingManagerUnbondingRequested)
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
func (it *StakingManagerUnbondingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnbondingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnbondingRequested represents a UnbondingRequested event raised by the StakingManager contract.
type StakingManagerUnbondingRequested struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Readiness   *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondingRequested is a free log retrieval operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterUnbondingRequested(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerUnbondingRequestedIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnbondingRequestedIterator{contract: _StakingManager.contract, event: "UnbondingRequested", logs: logs, sub: sub}, nil
}

// WatchUnbondingRequested is a free log subscription operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchUnbondingRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnbondingRequested, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnbondingRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
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

// ParseUnbondingRequested is a log parse operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseUnbondingRequested(log types.Log) (*StakingManagerUnbondingRequested, error) {
	event := new(StakingManagerUnbondingRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnjailedIterator is returned from FilterUnjailed and is used to iterate over the raw logs and unpacked data for Unjailed events raised by the StakingManager contract.
type StakingManagerUnjailedIterator struct {
	Event *StakingManagerUnjailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnjailed)
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
		it.Event = new(StakingManagerUnjailed)
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
func (it *StakingManagerUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnjailed represents a Unjailed event raised by the StakingManager contract.
type StakingManagerUnjailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnjailed is a free log retrieval operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterUnjailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerUnjailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnjailedIterator{contract: _StakingManager.contract, event: "Unjailed", logs: logs, sub: sub}, nil
}

// WatchUnjailed is a free log subscription operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchUnjailed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnjailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnjailed)
				if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
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

// ParseUnjailed is a log parse operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseUnjailed(log types.Log) (*StakingManagerUnjailed, error) {
	event := new(StakingManagerUnjailed)
	if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
		return nil, err
	}
	return event, nil
}
