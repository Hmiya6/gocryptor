package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("invalid number of arguments: <target dir> [e] [d]")
	}
	root := os.Args[1]
	mode := os.Args[2][:1]
	passwd := setAESKey()

	var err error
	if mode == "e" {
		err = EncryptDir(root, passwd)
	} else if mode == "d" {
		err = DecryptDir(root, passwd)
	} else {
		log.Fatal("invalid method: [e] or [d]")
	}
	if err != nil {
		log.Fatal(err)
	}
}
