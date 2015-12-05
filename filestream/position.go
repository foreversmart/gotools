package filestream

import (
	"io"
	"os"
)

type Positioner interface {
	SeekInFile() (io.Reader, error)
	Move(int)
}

type Position struct {
	SeekPos int64  `json:"seek_position"`
	Path    string `json:"path"`
}

func (pos *Position) SeekInFile() (reader io.Reader, err error) {
	fd, err := os.Open(pos.Path)
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

func (pos *Position) Move(n int64) {
	pos.SeekPos += n
}
