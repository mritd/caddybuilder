package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mritd/caddybuilder/conf"

	"github.com/mritd/caddybuilder/builder"

	"github.com/sirupsen/logrus"

	"github.com/mritd/caddybuilder/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caddybuilder",
	Short: "A simple build tool for caddy",
	Long: `
A simple build tool for caddy`,
	Run: func(cmd *cobra.Command, args []string) {

		// init log
		initLog()

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
		err := utils.InitCaddyRepo(conf.Tag)
		utils.CheckAndExit(err)

		// get plugin list
		ps := strings.Split(conf.PluginList, ",")

		// merge plugin map
		if conf.ExtJson != "" {
			logrus.Info("use extended plugins json...")
			err = builder.Merge(conf.ExtJson)
			utils.CheckAndExit(err)
		}

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
		err = builder.Build(conf.OutPut)
		utils.CheckAndExit(err)

		logrus.Infof("success!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&conf.Tag, "tag", "t", "v1.0.0", "caddy tag")
	rootCmd.PersistentFlags().StringVarP(&conf.PluginList, "plugins", "p", "realip,cache,ipfilter", "comma separated list of caddy builder")
	rootCmd.PersistentFlags().StringVarP(&conf.OutPut, "output", "o", "", "caddy binary output path")
	rootCmd.PersistentFlags().StringVarP(&conf.ExtJson, "extjson", "j", "", "extended caddy plugins json file")
	rootCmd.PersistentFlags().BoolVarP(&conf.Debug, "debug", "", false, "debug mode")
}

func initLog() {

	if conf.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
