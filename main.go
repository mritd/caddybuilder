package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tag string
var plugins string
var pluginsJson string

var rootCmd = &cobra.Command{
	Use:   "caddybuilder",
	Short: "A simple build tool for caddy",
	Long: `
A simple build tool for caddy`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "v1.0.0", "caddy tag")
	rootCmd.PersistentFlags().StringVarP(&plugins, "plugins", "p", "", "comma separated list of caddy plugins")
	rootCmd.PersistentFlags().StringVarP(&pluginsJson, "pluginsjson", "j", "", "extended caddy plugin list json file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
