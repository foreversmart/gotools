package fileswitcher

import (
	"fmt"
)

type NumberSwitcher struct {
	Pattern      string
	IsAscend     bool
	Index        int
	SpecialIndex int
	SpecialValue string

	Max int
	Min int
}

func NewNumberSwitcher(pattern string, index int, isAscend bool) *NumberSwitcher {
	return &NumberSwitcher{
		Pattern:  pattern,
		Index:    index,
		IsAscend: isAscend,
	}
}

func (switcher *NumberSwitcher) CurrentFile() string {
	if switcher.SpecialValue != "" && switcher.Index == switcher.SpecialIndex {
		return switcher.SpecialValue
	}

	return fmt.Sprintf(switcher.Pattern, switcher.Index)
}

func (switcher *NumberSwitcher) OlderFile() (string, error) {
	newIndex := switcher.Index
	if switcher.IsAscend {
		newIndex++
	} else {
		newIndex--
	}

	if !switcher.IsValid(newIndex) {
		return "", FileInvalidError
	}

	switcher.Index = newIndex

	return switcher.CurrentFile(), nil
}

func (switcher *NumberSwitcher) NewerFile() (string, error) {
	newIndex := switcher.Index
	if switcher.IsAscend {
		newIndex--
	} else {
		newIndex++
	}

	if !switcher.IsValid(newIndex) {
		return "", FileInvalidError
	}

	switcher.Index = newIndex

	return switcher.CurrentFile(), nil
}

func (switcher *NumberSwitcher) SetMax(max int) *NumberSwitcher {
	switcher.Max = max
	return switcher
}

func (switcher *NumberSwitcher) SetMin(min int) *NumberSwitcher {
	switcher.Min = min
	return switcher
}

func (switcher *NumberSwitcher) SetSpecial(spec int, value string) {
	switcher.SpecialIndex = spec
	switcher.SpecialValue = value
}

func (switcher *NumberSwitcher) IsValid(newIndex int) bool {
	// if max and min is not set then return true
	if switcher.Max == 0 && switcher.Min == 0 {
		return true
	}

	if newIndex > switcher.Max || newIndex < switcher.Min {
		return false
	}

	return true
}

func (switcher *NumberSwitcher) String() string {
	return switcher.CurrentFile()
}
