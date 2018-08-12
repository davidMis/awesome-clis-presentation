package main

import (
	"fmt"
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
	"github.com/fsnotify/fsnotify"
)

var flagEat bool
var flagPrepare bool
var flagFood string
var flagKind string
var flagContinuous bool

func init() {
	var err error

	viper.AutomaticEnv()
	viper.SetConfigName("conf")

	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		initDictionary()
	})
	initDictionary()

	// Define and parse flags
	flag.BoolVar(&flagEat, "eat", false, "Perform action: Eat the food")
	flag.BoolVar(&flagPrepare, "prepare", false, "Perform action: Prepare the food")
	flag.BoolVar(&flagContinuous, "continuous", false, "Continually eat")
	flag.StringVar(&flagFood, "food", "cookie", "What food to eat")

	// Use viper to parse flags instead of th Go standard library
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func initDictionary() {
	// Build i18n dictionary
	err := i18n.UseLanguage(viper.GetString("language"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	if flagEat {
		if flagContinuous {
			for range time.Tick(time.Second) {
				eat(flagFood, flagKind)
			}
		} else {
			eat(flagFood, flagKind)
		}
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
		fmt.Printf(i18n.Messages["eatCookie"])
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