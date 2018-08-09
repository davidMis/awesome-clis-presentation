package main

import (
	"fmt"
	"os"
	. "github.com/davidMis/cookieMonster/lib"
)

// language is the 2-letter ISO language code provided from the environment
var language string

// messages is our i18n dictionary
var messages Messages

func init() {
	var ok bool
	var err error

	// Populate 'language' from the environment.
	language, ok = os.LookupEnv("LANGUAGE")
	if !ok {
		fmt.Println("LANGUAGE must be set")
		os.Exit(1)
	}

	// Build i18n dictionary
	messages, err = NewMessages(language)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println(messages["main"])
}
