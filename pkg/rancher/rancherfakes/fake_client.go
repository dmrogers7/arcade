// Code generated by counterfeiter. DO NOT EDIT.
package rancherfakes

import (
	"context"
	"net/http"
	"sync"

	"github.homedepot.com/cd/arcade/pkg/rancher"
)

type FakeClient struct {
	NewTokenStub        func(context.Context) (rancher.KubeconfigToken, error)
	newTokenMutex       sync.RWMutex
	newTokenArgsForCall []struct {
		arg1 context.Context
	}
	newTokenReturns struct {
		result1 rancher.KubeconfigToken
		result2 error
	}
	newTokenReturnsOnCall map[int]struct {
		result1 rancher.KubeconfigToken
		result2 error
	}
	WithPasswordStub        func(string)
	withPasswordMutex       sync.RWMutex
	withPasswordArgsForCall []struct {
		arg1 string
	}
	WithTransportStub        func(*http.Transport)
	withTransportMutex       sync.RWMutex
	withTransportArgsForCall []struct {
		arg1 *http.Transport
	}
	WithURLStub        func(string)
	withURLMutex       sync.RWMutex
	withURLArgsForCall []struct {
		arg1 string
	}
	WithUsernameStub        func(string)
	withUsernameMutex       sync.RWMutex
	withUsernameArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) NewToken(arg1 context.Context) (rancher.KubeconfigToken, error) {
	fake.newTokenMutex.Lock()
	ret, specificReturn := fake.newTokenReturnsOnCall[len(fake.newTokenArgsForCall)]
	fake.newTokenArgsForCall = append(fake.newTokenArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("NewToken", []interface{}{arg1})
	fake.newTokenMutex.Unlock()
	if fake.NewTokenStub != nil {
		return fake.NewTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) NewTokenCallCount() int {
	fake.newTokenMutex.RLock()
	defer fake.newTokenMutex.RUnlock()
	return len(fake.newTokenArgsForCall)
}

func (fake *FakeClient) NewTokenCalls(stub func(context.Context) (rancher.KubeconfigToken, error)) {
	fake.newTokenMutex.Lock()
	defer fake.newTokenMutex.Unlock()
	fake.NewTokenStub = stub
}

func (fake *FakeClient) NewTokenArgsForCall(i int) context.Context {
	fake.newTokenMutex.RLock()
	defer fake.newTokenMutex.RUnlock()
	argsForCall := fake.newTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) NewTokenReturns(result1 rancher.KubeconfigToken, result2 error) {
	fake.newTokenMutex.Lock()
	defer fake.newTokenMutex.Unlock()
	fake.NewTokenStub = nil
	fake.newTokenReturns = struct {
		result1 rancher.KubeconfigToken
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) NewTokenReturnsOnCall(i int, result1 rancher.KubeconfigToken, result2 error) {
	fake.newTokenMutex.Lock()
	defer fake.newTokenMutex.Unlock()
	fake.NewTokenStub = nil
	if fake.newTokenReturnsOnCall == nil {
		fake.newTokenReturnsOnCall = make(map[int]struct {
			result1 rancher.KubeconfigToken
			result2 error
		})
	}
	fake.newTokenReturnsOnCall[i] = struct {
		result1 rancher.KubeconfigToken
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) WithPassword(arg1 string) {
	fake.withPasswordMutex.Lock()
	fake.withPasswordArgsForCall = append(fake.withPasswordArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithPassword", []interface{}{arg1})
	fake.withPasswordMutex.Unlock()
	if fake.WithPasswordStub != nil {
		fake.WithPasswordStub(arg1)
	}
}

func (fake *FakeClient) WithPasswordCallCount() int {
	fake.withPasswordMutex.RLock()
	defer fake.withPasswordMutex.RUnlock()
	return len(fake.withPasswordArgsForCall)
}

func (fake *FakeClient) WithPasswordCalls(stub func(string)) {
	fake.withPasswordMutex.Lock()
	defer fake.withPasswordMutex.Unlock()
	fake.WithPasswordStub = stub
}

func (fake *FakeClient) WithPasswordArgsForCall(i int) string {
	fake.withPasswordMutex.RLock()
	defer fake.withPasswordMutex.RUnlock()
	argsForCall := fake.withPasswordArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) WithTransport(arg1 *http.Transport) {
	fake.withTransportMutex.Lock()
	fake.withTransportArgsForCall = append(fake.withTransportArgsForCall, struct {
		arg1 *http.Transport
	}{arg1})
	fake.recordInvocation("WithTransport", []interface{}{arg1})
	fake.withTransportMutex.Unlock()
	if fake.WithTransportStub != nil {
		fake.WithTransportStub(arg1)
	}
}

func (fake *FakeClient) WithTransportCallCount() int {
	fake.withTransportMutex.RLock()
	defer fake.withTransportMutex.RUnlock()
	return len(fake.withTransportArgsForCall)
}

func (fake *FakeClient) WithTransportCalls(stub func(*http.Transport)) {
	fake.withTransportMutex.Lock()
	defer fake.withTransportMutex.Unlock()
	fake.WithTransportStub = stub
}

func (fake *FakeClient) WithTransportArgsForCall(i int) *http.Transport {
	fake.withTransportMutex.RLock()
	defer fake.withTransportMutex.RUnlock()
	argsForCall := fake.withTransportArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) WithURL(arg1 string) {
	fake.withURLMutex.Lock()
	fake.withURLArgsForCall = append(fake.withURLArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithURL", []interface{}{arg1})
	fake.withURLMutex.Unlock()
	if fake.WithURLStub != nil {
		fake.WithURLStub(arg1)
	}
}

func (fake *FakeClient) WithURLCallCount() int {
	fake.withURLMutex.RLock()
	defer fake.withURLMutex.RUnlock()
	return len(fake.withURLArgsForCall)
}

func (fake *FakeClient) WithURLCalls(stub func(string)) {
	fake.withURLMutex.Lock()
	defer fake.withURLMutex.Unlock()
	fake.WithURLStub = stub
}

func (fake *FakeClient) WithURLArgsForCall(i int) string {
	fake.withURLMutex.RLock()
	defer fake.withURLMutex.RUnlock()
	argsForCall := fake.withURLArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) WithUsername(arg1 string) {
	fake.withUsernameMutex.Lock()
	fake.withUsernameArgsForCall = append(fake.withUsernameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithUsername", []interface{}{arg1})
	fake.withUsernameMutex.Unlock()
	if fake.WithUsernameStub != nil {
		fake.WithUsernameStub(arg1)
	}
}

func (fake *FakeClient) WithUsernameCallCount() int {
	fake.withUsernameMutex.RLock()
	defer fake.withUsernameMutex.RUnlock()
	return len(fake.withUsernameArgsForCall)
}

func (fake *FakeClient) WithUsernameCalls(stub func(string)) {
	fake.withUsernameMutex.Lock()
	defer fake.withUsernameMutex.Unlock()
	fake.WithUsernameStub = stub
}

func (fake *FakeClient) WithUsernameArgsForCall(i int) string {
	fake.withUsernameMutex.RLock()
	defer fake.withUsernameMutex.RUnlock()
	argsForCall := fake.withUsernameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newTokenMutex.RLock()
	defer fake.newTokenMutex.RUnlock()
	fake.withPasswordMutex.RLock()
	defer fake.withPasswordMutex.RUnlock()
	fake.withTransportMutex.RLock()
	defer fake.withTransportMutex.RUnlock()
	fake.withURLMutex.RLock()
	defer fake.withURLMutex.RUnlock()
	fake.withUsernameMutex.RLock()
	defer fake.withUsernameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ rancher.Client = new(FakeClient)
