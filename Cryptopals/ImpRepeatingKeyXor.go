package main

import (
	"fmt"
	"xorcipher"
)

func main() {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	TrueEnc := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	res := xorcipher.EncryptRepeatingKeyXOR(plaintext, []byte("ICE"))

	EncText, err := xorcipher.HexaEncode(res)
	if err != 0 {
		fmt.Printf("%s\n", EncText)
		if string(TrueEnc) == string(EncText) {
			fmt.Println("Encrypted Text is true")
		} else {
			fmt.Println("Encrypted Text is wrong")
		}
	}

}
