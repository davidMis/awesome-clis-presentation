package main

import "fmt"

var msg = map[string]string{
	"EN": "I love cookies",
	"ES": "Yo quiero cookies",
}

var language string

func init() {
	language = "ES"
}

func main() {
	fmt.Println(msg[language])
}
