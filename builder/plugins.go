package builder

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/emirpasic/gods/sets/treeset"

	"github.com/mritd/caddybuilder/conf"

	"github.com/mritd/caddybuilder/utils"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func Find(names ...string) []conf.Plugin {
	var ps []conf.Plugin

	if len(names) == 0 {
		logrus.Warn("no plugin name specified, skip find builder!")
		return ps
	}

	if len(names) == 1 && names[0] == "all" {
		for _, v := range conf.PluginsMap {
			ps = append(ps, v)
		}
	} else {
		for _, name := range names {
			p, ok := conf.PluginsMap[strings.ToLower(name)]
			if !ok {
				logrus.Errorf("could not found [%s] plugin", name)
			} else {
				ps = append(ps, p)
			}
		}
	}

	return ps
}

func Merge(extJson string) error {

	var tmpPlugins []conf.Plugin

	bs, err := ioutil.ReadFile(extJson)
	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(bs, &tmpPlugins)
	if err != nil {
		return err
	}

	for _, plugin := range tmpPlugins {
		_, ok := conf.PluginsMap[strings.ToLower(plugin.Name)]
		if ok {
			logrus.Warnf("plugin [%s] already exist, skip!", plugin.Name)
			continue
		}
		conf.PluginsMap[strings.ToLower(plugin.Name)] = plugin
		logrus.Infof("added plugin [%s]", plugin.Name)
	}
	return nil
}

func Sort() *treeset.Set {

	var plugins = treeset.NewWith(func(a, b interface{}) int {
		a1, b1 := a.(conf.Plugin), b.(conf.Plugin)

		switch {
		case a1.Type > b1.Type:
			return 1
		case a1.Type < b1.Type:
			return -1
		default:
			switch {
			case a1.Name > b1.Name:
				return 1
			case a1.Name < b1.Name:
				return -1
			default:
				return 0
			}
		}
	})

	for _, v := range conf.PluginsMap {
		plugins.Add(v)
	}

	return plugins
}

func GenerateCode(names ...string) error {

	if len(names) == 0 {
		logrus.Warn("no plugin name specified, skip generate builder code!")
		return nil
	}

	pluginCodeDir := filepath.Join(utils.GetCaddyRepoPath(), "caddy", "caddymain")

	tpl, err := template.New("").Parse(conf.PluginCodeTpl)
	if err != nil {
		return err
	}

	targetCodeFile, err := os.OpenFile(filepath.Join(pluginCodeDir, "builder.go"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = targetCodeFile.Close() }()

	err = tpl.Execute(targetCodeFile, Find(names...))
	if err != nil {
		return err
	}
	return nil
}
