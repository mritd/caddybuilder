package plugins

import (
	"strings"

	packr "github.com/gobuffalo/packr/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type Plugin struct {
	Name    string
	Package string
	Type    string
}

var jsonFiles = []string{"dns_plugins.json", "http_plugins.json", "others_plugins.json"}

var Plugins map[string]Plugin

func LoadPlugins(pluginsJson string) {
	var tmpPlugins []Plugin
	box := packr.New("resources", "../resources")

	if pluginsJson != "" {
		jsonFiles = append(jsonFiles, pluginsJson)
	}

	for _, j := range jsonFiles {

		bs, err := box.Find(j)
		if err != nil {
			logrus.Panic(err)
		}
		err = jsoniter.Unmarshal(bs, &tmpPlugins)
		if err != nil {
			logrus.Panic(err)
		}

		for _, p := range tmpPlugins {
			if _, ok := Plugins[strings.ToLower(p.Name)]; ok {
				logrus.Panicf("plugin [%s] already exist!", p.Name)
			}
			Plugins[strings.ToLower(p.Name)] = p
		}
	}
}
