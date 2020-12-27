package xorcipher

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FindXorCipherString(InpFile string) ([]byte, error) {
	f, err := ioutil.ReadFile(InpFile)
	if err != nil {
		fmt.Println("Failed to open File")
		return nil, err
	}
	msg := strings.Split(string(f), "\r\n")
	var res []byte
	var score float32
	for _, m := range msg {
		str, s, err := SingleXorDecrypt([]byte(m))
		if err != nil {
			return nil, err
		}
		if s > score {
			res = str
			score = s
		}
	}
	return res, nil
}
