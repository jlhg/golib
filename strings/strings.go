package strings

import (
	"regexp"
	"strings"
)

// ReduceSpace replaces all duplicated whitespaces with a single space character.
func ReduceSpace(s string) string {
	pattern := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(pattern.ReplaceAllString(s, " "))
}

// RemoveSpace removes all whitespaces.
func RemoveSpace(s string) string {
	pattern := regexp.MustCompile(`\s+`)
	return pattern.ReplaceAllString(s, "")
}
