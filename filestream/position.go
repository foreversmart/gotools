package filestream

import (
	"io"
	"os"

	"github.com/foreversmart/gotools/fileswitcher"
)

type Positioner interface {
	SeekInFile() (io.Reader, error)
	Move(int64)
}

type Position struct {
	SeekPos  int64                     `json:"seek_position"`
	FileSize int64                     `json:"file_size"`
	Switcher fileswitcher.FileSwitcher `json:"path"`
}

func NewPosition(switcher fileswitcher.FileSwitcher) *Position {
	return &Position{
		SeekPos:  0,
		FileSize: 0,
		Switcher: switcher,
	}
}

func (pos *Position) SeekInFile() (reader io.Reader, err error) {
	fileInfo, err := os.Stat(pos.Switcher.CurrentFile())
	if err != nil {
		return
	}

	// if current file Modtime is newer than hold file indicate files switched
	if fileInfo.Size() < pos.FileSize {
		_, err = pos.Switcher.OlderFile()
		if err != nil {
			return
		}

		// switch to older file success update fileinfo
		fileInfo, err = os.Stat(pos.Switcher.CurrentFile())
		if err != nil {
			return
		}
	}

	// if current pos is tail of the file move to new file
	if pos.SeekPos >= fileInfo.Size() {
		_, err = pos.Switcher.NewerFile()

		if err == nil {
			// switch to newer file success update fileinfo
			fileInfo, err = os.Stat(pos.Switcher.CurrentFile())
			if err != nil {
				return
			}

			// update seekPos to zero
			pos.SeekPos = 0
		}
	}

	// update file size
	pos.FileSize = fileInfo.Size()

	fd, err := os.Open(pos.Switcher.CurrentFile())
	if err != nil {
		return
	}

	reader = fd

	// Try to get to our seek position, if our seek is 0, then start at the
	// beginning.
	if pos.SeekPos == 0 {
		return
	}

	_, err = fd.Seek(pos.SeekPos, 0)
	if err != nil {
		return
	}

	return
}

func (pos *Position) UpdateFileInfo() error {
	fileInfo, err := os.Stat(pos.Switcher.CurrentFile())
	if err != nil {
		return err
	}

	pos.FileSize = fileInfo.Size()

	return nil
}

func (pos *Position) Move(n int64) {
	pos.SeekPos += n
}
