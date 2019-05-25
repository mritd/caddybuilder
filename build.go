package main

import (
	"strings"

	"github.com/mritd/caddybuilder/builder"
	"github.com/mritd/caddybuilder/conf"
	"github.com/mritd/caddybuilder/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build caddy",
	Long: `
Build caddy`,
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

		// patch go mod dependencies
		logrus.Info("patch go mod...")
		err = builder.PatchDep()
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
	buildCmd.PersistentFlags().StringVarP(&conf.Tag, "tag", "t", "v1.0.0", "caddy tag")
	buildCmd.PersistentFlags().StringVarP(&conf.PluginList, "plugins", "p", "all", "comma separated list of caddy builder")
	buildCmd.PersistentFlags().StringVarP(&conf.OutPut, "output", "o", "", "caddy binary output path")
	buildCmd.PersistentFlags().StringVarP(&conf.ExtJson, "extjson", "j", "", "extended caddy plugins json file")
	buildCmd.PersistentFlags().StringVarP(&conf.ModCmd, "modcmd", "", "", "custom go mod command file for handling special dependencies")
	rootCmd.AddCommand(buildCmd)
}
