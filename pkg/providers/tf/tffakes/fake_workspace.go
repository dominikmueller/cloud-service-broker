// Code generated by counterfeiter. DO NOT EDIT.
package tffakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry/cloud-service-broker/pkg/providers/tf"
	"github.com/cloudfoundry/cloud-service-broker/pkg/providers/tf/wrapper"
	version "github.com/hashicorp/go-version"
)

type FakeWorkspace struct {
	ExecuteStub        func(context.Context, wrapper.TerraformExecutor, ...wrapper.TerraformCommand) (wrapper.ExecutionOutput, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 context.Context
		arg2 wrapper.TerraformExecutor
		arg3 []wrapper.TerraformCommand
	}
	executeReturns struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}
	ModuleDefinitionsStub        func() []wrapper.ModuleDefinition
	moduleDefinitionsMutex       sync.RWMutex
	moduleDefinitionsArgsForCall []struct {
	}
	moduleDefinitionsReturns struct {
		result1 []wrapper.ModuleDefinition
	}
	moduleDefinitionsReturnsOnCall map[int]struct {
		result1 []wrapper.ModuleDefinition
	}
	ModuleInstancesStub        func() []wrapper.ModuleInstance
	moduleInstancesMutex       sync.RWMutex
	moduleInstancesArgsForCall []struct {
	}
	moduleInstancesReturns struct {
		result1 []wrapper.ModuleInstance
	}
	moduleInstancesReturnsOnCall map[int]struct {
		result1 []wrapper.ModuleInstance
	}
	OutputsStub        func(string) (map[string]interface{}, error)
	outputsMutex       sync.RWMutex
	outputsArgsForCall []struct {
		arg1 string
	}
	outputsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	outputsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	SerializeStub        func() (string, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
	}
	serializeReturns struct {
		result1 string
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	StateVersionStub        func() (*version.Version, error)
	stateVersionMutex       sync.RWMutex
	stateVersionArgsForCall []struct {
	}
	stateVersionReturns struct {
		result1 *version.Version
		result2 error
	}
	stateVersionReturnsOnCall map[int]struct {
		result1 *version.Version
		result2 error
	}
	UpdateInstanceConfigurationStub        func(map[string]interface{}) error
	updateInstanceConfigurationMutex       sync.RWMutex
	updateInstanceConfigurationArgsForCall []struct {
		arg1 map[string]interface{}
	}
	updateInstanceConfigurationReturns struct {
		result1 error
	}
	updateInstanceConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkspace) Execute(arg1 context.Context, arg2 wrapper.TerraformExecutor, arg3 ...wrapper.TerraformCommand) (wrapper.ExecutionOutput, error) {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 context.Context
		arg2 wrapper.TerraformExecutor
		arg3 []wrapper.TerraformCommand
	}{arg1, arg2, arg3})
	stub := fake.ExecuteStub
	fakeReturns := fake.executeReturns
	fake.recordInvocation("Execute", []interface{}{arg1, arg2, arg3})
	fake.executeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkspace) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeWorkspace) ExecuteCalls(stub func(context.Context, wrapper.TerraformExecutor, ...wrapper.TerraformCommand) (wrapper.ExecutionOutput, error)) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeWorkspace) ExecuteArgsForCall(i int) (context.Context, wrapper.TerraformExecutor, []wrapper.TerraformCommand) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeWorkspace) ExecuteReturns(result1 wrapper.ExecutionOutput, result2 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 wrapper.ExecutionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) ExecuteReturnsOnCall(i int, result1 wrapper.ExecutionOutput, result2 error) {
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

func (fake *FakeWorkspace) ModuleDefinitions() []wrapper.ModuleDefinition {
	fake.moduleDefinitionsMutex.Lock()
	ret, specificReturn := fake.moduleDefinitionsReturnsOnCall[len(fake.moduleDefinitionsArgsForCall)]
	fake.moduleDefinitionsArgsForCall = append(fake.moduleDefinitionsArgsForCall, struct {
	}{})
	stub := fake.ModuleDefinitionsStub
	fakeReturns := fake.moduleDefinitionsReturns
	fake.recordInvocation("ModuleDefinitions", []interface{}{})
	fake.moduleDefinitionsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWorkspace) ModuleDefinitionsCallCount() int {
	fake.moduleDefinitionsMutex.RLock()
	defer fake.moduleDefinitionsMutex.RUnlock()
	return len(fake.moduleDefinitionsArgsForCall)
}

func (fake *FakeWorkspace) ModuleDefinitionsCalls(stub func() []wrapper.ModuleDefinition) {
	fake.moduleDefinitionsMutex.Lock()
	defer fake.moduleDefinitionsMutex.Unlock()
	fake.ModuleDefinitionsStub = stub
}

func (fake *FakeWorkspace) ModuleDefinitionsReturns(result1 []wrapper.ModuleDefinition) {
	fake.moduleDefinitionsMutex.Lock()
	defer fake.moduleDefinitionsMutex.Unlock()
	fake.ModuleDefinitionsStub = nil
	fake.moduleDefinitionsReturns = struct {
		result1 []wrapper.ModuleDefinition
	}{result1}
}

func (fake *FakeWorkspace) ModuleDefinitionsReturnsOnCall(i int, result1 []wrapper.ModuleDefinition) {
	fake.moduleDefinitionsMutex.Lock()
	defer fake.moduleDefinitionsMutex.Unlock()
	fake.ModuleDefinitionsStub = nil
	if fake.moduleDefinitionsReturnsOnCall == nil {
		fake.moduleDefinitionsReturnsOnCall = make(map[int]struct {
			result1 []wrapper.ModuleDefinition
		})
	}
	fake.moduleDefinitionsReturnsOnCall[i] = struct {
		result1 []wrapper.ModuleDefinition
	}{result1}
}

func (fake *FakeWorkspace) ModuleInstances() []wrapper.ModuleInstance {
	fake.moduleInstancesMutex.Lock()
	ret, specificReturn := fake.moduleInstancesReturnsOnCall[len(fake.moduleInstancesArgsForCall)]
	fake.moduleInstancesArgsForCall = append(fake.moduleInstancesArgsForCall, struct {
	}{})
	stub := fake.ModuleInstancesStub
	fakeReturns := fake.moduleInstancesReturns
	fake.recordInvocation("ModuleInstances", []interface{}{})
	fake.moduleInstancesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWorkspace) ModuleInstancesCallCount() int {
	fake.moduleInstancesMutex.RLock()
	defer fake.moduleInstancesMutex.RUnlock()
	return len(fake.moduleInstancesArgsForCall)
}

func (fake *FakeWorkspace) ModuleInstancesCalls(stub func() []wrapper.ModuleInstance) {
	fake.moduleInstancesMutex.Lock()
	defer fake.moduleInstancesMutex.Unlock()
	fake.ModuleInstancesStub = stub
}

func (fake *FakeWorkspace) ModuleInstancesReturns(result1 []wrapper.ModuleInstance) {
	fake.moduleInstancesMutex.Lock()
	defer fake.moduleInstancesMutex.Unlock()
	fake.ModuleInstancesStub = nil
	fake.moduleInstancesReturns = struct {
		result1 []wrapper.ModuleInstance
	}{result1}
}

func (fake *FakeWorkspace) ModuleInstancesReturnsOnCall(i int, result1 []wrapper.ModuleInstance) {
	fake.moduleInstancesMutex.Lock()
	defer fake.moduleInstancesMutex.Unlock()
	fake.ModuleInstancesStub = nil
	if fake.moduleInstancesReturnsOnCall == nil {
		fake.moduleInstancesReturnsOnCall = make(map[int]struct {
			result1 []wrapper.ModuleInstance
		})
	}
	fake.moduleInstancesReturnsOnCall[i] = struct {
		result1 []wrapper.ModuleInstance
	}{result1}
}

func (fake *FakeWorkspace) Outputs(arg1 string) (map[string]interface{}, error) {
	fake.outputsMutex.Lock()
	ret, specificReturn := fake.outputsReturnsOnCall[len(fake.outputsArgsForCall)]
	fake.outputsArgsForCall = append(fake.outputsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OutputsStub
	fakeReturns := fake.outputsReturns
	fake.recordInvocation("Outputs", []interface{}{arg1})
	fake.outputsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkspace) OutputsCallCount() int {
	fake.outputsMutex.RLock()
	defer fake.outputsMutex.RUnlock()
	return len(fake.outputsArgsForCall)
}

func (fake *FakeWorkspace) OutputsCalls(stub func(string) (map[string]interface{}, error)) {
	fake.outputsMutex.Lock()
	defer fake.outputsMutex.Unlock()
	fake.OutputsStub = stub
}

func (fake *FakeWorkspace) OutputsArgsForCall(i int) string {
	fake.outputsMutex.RLock()
	defer fake.outputsMutex.RUnlock()
	argsForCall := fake.outputsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkspace) OutputsReturns(result1 map[string]interface{}, result2 error) {
	fake.outputsMutex.Lock()
	defer fake.outputsMutex.Unlock()
	fake.OutputsStub = nil
	fake.outputsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) OutputsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.outputsMutex.Lock()
	defer fake.outputsMutex.Unlock()
	fake.OutputsStub = nil
	if fake.outputsReturnsOnCall == nil {
		fake.outputsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.outputsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) Serialize() (string, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
	}{})
	stub := fake.SerializeStub
	fakeReturns := fake.serializeReturns
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkspace) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *FakeWorkspace) SerializeCalls(stub func() (string, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *FakeWorkspace) SerializeReturns(result1 string, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) SerializeReturnsOnCall(i int, result1 string, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) StateVersion() (*version.Version, error) {
	fake.stateVersionMutex.Lock()
	ret, specificReturn := fake.stateVersionReturnsOnCall[len(fake.stateVersionArgsForCall)]
	fake.stateVersionArgsForCall = append(fake.stateVersionArgsForCall, struct {
	}{})
	stub := fake.StateVersionStub
	fakeReturns := fake.stateVersionReturns
	fake.recordInvocation("StateVersion", []interface{}{})
	fake.stateVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkspace) StateVersionCallCount() int {
	fake.stateVersionMutex.RLock()
	defer fake.stateVersionMutex.RUnlock()
	return len(fake.stateVersionArgsForCall)
}

func (fake *FakeWorkspace) StateVersionCalls(stub func() (*version.Version, error)) {
	fake.stateVersionMutex.Lock()
	defer fake.stateVersionMutex.Unlock()
	fake.StateVersionStub = stub
}

func (fake *FakeWorkspace) StateVersionReturns(result1 *version.Version, result2 error) {
	fake.stateVersionMutex.Lock()
	defer fake.stateVersionMutex.Unlock()
	fake.StateVersionStub = nil
	fake.stateVersionReturns = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) StateVersionReturnsOnCall(i int, result1 *version.Version, result2 error) {
	fake.stateVersionMutex.Lock()
	defer fake.stateVersionMutex.Unlock()
	fake.StateVersionStub = nil
	if fake.stateVersionReturnsOnCall == nil {
		fake.stateVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Version
			result2 error
		})
	}
	fake.stateVersionReturnsOnCall[i] = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkspace) UpdateInstanceConfiguration(arg1 map[string]interface{}) error {
	fake.updateInstanceConfigurationMutex.Lock()
	ret, specificReturn := fake.updateInstanceConfigurationReturnsOnCall[len(fake.updateInstanceConfigurationArgsForCall)]
	fake.updateInstanceConfigurationArgsForCall = append(fake.updateInstanceConfigurationArgsForCall, struct {
		arg1 map[string]interface{}
	}{arg1})
	stub := fake.UpdateInstanceConfigurationStub
	fakeReturns := fake.updateInstanceConfigurationReturns
	fake.recordInvocation("UpdateInstanceConfiguration", []interface{}{arg1})
	fake.updateInstanceConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWorkspace) UpdateInstanceConfigurationCallCount() int {
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	return len(fake.updateInstanceConfigurationArgsForCall)
}

func (fake *FakeWorkspace) UpdateInstanceConfigurationCalls(stub func(map[string]interface{}) error) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = stub
}

func (fake *FakeWorkspace) UpdateInstanceConfigurationArgsForCall(i int) map[string]interface{} {
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	argsForCall := fake.updateInstanceConfigurationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkspace) UpdateInstanceConfigurationReturns(result1 error) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = nil
	fake.updateInstanceConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkspace) UpdateInstanceConfigurationReturnsOnCall(i int, result1 error) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = nil
	if fake.updateInstanceConfigurationReturnsOnCall == nil {
		fake.updateInstanceConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateInstanceConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkspace) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.moduleDefinitionsMutex.RLock()
	defer fake.moduleDefinitionsMutex.RUnlock()
	fake.moduleInstancesMutex.RLock()
	defer fake.moduleInstancesMutex.RUnlock()
	fake.outputsMutex.RLock()
	defer fake.outputsMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	fake.stateVersionMutex.RLock()
	defer fake.stateVersionMutex.RUnlock()
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkspace) recordInvocation(key string, args []interface{}) {
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

var _ tf.Workspace = new(FakeWorkspace)
