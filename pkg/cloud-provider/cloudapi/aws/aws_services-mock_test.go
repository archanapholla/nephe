// // Copyright 2022 Antrea Authors.
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //      http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud-provider/cloudapi/aws/aws_services.go

// Package aws is a generated GoMock package.
package aws

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockawsServiceClientCreateInterface is a mock of awsServiceClientCreateInterface interface.
type MockawsServiceClientCreateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockawsServiceClientCreateInterfaceMockRecorder
}

// MockawsServiceClientCreateInterfaceMockRecorder is the mock recorder for MockawsServiceClientCreateInterface.
type MockawsServiceClientCreateInterfaceMockRecorder struct {
	mock *MockawsServiceClientCreateInterface
}

// NewMockawsServiceClientCreateInterface creates a new mock instance.
func NewMockawsServiceClientCreateInterface(ctrl *gomock.Controller) *MockawsServiceClientCreateInterface {
	mock := &MockawsServiceClientCreateInterface{ctrl: ctrl}
	mock.recorder = &MockawsServiceClientCreateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockawsServiceClientCreateInterface) EXPECT() *MockawsServiceClientCreateInterfaceMockRecorder {
	return m.recorder
}

// compute mocks base method.
func (m *MockawsServiceClientCreateInterface) compute() (awsEC2Wrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "compute")
	ret0, _ := ret[0].(awsEC2Wrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// compute indicates an expected call of compute.
func (mr *MockawsServiceClientCreateInterfaceMockRecorder) compute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "compute", reflect.TypeOf((*MockawsServiceClientCreateInterface)(nil).compute))
}

// MockawsServicesHelper is a mock of awsServicesHelper interface.
type MockawsServicesHelper struct {
	ctrl     *gomock.Controller
	recorder *MockawsServicesHelperMockRecorder
}

// MockawsServicesHelperMockRecorder is the mock recorder for MockawsServicesHelper.
type MockawsServicesHelperMockRecorder struct {
	mock *MockawsServicesHelper
}

// NewMockawsServicesHelper creates a new mock instance.
func NewMockawsServicesHelper(ctrl *gomock.Controller) *MockawsServicesHelper {
	mock := &MockawsServicesHelper{ctrl: ctrl}
	mock.recorder = &MockawsServicesHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockawsServicesHelper) EXPECT() *MockawsServicesHelperMockRecorder {
	return m.recorder
}

// newServiceSdkConfigProvider mocks base method.
func (m *MockawsServicesHelper) newServiceSdkConfigProvider(accCfg *awsAccountConfig) (awsServiceClientCreateInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "newServiceSdkConfigProvider", accCfg)
	ret0, _ := ret[0].(awsServiceClientCreateInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// newServiceSdkConfigProvider indicates an expected call of newServiceSdkConfigProvider.
func (mr *MockawsServicesHelperMockRecorder) newServiceSdkConfigProvider(accCfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newServiceSdkConfigProvider", reflect.TypeOf((*MockawsServicesHelper)(nil).newServiceSdkConfigProvider), accCfg)
}