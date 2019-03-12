// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"context"
	"sync"

	apichaincode "github.com/hyperledger/udo/peer/chaincode/api"
	"github.com/hyperledger/udo/peer/common/api"
	"google.golang.org/grpc"
)

type PeerDeliverClient struct {
	DeliverStub        func(ctx context.Context, opts ...grpc.CallOption) (apichaincode.Deliver, error)
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		ctx  context.Context
		opts []grpc.CallOption
	}
	deliverReturns struct {
		result1 apichaincode.Deliver
		result2 error
	}
	deliverReturnsOnCall map[int]struct {
		result1 apichaincode.Deliver
		result2 error
	}
	DeliverFilteredStub        func(ctx context.Context, opts ...grpc.CallOption) (apichaincode.Deliver, error)
	deliverFilteredMutex       sync.RWMutex
	deliverFilteredArgsForCall []struct {
		ctx  context.Context
		opts []grpc.CallOption
	}
	deliverFilteredReturns struct {
		result1 apichaincode.Deliver
		result2 error
	}
	deliverFilteredReturnsOnCall map[int]struct {
		result1 apichaincode.Deliver
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PeerDeliverClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (apichaincode.Deliver, error) {
	fake.deliverMutex.Lock()
	ret, specificReturn := fake.deliverReturnsOnCall[len(fake.deliverArgsForCall)]
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		ctx  context.Context
		opts []grpc.CallOption
	}{ctx, opts})
	fake.recordInvocation("Deliver", []interface{}{ctx, opts})
	fake.deliverMutex.Unlock()
	if fake.DeliverStub != nil {
		return fake.DeliverStub(ctx, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deliverReturns.result1, fake.deliverReturns.result2
}

func (fake *PeerDeliverClient) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *PeerDeliverClient) DeliverArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return fake.deliverArgsForCall[i].ctx, fake.deliverArgsForCall[i].opts
}

func (fake *PeerDeliverClient) DeliverReturns(result1 apichaincode.Deliver, result2 error) {
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 apichaincode.Deliver
		result2 error
	}{result1, result2}
}

func (fake *PeerDeliverClient) DeliverReturnsOnCall(i int, result1 apichaincode.Deliver, result2 error) {
	fake.DeliverStub = nil
	if fake.deliverReturnsOnCall == nil {
		fake.deliverReturnsOnCall = make(map[int]struct {
			result1 apichaincode.Deliver
			result2 error
		})
	}
	fake.deliverReturnsOnCall[i] = struct {
		result1 apichaincode.Deliver
		result2 error
	}{result1, result2}
}

func (fake *PeerDeliverClient) DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (apichaincode.Deliver, error) {
	fake.deliverFilteredMutex.Lock()
	ret, specificReturn := fake.deliverFilteredReturnsOnCall[len(fake.deliverFilteredArgsForCall)]
	fake.deliverFilteredArgsForCall = append(fake.deliverFilteredArgsForCall, struct {
		ctx  context.Context
		opts []grpc.CallOption
	}{ctx, opts})
	fake.recordInvocation("DeliverFiltered", []interface{}{ctx, opts})
	fake.deliverFilteredMutex.Unlock()
	if fake.DeliverFilteredStub != nil {
		return fake.DeliverFilteredStub(ctx, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deliverFilteredReturns.result1, fake.deliverFilteredReturns.result2
}

func (fake *PeerDeliverClient) DeliverFilteredCallCount() int {
	fake.deliverFilteredMutex.RLock()
	defer fake.deliverFilteredMutex.RUnlock()
	return len(fake.deliverFilteredArgsForCall)
}

func (fake *PeerDeliverClient) DeliverFilteredArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.deliverFilteredMutex.RLock()
	defer fake.deliverFilteredMutex.RUnlock()
	return fake.deliverFilteredArgsForCall[i].ctx, fake.deliverFilteredArgsForCall[i].opts
}

func (fake *PeerDeliverClient) DeliverFilteredReturns(result1 apichaincode.Deliver, result2 error) {
	fake.DeliverFilteredStub = nil
	fake.deliverFilteredReturns = struct {
		result1 apichaincode.Deliver
		result2 error
	}{result1, result2}
}

func (fake *PeerDeliverClient) DeliverFilteredReturnsOnCall(i int, result1 apichaincode.Deliver, result2 error) {
	fake.DeliverFilteredStub = nil
	if fake.deliverFilteredReturnsOnCall == nil {
		fake.deliverFilteredReturnsOnCall = make(map[int]struct {
			result1 apichaincode.Deliver
			result2 error
		})
	}
	fake.deliverFilteredReturnsOnCall[i] = struct {
		result1 apichaincode.Deliver
		result2 error
	}{result1, result2}
}

func (fake *PeerDeliverClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	fake.deliverFilteredMutex.RLock()
	defer fake.deliverFilteredMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PeerDeliverClient) recordInvocation(key string, args []interface{}) {
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

var _ api.PeerDeliverClient = new(PeerDeliverClient)
