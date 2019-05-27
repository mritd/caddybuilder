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

var (
	Tag        string
	PluginList string
	ModCmd     string
	ExtJson    string
	OutPut     string
	Debug      bool
)
