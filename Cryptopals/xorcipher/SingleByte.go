package xorcipher

import (
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	DecMsg, score, err := SingleXorDecrypt(input)
	if err != nil {
		fmt.Println("Failed to Decrypt")
	} else {
		fmt.Printf("Derypted message: %s\nScore: %f\n", DecMsg, score)
	}
}

func SingleXorDecrypt(codedMsg []byte) ([]byte, float32, error) {
	res, err := DecodeHex(codedMsg)
	if err != nil {
		return nil, 0, err
	}
	var answer []byte
	var score float32
	for i := 0; i < 256; i++ {
		r := make([]byte, len(res))
		var s float32
		for j := 0; j < len(res); j++ {
			c := res[j] ^ byte(i)
			s += GetCharFrequency(c)
			r[j] = c
		}
		if s > score {
			answer = r
			score = s
		}
		s = 0
	}
	return answer, score, nil
}

func SingleXorDecryptBytes(codedMsg []byte) ([]byte, float32, error) { // Xor with input is normal Slice of Bytes
	var answer []byte
	var score float32
	for i := 0; i < 256; i++ {
		r := make([]byte, len(codedMsg))
		var s float32
		for j := 0; j < len(codedMsg); j++ {
			c := codedMsg[j] ^ byte(i)
			s += GetCharFrequency(c)
			r[j] = c
		}
		if s > score {
			answer = r
			score = s
		}
		s = 0
	}
	return answer, score, nil
}

func GetCharFrequency(char byte) float32 {
	wf := map[byte]float32{ //get frequencies from Wikipedia, except ' ' estimated equal 6.5
		byte('A'): 8.2, byte('J'): 0.15, byte('S'): 6.3,
		byte('a'): 8.2, byte('j'): 0.15, byte('s'): 8.3,
		byte('B'): 1.5, byte('K'): 0.77, byte('T'): 9.1,
		byte('b'): 1.5, byte('k'): 0.77, byte('t'): 9.1,
		byte('C'): 2.8, byte('L'): 4, byte('U'): 2.8,
		byte('c'): 2.8, byte('l'): 4, byte('u'): 2.8,
		byte('D'): 4.3, byte('M'): 2.4, byte('V'): 0.98,
		byte('d'): 4.3, byte('m'): 2.4, byte('v'): 0.98,
		byte('E'): 13, byte('N'): 6.7, byte('W'): 2.4,
		byte('e'): 13, byte('n'): 6.7, byte('w'): 2.4,
		byte('F'): 2.2, byte('O'): 7.5, byte('X'): 0.15,
		byte('f'): 2.2, byte('o'): 7.5, byte('x'): 0.15,
		byte('G'): 2, byte('P'): 1.9, byte('Y'): 2,
		byte('g'): 2, byte('p'): 1.9, byte('y'): 2,
		byte('H'): 6.1, byte('Q'): 0.095, byte('Z'): 0.074,
		byte('h'): 6.1, byte('q'): 0.095, byte('z'): 0.074,
		byte('I'): 7, byte('R'): 6, byte(' '): 6.5,
		byte('i'): 7, byte('r'): 6,
	}
	return wf[char]
}

func DecodeHex(hexBytes []byte) ([]byte, error) {
	res := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(res, hexBytes)
	if err != nil {
		fmt.Println("Failed Decode")
		return nil, err
	}
	return res, nil
}
