package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(prepareCmd)
}

var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		prepare(viper.GetString("food"))
	},
}

func prepare(food string) {
	if food == "cookie" {
		fmt.Println(i18n.Messages["prepareCookie"])
	} else {
		fmt.Printf(i18n.Messages["prepareOther"], food)
	}
}
