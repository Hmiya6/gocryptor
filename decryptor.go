package main

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io/ioutil"
)

func DecryptDir(root string, key string) error {
	files, err := enumFiles(root)
	for _, file := range files {
		err = DecryptFile(file, key)
		if err != nil {
			return err
		}
	}
	return nil
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
