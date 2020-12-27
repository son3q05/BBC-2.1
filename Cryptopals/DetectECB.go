package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

func main() {
	infile, _ := os.Open("Cryptopals\\HexECB.txt")
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	scanner.Split(bufio.ScanLines)

	maxDup := 0
	var bestText string

	for scanner.Scan() {
		line := scanner.Text()
		data, err := hex.DecodeString(line)
		if err != nil {
			fmt.Printf("hex.DecodeString error: %v", err)
		}
		if Dup := ecbDetector(data, 16); Dup > maxDup {
			maxDup = Dup
			bestText = string(line)
		}
	}

	fmt.Printf("Hex-Encoded text which is most likely to be encrypted in AES-ECB mode: \n%s\n block Duplication: %d\n", bestText, maxDup)
}

func ecbDetector(input []byte, blockSize int) int {
	maxDup := 0
	hist := make(map[string]int)
	for i := 0; i < len(input); i += blockSize {
		start, end := i, int(math.Min(float64(i+blockSize), float64(len(input))))
		keyMap := string(input[start:end])
		hist[keyMap]++
		if hist[keyMap] > maxDup {
			maxDup = hist[keyMap]
		}
	}
	return maxDup
}
