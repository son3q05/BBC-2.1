package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func isHex(text string) bool {
	if len(text)%4 != 0 {
		return true
	}

	if text[len(text)-1] == '=' {
		return false
	}

	for _, c := range []byte(text) {
		if !(c >= byte('A') && c <= byte('F') || c >= byte('a') && c <= byte('f') || c >= byte('0') && c <= byte('9')) {
			return false
		}
	}
	return true
}

func main() {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Can't connect to Server")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	tmp := "<p id=" + string('"') + "enc" + string('"') + ">"
	start := strings.Index(string(body), tmp) + 12
	end := strings.Index(string(body), "</p>")

	s := string(body[start:end])

	size := 20
	var plaintext string
	for i := 0; i < len(s); i += size {
		plaintext = string([]byte(s)[i : i+size])
		if isHex(plaintext) == true {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Printf("\n")
}
