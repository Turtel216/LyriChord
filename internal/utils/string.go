package utils

func SplitStringIntoChunks(s string, maxLen int) []string {
	if maxLen <= 0 {
		return []string{"Invalid max length"}
	}

	if s == "" {
		return []string{""}
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
