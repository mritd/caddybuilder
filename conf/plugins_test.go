package conf

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
	"text/template"

	jsoniter "github.com/json-iterator/go"
)

var pluginsTpl = `package conf

var DNSPlugins = []Plugin{
{{ range .DNSPlugins }}	{
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
}

var HttpPlugins = []Plugin{
{{ range .HttpPlugins }}	{
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
}

var OtherPlugins = []Plugin{
{{ range .OtherPlugins }}	{
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
}

var PluginsMap = map[string]Plugin{
{{ range .DNSPlugins }}	"{{ .Name | ToLower }}": {
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
{{ range .HttpPlugins }}	"{{ .Name | ToLower }}": {
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
{{ range .OtherPlugins }}	"{{ .Name | ToLower }}": {
		Name:    "{{ .Name }}",
		Type:    "{{ .Type }}",
		Package: "{{ .Package }}",
	},
{{ end }}
}`

func TestPluginsCodeGen(t *testing.T) {

	var ps struct{ DNSPlugins, HttpPlugins, OtherPlugins []Plugin }
	var bs []byte

	bs, _ = ioutil.ReadFile("../resources/dns_plugins.json")
	_ = jsoniter.Unmarshal(bs, &ps.DNSPlugins)

	bs, _ = ioutil.ReadFile("../resources/http_plugins.json")
	_ = jsoniter.Unmarshal(bs, &ps.HttpPlugins)

	bs, _ = ioutil.ReadFile("../resources/other_plugins.json")
	_ = jsoniter.Unmarshal(bs, &ps.OtherPlugins)

	tpl, _ := template.New("").Funcs(map[string]interface{}{"ToLower": strings.ToLower}).Parse(pluginsTpl)
	var buf bytes.Buffer
	_ = tpl.Execute(&buf, ps)

	err := ioutil.WriteFile("plugins.go", buf.Bytes(), 0644)
	if err != nil {
		t.Fatal(err)
	}

}
