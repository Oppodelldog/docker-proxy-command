package helper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExecutablePath(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		t.Fatalf("Did not expect os.Executable() to return an error, but got: %v", err)
	}

	executablePath, err := GetExecutablePath()
	if err != nil {
		t.Fatalf("Did not expect GetExecutablePath() to return an error, but got: %v", err)
	}

	assert.Equal(t, filepath.Dir(ex), executablePath)
}

func TestGetExecutableFilePath_smoketest(t *testing.T) {
	GetExecutableFilePath()
}
