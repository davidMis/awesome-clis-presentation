package main

import (
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"github.com/pkg/errors"
	"flag"
	"fmt"
)

var foodFlag string
var eatFlag bool

func init() {
	// Get the configuration from the environment
	language, ok := os.LookupEnv("LANGUAGE")
	if !ok {
		panic(errors.New("Expected language in env"))
	}

	// Configure the i18 library
	i18n.UseLanguage(language)

	// Define and parse flags
	flag.BoolVar(&eatFlag, "eat", false, "Perform the eat action")
	flag.StringVar(&foodFlag, "food", "cookie", "The food to eat")
	flag.Parse()
}

func main() {
	if eatFlag {
		eat(foodFlag)
	} else {
		fmt.Printf(i18n.Messages["main"])
	}
}

// eat instructs cookie monster to consume the specified food. If
// food is "cookie", then cookie monster is happy, otherwise he gets
// upset.
func eat(food string) {
	if food == "cookie" {
		fmt.Printf(i18n.Messages["eatCookie"])
	} else {
		fmt.Printf(i18n.Messages["eatOther"], food)
	}
}


