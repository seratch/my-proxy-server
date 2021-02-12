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
var username string
var password string

var rootCmd = &cobra.Command{
	Use:   "my-proxy-server",
	Short: "Run a simple HTTP proxy server",
	Run: func (cmd *cobra.Command, _args []string) {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		if authEnabled {
			goproxyAuth.ProxyBasic(proxy, "my-realm", func(user, pwd string) bool {
				return user == username && password == pwd
			})
			log.Print(fmt.Sprintf("localhost:%s (auth: %s:%s)", port, username, password))
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
		&username,
		"username",
		"u",
		"user",
		"Basic auth username (default: user)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&password,
		"password",
		"p",
		"pass",
		"Basic auth password (default: pass)",
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}
}