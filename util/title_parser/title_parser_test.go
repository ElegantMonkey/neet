package title_parser

import (
	"testing"
)

type episodeTestData struct {
	Path     string
	Expected int
}

func TestEpisode(T *testing.T) {
	tests := []episodeTestData{
		{"[AwfulSubs] Animeme 09.mp4", 9},
		{"[Subs4Life] Cool Anime 01.mkv", 1},
		{"[Snowflakes]We_Like_Underscores_04_End.mkv", 4},
		{"[Snowflakes]CRC_Code_05[ABC323F].mp4", 5},
		{"Hello.mp4", -1},
	}

	for _, test := range tests {
		result := Episode(test.Path)
		if result != test.Expected {
			T.Fatalf("When parsing episode from '%s', got %d, expected %d.", test.Path, result, test.Expected)
			return
		}
	}
}
