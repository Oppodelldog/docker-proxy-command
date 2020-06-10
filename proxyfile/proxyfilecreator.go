package proxyfile

import (
	"os"

	"github.com/Oppodelldog/droxy/crossplatform"

	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
)

// New creates a new proxy file creator.
func New(creationStrategy FileCreationStrategy, configLoader config.Loader) Creator {
	return Creator{
		creationStrategy:          creationStrategy,
		configLoader:              configLoader,
		getExecutableFilePathFunc: getExecutableFilePath,
	}
}

// Creator creates commands.
type Creator struct {
	creationStrategy          FileCreationStrategy
	configLoader              config.Loader
	getExecutableFilePathFunc getExecutableFilePathFuncDef
}

type getExecutableFilePathFuncDef func() (string, error)

// CreateProxyFiles creates droxy commands.
func (pfc Creator) CreateProxyFiles(isForced bool) error {
	cfg := pfc.configLoader.Load()

	commandBinaryFilePath, err := pfc.getExecutableFilePathFunc()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	for _, command := range cfg.Command {
		if !command.HasName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if isTemplate, ok := command.GetIsTemplate(); isTemplate && ok {
			continue
		}

		commandName, ok := command.GetName()
		if !ok {
			continue
		}

		commandNameFileName := crossplatform.GetCommandNameFilename(commandName)

		if fileExistsAsDir(commandNameFileName) {
			logrus.Warnf("droxy command file already exists as a directory '%s'", commandNameFileName)
			return nil
		}

		if isForced {
			removeFile(commandNameFileName)
		} else if fileExists(commandNameFileName) {
			logrus.Warnf("droxy command file (%s) already exists for command '%s'", commandNameFileName, commandName)
			continue
		}

		err = pfc.creationStrategy.CreateProxyFile(commandBinaryFilePath, commandNameFileName)
		if err != nil {
			logrus.Errorf("error creating symlink '%s': %v", commandName, err)
			continue
		}

		logrus.Infof("created '%s'", commandName)
	}

	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

func removeFile(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		logrus.Warnf("cannot delete droxy command file (%s): %v", filePath, err)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		panic(err)
	}
}

func fileExistsAsDir(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
