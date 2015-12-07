package filestream

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/foreversmart/gotools/fileswitcher"
	"github.com/stretchr/testify/assert"
)

func Test_Filestream(t *testing.T) {
	var (
		assertion   = assert.New(t)
		pattern     = "message.log.%d"
		specialFile = "message.log"
	)
	p := make([]byte, 15)

	abs, _ := os.Getwd()
	switcher := fileswitcher.NewNumberSwitcher(filepath.Join(abs, pattern), 0, true)
	switcher.SetMax(20).SetMin(0).SetSpecial(0, filepath.Join(abs, specialFile))

	pos := NewPosition(switcher)

	fileStream := NewFileStream(pos)

	CreateAbsFile(specialFile)

	// init test
	n, err := fileStream.Read(p)
	assertion.Equal(err.Error(), "EOF")
	assertion.Equal(0, n)

	// base test
	AppendFile(specialFile, "sssssssss")
	n, err = fileStream.Read(p)
	assertion.Nil(err)
	assertion.Equal(9, n)

	AppendFile(specialFile, "aaaaa")
	n, err = fileStream.Read(p)
	assertion.Nil(err)
	assertion.Equal(5, n)

	// switch file test
	AppendFile(specialFile, "bbbbbb")
	RenameFile(specialFile, specialFile+".1")
	CreateAbsFile(specialFile)
	AppendFile(specialFile, "nnnnnnn")
	n, err = fileStream.Read(p)
	assertion.Nil(err)
	assertion.Equal(6, n)
	assertion.Equal("bbbbbb", string(p[:n]))

	n, err = fileStream.Read(p)
	assertion.Nil(err)
	assertion.Equal(7, n)
	assertion.Equal("nnnnnnn", string(p[:n]))

	// clear test files
	RemoveFile(specialFile)
	RemoveFile(specialFile + ".1")
}
