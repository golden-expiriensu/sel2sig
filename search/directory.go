package search

import (
	"errors"
	"os"
	"path"
)

func SearchDirectory(dirPath string, selector [4]byte) (Result, error) {
	baseDir, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range baseDir {
		path := path.Join(dirPath, entry.Name())

		var result Result
		if entry.IsDir() {
			result, err = SearchDirectory(path, selector)
		} else {
			if !IsArtifactFile(entry.Name()) {
				continue
			}
			result, err = SearchFile(path, selector)
		}

		if errors.Is(err, ErrNotFound) {
			continue
		}
		return result, err
	}

	return nil, ErrNotFound
}
