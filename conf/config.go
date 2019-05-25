package conf

type Plugin struct {
	Name    string
	Package string
	Type    string
}

const PluginCodeTpl = `package caddymain

import ({{ range . }}
	_ "{{ .Package }}"{{ end }}
)
`

type Plugins []Plugin

func (plugins Plugins) Len() int {
	return len(plugins)
}
func (plugins Plugins) Less(i, j int) bool {
	return plugins[i].Name < plugins[j].Name
}
func (plugins Plugins) Swap(i, j int) {
	plugins[i], plugins[j] = plugins[j], plugins[i]
}

var PluginMap = make(map[string]Plugin)
var PluginJsonFiles = []string{"dns_plugins.json", "http_plugins.json", "other_plugins.json"}

var (
	Tag        string
	PluginList string
	ModCmd     string
	ExtJson    string
	OutPut     string
	Debug      bool
)
