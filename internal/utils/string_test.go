package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitStringIntoChunks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected []string
		maxLen   int
	}{
		{"HelloWorld", []string{"Hello", "World"}, 5},
		{"GolangTest", []string{"Gol", "ang", "Tes", "t"}, 3},
		{"Short", []string{"Short"}, 10},
		{"", []string{""}, 5},
		{"NegativeTest", []string{"Invalid max length"}, -1},
	}

	for _, test := range tests {
		result := SplitStringIntoChunks(test.input, test.maxLen)
		assert.Equal(test.expected, result, "Test failed for input: %s", test.input)
	}
}
