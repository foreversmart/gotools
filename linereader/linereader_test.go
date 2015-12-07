package linereader

import (
	"bytes"
	// "log"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LineReader(t *testing.T) {
	var (
		assertion = assert.New(t)
		content   = "odododowuidiiooooa\n"
	)

	stream := bytes.NewBufferString("")
	lineReader := NewLineReader(stream)
	go lineReader.Reading()

	// write test data
	go func() {
		for {
			stream.WriteString(content)
			// log.Println("writing")
			runtime.Gosched()
		}
	}()

	line := lineReader.Line()

	assertion.Equal(content, line)

	line = lineReader.Line()

	assertion.Equal(content, line)
}
