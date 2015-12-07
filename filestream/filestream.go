package filestream

import (
	"log"
)

type FileStream struct {
	Pos Positioner
}

func NewFileStream(pos Positioner) *FileStream {
	return &FileStream{
		Pos: pos,
	}
}

func (fileStream *FileStream) Read(p []byte) (n int, err error) {
	reader, err := fileStream.Pos.SeekInFile()

	if err != nil {
		log.Println(err.Error())
		return
	}

	n, err = reader.Read(p)

	if err != nil {
		return
	}

	// if read success move the position to new location
	fileStream.Pos.Move(int64(n))

	return
}
