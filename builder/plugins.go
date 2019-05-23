package builder

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/mritd/caddybuilder/conf"

	"github.com/mritd/caddybuilder/utils"

	"github.com/gobuffalo/packr/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var initPluginOnce sync.Once

func init() {
	initPluginOnce.Do(func() {
		var tmpPlugins []conf.Plugin
		box := packr.New("resources", "../resources")

		for _, j := range conf.PluginJsonFiles {

			bs, err := box.Find(j)
			utils.CheckAndExit(err)
			err = jsoniter.Unmarshal(bs, &tmpPlugins)
			utils.CheckAndExit(err)

			for _, p := range tmpPlugins {
				logrus.Debugf("load plugin [%s]", p.Name)
				conf.PluginMap[strings.ToLower(p.Name)] = p
			}
		}
	})
}

func Find(names ...string) []conf.Plugin {
	var ps []conf.Plugin

	if len(names) == 0 {
		logrus.Warn("no plugin name specified, skip find builder!")
		return ps
	}

	if len(names) == 1 && names[0] == "all" {
		for _, v := range conf.PluginMap {
			ps = append(ps, v)
		}
	} else {
		for _, name := range names {
			p, ok := conf.PluginMap[strings.ToLower(name)]
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
		_, ok := conf.PluginMap[strings.ToLower(plugin.Name)]
		if ok {
			logrus.Warnf("plugin [%s] already exist, skip!")
			continue
		}
		conf.PluginMap[strings.ToLower(plugin.Name)] = plugin
		logrus.Infof("added plugin [%s]", plugin.Name)
	}
	return nil
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
