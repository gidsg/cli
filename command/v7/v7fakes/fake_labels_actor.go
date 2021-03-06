// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/types"
)

type FakeLabelsActor struct {
	GetApplicationLabelsStub        func(string, string) (map[string]types.NullString, v7action.Warnings, error)
	getApplicationLabelsMutex       sync.RWMutex
	getApplicationLabelsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getApplicationLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetBuildpackLabelsStub        func(string, string) (map[string]types.NullString, v7action.Warnings, error)
	getBuildpackLabelsMutex       sync.RWMutex
	getBuildpackLabelsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getBuildpackLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getBuildpackLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetDomainLabelsStub        func(string) (map[string]types.NullString, v7action.Warnings, error)
	getDomainLabelsMutex       sync.RWMutex
	getDomainLabelsArgsForCall []struct {
		arg1 string
	}
	getDomainLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getDomainLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationLabelsStub        func(string) (map[string]types.NullString, v7action.Warnings, error)
	getOrganizationLabelsMutex       sync.RWMutex
	getOrganizationLabelsArgsForCall []struct {
		arg1 string
	}
	getOrganizationLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetRouteLabelsStub        func(string, string) (map[string]types.NullString, v7action.Warnings, error)
	getRouteLabelsMutex       sync.RWMutex
	getRouteLabelsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getRouteLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getRouteLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetServiceBrokerLabelsStub        func(string) (map[string]types.NullString, v7action.Warnings, error)
	getServiceBrokerLabelsMutex       sync.RWMutex
	getServiceBrokerLabelsArgsForCall []struct {
		arg1 string
	}
	getServiceBrokerLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getServiceBrokerLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceLabelsStub        func(string, string) (map[string]types.NullString, v7action.Warnings, error)
	getSpaceLabelsMutex       sync.RWMutex
	getSpaceLabelsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getSpaceLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	GetStackLabelsStub        func(string) (map[string]types.NullString, v7action.Warnings, error)
	getStackLabelsMutex       sync.RWMutex
	getStackLabelsArgsForCall []struct {
		arg1 string
	}
	getStackLabelsReturns struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	getStackLabelsReturnsOnCall map[int]struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLabelsActor) GetApplicationLabels(arg1 string, arg2 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getApplicationLabelsMutex.Lock()
	ret, specificReturn := fake.getApplicationLabelsReturnsOnCall[len(fake.getApplicationLabelsArgsForCall)]
	fake.getApplicationLabelsArgsForCall = append(fake.getApplicationLabelsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationLabels", []interface{}{arg1, arg2})
	fake.getApplicationLabelsMutex.Unlock()
	if fake.GetApplicationLabelsStub != nil {
		return fake.GetApplicationLabelsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetApplicationLabelsCallCount() int {
	fake.getApplicationLabelsMutex.RLock()
	defer fake.getApplicationLabelsMutex.RUnlock()
	return len(fake.getApplicationLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetApplicationLabelsCalls(stub func(string, string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getApplicationLabelsMutex.Lock()
	defer fake.getApplicationLabelsMutex.Unlock()
	fake.GetApplicationLabelsStub = stub
}

func (fake *FakeLabelsActor) GetApplicationLabelsArgsForCall(i int) (string, string) {
	fake.getApplicationLabelsMutex.RLock()
	defer fake.getApplicationLabelsMutex.RUnlock()
	argsForCall := fake.getApplicationLabelsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLabelsActor) GetApplicationLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getApplicationLabelsMutex.Lock()
	defer fake.getApplicationLabelsMutex.Unlock()
	fake.GetApplicationLabelsStub = nil
	fake.getApplicationLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetApplicationLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getApplicationLabelsMutex.Lock()
	defer fake.getApplicationLabelsMutex.Unlock()
	fake.GetApplicationLabelsStub = nil
	if fake.getApplicationLabelsReturnsOnCall == nil {
		fake.getApplicationLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetBuildpackLabels(arg1 string, arg2 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getBuildpackLabelsMutex.Lock()
	ret, specificReturn := fake.getBuildpackLabelsReturnsOnCall[len(fake.getBuildpackLabelsArgsForCall)]
	fake.getBuildpackLabelsArgsForCall = append(fake.getBuildpackLabelsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetBuildpackLabels", []interface{}{arg1, arg2})
	fake.getBuildpackLabelsMutex.Unlock()
	if fake.GetBuildpackLabelsStub != nil {
		return fake.GetBuildpackLabelsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getBuildpackLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetBuildpackLabelsCallCount() int {
	fake.getBuildpackLabelsMutex.RLock()
	defer fake.getBuildpackLabelsMutex.RUnlock()
	return len(fake.getBuildpackLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetBuildpackLabelsCalls(stub func(string, string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getBuildpackLabelsMutex.Lock()
	defer fake.getBuildpackLabelsMutex.Unlock()
	fake.GetBuildpackLabelsStub = stub
}

func (fake *FakeLabelsActor) GetBuildpackLabelsArgsForCall(i int) (string, string) {
	fake.getBuildpackLabelsMutex.RLock()
	defer fake.getBuildpackLabelsMutex.RUnlock()
	argsForCall := fake.getBuildpackLabelsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLabelsActor) GetBuildpackLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getBuildpackLabelsMutex.Lock()
	defer fake.getBuildpackLabelsMutex.Unlock()
	fake.GetBuildpackLabelsStub = nil
	fake.getBuildpackLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetBuildpackLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getBuildpackLabelsMutex.Lock()
	defer fake.getBuildpackLabelsMutex.Unlock()
	fake.GetBuildpackLabelsStub = nil
	if fake.getBuildpackLabelsReturnsOnCall == nil {
		fake.getBuildpackLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getBuildpackLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetDomainLabels(arg1 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getDomainLabelsMutex.Lock()
	ret, specificReturn := fake.getDomainLabelsReturnsOnCall[len(fake.getDomainLabelsArgsForCall)]
	fake.getDomainLabelsArgsForCall = append(fake.getDomainLabelsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetDomainLabels", []interface{}{arg1})
	fake.getDomainLabelsMutex.Unlock()
	if fake.GetDomainLabelsStub != nil {
		return fake.GetDomainLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDomainLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetDomainLabelsCallCount() int {
	fake.getDomainLabelsMutex.RLock()
	defer fake.getDomainLabelsMutex.RUnlock()
	return len(fake.getDomainLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetDomainLabelsCalls(stub func(string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getDomainLabelsMutex.Lock()
	defer fake.getDomainLabelsMutex.Unlock()
	fake.GetDomainLabelsStub = stub
}

func (fake *FakeLabelsActor) GetDomainLabelsArgsForCall(i int) string {
	fake.getDomainLabelsMutex.RLock()
	defer fake.getDomainLabelsMutex.RUnlock()
	argsForCall := fake.getDomainLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLabelsActor) GetDomainLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getDomainLabelsMutex.Lock()
	defer fake.getDomainLabelsMutex.Unlock()
	fake.GetDomainLabelsStub = nil
	fake.getDomainLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetDomainLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getDomainLabelsMutex.Lock()
	defer fake.getDomainLabelsMutex.Unlock()
	fake.GetDomainLabelsStub = nil
	if fake.getDomainLabelsReturnsOnCall == nil {
		fake.getDomainLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getDomainLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetOrganizationLabels(arg1 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getOrganizationLabelsMutex.Lock()
	ret, specificReturn := fake.getOrganizationLabelsReturnsOnCall[len(fake.getOrganizationLabelsArgsForCall)]
	fake.getOrganizationLabelsArgsForCall = append(fake.getOrganizationLabelsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationLabels", []interface{}{arg1})
	fake.getOrganizationLabelsMutex.Unlock()
	if fake.GetOrganizationLabelsStub != nil {
		return fake.GetOrganizationLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetOrganizationLabelsCallCount() int {
	fake.getOrganizationLabelsMutex.RLock()
	defer fake.getOrganizationLabelsMutex.RUnlock()
	return len(fake.getOrganizationLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetOrganizationLabelsCalls(stub func(string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getOrganizationLabelsMutex.Lock()
	defer fake.getOrganizationLabelsMutex.Unlock()
	fake.GetOrganizationLabelsStub = stub
}

func (fake *FakeLabelsActor) GetOrganizationLabelsArgsForCall(i int) string {
	fake.getOrganizationLabelsMutex.RLock()
	defer fake.getOrganizationLabelsMutex.RUnlock()
	argsForCall := fake.getOrganizationLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLabelsActor) GetOrganizationLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationLabelsMutex.Lock()
	defer fake.getOrganizationLabelsMutex.Unlock()
	fake.GetOrganizationLabelsStub = nil
	fake.getOrganizationLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetOrganizationLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationLabelsMutex.Lock()
	defer fake.getOrganizationLabelsMutex.Unlock()
	fake.GetOrganizationLabelsStub = nil
	if fake.getOrganizationLabelsReturnsOnCall == nil {
		fake.getOrganizationLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetRouteLabels(arg1 string, arg2 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getRouteLabelsMutex.Lock()
	ret, specificReturn := fake.getRouteLabelsReturnsOnCall[len(fake.getRouteLabelsArgsForCall)]
	fake.getRouteLabelsArgsForCall = append(fake.getRouteLabelsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetRouteLabels", []interface{}{arg1, arg2})
	fake.getRouteLabelsMutex.Unlock()
	if fake.GetRouteLabelsStub != nil {
		return fake.GetRouteLabelsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getRouteLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetRouteLabelsCallCount() int {
	fake.getRouteLabelsMutex.RLock()
	defer fake.getRouteLabelsMutex.RUnlock()
	return len(fake.getRouteLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetRouteLabelsCalls(stub func(string, string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getRouteLabelsMutex.Lock()
	defer fake.getRouteLabelsMutex.Unlock()
	fake.GetRouteLabelsStub = stub
}

func (fake *FakeLabelsActor) GetRouteLabelsArgsForCall(i int) (string, string) {
	fake.getRouteLabelsMutex.RLock()
	defer fake.getRouteLabelsMutex.RUnlock()
	argsForCall := fake.getRouteLabelsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLabelsActor) GetRouteLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getRouteLabelsMutex.Lock()
	defer fake.getRouteLabelsMutex.Unlock()
	fake.GetRouteLabelsStub = nil
	fake.getRouteLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetRouteLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getRouteLabelsMutex.Lock()
	defer fake.getRouteLabelsMutex.Unlock()
	fake.GetRouteLabelsStub = nil
	if fake.getRouteLabelsReturnsOnCall == nil {
		fake.getRouteLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getRouteLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetServiceBrokerLabels(arg1 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getServiceBrokerLabelsMutex.Lock()
	ret, specificReturn := fake.getServiceBrokerLabelsReturnsOnCall[len(fake.getServiceBrokerLabelsArgsForCall)]
	fake.getServiceBrokerLabelsArgsForCall = append(fake.getServiceBrokerLabelsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceBrokerLabels", []interface{}{arg1})
	fake.getServiceBrokerLabelsMutex.Unlock()
	if fake.GetServiceBrokerLabelsStub != nil {
		return fake.GetServiceBrokerLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServiceBrokerLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetServiceBrokerLabelsCallCount() int {
	fake.getServiceBrokerLabelsMutex.RLock()
	defer fake.getServiceBrokerLabelsMutex.RUnlock()
	return len(fake.getServiceBrokerLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetServiceBrokerLabelsCalls(stub func(string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getServiceBrokerLabelsMutex.Lock()
	defer fake.getServiceBrokerLabelsMutex.Unlock()
	fake.GetServiceBrokerLabelsStub = stub
}

func (fake *FakeLabelsActor) GetServiceBrokerLabelsArgsForCall(i int) string {
	fake.getServiceBrokerLabelsMutex.RLock()
	defer fake.getServiceBrokerLabelsMutex.RUnlock()
	argsForCall := fake.getServiceBrokerLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLabelsActor) GetServiceBrokerLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getServiceBrokerLabelsMutex.Lock()
	defer fake.getServiceBrokerLabelsMutex.Unlock()
	fake.GetServiceBrokerLabelsStub = nil
	fake.getServiceBrokerLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetServiceBrokerLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getServiceBrokerLabelsMutex.Lock()
	defer fake.getServiceBrokerLabelsMutex.Unlock()
	fake.GetServiceBrokerLabelsStub = nil
	if fake.getServiceBrokerLabelsReturnsOnCall == nil {
		fake.getServiceBrokerLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getServiceBrokerLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetSpaceLabels(arg1 string, arg2 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getSpaceLabelsMutex.Lock()
	ret, specificReturn := fake.getSpaceLabelsReturnsOnCall[len(fake.getSpaceLabelsArgsForCall)]
	fake.getSpaceLabelsArgsForCall = append(fake.getSpaceLabelsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceLabels", []interface{}{arg1, arg2})
	fake.getSpaceLabelsMutex.Unlock()
	if fake.GetSpaceLabelsStub != nil {
		return fake.GetSpaceLabelsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetSpaceLabelsCallCount() int {
	fake.getSpaceLabelsMutex.RLock()
	defer fake.getSpaceLabelsMutex.RUnlock()
	return len(fake.getSpaceLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetSpaceLabelsCalls(stub func(string, string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getSpaceLabelsMutex.Lock()
	defer fake.getSpaceLabelsMutex.Unlock()
	fake.GetSpaceLabelsStub = stub
}

func (fake *FakeLabelsActor) GetSpaceLabelsArgsForCall(i int) (string, string) {
	fake.getSpaceLabelsMutex.RLock()
	defer fake.getSpaceLabelsMutex.RUnlock()
	argsForCall := fake.getSpaceLabelsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLabelsActor) GetSpaceLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getSpaceLabelsMutex.Lock()
	defer fake.getSpaceLabelsMutex.Unlock()
	fake.GetSpaceLabelsStub = nil
	fake.getSpaceLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetSpaceLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getSpaceLabelsMutex.Lock()
	defer fake.getSpaceLabelsMutex.Unlock()
	fake.GetSpaceLabelsStub = nil
	if fake.getSpaceLabelsReturnsOnCall == nil {
		fake.getSpaceLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetStackLabels(arg1 string) (map[string]types.NullString, v7action.Warnings, error) {
	fake.getStackLabelsMutex.Lock()
	ret, specificReturn := fake.getStackLabelsReturnsOnCall[len(fake.getStackLabelsArgsForCall)]
	fake.getStackLabelsArgsForCall = append(fake.getStackLabelsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStackLabels", []interface{}{arg1})
	fake.getStackLabelsMutex.Unlock()
	if fake.GetStackLabelsStub != nil {
		return fake.GetStackLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getStackLabelsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLabelsActor) GetStackLabelsCallCount() int {
	fake.getStackLabelsMutex.RLock()
	defer fake.getStackLabelsMutex.RUnlock()
	return len(fake.getStackLabelsArgsForCall)
}

func (fake *FakeLabelsActor) GetStackLabelsCalls(stub func(string) (map[string]types.NullString, v7action.Warnings, error)) {
	fake.getStackLabelsMutex.Lock()
	defer fake.getStackLabelsMutex.Unlock()
	fake.GetStackLabelsStub = stub
}

func (fake *FakeLabelsActor) GetStackLabelsArgsForCall(i int) string {
	fake.getStackLabelsMutex.RLock()
	defer fake.getStackLabelsMutex.RUnlock()
	argsForCall := fake.getStackLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLabelsActor) GetStackLabelsReturns(result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getStackLabelsMutex.Lock()
	defer fake.getStackLabelsMutex.Unlock()
	fake.GetStackLabelsStub = nil
	fake.getStackLabelsReturns = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) GetStackLabelsReturnsOnCall(i int, result1 map[string]types.NullString, result2 v7action.Warnings, result3 error) {
	fake.getStackLabelsMutex.Lock()
	defer fake.getStackLabelsMutex.Unlock()
	fake.GetStackLabelsStub = nil
	if fake.getStackLabelsReturnsOnCall == nil {
		fake.getStackLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]types.NullString
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getStackLabelsReturnsOnCall[i] = struct {
		result1 map[string]types.NullString
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLabelsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationLabelsMutex.RLock()
	defer fake.getApplicationLabelsMutex.RUnlock()
	fake.getBuildpackLabelsMutex.RLock()
	defer fake.getBuildpackLabelsMutex.RUnlock()
	fake.getDomainLabelsMutex.RLock()
	defer fake.getDomainLabelsMutex.RUnlock()
	fake.getOrganizationLabelsMutex.RLock()
	defer fake.getOrganizationLabelsMutex.RUnlock()
	fake.getRouteLabelsMutex.RLock()
	defer fake.getRouteLabelsMutex.RUnlock()
	fake.getServiceBrokerLabelsMutex.RLock()
	defer fake.getServiceBrokerLabelsMutex.RUnlock()
	fake.getSpaceLabelsMutex.RLock()
	defer fake.getSpaceLabelsMutex.RUnlock()
	fake.getStackLabelsMutex.RLock()
	defer fake.getStackLabelsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLabelsActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.LabelsActor = new(FakeLabelsActor)
