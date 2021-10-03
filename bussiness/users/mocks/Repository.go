// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import users "EzMusix/bussiness/users"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Login provides a mock function with given fields: _a0
func (_m *Repository) Login(_a0 users.Users) (users.Users, error) {
	ret := _m.Called(_a0)

	var r0 users.Users
	if rf, ok := ret.Get(0).(func(users.Users) users.Users); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(users.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Users) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *Repository) Register(_a0 users.Users) (users.Users, error) {
	ret := _m.Called(_a0)

	var r0 users.Users
	if rf, ok := ret.Get(0).(func(users.Users) users.Users); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(users.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Users) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}