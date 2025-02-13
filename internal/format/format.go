// Package format provides utility functions for formatting strings as Discord markdown messages.
// These functions help structure text in a visually appealing way for Discord messages,
// such as formatting song details and error messages with appropriate markdown headers.
package format

import "fmt"

// FormatSong formats a song title, artist, and lyrics into a Discord markdown message.
// It returns a formatted string where:
//   - The title is prefixed with a level 1 markdown header (#)
//   - The artist is prefixed with a level 2 markdown header (##)
//   - The lyrics follow as plain text.
//
// Example:
//
//	formatted := FormatSong("Song Title", "Artist Name", "Song lyrics...")
//	// Output:
//	// # Song Title
//	// ## By Artist Name
//	//
//	// Song lyrics...
func FormatSong(title, artist, lyrics string) string {
	return fmt.Sprintf("# %s\n## By %s\n\n%s", title, artist, lyrics)
}

// FormatError formats an error message for Discord markdown.
// It returns a formatted string where:
//   - The error message is prefixed with a level 1 markdown header (# Error!!!)
//   - The error message follows as plain text.
//
// Example:
//
//	formatted := FormatError("Something went wrong")
//	// Output:
//	// # Error!!!
//	// Something went wrong
func FormatError(msg string) string {
	return fmt.Sprintf("# Error!!!\n%s", msg)
}
