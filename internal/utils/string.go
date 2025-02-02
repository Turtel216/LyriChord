package utils

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
