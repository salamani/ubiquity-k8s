/**
 * Copyright 2017 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/resources"
)

type FakeMounter struct {
	MountStub        func(mountRequest resources.MountRequest) (string, error)
	mountMutex       sync.RWMutex
	mountArgsForCall []struct {
		mountRequest resources.MountRequest
	}
	mountReturns struct {
		result1 string
		result2 error
	}
	mountReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UnmountStub        func(unmountRequest resources.UnmountRequest) error
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		unmountRequest resources.UnmountRequest
	}
	unmountReturns struct {
		result1 error
	}
	unmountReturnsOnCall map[int]struct {
		result1 error
	}
	ActionAfterDetachStub        func(request resources.AfterDetachRequest) error
	actionAfterDetachMutex       sync.RWMutex
	actionAfterDetachArgsForCall []struct {
		request resources.AfterDetachRequest
	}
	actionAfterDetachReturns struct {
		result1 error
	}
	actionAfterDetachReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMounter) Mount(mountRequest resources.MountRequest) (string, error) {
	fake.mountMutex.Lock()
	ret, specificReturn := fake.mountReturnsOnCall[len(fake.mountArgsForCall)]
	fake.mountArgsForCall = append(fake.mountArgsForCall, struct {
		mountRequest resources.MountRequest
	}{mountRequest})
	fake.recordInvocation("Mount", []interface{}{mountRequest})
	fake.mountMutex.Unlock()
	if fake.MountStub != nil {
		return fake.MountStub(mountRequest)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.mountReturns.result1, fake.mountReturns.result2
}

func (fake *FakeMounter) MountCallCount() int {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return len(fake.mountArgsForCall)
}

func (fake *FakeMounter) MountArgsForCall(i int) resources.MountRequest {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return fake.mountArgsForCall[i].mountRequest
}

func (fake *FakeMounter) MountReturns(result1 string, result2 error) {
	fake.MountStub = nil
	fake.mountReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) MountReturnsOnCall(i int, result1 string, result2 error) {
	fake.MountStub = nil
	if fake.mountReturnsOnCall == nil {
		fake.mountReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.mountReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMounter) Unmount(unmountRequest resources.UnmountRequest) error {
	fake.unmountMutex.Lock()
	ret, specificReturn := fake.unmountReturnsOnCall[len(fake.unmountArgsForCall)]
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		unmountRequest resources.UnmountRequest
	}{unmountRequest})
	fake.recordInvocation("Unmount", []interface{}{unmountRequest})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(unmountRequest)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unmountReturns.result1
}

func (fake *FakeMounter) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeMounter) UnmountArgsForCall(i int) resources.UnmountRequest {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].unmountRequest
}

func (fake *FakeMounter) UnmountReturns(result1 error) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) UnmountReturnsOnCall(i int, result1 error) {
	fake.UnmountStub = nil
	if fake.unmountReturnsOnCall == nil {
		fake.unmountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) ActionAfterDetach(request resources.AfterDetachRequest) error {
	fake.actionAfterDetachMutex.Lock()
	ret, specificReturn := fake.actionAfterDetachReturnsOnCall[len(fake.actionAfterDetachArgsForCall)]
	fake.actionAfterDetachArgsForCall = append(fake.actionAfterDetachArgsForCall, struct {
		request resources.AfterDetachRequest
	}{request})
	fake.recordInvocation("ActionAfterDetach", []interface{}{request})
	fake.actionAfterDetachMutex.Unlock()
	if fake.ActionAfterDetachStub != nil {
		return fake.ActionAfterDetachStub(request)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.actionAfterDetachReturns.result1
}

func (fake *FakeMounter) ActionAfterDetachCallCount() int {
	fake.actionAfterDetachMutex.RLock()
	defer fake.actionAfterDetachMutex.RUnlock()
	return len(fake.actionAfterDetachArgsForCall)
}

func (fake *FakeMounter) ActionAfterDetachArgsForCall(i int) resources.AfterDetachRequest {
	fake.actionAfterDetachMutex.RLock()
	defer fake.actionAfterDetachMutex.RUnlock()
	return fake.actionAfterDetachArgsForCall[i].request
}

func (fake *FakeMounter) ActionAfterDetachReturns(result1 error) {
	fake.ActionAfterDetachStub = nil
	fake.actionAfterDetachReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) ActionAfterDetachReturnsOnCall(i int, result1 error) {
	fake.ActionAfterDetachStub = nil
	if fake.actionAfterDetachReturnsOnCall == nil {
		fake.actionAfterDetachReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.actionAfterDetachReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	fake.actionAfterDetachMutex.RLock()
	defer fake.actionAfterDetachMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMounter) recordInvocation(key string, args []interface{}) {
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

var _ resources.Mounter = new(FakeMounter)