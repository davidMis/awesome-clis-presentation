package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"time"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(eatCmd)

	eatCmd.PersistentFlags().Bool("continuous", false, "Eat continuously")
	viper.BindPFlags(eatCmd.PersistentFlags())
}

var eatCmd = &cobra.Command{
	Use:   "eat",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		food := viper.GetString("food")

		if viper.GetBool("continuous") {
			for range time.Tick(time.Second) {
				eat(food)
			}
		} else {
			eat(food)
		}
	},
}

func eat(food string) {
	if food == "cookie" {
		fmt.Printf(i18n.Messages["eatCookie"])
	} else {
		fmt.Printf(i18n.Messages["eatOther"], food)
	}
}
