// This file was generated by counterfeiter
package pluginactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakeConfig struct {
	PluginHomeStub        func() string
	pluginHomeMutex       sync.RWMutex
	pluginHomeArgsForCall []struct{}
	pluginHomeReturns     struct {
		result1 string
	}
	pluginHomeReturnsOnCall map[int]struct {
		result1 string
	}
	PluginsStub        func() map[string]configv3.Plugin
	pluginsMutex       sync.RWMutex
	pluginsArgsForCall []struct{}
	pluginsReturns     struct {
		result1 map[string]configv3.Plugin
	}
	pluginsReturnsOnCall map[int]struct {
		result1 map[string]configv3.Plugin
	}
	RemovePluginStub        func(string)
	removePluginMutex       sync.RWMutex
	removePluginArgsForCall []struct {
		arg1 string
	}
	WritePluginConfigStub        func() error
	writePluginConfigMutex       sync.RWMutex
	writePluginConfigArgsForCall []struct{}
	writePluginConfigReturns     struct {
		result1 error
	}
	writePluginConfigReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) PluginHome() string {
	fake.pluginHomeMutex.Lock()
	ret, specificReturn := fake.pluginHomeReturnsOnCall[len(fake.pluginHomeArgsForCall)]
	fake.pluginHomeArgsForCall = append(fake.pluginHomeArgsForCall, struct{}{})
	fake.recordInvocation("PluginHome", []interface{}{})
	fake.pluginHomeMutex.Unlock()
	if fake.PluginHomeStub != nil {
		return fake.PluginHomeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pluginHomeReturns.result1
}

func (fake *FakeConfig) PluginHomeCallCount() int {
	fake.pluginHomeMutex.RLock()
	defer fake.pluginHomeMutex.RUnlock()
	return len(fake.pluginHomeArgsForCall)
}

func (fake *FakeConfig) PluginHomeReturns(result1 string) {
	fake.PluginHomeStub = nil
	fake.pluginHomeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) PluginHomeReturnsOnCall(i int, result1 string) {
	fake.PluginHomeStub = nil
	if fake.pluginHomeReturnsOnCall == nil {
		fake.pluginHomeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pluginHomeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Plugins() map[string]configv3.Plugin {
	fake.pluginsMutex.Lock()
	ret, specificReturn := fake.pluginsReturnsOnCall[len(fake.pluginsArgsForCall)]
	fake.pluginsArgsForCall = append(fake.pluginsArgsForCall, struct{}{})
	fake.recordInvocation("Plugins", []interface{}{})
	fake.pluginsMutex.Unlock()
	if fake.PluginsStub != nil {
		return fake.PluginsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pluginsReturns.result1
}

func (fake *FakeConfig) PluginsCallCount() int {
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	return len(fake.pluginsArgsForCall)
}

func (fake *FakeConfig) PluginsReturns(result1 map[string]configv3.Plugin) {
	fake.PluginsStub = nil
	fake.pluginsReturns = struct {
		result1 map[string]configv3.Plugin
	}{result1}
}

func (fake *FakeConfig) PluginsReturnsOnCall(i int, result1 map[string]configv3.Plugin) {
	fake.PluginsStub = nil
	if fake.pluginsReturnsOnCall == nil {
		fake.pluginsReturnsOnCall = make(map[int]struct {
			result1 map[string]configv3.Plugin
		})
	}
	fake.pluginsReturnsOnCall[i] = struct {
		result1 map[string]configv3.Plugin
	}{result1}
}

func (fake *FakeConfig) RemovePlugin(arg1 string) {
	fake.removePluginMutex.Lock()
	fake.removePluginArgsForCall = append(fake.removePluginArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemovePlugin", []interface{}{arg1})
	fake.removePluginMutex.Unlock()
	if fake.RemovePluginStub != nil {
		fake.RemovePluginStub(arg1)
	}
}

func (fake *FakeConfig) RemovePluginCallCount() int {
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	return len(fake.removePluginArgsForCall)
}

func (fake *FakeConfig) RemovePluginArgsForCall(i int) string {
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	return fake.removePluginArgsForCall[i].arg1
}

func (fake *FakeConfig) WritePluginConfig() error {
	fake.writePluginConfigMutex.Lock()
	ret, specificReturn := fake.writePluginConfigReturnsOnCall[len(fake.writePluginConfigArgsForCall)]
	fake.writePluginConfigArgsForCall = append(fake.writePluginConfigArgsForCall, struct{}{})
	fake.recordInvocation("WritePluginConfig", []interface{}{})
	fake.writePluginConfigMutex.Unlock()
	if fake.WritePluginConfigStub != nil {
		return fake.WritePluginConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writePluginConfigReturns.result1
}

func (fake *FakeConfig) WritePluginConfigCallCount() int {
	fake.writePluginConfigMutex.RLock()
	defer fake.writePluginConfigMutex.RUnlock()
	return len(fake.writePluginConfigArgsForCall)
}

func (fake *FakeConfig) WritePluginConfigReturns(result1 error) {
	fake.WritePluginConfigStub = nil
	fake.writePluginConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) WritePluginConfigReturnsOnCall(i int, result1 error) {
	fake.WritePluginConfigStub = nil
	if fake.writePluginConfigReturnsOnCall == nil {
		fake.writePluginConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writePluginConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pluginHomeMutex.RLock()
	defer fake.pluginHomeMutex.RUnlock()
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	fake.writePluginConfigMutex.RLock()
	defer fake.writePluginConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ pluginaction.Config = new(FakeConfig)
