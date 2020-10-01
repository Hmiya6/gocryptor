package main

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func DecryptDir(root string, key string) error {
	files, err := enumFiles(root)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	pathCh := make(chan string, workerNum)
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go decryptWorker(pathCh, key, &wg)
	}

	for _, file := range files {
		if filepath.Ext(file) != newExtention {
			continue
		}
		pathCh <- file
	}
	close(pathCh)
	wg.Wait()
	return nil
}

func decryptWorker(ch chan string, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		path, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(path)
		err := DecryptFile(path, key)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func checkCry(path string) {

}

func DecryptFile(filename string, key string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	data, err = AESdecode(data, key)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 555)
	if err != nil {
		return err
	}

	os.Rename(filename, strings.Trim(filename, newExtention))
	return nil
}

func AESdecode(data []byte, key string) ([]byte, error) {
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		err = errors.New("Given data too short")
		return nil, err
	}

	initVec := data[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, initVec)
	stream.XORKeyStream(data, data)
	return data[aes.BlockSize:], nil
}
