// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// MockRevoker is an autogenerated mock type for the Revoker type
type MockRevoker struct {
	mock.Mock
}

// RevokeToken provides a mock function with given fields: _a0
func (_m *MockRevoker) RevokeToken(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}