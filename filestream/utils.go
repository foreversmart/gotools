package filestream

import (
	"os"
	"path/filepath"
)

func CreateAbsFile(name string) {
	os.Create(Abs(name))
}

func RenameFile(old, new string) {
	old = Abs(old)
	new = Abs(new)
	os.Rename(old, new)
}

func RemoveFile(name string) {
	name = Abs(name)
	os.Remove(name)
}

func AppendFile(name, content string) {
	name = Abs(name)
	fileInfo, _ := os.Stat(name)
	tail := fileInfo.Size()

	file, _ := os.OpenFile(name, os.O_WRONLY, os.ModeAppend)
	file.WriteAt([]byte(content), tail)
	file.Close()
}

func IsExsit(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func Abs(name string) string {
	abs, _ := filepath.Abs(name)
	return abs
}
