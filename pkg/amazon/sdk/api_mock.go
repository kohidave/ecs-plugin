// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/docker/ecs-plugin/pkg/amazon (interfaces: API)

// Package amazon is a generated GoMock package.
package sdk

import (
	context "context"
	reflect "reflect"

	"github.com/docker/ecs-plugin/pkg/compose"

	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	cloudformation0 "github.com/awslabs/goformation/v4/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ClusterExists mocks base method
func (m *MockAPI) ClusterExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterExists indicates an expected call of ClusterExists
func (mr *MockAPIMockRecorder) ClusterExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterExists", reflect.TypeOf((*MockAPI)(nil).ClusterExists), arg0, arg1)
}

// CreateSecret mocks base method
func (m *MockAPI) CreateSecret(arg0 context.Context, arg1 compose.Secret) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockAPIMockRecorder) CreateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockAPI)(nil).CreateSecret), arg0, arg1)
}

// CreateStack mocks base method
func (m *MockAPI) CreateStack(arg0 context.Context, arg1 string, arg2 *cloudformation0.Template, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStack", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStack indicates an expected call of CreateStack
func (mr *MockAPIMockRecorder) CreateStack(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStack", reflect.TypeOf((*MockAPI)(nil).CreateStack), arg0, arg1, arg2, arg3)
}

// DeleteCluster mocks base method
func (m *MockAPI) DeleteCluster(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockAPIMockRecorder) DeleteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockAPI)(nil).DeleteCluster), arg0, arg1)
}

// DeleteSecret mocks base method
func (m *MockAPI) DeleteSecret(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockAPIMockRecorder) DeleteSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockAPI)(nil).DeleteSecret), arg0, arg1, arg2)
}

// DeleteStack mocks base method
func (m *MockAPI) DeleteStack(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockAPIMockRecorder) DeleteStack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*MockAPI)(nil).DeleteStack), arg0, arg1)
}

// DescribeStackEvents mocks base method
func (m *MockAPI) DescribeStackEvents(arg0 context.Context, arg1 string) ([]*cloudformation.StackEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStackEvents", arg0, arg1)
	ret0, _ := ret[0].([]*cloudformation.StackEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStackEvents indicates an expected call of DescribeStackEvents
func (mr *MockAPIMockRecorder) DescribeStackEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStackEvents", reflect.TypeOf((*MockAPI)(nil).DescribeStackEvents), arg0, arg1)
}

// DescribeTasks mocks base method
func (m *MockAPI) DescribeTasks(arg0 context.Context, arg1 string, arg2 ...string) ([]compose.TaskStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTasks", varargs...)
	ret0, _ := ret[0].([]compose.TaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTasks indicates an expected call of DescribeTasks
func (mr *MockAPIMockRecorder) DescribeTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTasks", reflect.TypeOf((*MockAPI)(nil).DescribeTasks), varargs...)
}

// GetDefaultVPC mocks base method
func (m *MockAPI) GetDefaultVPC(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultVPC", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultVPC indicates an expected call of GetDefaultVPC
func (mr *MockAPIMockRecorder) GetDefaultVPC(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultVPC", reflect.TypeOf((*MockAPI)(nil).GetDefaultVPC), arg0)
}

// GetLogs mocks base method
func (m *MockAPI) GetLogs(arg0 context.Context, arg1 string, arg2 compose.LogConsumer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockAPIMockRecorder) GetLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockAPI)(nil).GetLogs), arg0, arg1, arg2)
}

// GetPublicIPs mocks base method
func (m *MockAPI) GetPublicIPs(arg0 context.Context, arg1 ...string) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicIPs", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicIPs indicates an expected call of GetPublicIPs
func (mr *MockAPIMockRecorder) GetPublicIPs(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPs", reflect.TypeOf((*MockAPI)(nil).GetPublicIPs), varargs...)
}

// GetStackID mocks base method
func (m *MockAPI) GetStackID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStackID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStackID indicates an expected call of GetStackID
func (mr *MockAPIMockRecorder) GetStackID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStackID", reflect.TypeOf((*MockAPI)(nil).GetStackID), arg0, arg1)
}

// GetSubNets mocks base method
func (m *MockAPI) GetSubNets(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubNets", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubNets indicates an expected call of GetSubNets
func (mr *MockAPIMockRecorder) GetSubNets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubNets", reflect.TypeOf((*MockAPI)(nil).GetSubNets), arg0, arg1)
}

// InspectSecret mocks base method
func (m *MockAPI) InspectSecret(arg0 context.Context, arg1 string) (compose.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectSecret", arg0, arg1)
	ret0, _ := ret[0].(compose.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectSecret indicates an expected call of InspectSecret
func (mr *MockAPIMockRecorder) InspectSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectSecret", reflect.TypeOf((*MockAPI)(nil).InspectSecret), arg0, arg1)
}

// ListSecrets mocks base method
func (m *MockAPI) ListSecrets(arg0 context.Context) ([]compose.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].([]compose.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockAPIMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockAPI)(nil).ListSecrets), arg0)
}

// ListTasks mocks base method
func (m *MockAPI) ListTasks(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockAPIMockRecorder) ListTasks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockAPI)(nil).ListTasks), arg0, arg1, arg2)
}

// StackExists mocks base method
func (m *MockAPI) StackExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackExists indicates an expected call of StackExists
func (mr *MockAPIMockRecorder) StackExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackExists", reflect.TypeOf((*MockAPI)(nil).StackExists), arg0, arg1)
}

// VpcExists mocks base method
func (m *MockAPI) VpcExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VpcExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VpcExists indicates an expected call of VpcExists
func (mr *MockAPIMockRecorder) VpcExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VpcExists", reflect.TypeOf((*MockAPI)(nil).VpcExists), arg0, arg1)
}

// WaitStackComplete mocks base method
func (m *MockAPI) WaitStackComplete(arg0 context.Context, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitStackComplete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitStackComplete indicates an expected call of WaitStackComplete
func (mr *MockAPIMockRecorder) WaitStackComplete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitStackComplete", reflect.TypeOf((*MockAPI)(nil).WaitStackComplete), arg0, arg1, arg2)
}

// LoadBalancerExists mocks base method
func (m *MockAPI) LoadBalancerExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadBalancerExists indicates an expected call of VpcExists
func (mr *MockAPIMockRecorder) LoadBalancerExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerExists", reflect.TypeOf((*MockAPI)(nil).LoadBalancerExists), arg0, arg1)
}

// GetLoadBalancerARN mocks base method
func (m *MockAPI) GetLoadBalancerARN(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancerARN", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
