package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// wrapper function to get password
func setAESKey() string {
	var (
		envKey string
		pass   string
	)
	flag.StringVar(&envKey, "e", "", "environment key")
	flag.StringVar(&pass, "p", "", "password")
	flag.Parse()

	if pass != "" {
		return pass
	} else if envKey != "" {
		pass = getKeyFromEnv(envKey)
		if pass != "" {
			return pass
		}
	}
	pass, err := getKeyFromStdin()
	if err != nil {
		log.Fatal("Error setting encryption key:", err)
	}
	return padKey(pass)
}

// set password from environment variables
func getKeyFromEnv(envKey string) string {
	key := os.Getenv(envKey)
	return key
}

func getKeyFromStdin() (string, error) {
	var key string
	for {
		// set password
		fmt.Print("Set your encryption key: ")
		key = scanPasswd()
		if len(key) == 0 {
			continue
		}

		// password confirmation
		fmt.Print("Enter your key again: ")
		rekey := scanPasswd()
		if key != rekey {
			fmt.Println("Error: different input detected")
			continue
		} else {
			break
		}
	}
	return key, nil
}

// scan user input as password
func scanPasswd() string {
	var passwd string

	fmt.Print("\033[8m") // hide user input

	// scan input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		passwd = scanner.Text()
	}

	fmt.Print("\033[28m") // show user input again

	return passwd
}

func padKey(key string) string {
	keyLen := len(key)
	blockLen := 32
	if keyLen < blockLen {
		key += strings.Repeat("X", blockLen-keyLen)
	} else if keyLen > blockLen {
		key = key[:blockLen]
	}
	return key
}
