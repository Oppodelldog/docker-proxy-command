package proxyfile

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/Oppodelldog/droxy/config"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const testFolder = "/tmp/droxy/createProxyFilesTest"

func TestCreator_New(t *testing.T) {
	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{}
	creator := New(fileCreatorMock, configLoaderMock)

	assert.IsType(t, new(Creator), creator)
	assert.Exactly(t, fileCreatorMock, creator.creationStrategy)
	assert.Exactly(t, configLoaderMock, creator.configLoader)

	if reflect.ValueOf(creator.getExecutableFilePathFunc).Pointer() != reflect.ValueOf(getExecutableFilePath).Pointer() {
		t.Fatal("expected 'getExecutableFilePath' to be configred as getExecutableFilePathFunc, but was not")
	}
}

func getTestConfig() *config.Configuration {
	commandNameStub := "some-command-name"
	return &config.Configuration{
		Command: []config.CommandDefinition{
			{Name: &commandNameStub},
		},
	}
}

func getTestConfigWithEmptyCommand() *config.Configuration {
	return &config.Configuration{Command: []config.CommandDefinition{{}}}
}

func TestCreator_CreateProxyFiles(t *testing.T) {
	prepareTest(t)

	commandBinaryFilePathStub := "/tmp/droxy"

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return commandBinaryFilePathStub, nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	expectedCommandFilename := *configLoaderMock.stubbedConfig.Command[0].Name

	assert.Equal(t, 1, fileCreatorMock.calls)
	assert.Equal(t, commandBinaryFilePathStub, fileCreatorMock.parmCommandBinaryFilePath)
	assert.Equal(t, expectedCommandFilename, fileCreatorMock.parmCommandNameFileName)
}

func TestCreator_CreateProxyFiles_commandHasNoName_noFileWillBeCreated(t *testing.T) {
	prepareTest(t)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfigWithEmptyCommand()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_commandIsTemplate_noFileWillBeCreated(t *testing.T) {
	prepareTest(t)

	fileCreatorMock := &mockFileCreationStrategy{}
	testConfig := getTestConfig()
	isTemplate := true
	testConfig.Command[0].IsTemplate = &isTemplate
	configLoaderMock := &configLoaderMock{stubbedConfig: testConfig}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	err := creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	assert.Equal(t, 0, fileCreatorMock.calls)
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsNotForced_existingFileWillNotBeReplaced(t *testing.T) {
	prepareTest(t)

	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	commandNameStub := *configLoaderMock.stubbedConfig.Command[0].Name
	fileThatShouldNotBeDeleted := commandNameStub
	err := ioutil.WriteFile(fileThatShouldNotBeDeleted, []byte("TEST"), 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	err = creator.CreateProxyFiles(false)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	_, err = os.Stat(fileThatShouldNotBeDeleted)
	assert.Nil(t, err, "Expect no error, since file should not have been deleted")

	err = os.Remove(fileThatShouldNotBeDeleted)
	if err != nil {
		t.Fatalf("Did not expect os.Remove to return an error, but got: %v", err)
	}
}

func TestCreator_CreateProxyFiles_fileAlreadyExistsAndCreationIsForced_existingFileWillBeReplaced(t *testing.T) {
	prepareTest(t)

	logrus.SetOutput(ioutil.Discard)

	fileCreatorMock := &mockFileCreationStrategy{}
	configLoaderMock := &configLoaderMock{stubbedConfig: getTestConfig()}
	creator := &Creator{
		creationStrategy:          fileCreatorMock,
		configLoader:              configLoaderMock,
		getExecutableFilePathFunc: func() (string, error) { return "", nil },
	}

	commandNameStub := *configLoaderMock.stubbedConfig.Command[0].Name
	fileThatShouldBeDeleted := commandNameStub
	err := ioutil.WriteFile(fileThatShouldBeDeleted, []byte("TEST"), 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	err = creator.CreateProxyFiles(true)
	if err != nil {
		t.Fatalf("Did not expect CreateProxyFiles to return an error, but got: %v", err)
	}

	_, err = os.Stat(fileThatShouldBeDeleted)
	assert.Error(t, err, "Expect error, since file should be deleted")
}

func prepareTest(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	err := os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.RemoveAll to return an error, but got: %v", err)
	}

	err = os.MkdirAll(testFolder, 0776)
	if err != nil {
		t.Fatalf("Did not expect os.MkdirAll to return an error, but got: %v", err)
	}
}
