package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/mritd/caddybuilder/builder"
	"github.com/mritd/caddybuilder/conf"
	"github.com/mritd/caddybuilder/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listTpl = `Name              Type              Repo
---------------------------------------------------------------------------------------
{{ range . }}{{ .Name | ListLayout }}{{ .Type | ListLayout }}{{ .Repo }}
{{ end }}`

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "List plugins",
	Long: `
List plugins.`,
	Run: func(cmd *cobra.Command, args []string) {
		// merge plugin map
		if conf.ExtJson != "" {
			logrus.Info("use extended plugins json...")
			err := builder.Merge(conf.ExtJson)
			utils.CheckAndExit(err)
		}

		var plugins []struct{ Name, Type, Repo string }

		builder.Sort().Each(func(index int, value interface{}) {
			v := value.(conf.Plugin)
			plugins = append(plugins, struct{ Name, Type, Repo string }{Name: v.Name, Type: v.Type, Repo: v.Package})
		})

		var buf bytes.Buffer

		t, err := template.New("").Funcs(map[string]interface{}{"ListLayout": listLayout}).Parse(listTpl)
		utils.CheckAndExit(err)

		err = t.Execute(&buf, plugins)
		utils.CheckAndExit(err)

		fmt.Println(buf.String())

	},
}

func listLayout(name string) string {
	if len(name) < 18 {
		return fmt.Sprintf("%-18s", name)
	} else {
		return fmt.Sprintf("%-18s", utils.ShortenString(name, 18))
	}
}

func init() {
	pluginsCmd.PersistentFlags().StringVarP(&conf.ExtJson, "extjson", "j", "", "extended caddy plugins json file")
	rootCmd.AddCommand(pluginsCmd)
}
