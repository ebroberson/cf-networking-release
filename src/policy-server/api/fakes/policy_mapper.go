// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/api"
	"policy-server/store"
	"sync"
)

type PolicyMapper struct {
	AsStorePolicyStub        func([]byte) (store.PolicyCollection, error)
	asStorePolicyMutex       sync.RWMutex
	asStorePolicyArgsForCall []struct {
		arg1 []byte
	}
	asStorePolicyReturns struct {
		result1 store.PolicyCollection
		result2 error
	}
	asStorePolicyReturnsOnCall map[int]struct {
		result1 store.PolicyCollection
		result2 error
	}
	AsBytesStub        func([]store.Policy) ([]byte, error)
	asBytesMutex       sync.RWMutex
	asBytesArgsForCall []struct {
		arg1 []store.Policy
	}
	asBytesReturns struct {
		result1 []byte
		result2 error
	}
	asBytesReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyMapper) AsStorePolicy(arg1 []byte) (store.PolicyCollection, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.asStorePolicyMutex.Lock()
	ret, specificReturn := fake.asStorePolicyReturnsOnCall[len(fake.asStorePolicyArgsForCall)]
	fake.asStorePolicyArgsForCall = append(fake.asStorePolicyArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("AsStorePolicy", []interface{}{arg1Copy})
	fake.asStorePolicyMutex.Unlock()
	if fake.AsStorePolicyStub != nil {
		return fake.AsStorePolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.asStorePolicyReturns.result1, fake.asStorePolicyReturns.result2
}

func (fake *PolicyMapper) AsStorePolicyCallCount() int {
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	return len(fake.asStorePolicyArgsForCall)
}

func (fake *PolicyMapper) AsStorePolicyArgsForCall(i int) []byte {
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	return fake.asStorePolicyArgsForCall[i].arg1
}

func (fake *PolicyMapper) AsStorePolicyReturns(result1 store.PolicyCollection, result2 error) {
	fake.AsStorePolicyStub = nil
	fake.asStorePolicyReturns = struct {
		result1 store.PolicyCollection
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsStorePolicyReturnsOnCall(i int, result1 store.PolicyCollection, result2 error) {
	fake.AsStorePolicyStub = nil
	if fake.asStorePolicyReturnsOnCall == nil {
		fake.asStorePolicyReturnsOnCall = make(map[int]struct {
			result1 store.PolicyCollection
			result2 error
		})
	}
	fake.asStorePolicyReturnsOnCall[i] = struct {
		result1 store.PolicyCollection
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsBytes(arg1 []store.Policy) ([]byte, error) {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.asBytesMutex.Lock()
	ret, specificReturn := fake.asBytesReturnsOnCall[len(fake.asBytesArgsForCall)]
	fake.asBytesArgsForCall = append(fake.asBytesArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	fake.recordInvocation("AsBytes", []interface{}{arg1Copy})
	fake.asBytesMutex.Unlock()
	if fake.AsBytesStub != nil {
		return fake.AsBytesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.asBytesReturns.result1, fake.asBytesReturns.result2
}

func (fake *PolicyMapper) AsBytesCallCount() int {
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	return len(fake.asBytesArgsForCall)
}

func (fake *PolicyMapper) AsBytesArgsForCall(i int) []store.Policy {
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	return fake.asBytesArgsForCall[i].arg1
}

func (fake *PolicyMapper) AsBytesReturns(result1 []byte, result2 error) {
	fake.AsBytesStub = nil
	fake.asBytesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsBytesReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.AsBytesStub = nil
	if fake.asBytesReturnsOnCall == nil {
		fake.asBytesReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.asBytesReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyMapper) recordInvocation(key string, args []interface{}) {
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

var _ api.PolicyMapper = new(PolicyMapper)
