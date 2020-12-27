package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := ioutil.ReadFile("Cryptopals\\EncB64aes.txt")
	if err != nil {
		fmt.Println("Failed to open File")
		return
	}
	var cipherb []byte
	for _, line := range bytes.Split(f, []byte{'\n'}) {
		data, err := base64.StdEncoding.DecodeString(string(line))
		if err != nil {
			fmt.Println(err)
		}
		cipherb = append(cipherb, data...)
	}

	decOutput, _ := DecryptAESECB(cipherb, []byte("YELLOW SUBMARINE"))

	file, err := os.OpenFile("Cryptopals\\DecFileAes.txt", os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(string(decOutput)); err != nil {
		panic(err)
	} else {
		fmt.Println("Write to File succeed")
	}
}

func DecryptAESECB(ciphertext, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	plaintext = make([]byte, len(ciphertext))
	ptblock := plaintext
	for len(ciphertext) > 0 {
		block.Decrypt(ptblock, ciphertext[:bs])
		ptblock = ptblock[bs:]
		ciphertext = ciphertext[bs:]
	}
	return plaintext, nil
}
