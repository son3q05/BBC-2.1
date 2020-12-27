package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	in1, _ := decodeHex([]byte("1c0111001f010100061a024b53535009181c"))

	in2, _ := decodeHex([]byte("686974207468652062756c6c277320657965"))

	trueOut := "746865206b696420646f6e277420706c6179"

	afterXOR, err := fixedXor(in1, in2)
	if err != nil {
		fmt.Println(err)
		return
	}

	out := encodeHex(afterXOR)
	fmt.Println(out)

	if out != trueOut {
		fmt.Println("Output is Wrong")
	} else {
		fmt.Println("Output is True")
	}

}
func decodeHex(hexBytes []byte) ([]byte, error) {
	res := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(res, hexBytes)
	if err != nil {
		fmt.Println("Failed Decode")
		return nil, err
	}
	return res, nil
}

func encodeHex(input []byte) string {
	res := hex.EncodeToString(input)
	return res
}

func fixedXor(input1, input2 []byte) ([]byte, error) {
	if len(input1) != len(input2) {
		return nil, errors.New("Two input have different length")
	}
	res := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		res[i] = input1[i] ^ input2[i]
	}
	return res, nil
}
