package helper

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCopyFile(t *testing.T) {

	testFolder := "/tmp/droxy/fileCopyTest/"
	os.MkdirAll(testFolder, 0776)

	srcBytes := []byte("HELLO DROXY!!")
	src := path.Join(testFolder, "fileToCopy")
	err := ioutil.WriteFile(src, srcBytes, 0666)
	if err != nil {
		t.Fatalf("Did not expect ioutil.WriteFile to return an error, but got: %v", err)
	}

	dst := path.Join(testFolder, "fileCopied")

	err = CopyFile(src, dst)
	if err != nil {
		t.Fatalf("Did not expect CopyFile to return an error, but got: %v", err)
	}

	dstBytes, err := ioutil.ReadFile(dst)
	if err != nil {
		t.Fatalf("Did not expect ioutil.ReadFile to return an error, but got: %v", err)
	}

	assert.Equal(t, dstBytes, srcBytes)

	err = os.RemoveAll(testFolder)
	if err != nil {
		t.Fatalf("Did not expect os.Remove to return an error, but got: %v", err)
	}
}

func TestCopyFile_srcFileDoesNotExist_expectError(t *testing.T) {

	src := "/tmp/THIS_FILE_DOES_NOT_EXIST"
	dst := ""
	err := CopyFile(src, dst)

	assert.Error(t, err)
}