// Code generated by mockery v1.0.0
package mocks

import builder "github.com/Oppodelldog/droxy/dockercmd/builder"
import exec "os/exec"
import io "io"
import mock "github.com/stretchr/testify/mock"

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

// AddArgument provides a mock function with given fields: argument
func (_m *Builder) AddArgument(argument string) builder.Builder {
	ret := _m.Called(argument)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(argument)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AddCmdArguments provides a mock function with given fields: arguments
func (_m *Builder) AddCmdArguments(arguments []string) builder.Builder {
	ret := _m.Called(arguments)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func([]string) builder.Builder); ok {
		r0 = rf(arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AddEnvVar provides a mock function with given fields: envVarDeclaration
func (_m *Builder) AddEnvVar(envVarDeclaration string) builder.Builder {
	ret := _m.Called(envVarDeclaration)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(envVarDeclaration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AddGroup provides a mock function with given fields: groupName
func (_m *Builder) AddGroup(groupName string) builder.Builder {
	ret := _m.Called(groupName)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AddPortMapping provides a mock function with given fields: portMapping
func (_m *Builder) AddPortMapping(portMapping string) builder.Builder {
	ret := _m.Called(portMapping)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(portMapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AddVolumeMapping provides a mock function with given fields: hostPath, containerPath, options
func (_m *Builder) AddVolumeMapping(hostPath string, containerPath string, options string) builder.Builder {
	ret := _m.Called(hostPath, containerPath, options)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string, string, string) builder.Builder); ok {
		r0 = rf(hostPath, containerPath, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// AttachTo provides a mock function with given fields: stream
func (_m *Builder) AttachTo(stream string) builder.Builder {
	ret := _m.Called(stream)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(stream)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// Build provides a mock function with given fields:
func (_m *Builder) Build() *exec.Cmd {
	ret := _m.Called()

	var r0 *exec.Cmd
	if rf, ok := ret.Get(0).(func() *exec.Cmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exec.Cmd)
		}
	}

	return r0
}

// SetContainerName provides a mock function with given fields: containerName
func (_m *Builder) SetContainerName(containerName string) builder.Builder {
	ret := _m.Called(containerName)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(containerName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetContainerUserAndGroup provides a mock function with given fields: userID, groupID
func (_m *Builder) SetContainerUserAndGroup(userID string, groupID string) builder.Builder {
	ret := _m.Called(userID, groupID)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string, string) builder.Builder); ok {
		r0 = rf(userID, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetEntryPoint provides a mock function with given fields: entryPoint
func (_m *Builder) SetEntryPoint(entryPoint string) builder.Builder {
	ret := _m.Called(entryPoint)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(entryPoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetImageName provides a mock function with given fields: imageName
func (_m *Builder) SetImageName(imageName string) builder.Builder {
	ret := _m.Called(imageName)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(imageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetNetwork provides a mock function with given fields: network
func (_m *Builder) SetNetwork(network string) builder.Builder {
	ret := _m.Called(network)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(network)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetStdErr provides a mock function with given fields: w
func (_m *Builder) SetStdErr(w io.Writer) builder.Builder {
	ret := _m.Called(w)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(io.Writer) builder.Builder); ok {
		r0 = rf(w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetStdIn provides a mock function with given fields: r
func (_m *Builder) SetStdIn(r io.Reader) builder.Builder {
	ret := _m.Called(r)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(io.Reader) builder.Builder); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetStdOut provides a mock function with given fields: w
func (_m *Builder) SetStdOut(w io.Writer) builder.Builder {
	ret := _m.Called(w)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(io.Writer) builder.Builder); ok {
		r0 = rf(w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}

// SetWorkingDir provides a mock function with given fields: workingDir
func (_m *Builder) SetWorkingDir(workingDir string) builder.Builder {
	ret := _m.Called(workingDir)

	var r0 builder.Builder
	if rf, ok := ret.Get(0).(func(string) builder.Builder); ok {
		r0 = rf(workingDir)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(builder.Builder)
		}
	}

	return r0
}
