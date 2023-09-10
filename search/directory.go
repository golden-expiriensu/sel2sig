package search

import (
	"errors"
	"os"
	"path"
)

func SearchDirectory(dirPath string, selector [4]byte) (string, error) {
	baseDir, err := os.ReadDir(dirPath)
	if err != nil {
		return "", err
	}

	for _, entry := range baseDir {
		path := path.Join(dirPath, entry.Name())

		var result string
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
		if err != nil {
			return "", err
		}
		return result, nil
	}

	return "", ErrNotFound
}
