package proxyexecution

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/Oppodelldog/droxy/cmd/mocks"
	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const someCommandName = "some-command-name"

func TestExecuteCommand_LoadsConfigFromLoader(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderMock := &mocks.ConfigLoader{}
	configLoaderMock.On("Load").Once().
		Return(configStub)

	executableNameParserStub := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName

	executableNameParserStub.On("ParseCommandNameFromCommandLine").
		Return(commandNameStub)

	commandBuilderStub := &mocks.CommandBuilder{}

	commandBuilderStub.On("BuildCommandFromConfig", mock.Anything, mock.Anything).
		Return(&exec.Cmd{Args: []string{"echo", "1"}}, nil)

	commandRunnerStub := &mocks.CommandRunner{}
	commandRunnerStub.On("RunCommand", mock.Anything).
		Return(nil)

	commandResultHandlerStub := &mocks.CommandResultHandler{}
	commandResultHandlerStub.On("HandleCommandResult", mock.Anything, mock.Anything).
		Return(4711)

	var args []string

	executeCommand(
		args,
		commandBuilderStub,
		configLoaderMock,
		commandResultHandlerStub,
		commandRunnerStub,
		executableNameParserStub,
	)

	configLoaderMock.AssertExpectations(t)
}

func TestExecuteCommand_ExecutableNameIsParsed(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderStub := &mocks.ConfigLoader{}
	configLoaderStub.On("Load").
		Return(configStub)

	executableNameParserMock := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName
	executableNameParserMock.On("ParseCommandNameFromCommandLine").Once().
		Return(commandNameStub)

	commandBuilderStub := &mocks.CommandBuilder{}

	commandBuilderStub.On("BuildCommandFromConfig", mock.Anything, mock.Anything).
		Return(&exec.Cmd{Args: []string{"echo", "1"}}, nil)

	commandRunnerStub := &mocks.CommandRunner{}
	commandRunnerStub.On("RunCommand", mock.Anything).Return(nil)

	commandResultHandlerStub := &mocks.CommandResultHandler{}
	commandResultHandlerStub.On("HandleCommandResult", mock.Anything, mock.Anything).
		Return(4711)

	var args []string

	executeCommand(
		args,
		commandBuilderStub,
		configLoaderStub,
		commandResultHandlerStub,
		commandRunnerStub,
		executableNameParserMock,
	)

	executableNameParserMock.AssertExpectations(t)
}

func TestExecuteCommand_CommandIsBuild(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderStub := &mocks.ConfigLoader{}
	configLoaderStub.On("Load").
		Return(configStub)

	executableNameParserStub := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName
	executableNameParserStub.On("ParseCommandNameFromCommandLine").
		Return(commandNameStub)

	commandBuilderMock := &mocks.CommandBuilder{}

	commandBuilderMock.On("BuildCommandFromConfig", commandNameStub, configStub).
		Return(&exec.Cmd{Args: []string{"echo", "1"}}, nil)

	commandRunnerStub := &mocks.CommandRunner{}
	commandRunnerStub.On("RunCommand", mock.Anything).
		Return(nil)

	commandResultHandlerStub := &mocks.CommandResultHandler{}
	commandResultHandlerStub.On("HandleCommandResult", mock.Anything, mock.Anything).
		Return(4711)

	var args []string

	executeCommand(
		args,
		commandBuilderMock,
		configLoaderStub,
		commandResultHandlerStub,
		commandRunnerStub,
		executableNameParserStub,
	)

	commandBuilderMock.AssertExpectations(t)
}

func TestExecuteCommand_CommandIsRun(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderStub := &mocks.ConfigLoader{}
	configLoaderStub.On("Load").Return(configStub)

	executableNameParserStub := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName
	executableNameParserStub.On("ParseCommandNameFromCommandLine").Return(commandNameStub)

	commandBuilderStub := &mocks.CommandBuilder{}
	cmdStub := &exec.Cmd{Args: []string{}}

	var errStub error

	commandBuilderStub.On("BuildCommandFromConfig", commandNameStub, configStub).Return(cmdStub, errStub)

	commandRunnerMock := &mocks.CommandRunner{}
	commandRunnerMock.On("RunCommand", mock.Anything).Once().Return(nil)

	commandResultHandlerStub := &mocks.CommandResultHandler{}
	commandResultHandlerStub.On("HandleCommandResult", cmdStub, errStub).Return(4711)

	var args []string

	executeCommand(
		args,
		commandBuilderStub,
		configLoaderStub,
		commandResultHandlerStub,
		commandRunnerMock,
		executableNameParserStub,
	)

	commandRunnerMock.AssertExpectations(t)
}

func TestExecuteCommand_CommandResultIsHandled(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderStub := &mocks.ConfigLoader{}
	configLoaderStub.On("Load").Return(configStub)

	executableNameParserStub := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName
	executableNameParserStub.On("ParseCommandNameFromCommandLine").Return(commandNameStub)

	commandBuilderStub := &mocks.CommandBuilder{}
	cmdStub := &exec.Cmd{Args: []string{}}

	var errStub error

	commandBuilderStub.On("BuildCommandFromConfig", commandNameStub, configStub).Return(cmdStub, errStub)

	commandRunnerStub := &mocks.CommandRunner{}
	commandRunnerStub.On("RunCommand", mock.Anything).Return(nil)

	commandResultHandlerMock := &mocks.CommandResultHandler{}
	commandResultHandlerMock.On("HandleCommandResult", cmdStub, errStub).Return(4711)

	var args []string

	executeCommand(
		args,
		commandBuilderStub,
		configLoaderStub,
		commandResultHandlerMock,
		commandRunnerStub,
		executableNameParserStub,
	)

	commandResultHandlerMock.AssertExpectations(t)
}

func TestExecuteCommand_ErrorFromCommandBuild_ExitCode900Returned(t *testing.T) {
	configStub := &config.Configuration{}
	configLoaderStub := &mocks.ConfigLoader{}
	configLoaderStub.On("Load").Return(configStub)

	executableNameParserStub := &mocks.ExecutableNameParser{}
	commandNameStub := someCommandName
	executableNameParserStub.On("ParseCommandNameFromCommandLine").Return(commandNameStub)

	commandBuilderStub := &mocks.CommandBuilder{}
	commandBuilderStub.On("BuildCommandFromConfig", commandNameStub, configStub).Return(nil, errors.New("some-error"))

	commandRunnerStub := &mocks.CommandRunner{}
	commandRunnerStub.On("RunCommand", mock.Anything).Return(nil)

	commandResultHandlerStub := &mocks.CommandResultHandler{}
	commandResultHandlerStub.On("HandleCommandResult", mock.Anything, mock.Anything).Return(4711)

	var args []string
	exitCode := executeCommand(args,
		commandBuilderStub,
		configLoaderStub,
		commandResultHandlerStub,
		commandRunnerStub,
		executableNameParserStub,
	)

	assert.Equal(t, 900, exitCode)
}
