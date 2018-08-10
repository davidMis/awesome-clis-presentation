package main

import (
	"fmt"
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"flag"
)

var flagEat bool
var flagPrepare bool
var flagFood string
var flagKind string

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

	// Define and parse flags
	flag.BoolVar(&flagEat, "eat", false, "Perform action: Eat the food")
	flag.BoolVar(&flagPrepare, "prepare", false, "Perform action: Prepare the food")
	flag.StringVar(&flagFood, "food", "cookie", "What food to eat")
	flag.StringVar(&flagKind, "kind", "chocolate chip", "What kind of cookie to eat")
	flag.Parse()
}

func main() {
	if flagEat {
		eat(flagFood, flagKind)
		return
	}

	if flagPrepare {
		prepare(flagFood)
		return
	}

	fmt.Println(i18n.Messages["main"])
}

func eat(food, kind string) {
	if food == "cookie" {
		fmt.Printf(i18n.Messages["eatCookie"], kind)
	} else {
		fmt.Printf(i18n.Messages["eatOther"], food)
	}
}

func prepare(food string) {
	if food == "cookie" {
		fmt.Println(i18n.Messages["prepareCookie"])
	} else {
		fmt.Printf(i18n.Messages["prepareOther"], food)
	}
}