package linereader

import (
	"bytes"
	"io"
	"log"
	"sync"
	"time"
)

type LineReader struct {
	Reader     io.Reader
	Buffer     *bytes.Buffer
	CacheBytes []byte
	LeftBytes  []byte
	Mutex      sync.Mutex
}

func NewLineReader(reader io.Reader) *LineReader {
	return &LineReader{
		Reader:     reader,
		Buffer:     bytes.NewBufferString(""),
		CacheBytes: make([]byte, 0, 500),
		LeftBytes:  make([]byte, 0, 50),
	}
}

func (lineReader *LineReader) Line() string {
	line, err := lineReader.Buffer.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			lineReader.Mutex.Lock()
			lineReader.LeftBytes = append(lineReader.LeftBytes, line...)
			log.Println(line)
			lineReader.Mutex.Unlock()
		}
		log.Println("line string error:", err.Error())
		time.Sleep(time.Second * 1)
		return lineReader.Line()
	}

	return string(line)
}

func (lineReader *LineReader) Reading() {
	p := make([]byte, 1000)
	for {
		n, err := lineReader.Reader.Read(p)
		if err != nil {
			log.Println("line reading error%v", err)
		}

		if n > 0 {
			lineReader.WriteBuffer(p[:n])
		}

		// if there is not enough content to read
		if n < 300 {
			time.Sleep(time.Second * 1)
		}
	}
}

func (lineReader *LineReader) WriteBuffer(b []byte) {
	n := len(b)

	// Enlarge our cache bytes to match what we were handed, and copy
	// our current buffer over
	cbLen := len(lineReader.CacheBytes)
	if cap(lineReader.CacheBytes) < cap(b) {
		newBuf := make([]byte, cbLen, cap(b))
		copy(newBuf, lineReader.CacheBytes)
		lineReader.CacheBytes = newBuf
	}

	//TODO the Cache will lose by unexpect method log lose < 2 item
	// Write the current buffer if the new data won't fit
	if cbLen+n > cap(lineReader.CacheBytes) {
		lineReader.Mutex.Lock()
		lineReader.Buffer.Write(lineReader.LeftBytes)
		lineReader.LeftBytes = make([]byte, 0, 50)
		lineReader.Buffer.Write(lineReader.CacheBytes[:cbLen])
		lineReader.CacheBytes = make([]byte, 0, cap(b))
		cbLen = 0
		lineReader.Mutex.Unlock()
	}

	// Copy the new data into our saveBuffer
	lineReader.CacheBytes = lineReader.CacheBytes[:cbLen+n]
	copy(lineReader.CacheBytes[cbLen:], b)
}
