package xorcipher

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
)

func EncryptRepeatingKeyXOR(input, key []byte) []byte { //Encrypt with the given key
	eb := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		eb[i] = input[i] ^ key[i%len(key)]
	}
	return eb
}

func HexaEncode(text []byte) ([]byte, int) { //Encode output for Human readable
	ret := make([]byte, hex.EncodedLen(len(text)))
	err := hex.Encode(ret, text)
	if err == 0 {
		fmt.Println("Error to Encode")
		return nil, err
	}
	return ret, err
}

func DecodeBase64(encodedBytes []byte) ([]byte, error) { //Decode input message
	retLen := base64.StdEncoding.DecodedLen(len(encodedBytes))
	ret := make([]byte, retLen)
	_, err := base64.StdEncoding.Decode(ret, encodedBytes)
	if err != nil {
		fmt.Println("Failed to decode Base64")
		return nil, err
	}
	return ret, nil
}

func getHamming(in1, in2 []byte) int {
	var hd int
	for a := 0; a < len(in1); a++ {
		for b := 0; b < 8; b++ {
			if in1[a]&(1<<uint(b)) != in2[a]&(1<<uint(b)) { //Compare bit to get number of differencies
				hd++
			}
		}
	}
	return hd
}

type keyLen struct { //Create Struct to guest Key Length
	length int
	score  float64
}

func getKeyLens(input []byte) []keyLen { //Create Slice of Struct KeyLen which sorted
	kl := make([]keyLen, 0)
	for len := 2; len <= 40; len++ { //Known key length [2,40]
		var hd float64           //Haming Distance
		for i := 0; i < 4; i++ { //get total HD of 4 block
			si := len * i
			hd += float64(getHamming(input[si:si+len], input[si+len:si+(len*2)]))
		}
		hd = hd / float64(len) //get Normalized Hamming Distance
		kl = append(kl, keyLen{score: hd, length: len})
	}
	sort.Slice(kl, func(i, j int) bool {
		return kl[i].score < kl[j].score
	})
	return kl
}

func breakTheSlice(keyLength int, encryptedSlBytes []byte) [][]byte { //Break the Slice, all byte xor with same single key will be in same group
	groups := make([][]byte, keyLength)
	for i, c := range encryptedSlBytes {
		groups[i%keyLength] = append(groups[i%keyLength], c)
	}
	return groups
}

func rebuildTheSlice(decGroups [][]byte) []byte { //Rebuild group into Slice
	reSlice := make([]byte, 0)
	for i := 0; i < len(decGroups[0]); i++ {
		for _, sl := range decGroups {
			if i < len(sl) {
				reSlice = append(reSlice, sl[i])
			}
		}
	}
	return reSlice
}

func XORCipher(groups [][]byte) [][]byte {
	res := make([][]byte, len(groups))
	for i, _ := range groups {
		res[i], _, _ = SingleXorDecryptBytes(groups[i])
	}
	return res
}

func DecryptRepeatingKeyXOR(encryptedInput []byte) []byte {
	kls := getKeyLens(encryptedInput)
	var ret []byte
	var retStrength float32
	for kli := 0; kli < 3; kli++ { //Try 3 Best KeyLengths
		kl := kls[kli].length
		groups := breakTheSlice(kl, encryptedInput)
		decGroups := XORCipher(groups)
		r := rebuildTheSlice(decGroups)
		var score float32
		for _, b := range r {
			score += GetCharFrequency(b)
		}
		if score > retStrength {
			ret = r
			retStrength = score
		}
	}
	return ret
}
