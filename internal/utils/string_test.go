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

func TestParseLyricsCommand(t *testing.T) {
	tests := []struct {
		input          string
		expectedPing   string
		expectedSong   string
		expectedArtist string
		expectError    bool
	}{
		{"!lyrics Bohemian Rhapsody by Queen", "!lyrics", "Bohemian Rhapsody", "Queen", false},
		{"!lyrics Shape of You by Ed Sheeran", "!lyrics", "Shape of You", "Ed Sheeran", false},
		{"!lyrics Hotel California", "", "", "", true},
		{"Bohemian Rhapsody by Queen", "", "", "", true},
		{"!lyrics by Queen", "", "", "", true},
		{"!lyrics Imagine by John Lennon", "!lyrics", "Imagine", "John Lennon", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ping, song, artist, err := parseLyricsCommand(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPing, ping)
				assert.Equal(t, tt.expectedSong, song)
				assert.Equal(t, tt.expectedArtist, artist)
			}
		})
	}
}
