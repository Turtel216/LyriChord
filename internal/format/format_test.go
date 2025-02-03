package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSong(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		artist   string
		lyrics   string
		expected string
	}{
		{
			name:     "All fields provided",
			title:    "Bohemian Rhapsody",
			artist:   "Queen",
			lyrics:   "Is this the real life? Is this just fantasy?",
			expected: "# Bohemian Rhapsody\n## By Queen\n\nIs this the real life? Is this just fantasy?",
		},
		{
			name:     "Empty title",
			title:    "",
			artist:   "Queen",
			lyrics:   "Is this the real life? Is this just fantasy?",
			expected: "# \n## By Queen\n\nIs this the real life? Is this just fantasy?",
		},
		{
			name:     "Empty artist",
			title:    "Bohemian Rhapsody",
			artist:   "",
			lyrics:   "Is this the real life? Is this just fantasy?",
			expected: "# Bohemian Rhapsody\n## By \n\nIs this the real life? Is this just fantasy?",
		},
		{
			name:     "Empty lyrics",
			title:    "Bohemian Rhapsody",
			artist:   "Queen",
			lyrics:   "",
			expected: "# Bohemian Rhapsody\n## By Queen\n\n",
		},
		{
			name:     "All fields empty",
			title:    "",
			artist:   "",
			lyrics:   "",
			expected: "# \n## By \n\n",
		},
		{
			name:     "Special characters in title and artist",
			title:    "Ænima",
			artist:   "Tool",
			lyrics:   "Learn to swim.",
			expected: "# Ænima\n## By Tool\n\nLearn to swim.",
		},
		{
			name:     "Multiline lyrics",
			title:    "Stairway to Heaven",
			artist:   "Led Zeppelin",
			lyrics:   "There's a lady who's sure\nAll that glitters is gold\nAnd she's buying a stairway to heaven.",
			expected: "# Stairway to Heaven\n## By Led Zeppelin\n\nThere's a lady who's sure\nAll that glitters is gold\nAnd she's buying a stairway to heaven.",
		},
		{
			name:     "Long title and artist",
			title:    "This Is a Very Long Song Title That Goes On and On",
			artist:   "This Is a Very Long Artist Name That Also Goes On and On",
			lyrics:   "Short lyrics.",
			expected: "# This Is a Very Long Song Title That Goes On and On\n## By This Is a Very Long Artist Name That Also Goes On and On\n\nShort lyrics.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSong(tt.title, tt.artist, tt.lyrics)
			assert.Equal(t, tt.expected, result, "FormatSong() = %v, want %v", result, tt.expected)
		})
	}
}

func TestFormatError(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected string
	}{
		{
			name:     "Given Message",
			msg:      "Some Error Message",
			expected: "# Error!!!\nSome Error Message",
		},
		{
			name:     "No Message given",
			msg:      "",
			expected: "# Error!!!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatError(tt.msg)
			assert.Equal(t, tt.expected, result, "FormatError() = %v, want %v", result, tt.expected)
		})
	}
}
