// Code generated by mockery v2.53.4. DO NOT EDIT.

package credentialmanagement_mocks

import (
	context "context"

	domain "github.com/context-space/context-space/backend/internal/credentialmanagement/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockVaultService is an autogenerated mock type for the VaultService type
type MockVaultService struct {
	mock.Mock
}

type MockVaultService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVaultService) EXPECT() *MockVaultService_Expecter {
	return &MockVaultService_Expecter{mock: &_m.Mock}
}

// DecryptData provides a mock function with given fields: ctx, metadata
func (_m *MockVaultService) DecryptData(ctx context.Context, metadata *domain.EncryptionMetadata) (string, error) {
	ret := _m.Called(ctx, metadata)

	if len(ret) == 0 {
		panic("no return value specified for DecryptData")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.EncryptionMetadata) (string, error)); ok {
		return rf(ctx, metadata)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.EncryptionMetadata) string); ok {
		r0 = rf(ctx, metadata)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.EncryptionMetadata) error); ok {
		r1 = rf(ctx, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVaultService_DecryptData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecryptData'
type MockVaultService_DecryptData_Call struct {
	*mock.Call
}

// DecryptData is a helper method to define mock.On call
//   - ctx context.Context
//   - metadata *domain.EncryptionMetadata
func (_e *MockVaultService_Expecter) DecryptData(ctx interface{}, metadata interface{}) *MockVaultService_DecryptData_Call {
	return &MockVaultService_DecryptData_Call{Call: _e.mock.On("DecryptData", ctx, metadata)}
}

func (_c *MockVaultService_DecryptData_Call) Run(run func(ctx context.Context, metadata *domain.EncryptionMetadata)) *MockVaultService_DecryptData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.EncryptionMetadata))
	})
	return _c
}

func (_c *MockVaultService_DecryptData_Call) Return(_a0 string, _a1 error) *MockVaultService_DecryptData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVaultService_DecryptData_Call) RunAndReturn(run func(context.Context, *domain.EncryptionMetadata) (string, error)) *MockVaultService_DecryptData_Call {
	_c.Call.Return(run)
	return _c
}

// DecryptJSON provides a mock function with given fields: ctx, metadata, target
func (_m *MockVaultService) DecryptJSON(ctx context.Context, metadata *domain.EncryptionMetadata, target interface{}) error {
	ret := _m.Called(ctx, metadata, target)

	if len(ret) == 0 {
		panic("no return value specified for DecryptJSON")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.EncryptionMetadata, interface{}) error); ok {
		r0 = rf(ctx, metadata, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVaultService_DecryptJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecryptJSON'
type MockVaultService_DecryptJSON_Call struct {
	*mock.Call
}

// DecryptJSON is a helper method to define mock.On call
//   - ctx context.Context
//   - metadata *domain.EncryptionMetadata
//   - target interface{}
func (_e *MockVaultService_Expecter) DecryptJSON(ctx interface{}, metadata interface{}, target interface{}) *MockVaultService_DecryptJSON_Call {
	return &MockVaultService_DecryptJSON_Call{Call: _e.mock.On("DecryptJSON", ctx, metadata, target)}
}

func (_c *MockVaultService_DecryptJSON_Call) Run(run func(ctx context.Context, metadata *domain.EncryptionMetadata, target interface{})) *MockVaultService_DecryptJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.EncryptionMetadata), args[2].(interface{}))
	})
	return _c
}

func (_c *MockVaultService_DecryptJSON_Call) Return(_a0 error) *MockVaultService_DecryptJSON_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVaultService_DecryptJSON_Call) RunAndReturn(run func(context.Context, *domain.EncryptionMetadata, interface{}) error) *MockVaultService_DecryptJSON_Call {
	_c.Call.Return(run)
	return _c
}

// EncryptData provides a mock function with given fields: ctx, plaintext, region, credentialType
func (_m *MockVaultService) EncryptData(ctx context.Context, plaintext string, region domain.VaultRegion, credentialType domain.CredentialType) (*domain.EncryptionMetadata, error) {
	ret := _m.Called(ctx, plaintext, region, credentialType)

	if len(ret) == 0 {
		panic("no return value specified for EncryptData")
	}

	var r0 *domain.EncryptionMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.VaultRegion, domain.CredentialType) (*domain.EncryptionMetadata, error)); ok {
		return rf(ctx, plaintext, region, credentialType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.VaultRegion, domain.CredentialType) *domain.EncryptionMetadata); ok {
		r0 = rf(ctx, plaintext, region, credentialType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.EncryptionMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.VaultRegion, domain.CredentialType) error); ok {
		r1 = rf(ctx, plaintext, region, credentialType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVaultService_EncryptData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EncryptData'
type MockVaultService_EncryptData_Call struct {
	*mock.Call
}

// EncryptData is a helper method to define mock.On call
//   - ctx context.Context
//   - plaintext string
//   - region domain.VaultRegion
//   - credentialType domain.CredentialType
func (_e *MockVaultService_Expecter) EncryptData(ctx interface{}, plaintext interface{}, region interface{}, credentialType interface{}) *MockVaultService_EncryptData_Call {
	return &MockVaultService_EncryptData_Call{Call: _e.mock.On("EncryptData", ctx, plaintext, region, credentialType)}
}

func (_c *MockVaultService_EncryptData_Call) Run(run func(ctx context.Context, plaintext string, region domain.VaultRegion, credentialType domain.CredentialType)) *MockVaultService_EncryptData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.VaultRegion), args[3].(domain.CredentialType))
	})
	return _c
}

func (_c *MockVaultService_EncryptData_Call) Return(_a0 *domain.EncryptionMetadata, _a1 error) *MockVaultService_EncryptData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVaultService_EncryptData_Call) RunAndReturn(run func(context.Context, string, domain.VaultRegion, domain.CredentialType) (*domain.EncryptionMetadata, error)) *MockVaultService_EncryptData_Call {
	_c.Call.Return(run)
	return _c
}

// EncryptJSON provides a mock function with given fields: ctx, data, region, credentialType
func (_m *MockVaultService) EncryptJSON(ctx context.Context, data interface{}, region domain.VaultRegion, credentialType domain.CredentialType) (*domain.EncryptionMetadata, error) {
	ret := _m.Called(ctx, data, region, credentialType)

	if len(ret) == 0 {
		panic("no return value specified for EncryptJSON")
	}

	var r0 *domain.EncryptionMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, domain.VaultRegion, domain.CredentialType) (*domain.EncryptionMetadata, error)); ok {
		return rf(ctx, data, region, credentialType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, domain.VaultRegion, domain.CredentialType) *domain.EncryptionMetadata); ok {
		r0 = rf(ctx, data, region, credentialType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.EncryptionMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, domain.VaultRegion, domain.CredentialType) error); ok {
		r1 = rf(ctx, data, region, credentialType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVaultService_EncryptJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EncryptJSON'
type MockVaultService_EncryptJSON_Call struct {
	*mock.Call
}

// EncryptJSON is a helper method to define mock.On call
//   - ctx context.Context
//   - data interface{}
//   - region domain.VaultRegion
//   - credentialType domain.CredentialType
func (_e *MockVaultService_Expecter) EncryptJSON(ctx interface{}, data interface{}, region interface{}, credentialType interface{}) *MockVaultService_EncryptJSON_Call {
	return &MockVaultService_EncryptJSON_Call{Call: _e.mock.On("EncryptJSON", ctx, data, region, credentialType)}
}

func (_c *MockVaultService_EncryptJSON_Call) Run(run func(ctx context.Context, data interface{}, region domain.VaultRegion, credentialType domain.CredentialType)) *MockVaultService_EncryptJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(domain.VaultRegion), args[3].(domain.CredentialType))
	})
	return _c
}

func (_c *MockVaultService_EncryptJSON_Call) Return(_a0 *domain.EncryptionMetadata, _a1 error) *MockVaultService_EncryptJSON_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVaultService_EncryptJSON_Call) RunAndReturn(run func(context.Context, interface{}, domain.VaultRegion, domain.CredentialType) (*domain.EncryptionMetadata, error)) *MockVaultService_EncryptJSON_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVaultService creates a new instance of MockVaultService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVaultService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVaultService {
	mock := &MockVaultService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
