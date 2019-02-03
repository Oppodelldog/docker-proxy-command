package cmd

import (
	"github.com/Oppodelldog/droxy/cmd/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_fileCreationSubCommandWrapper_createCommand(t *testing.T) {

	proxyFilesCreatorMock := &mocks.ProxyFilesCreator{}
	proxyFilesCreatorMock.On("CreateProxyFiles", mock.Anything).Return(nil)
	filecreator := newFileCreationSubCommand("symlinks", proxyFilesCreatorMock)

	filecreator.cobraCommand.Run(nil, nil)

	proxyFilesCreatorMock.AssertExpectations(t)
}