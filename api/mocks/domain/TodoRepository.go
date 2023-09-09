// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	domain "example.com/mod/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

type TodoRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TodoRepository) EXPECT() *TodoRepository_Expecter {
	return &TodoRepository_Expecter{mock: &_m.Mock}
}

// AllGet provides a mock function with given fields:
func (_m *TodoRepository) AllGet() ([]domain.Todo, error) {
	ret := _m.Called()

	var r0 []domain.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TodoRepository_AllGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllGet'
type TodoRepository_AllGet_Call struct {
	*mock.Call
}

// AllGet is a helper method to define mock.On call
func (_e *TodoRepository_Expecter) AllGet() *TodoRepository_AllGet_Call {
	return &TodoRepository_AllGet_Call{Call: _e.mock.On("AllGet")}
}

func (_c *TodoRepository_AllGet_Call) Run(run func()) *TodoRepository_AllGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TodoRepository_AllGet_Call) Return(_a0 []domain.Todo, _a1 error) *TodoRepository_AllGet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TodoRepository_AllGet_Call) RunAndReturn(run func() ([]domain.Todo, error)) *TodoRepository_AllGet_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *TodoRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TodoRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int
func (_e *TodoRepository_Expecter) Delete(id interface{}) *TodoRepository_Delete_Call {
	return &TodoRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *TodoRepository_Delete_Call) Run(run func(id int)) *TodoRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *TodoRepository_Delete_Call) Return(_a0 error) *TodoRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_Delete_Call) RunAndReturn(run func(int) error) *TodoRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// StatusUpdate provides a mock function with given fields: id
func (_m *TodoRepository) StatusUpdate(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_StatusUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatusUpdate'
type TodoRepository_StatusUpdate_Call struct {
	*mock.Call
}

// StatusUpdate is a helper method to define mock.On call
//   - id int
func (_e *TodoRepository_Expecter) StatusUpdate(id interface{}) *TodoRepository_StatusUpdate_Call {
	return &TodoRepository_StatusUpdate_Call{Call: _e.mock.On("StatusUpdate", id)}
}

func (_c *TodoRepository_StatusUpdate_Call) Run(run func(id int)) *TodoRepository_StatusUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *TodoRepository_StatusUpdate_Call) Return(_a0 error) *TodoRepository_StatusUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_StatusUpdate_Call) RunAndReturn(run func(int) error) *TodoRepository_StatusUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function with given fields: todo
func (_m *TodoRepository) Store(todo domain.Todo) error {
	ret := _m.Called(todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type TodoRepository_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - todo domain.Todo
func (_e *TodoRepository_Expecter) Store(todo interface{}) *TodoRepository_Store_Call {
	return &TodoRepository_Store_Call{Call: _e.mock.On("Store", todo)}
}

func (_c *TodoRepository_Store_Call) Run(run func(todo domain.Todo)) *TodoRepository_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Todo))
	})
	return _c
}

func (_c *TodoRepository_Store_Call) Return(_a0 error) *TodoRepository_Store_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_Store_Call) RunAndReturn(run func(domain.Todo) error) *TodoRepository_Store_Call {
	_c.Call.Return(run)
	return _c
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}