package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/srijilv/go-api-template.git/api/pkg/interfaces/openapi"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "API",
		Short: "API Services",
		Run: func(cmd *cobra.Command, args []string) { // OnInitialize is called first
			s := openapi.NewServer(time.Now())
			err := openapi.RunHTTPServer(":8081", s, "/services")
			if err != nil {
				log.Fatal("Unable to start HTTP server:", err)
			}
			fmt.Println(viper.AllKeys())
		},
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
