package config

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	errCommandNotDefined    = errors.New("command not defined")
	errCouldNotFindTemplate = errors.New("could not find template")
)

// Configuration defines the fields/types of the configuration file.
type Configuration struct {
	Command          []CommandDefinition
	Version          string
	ConfigFilePath   string
	Logging          bool
	EnvVarOverwrites *[]string
	osNameMatcher    func(string) bool
}

// FindCommandByName finds a command by the given name.
func (c Configuration) FindCommandByName(commandName string) (CommandDefinition, error) {
	var commandDef *CommandDefinition

	for _, command := range c.Command {
		if !c.match(commandName, command) {
			continue
		}

		cd, err := c.resolveConfig(command)
		if err != nil {
			return CommandDefinition{}, fmt.Errorf("error resolving config '%s': %w", commandName, err)
		}

		commandDef = &cd
	}

	if commandDef != nil {
		c.applyGlobalConfig(commandDef)

		return *commandDef, nil
	}

	return CommandDefinition{}, fmt.Errorf("%w: '%s'", errCommandNotDefined, commandName)
}

func (c Configuration) applyGlobalConfig(commandDef *CommandDefinition) {
	if globalEnvOverwrites, ok := c.GetEnvVarOverwrites(); ok {
		if commandDef.EnvVarOverwrites != nil && len(*commandDef.EnvVarOverwrites) > 0 {
			var newOverwrites []string
			newOverwrites = append(newOverwrites, *c.EnvVarOverwrites...)
			newOverwrites = append(newOverwrites, *commandDef.EnvVarOverwrites...)
			commandDef.EnvVarOverwrites = &newOverwrites
		} else {
			commandDef.EnvVarOverwrites = &globalEnvOverwrites
		}
	}
}

func (c Configuration) match(commandName string, command CommandDefinition) bool {
	configCommandName, ok := command.GetName()
	if !ok {
		return false
	}

	if configCommandName != commandName {
		return false
	}

	if !c.matchesOS(command) {
		return false
	}

	return true
}

func (c Configuration) matchesOS(command CommandDefinition) bool {
	osName, _ := command.GetOS()

	if osName == "" {
		return true
	}

	return c.osNameMatcher(osName)
}

// GetConfigurationFilePath returns the path the configuration was load from. this is for debugging purpose.
func (c Configuration) GetConfigurationFilePath() string {
	return c.ConfigFilePath
}

func (c Configuration) resolveConfig(command CommandDefinition) (CommandDefinition, error) {
	if !command.HasTemplate() {
		return command, nil
	}

	templateDefinition, err := c.FindCommandByName(*command.Template)
	if err != nil {
		return CommandDefinition{},
			fmt.Errorf(
				"%w '%s' to resolve config of '%s'",
				errCouldNotFindTemplate,
				*command.Template,
				*command.Name,
			)
	}

	return mergeCommand(templateDefinition, command), nil
}

// GetEnvVarOverwrites returns env var mappings that will be passed to all commands EnvVarOverwrites.
func (c Configuration) GetEnvVarOverwrites() ([]string, bool) {
	return getStringSlice(c.EnvVarOverwrites)
}

func defaultOSNameMatcher(osName string) bool {
	return runtime.GOOS == osName
}
