package fileswitcher

import (
	"errors"
)

var (
	FileInvalidError = errors.New("file name is invalid")
	FileNoExsitError = errors.New("file no exsit")
)

type FileSwitcher interface {
	CurrentFile() string
	OlderFile() (string, error)
	NewerFile() (string, error)
}
