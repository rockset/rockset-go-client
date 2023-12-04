// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"net/http"
	"sync"

	"github.com/rockset/rockset-go-client/openapi"
)

type FakeSharedLambdasApi struct {
	ExecutePublicQueryLambdaWithParamsStub        func(context.Context, string) openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	executePublicQueryLambdaWithParamsMutex       sync.RWMutex
	executePublicQueryLambdaWithParamsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	executePublicQueryLambdaWithParamsReturns struct {
		result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}
	executePublicQueryLambdaWithParamsReturnsOnCall map[int]struct {
		result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}
	ExecutePublicQueryLambdaWithParamsExecuteStub        func(openapi.ApiExecutePublicQueryLambdaWithParamsRequest) (*openapi.QueryResponse, *http.Response, error)
	executePublicQueryLambdaWithParamsExecuteMutex       sync.RWMutex
	executePublicQueryLambdaWithParamsExecuteArgsForCall []struct {
		arg1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}
	executePublicQueryLambdaWithParamsExecuteReturns struct {
		result1 *openapi.QueryResponse
		result2 *http.Response
		result3 error
	}
	executePublicQueryLambdaWithParamsExecuteReturnsOnCall map[int]struct {
		result1 *openapi.QueryResponse
		result2 *http.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParams(arg1 context.Context, arg2 string) openapi.ApiExecutePublicQueryLambdaWithParamsRequest {
	fake.executePublicQueryLambdaWithParamsMutex.Lock()
	ret, specificReturn := fake.executePublicQueryLambdaWithParamsReturnsOnCall[len(fake.executePublicQueryLambdaWithParamsArgsForCall)]
	fake.executePublicQueryLambdaWithParamsArgsForCall = append(fake.executePublicQueryLambdaWithParamsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.ExecutePublicQueryLambdaWithParamsStub
	fakeReturns := fake.executePublicQueryLambdaWithParamsReturns
	fake.recordInvocation("ExecutePublicQueryLambdaWithParams", []interface{}{arg1, arg2})
	fake.executePublicQueryLambdaWithParamsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsCallCount() int {
	fake.executePublicQueryLambdaWithParamsMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsMutex.RUnlock()
	return len(fake.executePublicQueryLambdaWithParamsArgsForCall)
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsCalls(stub func(context.Context, string) openapi.ApiExecutePublicQueryLambdaWithParamsRequest) {
	fake.executePublicQueryLambdaWithParamsMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsStub = stub
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsArgsForCall(i int) (context.Context, string) {
	fake.executePublicQueryLambdaWithParamsMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsMutex.RUnlock()
	argsForCall := fake.executePublicQueryLambdaWithParamsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsReturns(result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest) {
	fake.executePublicQueryLambdaWithParamsMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsStub = nil
	fake.executePublicQueryLambdaWithParamsReturns = struct {
		result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}{result1}
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsReturnsOnCall(i int, result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest) {
	fake.executePublicQueryLambdaWithParamsMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsStub = nil
	if fake.executePublicQueryLambdaWithParamsReturnsOnCall == nil {
		fake.executePublicQueryLambdaWithParamsReturnsOnCall = make(map[int]struct {
			result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
		})
	}
	fake.executePublicQueryLambdaWithParamsReturnsOnCall[i] = struct {
		result1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}{result1}
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecute(arg1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest) (*openapi.QueryResponse, *http.Response, error) {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.Lock()
	ret, specificReturn := fake.executePublicQueryLambdaWithParamsExecuteReturnsOnCall[len(fake.executePublicQueryLambdaWithParamsExecuteArgsForCall)]
	fake.executePublicQueryLambdaWithParamsExecuteArgsForCall = append(fake.executePublicQueryLambdaWithParamsExecuteArgsForCall, struct {
		arg1 openapi.ApiExecutePublicQueryLambdaWithParamsRequest
	}{arg1})
	stub := fake.ExecutePublicQueryLambdaWithParamsExecuteStub
	fakeReturns := fake.executePublicQueryLambdaWithParamsExecuteReturns
	fake.recordInvocation("ExecutePublicQueryLambdaWithParamsExecute", []interface{}{arg1})
	fake.executePublicQueryLambdaWithParamsExecuteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecuteCallCount() int {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.RUnlock()
	return len(fake.executePublicQueryLambdaWithParamsExecuteArgsForCall)
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecuteCalls(stub func(openapi.ApiExecutePublicQueryLambdaWithParamsRequest) (*openapi.QueryResponse, *http.Response, error)) {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsExecuteStub = stub
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecuteArgsForCall(i int) openapi.ApiExecutePublicQueryLambdaWithParamsRequest {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.RUnlock()
	argsForCall := fake.executePublicQueryLambdaWithParamsExecuteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecuteReturns(result1 *openapi.QueryResponse, result2 *http.Response, result3 error) {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsExecuteStub = nil
	fake.executePublicQueryLambdaWithParamsExecuteReturns = struct {
		result1 *openapi.QueryResponse
		result2 *http.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSharedLambdasApi) ExecutePublicQueryLambdaWithParamsExecuteReturnsOnCall(i int, result1 *openapi.QueryResponse, result2 *http.Response, result3 error) {
	fake.executePublicQueryLambdaWithParamsExecuteMutex.Lock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.Unlock()
	fake.ExecutePublicQueryLambdaWithParamsExecuteStub = nil
	if fake.executePublicQueryLambdaWithParamsExecuteReturnsOnCall == nil {
		fake.executePublicQueryLambdaWithParamsExecuteReturnsOnCall = make(map[int]struct {
			result1 *openapi.QueryResponse
			result2 *http.Response
			result3 error
		})
	}
	fake.executePublicQueryLambdaWithParamsExecuteReturnsOnCall[i] = struct {
		result1 *openapi.QueryResponse
		result2 *http.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSharedLambdasApi) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executePublicQueryLambdaWithParamsMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsMutex.RUnlock()
	fake.executePublicQueryLambdaWithParamsExecuteMutex.RLock()
	defer fake.executePublicQueryLambdaWithParamsExecuteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSharedLambdasApi) recordInvocation(key string, args []interface{}) {
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

var _ openapi.SharedLambdasApi = new(FakeSharedLambdasApi)