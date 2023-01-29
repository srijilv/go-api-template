package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/srijilv/go-api-template.git/pkg/interfaces/openapi"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the api server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v\n", args)
			initialiseServer()
		},
	}
)

func init() {
	// rootCmd.AddCommand(serveCmd)
}

func initialiseServer() {
	fmt.Println("starting server...")
	s := openapi.NewServer(time.Now())
	fmt.Printf("server is now listening to port: %d...", config.HTTP.Port)
	err := openapi.RunHTTPServer(fmt.Sprintf(":%d", config.HTTP.Port), s, config.HTTP.Prefix)
	if err != nil {
		log.Fatal("Unable to start HTTP server:", err)
	}

}
