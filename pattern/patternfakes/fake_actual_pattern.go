// Code generated by counterfeiter. DO NOT EDIT.
package patternfakes

import (
	"sync"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/tuple"
)

type FakeActualPattern struct {
	PatternAtStub        func(tuple.Tuple) color.Color
	patternAtMutex       sync.RWMutex
	patternAtArgsForCall []struct {
		arg1 tuple.Tuple
	}
	patternAtReturns struct {
		result1 color.Color
	}
	patternAtReturnsOnCall map[int]struct {
		result1 color.Color
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActualPattern) PatternAt(arg1 tuple.Tuple) color.Color {
	fake.patternAtMutex.Lock()
	ret, specificReturn := fake.patternAtReturnsOnCall[len(fake.patternAtArgsForCall)]
	fake.patternAtArgsForCall = append(fake.patternAtArgsForCall, struct {
		arg1 tuple.Tuple
	}{arg1})
	fake.recordInvocation("PatternAt", []interface{}{arg1})
	fake.patternAtMutex.Unlock()
	if fake.PatternAtStub != nil {
		return fake.PatternAtStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.patternAtReturns
	return fakeReturns.result1
}

func (fake *FakeActualPattern) PatternAtCallCount() int {
	fake.patternAtMutex.RLock()
	defer fake.patternAtMutex.RUnlock()
	return len(fake.patternAtArgsForCall)
}

func (fake *FakeActualPattern) PatternAtCalls(stub func(tuple.Tuple) color.Color) {
	fake.patternAtMutex.Lock()
	defer fake.patternAtMutex.Unlock()
	fake.PatternAtStub = stub
}

func (fake *FakeActualPattern) PatternAtArgsForCall(i int) tuple.Tuple {
	fake.patternAtMutex.RLock()
	defer fake.patternAtMutex.RUnlock()
	argsForCall := fake.patternAtArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActualPattern) PatternAtReturns(result1 color.Color) {
	fake.patternAtMutex.Lock()
	defer fake.patternAtMutex.Unlock()
	fake.PatternAtStub = nil
	fake.patternAtReturns = struct {
		result1 color.Color
	}{result1}
}

func (fake *FakeActualPattern) PatternAtReturnsOnCall(i int, result1 color.Color) {
	fake.patternAtMutex.Lock()
	defer fake.patternAtMutex.Unlock()
	fake.PatternAtStub = nil
	if fake.patternAtReturnsOnCall == nil {
		fake.patternAtReturnsOnCall = make(map[int]struct {
			result1 color.Color
		})
	}
	fake.patternAtReturnsOnCall[i] = struct {
		result1 color.Color
	}{result1}
}

func (fake *FakeActualPattern) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.patternAtMutex.RLock()
	defer fake.patternAtMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeActualPattern) recordInvocation(key string, args []interface{}) {
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

var _ pattern.ActualPattern = new(FakeActualPattern)