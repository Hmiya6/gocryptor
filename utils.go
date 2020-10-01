package main

import (
	"os"
	"path/filepath"
)

const (
	newExtention = ".cry"
	workerNum    = 20
)

func enumFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// pass directory
		if info.IsDir() {
			return nil
		}
		// pass itself
		// [2:] strips "./" in "./yourExecutiveFile"
		if info.Name() == os.Args[0][2:] {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
