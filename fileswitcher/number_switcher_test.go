package fileswitcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberSwitcher(t *testing.T) {
	var (
		assertion = assert.New(t)
		pattern   = "message.log.%d"
	)

	switcher := NewNumberSwitcher(pattern, 0, true)
	switcher.SetMax(10).SetMin(0)

	olderFile, err := switcher.OlderFile()
	assertion.Nil(err)
	assertion.Equal("message.log.1", olderFile)

	olderFile, err = switcher.OlderFile()
	assertion.Nil(err)
	assertion.Equal("message.log.2", olderFile)

	currentFile := switcher.CurrentFile()
	assertion.Equal("message.log.2", currentFile)

	newerFile, err := switcher.NewerFile()
	assertion.Nil(err)
	assertion.Equal("message.log.1", newerFile)

}

func Test_NumberSwitcher_Invalid(t *testing.T) {
	var (
		assertion = assert.New(t)
		pattern   = "message.log.%d"
	)

	switcher := NewNumberSwitcher(pattern, 0, true)
	switcher.SetMax(10).SetMin(0)

	assertion.True(switcher.IsValid(5))
	assertion.True(switcher.IsValid(10))
	assertion.True(switcher.IsValid(0))
	assertion.False(switcher.IsValid(-1))
	assertion.False(switcher.IsValid(11))
}
