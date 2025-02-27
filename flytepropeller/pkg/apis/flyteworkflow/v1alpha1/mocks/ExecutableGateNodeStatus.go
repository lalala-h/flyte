// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// ExecutableGateNodeStatus is an autogenerated mock type for the ExecutableGateNodeStatus type
type ExecutableGateNodeStatus struct {
	mock.Mock
}

type ExecutableGateNodeStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *ExecutableGateNodeStatus) EXPECT() *ExecutableGateNodeStatus_Expecter {
	return &ExecutableGateNodeStatus_Expecter{mock: &_m.Mock}
}

// GetGateNodePhase provides a mock function with given fields:
func (_m *ExecutableGateNodeStatus) GetGateNodePhase() v1alpha1.GateNodePhase {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGateNodePhase")
	}

	var r0 v1alpha1.GateNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.GateNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.GateNodePhase)
	}

	return r0
}

// ExecutableGateNodeStatus_GetGateNodePhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGateNodePhase'
type ExecutableGateNodeStatus_GetGateNodePhase_Call struct {
	*mock.Call
}

// GetGateNodePhase is a helper method to define mock.On call
func (_e *ExecutableGateNodeStatus_Expecter) GetGateNodePhase() *ExecutableGateNodeStatus_GetGateNodePhase_Call {
	return &ExecutableGateNodeStatus_GetGateNodePhase_Call{Call: _e.mock.On("GetGateNodePhase")}
}

func (_c *ExecutableGateNodeStatus_GetGateNodePhase_Call) Run(run func()) *ExecutableGateNodeStatus_GetGateNodePhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ExecutableGateNodeStatus_GetGateNodePhase_Call) Return(_a0 v1alpha1.GateNodePhase) *ExecutableGateNodeStatus_GetGateNodePhase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExecutableGateNodeStatus_GetGateNodePhase_Call) RunAndReturn(run func() v1alpha1.GateNodePhase) *ExecutableGateNodeStatus_GetGateNodePhase_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutableGateNodeStatus creates a new instance of ExecutableGateNodeStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutableGateNodeStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExecutableGateNodeStatus {
	mock := &ExecutableGateNodeStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
