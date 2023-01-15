package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "register",
		Short: "register an account",
		Long:  "",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initialiseConfigs)
	fmt.Println("initialising...")
}

func initialiseConfigs() {
	fmt.Println("initialiseConfigs...")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("here am i")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("configs")
	}
	viper.AutomaticEnv()
}
