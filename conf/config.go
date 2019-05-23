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

var PluginMap = make(map[string]Plugin)
var PluginJsonFiles = []string{"dns_plugins.json", "http_plugins.json", "other_plugins.json"}

var (
	Tag        string
	PluginList string
	ModPatch   string
	ExtJson    string
	OutPut     string
	Debug      bool
)
