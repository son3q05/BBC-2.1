package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz+-/=" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharset(length, charset)
}

func RandText(text string, value int) string {
	var encText string
	if value%2 == 0 {
		encText = hex.EncodeToString([]byte(text))
	} else {
		encText = base64.StdEncoding.EncodeToString([]byte(text))
	}
	return encText
}

type Page struct {
	Text string
	Time string
}

func main() {
	var plaintext string
	msg := "1001010011110000"
	for _, a := range msg {
		if a == '1' {
			plaintext += RandText(RandString(15), 1)
		} else {
			plaintext += RandText(RandString(10), 0)
		}
	}

	welcome := Page{plaintext, time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start web server, set port 8080
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
