package main

import (
	"crypto/aes"
	"fmt"
)

var key = []byte("1234567890123456")
var cipherText = make([]byte, 16)
var cipher, _ = aes.NewCipher(key)

func main() {
	Encryption([]byte("Hi There, World!"))
	Decryption()
}

func Encryption(data []byte) {
	cipher.Encrypt(cipherText, data)
	fmt.Printf("CipherText: %x\n", cipherText)
}

func Decryption() {
	decrypted := []byte("------------------")
	cipher.Decrypt(decrypted, cipherText)
	fmt.Printf("DecryptedText: %s\n", decrypted)
}
