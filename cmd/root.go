package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi"
)

var (
	rootCmd = &cobra.Command{
		Use:   "base",
		Short: "API Services",
		Long:  `boiler plate for API using OAPI Code gen`,
		Run: func(cmd *cobra.Command, args []string) {
			initialiseServer()
		},
	}
	env    string
	config *openapi.Config
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initialiseConfigs)
	rootCmd.PersistentFlags().StringVar(&env, "env", "", "environment to run the program")
}

func initialiseConfigs() {
	fmt.Println("initialising Configs...")
	var cfgFile = "config"
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	if strings.TrimSpace(env) != "" {
		cfgFile += "." + env
	}
	viper.SetConfigName(cfgFile)

	viper.AutomaticEnv()

	config = new(openapi.Config)
	fmt.Println("reading Configs...")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
	fmt.Println("unmarshalling Configs...")
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Sprintf("unable to unmarshal the configuration: %+v", err))
	}
}
