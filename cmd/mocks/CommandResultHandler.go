// Code generated by mockery v1.0.0
package mocks

import (
	"os/exec"

	"github.com/stretchr/testify/mock"
)

// CommandResultHandler is an autogenerated mock type for the CommandResultHandler type
type CommandResultHandler struct {
	mock.Mock
}

// HandleCommandResult provides a mock function with given fields: _a0, _a1
func (_m *CommandResultHandler) HandleCommandResult(_a0 *exec.Cmd, _a1 error) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*exec.Cmd, error) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
