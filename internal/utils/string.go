package utils

import (
	"errors"
	"strings"
)

// SplitStringIntoChunks splits a given string `s` into smaller chunks of at most `maxLen` characters.
// It returns a slice of strings containing the chunks.
// If `maxLen` is less than or equal to zero, it returns a slice with an error message.
// If `s` is empty, it returns a slice containing an empty string.
func SplitStringIntoChunks(s string, maxLen int) []string {
	if maxLen <= 0 {
		return []string{"Invalid max length"}
	}

	if s == "" {
		return []string{s}
	}

	var result []string
	for i := 0; i < len(s); i += maxLen {
		end := i + maxLen
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}

// parseLyricsCommand parses a lyrics command from the given input string.
// The expected format is: "!lyrics <song> by <artist>".
//
// It returns:
// - ping: The command prefix (e.g., "!lyrics").
// - song: The extracted song title.
// - artist: The extracted artist name.
// - err: An error if the input format is invalid.
//
// Possible errors include:
// - Missing or improperly formatted command prefix.
// - Absence of the "by" separator between song and artist.
// - Missing song title or artist name.
func ParseLyricsCommand(input string) (ping, song, artist string, err error) {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) < 2 || !strings.HasPrefix(parts[0], "!lyrics") {
		return "", "", "", errors.New("invalid format, expected '!lyrics song by artist'")
	}

	ping = parts[0]

	remaining := strings.TrimSpace(parts[1])
	separator := " by "
	index := strings.LastIndex(remaining, separator)
	if index == -1 {
		return "", "", "", errors.New("missing 'by' separator between song and artist")
	}

	song = strings.TrimSpace(remaining[:index])
	artist = strings.TrimSpace(remaining[index+len(separator):])

	if song == "" || artist == "" {
		return "", "", "", errors.New("song or artist is missing")
	}

	return ping, song, artist, nil
}
