package plugins

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/mritd/caddybuilder/utils"

	"github.com/gobuffalo/packr/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const pluginCodeTpl = `package caddybuilder

import ({{ range . }}
	_ "{{ .Package }}"{{ end }}
)
`

type Plugin struct {
	Name    string
	Package string
	Type    string
}

var jsonFiles = []string{"dns_plugins.json", "http_plugins.json", "others_plugins.json"}
var initPluginOnce sync.Once

var Plugins = make(map[string]Plugin)

func init() {
	initPluginOnce.Do(func() {
		var tmpPlugins []Plugin
		box := packr.New("resources", "../resources")

		for _, j := range jsonFiles {

			bs, err := box.Find(j)
			utils.CheckAndExit(err)
			err = jsoniter.Unmarshal(bs, &tmpPlugins)
			utils.CheckAndExit(err)

			for _, p := range tmpPlugins {
				logrus.Debugf("load plugin [%s]", p.Name)
				Plugins[strings.ToLower(p.Name)] = p
			}
		}
	})
}

func FindPlugins(names ...string) []Plugin {
	var ps []Plugin
	for _, name := range names {
		p, ok := Plugins[strings.ToLower(name)]
		if !ok {
			logrus.Errorf("could not found [%s] plugin", name)
		} else {
			ps = append(ps, p)
		}
	}
	return ps
}

func GeneratePluginsCode(names ...string) error {

	if len(names) == 0 {
		logrus.Warn("no plugin name specified, skip generate plugins code!")
		return nil
	}

	pluginCodeDir := filepath.Join(utils.GetCaddyRepoPath(), "caddybuilder")
	err := os.RemoveAll(pluginCodeDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(pluginCodeDir, 0755)
	if err != nil {
		return err
	}

	tpl, err := template.New("").Parse(pluginCodeTpl)
	if err != nil {
		return err
	}

	targetCodeFile, err := os.OpenFile(filepath.Join(pluginCodeDir, "plugins.go"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = targetCodeFile.Close() }()

	err = tpl.Execute(targetCodeFile, FindPlugins(names...))
	if err != nil {
		return err
	}
	return nil
}
