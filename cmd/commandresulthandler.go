package cmd

import (
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

//ExtCodeError general error occurred when executing cmd
const ExtCodeError = 993

//ExitError ExitCode of excuted cmd could not be determined
const ExitCodeExitError = 990

//ExitSuccessError ExitCode of successfully executed cmd could not be determined
const ExitSuccessError = 991

type (
	commandResultHandler struct{}
)

//NewCommandResultHandler handles an executed command and it's error code to get the executed commands exit code.
func NewCommandResultHandler() *commandResultHandler {
	return &commandResultHandler{}
}

// HandleCommandResult tries to get to ExitCode of and already run cmd. Returns the exit code or a custom one if original exitcode could not be determined.
func (rh *commandResultHandler) HandleCommandResult(cmd *exec.Cmd, err error) int {

	switch exitErr := err.(type) {
	case *exec.Error:
		logrus.Warning("Could execute command")

		return ExtCodeError
	case *exec.ExitError:
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
			return status.ExitStatus()
		}

		logrus.Warning("Could not get exit code")

		return ExitCodeExitError

	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
		return status.ExitStatus()
	}

	logrus.Warning("Could not get exit code")
	return ExitSuccessError
}
