package title_parser

import (
	"testing"
)

type episodeTestData struct {
	Path string
	Expected int
}

func TestEpisode(T *testing.T) {
	tests := []episodeTestData{
		episodeTestData{"[AwfulSubs] Animeme 09.mp4", 9},
		episodeTestData{"[Subs4Life] Cool Anime 01.mkv", 1},
		episodeTestData{"[Snowflakes]We_Like_Underscores_04_End.mkv", 4},
		episodeTestData{"[Snowflakes]CRC_Code_05[ABC323F].mp4", 5},
		episodeTestData{"Hello.mp4", -1},
	}

	for _, test := range tests {
		result := Episode(test.Path)
		if result != test.Expected {
			T.Fatalf("When parsing episode from '%s', got %d, expected %d.", test.Path, result, test.Expected)
			return
		}
	}
}