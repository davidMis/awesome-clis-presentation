package main

import (
	"fmt"
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"github.com/pkg/errors"
)

func init() {
	// Get the configuration from the environment
	language, ok := os.LookupEnv("LANGUAGE")
	if !ok {
		panic(errors.New("Expected language in env"))
	}

	// Configure the i18 library
	i18n.UseLanguage(language)
}

func main() {
	fmt.Println(i18n.Messages["main"])
}


