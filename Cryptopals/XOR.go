package main

import (
	"fmt"
	"xorcipher"
)

func main() {
	test := "132f175e4c4af1120138e1f2085a3804471f5824555d083de6123f533123"
	_, err := xorcipher.DecodeHex([]byte(test))
	if err != nil {
		fmt.Println("Failed")
	}
	answer, err := xorcipher.FindXorCipherString("InputXOR.txt")
	if err != nil {
		fmt.Println("Failed to Find")
	} else {
		fmt.Printf("Derypted message found: %s\n", answer)
	}
}
