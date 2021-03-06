// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeRestartAppInstanceActor struct {
	DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub        func(string, string, string, int) (v7action.Warnings, error)
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex       sync.RWMutex
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 int
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex(arg1 string, arg2 string, arg3 string, arg4 int) (v7action.Warnings, error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	ret, specificReturn := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)]
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall = append(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 int
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex", []interface{}{arg1, arg2, arg3, arg4})
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	if fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub != nil {
		return fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCallCount() int {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCalls(stub func(string, string, string, int) (v7action.Warnings, error)) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = stub
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall(i int) (string, string, string, int) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	argsForCall := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	if fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall == nil {
		fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRestartAppInstanceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRestartAppInstanceActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.RestartAppInstanceActor = new(FakeRestartAppInstanceActor)
