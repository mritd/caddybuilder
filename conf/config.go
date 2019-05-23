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
var PluginJsonFiles = []string{"dns_plugins.json", "http_plugins.json", "others_plugins.json"}

var Tag string
var PluginList string
var ModPatch string
var ExtJson string
var OutPut string
var Debug bool
