package cmd

import (
	"log"
	"os"

	"github.com/Hmiya6/gocryptor"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("invalid number of arguments: <target dir> [e] [d]")
	}
	root := os.Args[1]
	mode := os.Args[2][:1]
	passwd := gocryptor.setAESKey()

	var err error
	if mode == "e" {
		err = gocryptor.EncryptDir(root, passwd)
	} else if mode == "d" {
		err = gocryptor.DecryptDir(root, passwd)
	} else {
		log.Fatal("invalid method: [e] or [d]")
	}
	if err != nil {
		log.Fatal(err)
	}
}
