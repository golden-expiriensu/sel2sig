package search

import (
	"errors"
	"os"
	"path"
	"sync"
)

type request struct {
	path     string
	selector [4]byte
	results  chan Result
	errors   chan error
}

func SearchDirectory(dir string, selector [4]byte) (Result, error) {
	req := request{dir, selector, make(chan Result), make(chan error)}

	done := make(chan struct{})
	go func() {
		search(req)
		close(done)
	}()

wait:
	for {
		select {
		case err := <-req.errors:
			if !errors.Is(err, ErrNotFound) {
				return nil, err
			}
		case res := <-req.results:
			return res, nil
		case <-done:
			break wait
		}
	}

	return nil, ErrNotFound
}

func search(req request) {
	baseDir, err := os.ReadDir(req.path)
	if err != nil {
		req.errors <- err
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(baseDir))

	for _, entry := range baseDir {
		go func(req request, entry os.DirEntry) {
			defer wg.Done()
			req.path = path.Join(req.path, entry.Name())

			if entry.IsDir() {
				search(req)
			} else if IsArtifactFile(req.path) {
				result, err := SearchFile(req.path, req.selector)
				if err != nil {
					req.errors <- err
					return
				}
				req.results <- result
			}
		}(req, entry)
	}

	wg.Wait()
}
