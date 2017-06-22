package title_parser

import (
	"regexp"
	"strconv"
)

// literally stolen from flexget
func RegexpNotInWord(pattern string) *regexp.Regexp {
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
