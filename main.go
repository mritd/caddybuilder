package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mritd/caddybuilder/builder"

	"github.com/sirupsen/logrus"

	"github.com/mritd/caddybuilder/utils"

	"github.com/spf13/cobra"
)

var tag string
var pluginList string
var outPut string
var pluginsJson string

var rootCmd = &cobra.Command{
	Use:   "caddybuilder",
	Short: "A simple build tool for caddy",
	Long: `
A simple build tool for caddy`,
	Run: func(cmd *cobra.Command, args []string) {

		// check golang is install
		logrus.Info("check go command...")
		if !utils.CheckCommand("go", "version") {
			logrus.Fatal("go command not found!")
		}

		// check git is install
		logrus.Info("check git command...")
		if !utils.CheckCommand("git", "version") {
			logrus.Fatal("git command not found!")
		}

		// clone caddy
		logrus.Info("clone caddy...")
		err := utils.InitCaddyRepo(tag)
		utils.CheckAndExit(err)

		// get builder list
		ps := strings.Split(pluginList, ",")

		// generate builder code
		logrus.Info("generate plugins code...")
		err = builder.GenerateCode(ps...)
		utils.CheckAndExit(err)

		// init go mod dependencies
		logrus.Info("init go mod...")
		err = builder.InitDep(ps...)
		utils.CheckAndExit(err)

		// build
		logrus.Info("building...")
		err = builder.Build(outPut)
		utils.CheckAndExit(err)

		logrus.Infof("build success!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "v1.0.0", "caddy tag")
	rootCmd.PersistentFlags().StringVarP(&pluginList, "plugins", "p", "realip,cache,ipfilter", "comma separated list of caddy builder")
	rootCmd.PersistentFlags().StringVarP(&outPut, "output", "o", "", "caddy binary output path")
	rootCmd.PersistentFlags().StringVarP(&pluginsJson, "pluginsjson", "j", "", "extended caddy plugin list json file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
