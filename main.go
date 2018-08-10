package main

import (
	"fmt"
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
)

func init() {
	// Populate 'language' from the environment.
	language, ok := os.LookupEnv("LANGUAGE")
	if !ok {
		fmt.Println("LANGUAGE must be set")
		os.Exit(1)
	}

	// Build i18n dictionary
	err := i18n.UseLanguage(language)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println(i18n.Messages["main"])
}
