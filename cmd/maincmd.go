package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Oppodelldog/droxy/proxyfile"

	"github.com/Oppodelldog/droxy/cmd/proxyexecution"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//Execute droxy
func Execute() int {
	var rootCmd = NewRoot()

	if len(os.Args) >= 2 && isSubCommand(os.Args[1], rootCmd.Commands()) {
		err := rootCmd.Execute()
		if err != nil {
			logrus.Info(err)
		}
	} else if len(os.Args) >= 1 && filepath.Base(os.Args[0]) == proxyfile.GetCommandName() {
		err := rootCmd.Help()
		if err != nil {
			logrus.Info(err)
		}
	} else {
		if shallRevealItsDroxy() {
			fmt.Println("YES-IT-IS")

			return 0
		}

		dockerRunCommandBuilder := dockercommand.NewCommandBuilder()
		configLoader := config.NewLoader()
		commandresultHandler := proxyexecution.NewCommandResultHandler()
		commandRunner := proxyexecution.NewCommandRunner()
		executableNameParser := proxyexecution.NewExecutableNameParser()

		return proxyexecution.ExecuteCommand(dockerRunCommandBuilder, configLoader, commandresultHandler, commandRunner, executableNameParser)
	}

	return 0
}

func shallRevealItsDroxy() bool {
	for _, arg := range os.Args {
		if arg == "--is-it-droxy" {

			return true
		}
	}

	return false
}

func isSubCommand(s string, commands []*cobra.Command) bool {
	var subCommandNames []string
	for _, subCommand := range commands {
		subCommandNames = append(subCommandNames, subCommand.Name())

	}
	return stringInSlice(s, subCommandNames)
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}

	return false
}
