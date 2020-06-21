package proxyfile

import "github.com/Oppodelldog/droxy/config"

// This file does not contain tests.
// Is shares some mocks used in several test files

type fileCreationFunctionMock struct {
	returnValue error
	parmSrc     string
	parmDst     string
	calls       int
}

func (m *fileCreationFunctionMock) FileCreationFunc(src, dst string) error {
	m.parmSrc = src
	m.parmDst = dst
	m.calls++

	return m.returnValue
}

type mockFileCreationStrategy struct {
	returnValue               error
	parmCommandBinaryFilePath string
	parmCommandNameFileName   string
	calls                     int
}

func (m *mockFileCreationStrategy) CreateProxyFile(commandBinaryFilePath string, commandNameFileName string) error {
	m.parmCommandBinaryFilePath = commandBinaryFilePath
	m.parmCommandNameFileName = commandNameFileName
	m.calls++

	return m.returnValue
}

type configLoaderMock struct {
	wasLoadCalled bool
	stubbedConfig config.Configuration
}

func (m *configLoaderMock) Load() config.Configuration {
	m.wasLoadCalled = true

	return m.stubbedConfig
}
