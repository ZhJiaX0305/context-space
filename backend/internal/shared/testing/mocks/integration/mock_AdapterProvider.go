// Code generated by mockery v2.53.4. DO NOT EDIT.

package integration_mocks

import (
	context "context"

	domain "github.com/context-space/context-space/backend/internal/provideradapter/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockAdapterProvider is an autogenerated mock type for the AdapterProvider type
type MockAdapterProvider struct {
	mock.Mock
}

type MockAdapterProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAdapterProvider) EXPECT() *MockAdapterProvider_Expecter {
	return &MockAdapterProvider_Expecter{mock: &_m.Mock}
}

// GetAdapter provides a mock function with given fields: ctx, providerIdentifier
func (_m *MockAdapterProvider) GetAdapter(ctx context.Context, providerIdentifier string) (domain.Adapter, error) {
	ret := _m.Called(ctx, providerIdentifier)

	if len(ret) == 0 {
		panic("no return value specified for GetAdapter")
	}

	var r0 domain.Adapter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Adapter, error)); ok {
		return rf(ctx, providerIdentifier)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Adapter); ok {
		r0 = rf(ctx, providerIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Adapter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, providerIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapterProvider_GetAdapter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAdapter'
type MockAdapterProvider_GetAdapter_Call struct {
	*mock.Call
}

// GetAdapter is a helper method to define mock.On call
//   - ctx context.Context
//   - providerIdentifier string
func (_e *MockAdapterProvider_Expecter) GetAdapter(ctx interface{}, providerIdentifier interface{}) *MockAdapterProvider_GetAdapter_Call {
	return &MockAdapterProvider_GetAdapter_Call{Call: _e.mock.On("GetAdapter", ctx, providerIdentifier)}
}

func (_c *MockAdapterProvider_GetAdapter_Call) Run(run func(ctx context.Context, providerIdentifier string)) *MockAdapterProvider_GetAdapter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAdapterProvider_GetAdapter_Call) Return(_a0 domain.Adapter, _a1 error) *MockAdapterProvider_GetAdapter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapterProvider_GetAdapter_Call) RunAndReturn(run func(context.Context, string) (domain.Adapter, error)) *MockAdapterProvider_GetAdapter_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAdapterProvider creates a new instance of MockAdapterProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAdapterProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAdapterProvider {
	mock := &MockAdapterProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
