package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/davidMis/cookieMonster/lib/i18n"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

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
		configChange()
	})
	configChange()

	rootCmd.PersistentFlags().StringP("food", "f", "cookie", "What food to eat")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func configChange() {
	// Build i18n dictionary
	err := i18n.UseLanguage(viper.GetString("language"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "cookieMonster",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(i18n.Messages["main"])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}