package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	TrueEnc := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	data, err := decodeHex(input)
	if err != nil {
		fmt.Printf("Failed to decode hex: %s", err)
		return
	}

	enc := base64Encode(data)
	fmt.Printf("%s\n", enc)
	sEnc := string(enc)
	if sEnc != TrueEnc {
		fmt.Println("Output is wrong!")
	} else {
		fmt.Println("Output is True!")
	}
}
func decodeHex(input []byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(data, input)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}
