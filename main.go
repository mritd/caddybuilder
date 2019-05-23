package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mritd/caddybuilder/config"

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
		err := utils.InitCaddyRepo(config.Tag)
		utils.CheckAndExit(err)

		// get builder list
		ps := strings.Split(config.PluginList, ",")

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
		err = builder.Build(config.OutPut)
		utils.CheckAndExit(err)

		logrus.Infof("success!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.Tag, "tag", "t", "v1.0.0", "caddy tag")
	rootCmd.PersistentFlags().StringVarP(&config.PluginList, "plugins", "p", "realip,cache,ipfilter", "comma separated list of caddy builder")
	rootCmd.PersistentFlags().StringVarP(&config.OutPut, "output", "o", "", "caddy binary output path")
	rootCmd.PersistentFlags().StringVarP(&config.PluginsJson, "pluginsjson", "j", "", "extended caddy plugin list json file")
	rootCmd.PersistentFlags().BoolVarP(&config.Debug, "debug", "", false, "debug mode")
}

func initLog() {

	if config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.Infof("GOMAXPROCS: %d", runtime.NumCPU())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
