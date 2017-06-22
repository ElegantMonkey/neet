package util

import (
	"path/filepath"
)

var (
	videoExtensions = []string{
		".mp4",
		".mkv",
		".avi",
		".rmvb",
	}
)

func IsVideoFile(path string) bool {
	ext := filepath.Ext(path)
	for _, videoExt := range videoExtensions {
		if ext == videoExt {
			return true
		}
	}

	return false
}
