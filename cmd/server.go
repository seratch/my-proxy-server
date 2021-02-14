package cmd

import (
	"fmt"
	"github.com/elazarl/goproxy"
	goproxyAuth "github.com/elazarl/goproxy/ext/auth"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var port string
var authEnabled bool
var validUsername string
var validPassword string

var rootCmd = &cobra.Command{
	Use:   "my-proxy-server",
	Short: "Run a simple HTTP proxy server",
	Run: func(cmd *cobra.Command, _args []string) {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		if authEnabled {
			goproxyAuth.ProxyBasic(proxy, "my-realm", func(username, password string) bool {
				return validUsername == username &&
					validPassword == password
			})
			log.Print(fmt.Sprintf("localhost:%s (auth: %s:%s)", port, validUsername, validPassword))
		} else {
			log.Print(fmt.Sprintf("localhost:%s", port))
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), proxy))
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(
		&port,
		"port",
		"P",
		"9000",
		"HTTP proxy port (default: 9000)",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&authEnabled,
		"auth",
		"a",
		false,
		"Enable proxy auth (Basic Auth)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&validUsername,
		"username",
		"u",
		"user",
		"Basic auth validUsername (default: user)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&validPassword,
		"password",
		"p",
		"pass",
		"Basic auth validPassword (default: pass)",
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}
}
