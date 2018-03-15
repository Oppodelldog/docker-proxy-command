// Code generated by mockery v1.0.0
package mocks

import config "github.com/Oppodelldog/droxy/config"
import exec "os/exec"
import mock "github.com/stretchr/testify/mock"

// CommandBuilder is an autogenerated mock type for the CommandBuilder type
type CommandBuilder struct {
	mock.Mock
}

// BuildCommandFromConfig provides a mock function with given fields: commandName, cfg
func (_m *CommandBuilder) BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error) {
	ret := _m.Called(commandName, cfg)

	var r0 *exec.Cmd
	if rf, ok := ret.Get(0).(func(string, *config.Configuration) *exec.Cmd); ok {
		r0 = rf(commandName, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exec.Cmd)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *config.Configuration) error); ok {
		r1 = rf(commandName, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}