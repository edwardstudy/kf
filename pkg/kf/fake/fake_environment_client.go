// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/kf/pkg/kf/fake (interfaces: EnvironmentClient)

// Package fake is a generated GoMock package.
package fake

import (
	kf "github.com/GoogleCloudPlatform/kf/pkg/kf"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// FakeEnvironmentClient is a mock of EnvironmentClient interface
type FakeEnvironmentClient struct {
	ctrl     *gomock.Controller
	recorder *FakeEnvironmentClientMockRecorder
}

// FakeEnvironmentClientMockRecorder is the mock recorder for FakeEnvironmentClient
type FakeEnvironmentClientMockRecorder struct {
	mock *FakeEnvironmentClient
}

// NewFakeEnvironmentClient creates a new mock instance
func NewFakeEnvironmentClient(ctrl *gomock.Controller) *FakeEnvironmentClient {
	mock := &FakeEnvironmentClient{ctrl: ctrl}
	mock.recorder = &FakeEnvironmentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeEnvironmentClient) EXPECT() *FakeEnvironmentClientMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *FakeEnvironmentClient) List(arg0 string, arg1 ...kf.ListEnvOption) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *FakeEnvironmentClientMockRecorder) List(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeEnvironmentClient)(nil).List), varargs...)
}

// Set mocks base method
func (m *FakeEnvironmentClient) Set(arg0 string, arg1 map[string]string, arg2 ...kf.SetEnvOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *FakeEnvironmentClientMockRecorder) Set(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*FakeEnvironmentClient)(nil).Set), varargs...)
}

// Unset mocks base method
func (m *FakeEnvironmentClient) Unset(arg0 string, arg1 []string, arg2 ...kf.UnsetEnvOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unset", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unset indicates an expected call of Unset
func (mr *FakeEnvironmentClientMockRecorder) Unset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unset", reflect.TypeOf((*FakeEnvironmentClient)(nil).Unset), varargs...)
}
