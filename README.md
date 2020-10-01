# gocryptor

A file encryptor/decryptor with golang standard library.  
AES (CFB) is used.

## Usage
Encrypting file may cause unpredictable system failure.  
Please execute it at your own risk.
```bash
$ go install github.com/Hmiya6/gocryptor/cmd/gocryptor
$ export PATH=~/go/bin:$PATH
$ gocryptor [target directory] [e: encryption]/[d: decryption]
```

## TODO
* concurrent en/de-cryption
* v2: Imprement RSA encryption
* Imprement read function
