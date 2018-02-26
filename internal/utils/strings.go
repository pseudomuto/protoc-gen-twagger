package utils

import (
	"strings"
)

// FirstParagraph returns everything up to the first line break, with its whitespace trimmed off.
func FirstParagraph(str string) string {
	return strings.TrimSpace(strings.SplitN(str, "\n", 2)[0])
}

// LastSubstring splits str by the separator and returns the last substring in the result with its whitespace trimmed
// off.
func LastSubstring(str, sep string) string {
	substrs := strings.Split(str, sep)
	return strings.TrimSpace(substrs[len(substrs)-1])
}

// Description return a copy of str with the prefix of `REQUIRED:` removed and its whitespace trimmed off.
func Description(str string) string {
	return strings.TrimSpace(strings.TrimPrefix(str, "REQUIRED:"))
}
