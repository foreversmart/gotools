package filestream

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Position(t *testing.T) {
	var (
		filename = "postion_file"
		content  = "lsl2oouiiemmcj"

		assertion = assert.New(t)
	)

	CreateAbsFile(filename)
	AppendFile(filename, content)

	fileInfo, _ := os.Stat(Abs(filename))

	p := make([]byte, 500)
	fd, _ := os.Open(Abs(filename))
	n, _ := fd.Read(p)

	assertion.EqualValues(n, fileInfo.Size())

	RemoveFile(filename)
}
