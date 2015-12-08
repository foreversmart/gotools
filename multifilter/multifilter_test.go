package multifilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multifilter(t *testing.T) {
	var (
		assertion = assert.New(t)
	)

	count := 0
	multifilter := NewMultiFilter()
	multifilter.AddFilter(func(...interface{}) {
		count += 2
	})

	multifilter.AddFilter(func(params ...interface{}) {
		count += params[0].(int)
	})

	multifilter.Call(10)

	assertion.Equal(12, count)

}
