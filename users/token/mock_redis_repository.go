// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package token

import mock "github.com/stretchr/testify/mock"

// mockRedisRepository is an autogenerated mock type for the redisRepository type
type mockRedisRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *mockRedisRepository) Add(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckExists provides a mock function with given fields: _a0
func (_m *mockRedisRepository) CheckExists(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
