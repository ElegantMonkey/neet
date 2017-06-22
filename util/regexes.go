package util

import (
	"regexp"
	"strings"
)

// MakeSearchRegex makes a *regexp.Regexp that behaves
// similar to a ILIKE %% SQL query when matched.
func MakeSearchRegex(search string) (*regexp.Regexp, error) {
	words := strings.Split(search, " ")
	// Add case-insensitiveness
	pattern := "(?i)"
	for _, word := range words {
		if word != "" {
			pattern += ".*" + word + ".*"
		}
	}

	return regexp.Compile(pattern)
}
