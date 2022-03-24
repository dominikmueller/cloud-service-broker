// Code generated by counterfeiter. DO NOT EDIT.
package wrapperfakes

import (
	"context"
	"os/exec"
	"sync"

	"github.com/cloudfoundry/cloud-service-broker/pkg/providers/tf/wrapper"
)

type FakeTerraformExecutor struct {
	ExecuteStub        func(context.Context, *exec.Cmd) (wrapper.ExecutionOutput, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 context.Context
		arg2 *exec.Cmd
	}
	executeReturns struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTerraformExecutor) Execute(arg1 context.Context, arg2 *exec.Cmd) (wrapper.ExecutionOutput, error) {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 context.Context
		arg2 *exec.Cmd
	}{arg1, arg2})
	stub := fake.ExecuteStub
	fakeReturns := fake.executeReturns
	fake.recordInvocation("Execute", []interface{}{arg1, arg2})
	fake.executeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTerraformExecutor) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeTerraformExecutor) ExecuteCalls(stub func(context.Context, *exec.Cmd) (wrapper.ExecutionOutput, error)) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeTerraformExecutor) ExecuteArgsForCall(i int) (context.Context, *exec.Cmd) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTerraformExecutor) ExecuteReturns(result1 wrapper.ExecutionOutput, result2 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeTerraformExecutor) ExecuteReturnsOnCall(i int, result1 wrapper.ExecutionOutput, result2 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 wrapper.ExecutionOutput
			result2 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeTerraformExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTerraformExecutor) recordInvocation(key string, args []interface{}) {
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

var _ wrapper.TerraformExecutor = new(FakeTerraformExecutor)