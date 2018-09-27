// Code generated by counterfeiter. DO NOT EDIT.
package uaafakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/uaa"
)

type FakeUAA struct {
	PromptsStub        func() ([]uaa.Prompt, error)
	promptsMutex       sync.RWMutex
	promptsArgsForCall []struct{}
	promptsReturns     struct {
		result1 []uaa.Prompt
		result2 error
	}
	promptsReturnsOnCall map[int]struct {
		result1 []uaa.Prompt
		result2 error
	}
	RefreshTokenGrantStub        func(string) (uaa.AccessToken, error)
	refreshTokenGrantMutex       sync.RWMutex
	refreshTokenGrantArgsForCall []struct {
		arg1 string
	}
	refreshTokenGrantReturns struct {
		result1 uaa.AccessToken
		result2 error
	}
	refreshTokenGrantReturnsOnCall map[int]struct {
		result1 uaa.AccessToken
		result2 error
	}
	ClientCredentialsGrantStub        func() (uaa.AccessToken, error)
	clientCredentialsGrantMutex       sync.RWMutex
	clientCredentialsGrantArgsForCall []struct{}
	clientCredentialsGrantReturns     struct {
		result1 uaa.AccessToken
		result2 error
	}
	clientCredentialsGrantReturnsOnCall map[int]struct {
		result1 uaa.AccessToken
		result2 error
	}
	OwnerPasswordCredentialsGrantStub        func([]uaa.PromptAnswer) (uaa.AccessToken, error)
	ownerPasswordCredentialsGrantMutex       sync.RWMutex
	ownerPasswordCredentialsGrantArgsForCall []struct {
		arg1 []uaa.PromptAnswer
	}
	ownerPasswordCredentialsGrantReturns struct {
		result1 uaa.AccessToken
		result2 error
	}
	ownerPasswordCredentialsGrantReturnsOnCall map[int]struct {
		result1 uaa.AccessToken
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAA) Prompts() ([]uaa.Prompt, error) {
	fake.promptsMutex.Lock()
	ret, specificReturn := fake.promptsReturnsOnCall[len(fake.promptsArgsForCall)]
	fake.promptsArgsForCall = append(fake.promptsArgsForCall, struct{}{})
	fake.recordInvocation("Prompts", []interface{}{})
	fake.promptsMutex.Unlock()
	if fake.PromptsStub != nil {
		return fake.PromptsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.promptsReturns.result1, fake.promptsReturns.result2
}

func (fake *FakeUAA) PromptsCallCount() int {
	fake.promptsMutex.RLock()
	defer fake.promptsMutex.RUnlock()
	return len(fake.promptsArgsForCall)
}

func (fake *FakeUAA) PromptsReturns(result1 []uaa.Prompt, result2 error) {
	fake.PromptsStub = nil
	fake.promptsReturns = struct {
		result1 []uaa.Prompt
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) PromptsReturnsOnCall(i int, result1 []uaa.Prompt, result2 error) {
	fake.PromptsStub = nil
	if fake.promptsReturnsOnCall == nil {
		fake.promptsReturnsOnCall = make(map[int]struct {
			result1 []uaa.Prompt
			result2 error
		})
	}
	fake.promptsReturnsOnCall[i] = struct {
		result1 []uaa.Prompt
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) RefreshTokenGrant(arg1 string) (uaa.AccessToken, error) {
	fake.refreshTokenGrantMutex.Lock()
	ret, specificReturn := fake.refreshTokenGrantReturnsOnCall[len(fake.refreshTokenGrantArgsForCall)]
	fake.refreshTokenGrantArgsForCall = append(fake.refreshTokenGrantArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RefreshTokenGrant", []interface{}{arg1})
	fake.refreshTokenGrantMutex.Unlock()
	if fake.RefreshTokenGrantStub != nil {
		return fake.RefreshTokenGrantStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.refreshTokenGrantReturns.result1, fake.refreshTokenGrantReturns.result2
}

func (fake *FakeUAA) RefreshTokenGrantCallCount() int {
	fake.refreshTokenGrantMutex.RLock()
	defer fake.refreshTokenGrantMutex.RUnlock()
	return len(fake.refreshTokenGrantArgsForCall)
}

func (fake *FakeUAA) RefreshTokenGrantArgsForCall(i int) string {
	fake.refreshTokenGrantMutex.RLock()
	defer fake.refreshTokenGrantMutex.RUnlock()
	return fake.refreshTokenGrantArgsForCall[i].arg1
}

func (fake *FakeUAA) RefreshTokenGrantReturns(result1 uaa.AccessToken, result2 error) {
	fake.RefreshTokenGrantStub = nil
	fake.refreshTokenGrantReturns = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) RefreshTokenGrantReturnsOnCall(i int, result1 uaa.AccessToken, result2 error) {
	fake.RefreshTokenGrantStub = nil
	if fake.refreshTokenGrantReturnsOnCall == nil {
		fake.refreshTokenGrantReturnsOnCall = make(map[int]struct {
			result1 uaa.AccessToken
			result2 error
		})
	}
	fake.refreshTokenGrantReturnsOnCall[i] = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) ClientCredentialsGrant() (uaa.AccessToken, error) {
	fake.clientCredentialsGrantMutex.Lock()
	ret, specificReturn := fake.clientCredentialsGrantReturnsOnCall[len(fake.clientCredentialsGrantArgsForCall)]
	fake.clientCredentialsGrantArgsForCall = append(fake.clientCredentialsGrantArgsForCall, struct{}{})
	fake.recordInvocation("ClientCredentialsGrant", []interface{}{})
	fake.clientCredentialsGrantMutex.Unlock()
	if fake.ClientCredentialsGrantStub != nil {
		return fake.ClientCredentialsGrantStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.clientCredentialsGrantReturns.result1, fake.clientCredentialsGrantReturns.result2
}

func (fake *FakeUAA) ClientCredentialsGrantCallCount() int {
	fake.clientCredentialsGrantMutex.RLock()
	defer fake.clientCredentialsGrantMutex.RUnlock()
	return len(fake.clientCredentialsGrantArgsForCall)
}

func (fake *FakeUAA) ClientCredentialsGrantReturns(result1 uaa.AccessToken, result2 error) {
	fake.ClientCredentialsGrantStub = nil
	fake.clientCredentialsGrantReturns = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) ClientCredentialsGrantReturnsOnCall(i int, result1 uaa.AccessToken, result2 error) {
	fake.ClientCredentialsGrantStub = nil
	if fake.clientCredentialsGrantReturnsOnCall == nil {
		fake.clientCredentialsGrantReturnsOnCall = make(map[int]struct {
			result1 uaa.AccessToken
			result2 error
		})
	}
	fake.clientCredentialsGrantReturnsOnCall[i] = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) OwnerPasswordCredentialsGrant(arg1 []uaa.PromptAnswer) (uaa.AccessToken, error) {
	var arg1Copy []uaa.PromptAnswer
	if arg1 != nil {
		arg1Copy = make([]uaa.PromptAnswer, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.ownerPasswordCredentialsGrantMutex.Lock()
	ret, specificReturn := fake.ownerPasswordCredentialsGrantReturnsOnCall[len(fake.ownerPasswordCredentialsGrantArgsForCall)]
	fake.ownerPasswordCredentialsGrantArgsForCall = append(fake.ownerPasswordCredentialsGrantArgsForCall, struct {
		arg1 []uaa.PromptAnswer
	}{arg1Copy})
	fake.recordInvocation("OwnerPasswordCredentialsGrant", []interface{}{arg1Copy})
	fake.ownerPasswordCredentialsGrantMutex.Unlock()
	if fake.OwnerPasswordCredentialsGrantStub != nil {
		return fake.OwnerPasswordCredentialsGrantStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ownerPasswordCredentialsGrantReturns.result1, fake.ownerPasswordCredentialsGrantReturns.result2
}

func (fake *FakeUAA) OwnerPasswordCredentialsGrantCallCount() int {
	fake.ownerPasswordCredentialsGrantMutex.RLock()
	defer fake.ownerPasswordCredentialsGrantMutex.RUnlock()
	return len(fake.ownerPasswordCredentialsGrantArgsForCall)
}

func (fake *FakeUAA) OwnerPasswordCredentialsGrantArgsForCall(i int) []uaa.PromptAnswer {
	fake.ownerPasswordCredentialsGrantMutex.RLock()
	defer fake.ownerPasswordCredentialsGrantMutex.RUnlock()
	return fake.ownerPasswordCredentialsGrantArgsForCall[i].arg1
}

func (fake *FakeUAA) OwnerPasswordCredentialsGrantReturns(result1 uaa.AccessToken, result2 error) {
	fake.OwnerPasswordCredentialsGrantStub = nil
	fake.ownerPasswordCredentialsGrantReturns = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) OwnerPasswordCredentialsGrantReturnsOnCall(i int, result1 uaa.AccessToken, result2 error) {
	fake.OwnerPasswordCredentialsGrantStub = nil
	if fake.ownerPasswordCredentialsGrantReturnsOnCall == nil {
		fake.ownerPasswordCredentialsGrantReturnsOnCall = make(map[int]struct {
			result1 uaa.AccessToken
			result2 error
		})
	}
	fake.ownerPasswordCredentialsGrantReturnsOnCall[i] = struct {
		result1 uaa.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.promptsMutex.RLock()
	defer fake.promptsMutex.RUnlock()
	fake.refreshTokenGrantMutex.RLock()
	defer fake.refreshTokenGrantMutex.RUnlock()
	fake.clientCredentialsGrantMutex.RLock()
	defer fake.clientCredentialsGrantMutex.RUnlock()
	fake.ownerPasswordCredentialsGrantMutex.RLock()
	defer fake.ownerPasswordCredentialsGrantMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAA) recordInvocation(key string, args []interface{}) {
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

var _ uaa.UAA = new(FakeUAA)
