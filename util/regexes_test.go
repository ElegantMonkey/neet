package util

import (
	"fmt"
	"regexp"
	"testing"
)

var (
	testStrings = []string{
		"k-on",
		"re:creators",
		"Shingeki no Kyojin",
		"shingeki no bahamut: virgin soul",
	}
)

func countMatches(re *regexp.Regexp) (n int) {
	for _, s := range testStrings {
		if re.MatchString(s) {
			n++
		}
	}

	return
}

func TestSearch(T *testing.T) {
	// search string -> expected num of matches
	searches := map[string]int{
		"re creators": 1,
		"shingeki": 2,
		"bahamut": 1,
	}

	for search, expected := range searches {
		re, err := MakeSearchRegex(search)
		if err != nil {
			T.Fatalf("Got error while making search regex for '%s': %v\n", search, err)
			return
		}
		matches := countMatches(re)

		if matches != expected {
			T.Fatalf("While searching for '%s', got %d matches, expected %d.", search, matches, expected)
		}
	}
}
