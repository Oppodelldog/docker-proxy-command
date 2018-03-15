package proxyfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreator_CreateProxyFiles(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandBinaryFilePathStub := "droxy-file-somewhere"
	commandNameStub := "some-command-name"

	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandNameStub,
			},
		},
	}
	creator.CreateProxyFiles(commandBinaryFilePathStub, cfg, false)

	expectedCommandFilename := commandNameStub

	assert.Equal(t, 1, fileCreatorMock.calls)
	assert.Equal(t, commandBinaryFilePathStub, fileCreatorMock.parmCommandBinaryFilePath)
	assert.Equal(t, expectedCommandFilename, fileCreatorMock.parmCommandNameFileName)
}

func TestCreator_CreateProxyFiles_Forced(t *testing.T) {

	testFolder := "/tmp/droxy/createProxyFilesTest/force"
	err := os.MkdirAll(testFolder, 0776)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}

	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	creator := New(fileCreatorMock)

	commandNameStub := "some-command-name"
	fileThatShouldBeDeleted := commandNameStub
	err = ioutil.WriteFile(fileThatShouldBeDeleted, []byte("TEST"), 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	cfg := &config.Configuration{
		Command: []config.CommandDefinition{
			{
				Name: &commandNameStub,
			},
		},
	}
	creator.CreateProxyFiles("", cfg, true)

	_, err = os.Stat(fileThatShouldBeDeleted)
	assert.Error(t, err, "Expect error, since file should be deleted")

	os.RemoveAll(testFolder)
}
