// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/weaveworks/eksctl/pkg/actions/irsa"
	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
)

type FakeStackManager struct {
	CreateStackStub        func(context.Context, string, builder.ResourceSetReader, map[string]string, map[string]string, chan error) error
	createStackMutex       sync.RWMutex
	createStackArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 builder.ResourceSetReader
		arg4 map[string]string
		arg5 map[string]string
		arg6 chan error
	}
	createStackReturns struct {
		result1 error
	}
	createStackReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStackBySpecStub        func(context.Context, *types.Stack) (*types.Stack, error)
	deleteStackBySpecMutex       sync.RWMutex
	deleteStackBySpecArgsForCall []struct {
		arg1 context.Context
		arg2 *types.Stack
	}
	deleteStackBySpecReturns struct {
		result1 *types.Stack
		result2 error
	}
	deleteStackBySpecReturnsOnCall map[int]struct {
		result1 *types.Stack
		result2 error
	}
	DeleteStackBySpecSyncStub        func(context.Context, *types.Stack, chan error) error
	deleteStackBySpecSyncMutex       sync.RWMutex
	deleteStackBySpecSyncArgsForCall []struct {
		arg1 context.Context
		arg2 *types.Stack
		arg3 chan error
	}
	deleteStackBySpecSyncReturns struct {
		result1 error
	}
	deleteStackBySpecSyncReturnsOnCall map[int]struct {
		result1 error
	}
	DescribeIAMServiceAccountStacksStub        func(context.Context) ([]*types.Stack, error)
	describeIAMServiceAccountStacksMutex       sync.RWMutex
	describeIAMServiceAccountStacksArgsForCall []struct {
		arg1 context.Context
	}
	describeIAMServiceAccountStacksReturns struct {
		result1 []*types.Stack
		result2 error
	}
	describeIAMServiceAccountStacksReturnsOnCall map[int]struct {
		result1 []*types.Stack
		result2 error
	}
	GetIAMServiceAccountsStub        func(context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error)
	getIAMServiceAccountsMutex       sync.RWMutex
	getIAMServiceAccountsArgsForCall []struct {
		arg1 context.Context
	}
	getIAMServiceAccountsReturns struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}
	getIAMServiceAccountsReturnsOnCall map[int]struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}
	GetStackTemplateStub        func(context.Context, string) (string, error)
	getStackTemplateMutex       sync.RWMutex
	getStackTemplateArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStackTemplateReturns struct {
		result1 string
		result2 error
	}
	getStackTemplateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UpdateStackStub        func(context.Context, manager.UpdateStackOptions) error
	updateStackMutex       sync.RWMutex
	updateStackArgsForCall []struct {
		arg1 context.Context
		arg2 manager.UpdateStackOptions
	}
	updateStackReturns struct {
		result1 error
	}
	updateStackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStackManager) CreateStack(arg1 context.Context, arg2 string, arg3 builder.ResourceSetReader, arg4 map[string]string, arg5 map[string]string, arg6 chan error) error {
	fake.createStackMutex.Lock()
	ret, specificReturn := fake.createStackReturnsOnCall[len(fake.createStackArgsForCall)]
	fake.createStackArgsForCall = append(fake.createStackArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 builder.ResourceSetReader
		arg4 map[string]string
		arg5 map[string]string
		arg6 chan error
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.CreateStackStub
	fakeReturns := fake.createStackReturns
	fake.recordInvocation("CreateStack", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createStackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStackManager) CreateStackCallCount() int {
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	return len(fake.createStackArgsForCall)
}

func (fake *FakeStackManager) CreateStackCalls(stub func(context.Context, string, builder.ResourceSetReader, map[string]string, map[string]string, chan error) error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = stub
}

func (fake *FakeStackManager) CreateStackArgsForCall(i int) (context.Context, string, builder.ResourceSetReader, map[string]string, map[string]string, chan error) {
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	argsForCall := fake.createStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeStackManager) CreateStackReturns(result1 error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = nil
	fake.createStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) CreateStackReturnsOnCall(i int, result1 error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = nil
	if fake.createStackReturnsOnCall == nil {
		fake.createStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) DeleteStackBySpec(arg1 context.Context, arg2 *types.Stack) (*types.Stack, error) {
	fake.deleteStackBySpecMutex.Lock()
	ret, specificReturn := fake.deleteStackBySpecReturnsOnCall[len(fake.deleteStackBySpecArgsForCall)]
	fake.deleteStackBySpecArgsForCall = append(fake.deleteStackBySpecArgsForCall, struct {
		arg1 context.Context
		arg2 *types.Stack
	}{arg1, arg2})
	stub := fake.DeleteStackBySpecStub
	fakeReturns := fake.deleteStackBySpecReturns
	fake.recordInvocation("DeleteStackBySpec", []interface{}{arg1, arg2})
	fake.deleteStackBySpecMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) DeleteStackBySpecCallCount() int {
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	return len(fake.deleteStackBySpecArgsForCall)
}

func (fake *FakeStackManager) DeleteStackBySpecCalls(stub func(context.Context, *types.Stack) (*types.Stack, error)) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = stub
}

func (fake *FakeStackManager) DeleteStackBySpecArgsForCall(i int) (context.Context, *types.Stack) {
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	argsForCall := fake.deleteStackBySpecArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackManager) DeleteStackBySpecReturns(result1 *types.Stack, result2 error) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = nil
	fake.deleteStackBySpecReturns = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) DeleteStackBySpecReturnsOnCall(i int, result1 *types.Stack, result2 error) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = nil
	if fake.deleteStackBySpecReturnsOnCall == nil {
		fake.deleteStackBySpecReturnsOnCall = make(map[int]struct {
			result1 *types.Stack
			result2 error
		})
	}
	fake.deleteStackBySpecReturnsOnCall[i] = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) DeleteStackBySpecSync(arg1 context.Context, arg2 *types.Stack, arg3 chan error) error {
	fake.deleteStackBySpecSyncMutex.Lock()
	ret, specificReturn := fake.deleteStackBySpecSyncReturnsOnCall[len(fake.deleteStackBySpecSyncArgsForCall)]
	fake.deleteStackBySpecSyncArgsForCall = append(fake.deleteStackBySpecSyncArgsForCall, struct {
		arg1 context.Context
		arg2 *types.Stack
		arg3 chan error
	}{arg1, arg2, arg3})
	stub := fake.DeleteStackBySpecSyncStub
	fakeReturns := fake.deleteStackBySpecSyncReturns
	fake.recordInvocation("DeleteStackBySpecSync", []interface{}{arg1, arg2, arg3})
	fake.deleteStackBySpecSyncMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStackManager) DeleteStackBySpecSyncCallCount() int {
	fake.deleteStackBySpecSyncMutex.RLock()
	defer fake.deleteStackBySpecSyncMutex.RUnlock()
	return len(fake.deleteStackBySpecSyncArgsForCall)
}

func (fake *FakeStackManager) DeleteStackBySpecSyncCalls(stub func(context.Context, *types.Stack, chan error) error) {
	fake.deleteStackBySpecSyncMutex.Lock()
	defer fake.deleteStackBySpecSyncMutex.Unlock()
	fake.DeleteStackBySpecSyncStub = stub
}

func (fake *FakeStackManager) DeleteStackBySpecSyncArgsForCall(i int) (context.Context, *types.Stack, chan error) {
	fake.deleteStackBySpecSyncMutex.RLock()
	defer fake.deleteStackBySpecSyncMutex.RUnlock()
	argsForCall := fake.deleteStackBySpecSyncArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStackManager) DeleteStackBySpecSyncReturns(result1 error) {
	fake.deleteStackBySpecSyncMutex.Lock()
	defer fake.deleteStackBySpecSyncMutex.Unlock()
	fake.DeleteStackBySpecSyncStub = nil
	fake.deleteStackBySpecSyncReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) DeleteStackBySpecSyncReturnsOnCall(i int, result1 error) {
	fake.deleteStackBySpecSyncMutex.Lock()
	defer fake.deleteStackBySpecSyncMutex.Unlock()
	fake.DeleteStackBySpecSyncStub = nil
	if fake.deleteStackBySpecSyncReturnsOnCall == nil {
		fake.deleteStackBySpecSyncReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteStackBySpecSyncReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacks(arg1 context.Context) ([]*types.Stack, error) {
	fake.describeIAMServiceAccountStacksMutex.Lock()
	ret, specificReturn := fake.describeIAMServiceAccountStacksReturnsOnCall[len(fake.describeIAMServiceAccountStacksArgsForCall)]
	fake.describeIAMServiceAccountStacksArgsForCall = append(fake.describeIAMServiceAccountStacksArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.DescribeIAMServiceAccountStacksStub
	fakeReturns := fake.describeIAMServiceAccountStacksReturns
	fake.recordInvocation("DescribeIAMServiceAccountStacks", []interface{}{arg1})
	fake.describeIAMServiceAccountStacksMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacksCallCount() int {
	fake.describeIAMServiceAccountStacksMutex.RLock()
	defer fake.describeIAMServiceAccountStacksMutex.RUnlock()
	return len(fake.describeIAMServiceAccountStacksArgsForCall)
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacksCalls(stub func(context.Context) ([]*types.Stack, error)) {
	fake.describeIAMServiceAccountStacksMutex.Lock()
	defer fake.describeIAMServiceAccountStacksMutex.Unlock()
	fake.DescribeIAMServiceAccountStacksStub = stub
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacksArgsForCall(i int) context.Context {
	fake.describeIAMServiceAccountStacksMutex.RLock()
	defer fake.describeIAMServiceAccountStacksMutex.RUnlock()
	argsForCall := fake.describeIAMServiceAccountStacksArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacksReturns(result1 []*types.Stack, result2 error) {
	fake.describeIAMServiceAccountStacksMutex.Lock()
	defer fake.describeIAMServiceAccountStacksMutex.Unlock()
	fake.DescribeIAMServiceAccountStacksStub = nil
	fake.describeIAMServiceAccountStacksReturns = struct {
		result1 []*types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) DescribeIAMServiceAccountStacksReturnsOnCall(i int, result1 []*types.Stack, result2 error) {
	fake.describeIAMServiceAccountStacksMutex.Lock()
	defer fake.describeIAMServiceAccountStacksMutex.Unlock()
	fake.DescribeIAMServiceAccountStacksStub = nil
	if fake.describeIAMServiceAccountStacksReturnsOnCall == nil {
		fake.describeIAMServiceAccountStacksReturnsOnCall = make(map[int]struct {
			result1 []*types.Stack
			result2 error
		})
	}
	fake.describeIAMServiceAccountStacksReturnsOnCall[i] = struct {
		result1 []*types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) GetIAMServiceAccounts(arg1 context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error) {
	fake.getIAMServiceAccountsMutex.Lock()
	ret, specificReturn := fake.getIAMServiceAccountsReturnsOnCall[len(fake.getIAMServiceAccountsArgsForCall)]
	fake.getIAMServiceAccountsArgsForCall = append(fake.getIAMServiceAccountsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetIAMServiceAccountsStub
	fakeReturns := fake.getIAMServiceAccountsReturns
	fake.recordInvocation("GetIAMServiceAccounts", []interface{}{arg1})
	fake.getIAMServiceAccountsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) GetIAMServiceAccountsCallCount() int {
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	return len(fake.getIAMServiceAccountsArgsForCall)
}

func (fake *FakeStackManager) GetIAMServiceAccountsCalls(stub func(context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error)) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = stub
}

func (fake *FakeStackManager) GetIAMServiceAccountsArgsForCall(i int) context.Context {
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	argsForCall := fake.getIAMServiceAccountsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackManager) GetIAMServiceAccountsReturns(result1 []*v1alpha5.ClusterIAMServiceAccount, result2 error) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = nil
	fake.getIAMServiceAccountsReturns = struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) GetIAMServiceAccountsReturnsOnCall(i int, result1 []*v1alpha5.ClusterIAMServiceAccount, result2 error) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = nil
	if fake.getIAMServiceAccountsReturnsOnCall == nil {
		fake.getIAMServiceAccountsReturnsOnCall = make(map[int]struct {
			result1 []*v1alpha5.ClusterIAMServiceAccount
			result2 error
		})
	}
	fake.getIAMServiceAccountsReturnsOnCall[i] = struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) GetStackTemplate(arg1 context.Context, arg2 string) (string, error) {
	fake.getStackTemplateMutex.Lock()
	ret, specificReturn := fake.getStackTemplateReturnsOnCall[len(fake.getStackTemplateArgsForCall)]
	fake.getStackTemplateArgsForCall = append(fake.getStackTemplateArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStackTemplateStub
	fakeReturns := fake.getStackTemplateReturns
	fake.recordInvocation("GetStackTemplate", []interface{}{arg1, arg2})
	fake.getStackTemplateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) GetStackTemplateCallCount() int {
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	return len(fake.getStackTemplateArgsForCall)
}

func (fake *FakeStackManager) GetStackTemplateCalls(stub func(context.Context, string) (string, error)) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = stub
}

func (fake *FakeStackManager) GetStackTemplateArgsForCall(i int) (context.Context, string) {
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	argsForCall := fake.getStackTemplateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackManager) GetStackTemplateReturns(result1 string, result2 error) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = nil
	fake.getStackTemplateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) GetStackTemplateReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = nil
	if fake.getStackTemplateReturnsOnCall == nil {
		fake.getStackTemplateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStackTemplateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) UpdateStack(arg1 context.Context, arg2 manager.UpdateStackOptions) error {
	fake.updateStackMutex.Lock()
	ret, specificReturn := fake.updateStackReturnsOnCall[len(fake.updateStackArgsForCall)]
	fake.updateStackArgsForCall = append(fake.updateStackArgsForCall, struct {
		arg1 context.Context
		arg2 manager.UpdateStackOptions
	}{arg1, arg2})
	stub := fake.UpdateStackStub
	fakeReturns := fake.updateStackReturns
	fake.recordInvocation("UpdateStack", []interface{}{arg1, arg2})
	fake.updateStackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStackManager) UpdateStackCallCount() int {
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	return len(fake.updateStackArgsForCall)
}

func (fake *FakeStackManager) UpdateStackCalls(stub func(context.Context, manager.UpdateStackOptions) error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = stub
}

func (fake *FakeStackManager) UpdateStackArgsForCall(i int) (context.Context, manager.UpdateStackOptions) {
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	argsForCall := fake.updateStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackManager) UpdateStackReturns(result1 error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = nil
	fake.updateStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) UpdateStackReturnsOnCall(i int, result1 error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = nil
	if fake.updateStackReturnsOnCall == nil {
		fake.updateStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	fake.deleteStackBySpecSyncMutex.RLock()
	defer fake.deleteStackBySpecSyncMutex.RUnlock()
	fake.describeIAMServiceAccountStacksMutex.RLock()
	defer fake.describeIAMServiceAccountStacksMutex.RUnlock()
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStackManager) recordInvocation(key string, args []interface{}) {
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

var _ irsa.StackManager = new(FakeStackManager)
