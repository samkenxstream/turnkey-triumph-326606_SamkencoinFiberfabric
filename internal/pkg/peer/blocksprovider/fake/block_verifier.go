// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	commona "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/internal/pkg/peer/blocksprovider"
)

type BlockVerifier struct {
	VerifyBlockStub        func(common.ChannelID, uint64, *commona.Block) error
	verifyBlockMutex       sync.RWMutex
	verifyBlockArgsForCall []struct {
		arg1 common.ChannelID
		arg2 uint64
		arg3 *commona.Block
	}
	verifyBlockReturns struct {
		result1 error
	}
	verifyBlockReturnsOnCall map[int]struct {
		result1 error
	}
	VerifyBlockAttestationStub        func(string, *commona.Block) error
	verifyBlockAttestationMutex       sync.RWMutex
	verifyBlockAttestationArgsForCall []struct {
		arg1 string
		arg2 *commona.Block
	}
	verifyBlockAttestationReturns struct {
		result1 error
	}
	verifyBlockAttestationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockVerifier) VerifyBlock(arg1 common.ChannelID, arg2 uint64, arg3 *commona.Block) error {
	fake.verifyBlockMutex.Lock()
	ret, specificReturn := fake.verifyBlockReturnsOnCall[len(fake.verifyBlockArgsForCall)]
	fake.verifyBlockArgsForCall = append(fake.verifyBlockArgsForCall, struct {
		arg1 common.ChannelID
		arg2 uint64
		arg3 *commona.Block
	}{arg1, arg2, arg3})
	stub := fake.VerifyBlockStub
	fakeReturns := fake.verifyBlockReturns
	fake.recordInvocation("VerifyBlock", []interface{}{arg1, arg2, arg3})
	fake.verifyBlockMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BlockVerifier) VerifyBlockCallCount() int {
	fake.verifyBlockMutex.RLock()
	defer fake.verifyBlockMutex.RUnlock()
	return len(fake.verifyBlockArgsForCall)
}

func (fake *BlockVerifier) VerifyBlockCalls(stub func(common.ChannelID, uint64, *commona.Block) error) {
	fake.verifyBlockMutex.Lock()
	defer fake.verifyBlockMutex.Unlock()
	fake.VerifyBlockStub = stub
}

func (fake *BlockVerifier) VerifyBlockArgsForCall(i int) (common.ChannelID, uint64, *commona.Block) {
	fake.verifyBlockMutex.RLock()
	defer fake.verifyBlockMutex.RUnlock()
	argsForCall := fake.verifyBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *BlockVerifier) VerifyBlockReturns(result1 error) {
	fake.verifyBlockMutex.Lock()
	defer fake.verifyBlockMutex.Unlock()
	fake.VerifyBlockStub = nil
	fake.verifyBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *BlockVerifier) VerifyBlockReturnsOnCall(i int, result1 error) {
	fake.verifyBlockMutex.Lock()
	defer fake.verifyBlockMutex.Unlock()
	fake.VerifyBlockStub = nil
	if fake.verifyBlockReturnsOnCall == nil {
		fake.verifyBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BlockVerifier) VerifyBlockAttestation(arg1 string, arg2 *commona.Block) error {
	fake.verifyBlockAttestationMutex.Lock()
	ret, specificReturn := fake.verifyBlockAttestationReturnsOnCall[len(fake.verifyBlockAttestationArgsForCall)]
	fake.verifyBlockAttestationArgsForCall = append(fake.verifyBlockAttestationArgsForCall, struct {
		arg1 string
		arg2 *commona.Block
	}{arg1, arg2})
	stub := fake.VerifyBlockAttestationStub
	fakeReturns := fake.verifyBlockAttestationReturns
	fake.recordInvocation("VerifyBlockAttestation", []interface{}{arg1, arg2})
	fake.verifyBlockAttestationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BlockVerifier) VerifyBlockAttestationCallCount() int {
	fake.verifyBlockAttestationMutex.RLock()
	defer fake.verifyBlockAttestationMutex.RUnlock()
	return len(fake.verifyBlockAttestationArgsForCall)
}

func (fake *BlockVerifier) VerifyBlockAttestationCalls(stub func(string, *commona.Block) error) {
	fake.verifyBlockAttestationMutex.Lock()
	defer fake.verifyBlockAttestationMutex.Unlock()
	fake.VerifyBlockAttestationStub = stub
}

func (fake *BlockVerifier) VerifyBlockAttestationArgsForCall(i int) (string, *commona.Block) {
	fake.verifyBlockAttestationMutex.RLock()
	defer fake.verifyBlockAttestationMutex.RUnlock()
	argsForCall := fake.verifyBlockAttestationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *BlockVerifier) VerifyBlockAttestationReturns(result1 error) {
	fake.verifyBlockAttestationMutex.Lock()
	defer fake.verifyBlockAttestationMutex.Unlock()
	fake.VerifyBlockAttestationStub = nil
	fake.verifyBlockAttestationReturns = struct {
		result1 error
	}{result1}
}

func (fake *BlockVerifier) VerifyBlockAttestationReturnsOnCall(i int, result1 error) {
	fake.verifyBlockAttestationMutex.Lock()
	defer fake.verifyBlockAttestationMutex.Unlock()
	fake.VerifyBlockAttestationStub = nil
	if fake.verifyBlockAttestationReturnsOnCall == nil {
		fake.verifyBlockAttestationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyBlockAttestationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BlockVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyBlockMutex.RLock()
	defer fake.verifyBlockMutex.RUnlock()
	fake.verifyBlockAttestationMutex.RLock()
	defer fake.verifyBlockAttestationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockVerifier) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ blocksprovider.BlockVerifier = new(BlockVerifier)
