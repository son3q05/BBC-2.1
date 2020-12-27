package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"xorcipher"
)

func main() {
	f, err := ioutil.ReadFile("Cryptopals\\EncB64.txt")
	if err != nil {
		fmt.Println("Failed to open File")
		return
	}

	decB64, _ := xorcipher.DecodeBase64(f)
	plaintext, score := xorcipher.DecryptRepeatingKeyXOR(decB64)

	if score != float32(0) {
		fmt.Println("Decrypted Successfuly")
	}
	file, err := os.OpenFile("Cryptopals\\DecFile.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(string(plaintext)); err != nil {
		panic(err)
	} else {
		fmt.Println("Write to File succeed")
	}
}
