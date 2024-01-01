// Copyright 2023 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package ecache is a generated GoMock package.
package ecache

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

func NewMockNamespaceCache(cache *MockCache, namespace string) *NamespaceCache {
	return &NamespaceCache{
		C:         cache,
		Namespace: namespace,
	}
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// DecrBy mocks base method.
func (m *MockCache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", ctx, key, value)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockCacheMockRecorder) DecrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockCache)(nil).DecrBy), ctx, key, value)
}

// Delete mocks base method.
func (m *MockCache) Delete(ctx context.Context, key ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range key {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(ctx interface{}, key ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, key...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockCache) Get(ctx context.Context, key string) Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(Value)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), ctx, key)
}

// GetSet mocks base method.
func (m *MockCache) GetSet(ctx context.Context, key, val string) Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSet", ctx, key, val)
	ret0, _ := ret[0].(Value)
	return ret0
}

// GetSet indicates an expected call of GetSet.
func (mr *MockCacheMockRecorder) GetSet(ctx, key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSet", reflect.TypeOf((*MockCache)(nil).GetSet), ctx, key, val)
}

// IncrBy mocks base method.
func (m *MockCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", ctx, key, value)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockCacheMockRecorder) IncrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockCache)(nil).IncrBy), ctx, key, value)
}

// IncrByFloat mocks base method.
func (m *MockCache) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrByFloat", ctx, key, value)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrByFloat indicates an expected call of IncrByFloat.
func (mr *MockCacheMockRecorder) IncrByFloat(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrByFloat", reflect.TypeOf((*MockCache)(nil).IncrByFloat), ctx, key, value)
}

// LPop mocks base method.
func (m *MockCache) LPop(ctx context.Context, key string) Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", ctx, key)
	ret0, _ := ret[0].(Value)
	return ret0
}

// LPop indicates an expected call of LPop.
func (mr *MockCacheMockRecorder) LPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockCache)(nil).LPop), ctx, key)
}

// LPush mocks base method.
func (m *MockCache) LPush(ctx context.Context, key string, val ...any) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range val {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LPush indicates an expected call of LPush.
func (mr *MockCacheMockRecorder) LPush(ctx, key interface{}, val ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, val...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockCache)(nil).LPush), varargs...)
}

// SAdd mocks base method.
func (m *MockCache) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SAdd indicates an expected call of SAdd.
func (mr *MockCacheMockRecorder) SAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockCache)(nil).SAdd), varargs...)
}

// SRem mocks base method.
func (m *MockCache) SRem(ctx context.Context, key string, members ...any) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SRem indicates an expected call of SRem.
func (mr *MockCacheMockRecorder) SRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockCache)(nil).SRem), varargs...)
}

// Set mocks base method.
func (m *MockCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, val, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(ctx, key, val, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), ctx, key, val, expiration)
}

// SetNX mocks base method.
func (m *MockCache) SetNX(ctx context.Context, key string, val any, expiration time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, val, expiration)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockCacheMockRecorder) SetNX(ctx, key, val, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockCache)(nil).SetNX), ctx, key, val, expiration)
}