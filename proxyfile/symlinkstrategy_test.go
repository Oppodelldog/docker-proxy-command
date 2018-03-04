package proxyfile

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSymlinkStrategy_callsTheAppripriateSystemFunction(t *testing.T) {
	strategy := NewSymlinkStrategy()

	strategyFunction := strategy.(*SymlinkStrategy).symlinkFunction
	expectedFunction := os.Symlink

	if reflect.ValueOf(expectedFunction).Pointer() != reflect.ValueOf(strategyFunction).Pointer() {
		t.Fail()
	}
}

func TestNewSymlinkStrategy_callsConfiguredSystemFunction(t *testing.T) {
	mock := fileCreationFunctionMock{}
	strategy := NewSymlinkStrategy()
	strategy.(*SymlinkStrategy).symlinkFunction = mock.FileCreationFunc

	expectedSrc := "A"
	expectedDst := "B"
	strategy.CreateProxyFile(expectedSrc, expectedDst)

	assert.Equal(t, expectedSrc, mock.parmSrc)
	assert.Equal(t, expectedDst, mock.parmDst)
	assert.Equal(t, 1, mock.calls)
}

func TestNewSymlinkStrategy_returnsErrorIfSystemFunctionReturnsError(t *testing.T) {
	expectedError := errors.New("error from configured system function")

	mock := fileCreationFunctionMock{returnValue: expectedError}
	strategy := NewSymlinkStrategy()
	strategy.(*SymlinkStrategy).symlinkFunction = mock.FileCreationFunc

	err := strategy.CreateProxyFile("A", "B")

	assert.Equal(t, expectedError, err)
}