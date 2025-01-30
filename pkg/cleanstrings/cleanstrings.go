package cleanstrings

import (
	"strings"
	"unicode"
)

func CleanWhitespace(text string) string {
	textArray := strings.FieldsFunc(text, func(r rune) bool {
		if r == '\n' {
			return false
		}

		return unicode.IsSpace(r)
	})

	textJoined := strings.Join(textArray, " ")

	return strings.ReplaceAll(textJoined, " \n", "\n")
}
