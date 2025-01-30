package truncate

import (
	"unicode"
)

// String truncates a string to a maximum length, adding "..." if the string is longer than the max length.
func String(text string, maxLen int) string {
	currentLen := len(text)
	lastSpaceIdx := maxLen
	len := 0

	if currentLen <= maxLen {
		return text
	}

	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIdx = i
		}

		len++
		if len == maxLen {
			break
		}
	}

	if lastSpaceIdx < maxLen {
		return text[:lastSpaceIdx] + "..."
	}

	return text
}
