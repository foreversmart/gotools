package filestream

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAndRemoveFile(t *testing.T) {
	var (
		assertion = assert.New(t)
		filename  = "aaaaaaaa"
	)

	CreateAbsFile(filename)
	assertion.True(IsExsit(filename))

	RemoveFile(filename)
	assertion.False(IsExsit(filename))
}

func Test_RenameFile(t *testing.T) {
	var (
		assertion   = assert.New(t)
		filename    = "bbbbb"
		newFilename = "cccccccc"
	)

	CreateAbsFile(filename)
	assertion.True(IsExsit(filename))

	RenameFile(filename, newFilename)
	assertion.False(IsExsit(filename))
	assertion.True(IsExsit(newFilename))

	RemoveFile(newFilename)
}

func Test_AppendFile(t *testing.T) {
	var (
		assertion = assert.New(t)
		filename  = "ddddd"
		content   = "dfakdfdkfkd"
		content1  = "toiyotyoto"
	)

	CreateAbsFile(filename)
	AppendFile(filename, content)
	reader, _ := os.Open(Abs(filename))
	res, _ := ioutil.ReadAll(reader)

	assertion.Equal(content, string(res))

	AppendFile(filename, content1)
	reader, _ = os.Open(Abs(filename))
	res, _ = ioutil.ReadAll(reader)

	assertion.Equal(content+content1, string(res))
	RemoveFile(filename)
}
