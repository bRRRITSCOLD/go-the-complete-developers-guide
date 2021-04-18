package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func (englishBot) getGreeting() string {
// 	return "Hi there!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Printf("%+v", resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
}
