package title_parser

import (
	"regexp"
	"strconv"
)

// RegexpNotInWord matches the pattern if and only if it
// isn't in a word (e.g. "\d" matches "A 1 B" but won't match
// "A1 B").
func RegexpNotInWord(pattern string) *regexp.Regexp {
	// literally stolen from flexget
	pattern = `(?:[\W_])` + pattern + `(?:[\W_])`
	return regexp.MustCompile(pattern)
}

var (
	episodeRegexes = []*regexp.Regexp{
		RegexpNotInWord(`(\d{1,3})`),
	}
)

// Episode parses the episode from the title. Returns -1 on failure.
func Episode(title string) int {
	for _, re := range episodeRegexes {
		matches := re.FindAllStringSubmatch(title, -1)
		if matches == nil {
			continue
		}

		for _, m := range matches {
			ep, err := strconv.Atoi(m[1])
			if err == nil {
				return ep
			}
		}
	}

	return -1
}
