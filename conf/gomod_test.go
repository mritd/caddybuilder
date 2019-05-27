package conf

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"text/template"
)

var goModTpl = `package conf

var GoModCmds = []string{
{{ range . }}	"{{ . }}",
{{ end }}
}`

func TestGoModCodeGen(t *testing.T) {

	f, err := os.Open("../resources/mod_command")
	if err != nil {
		t.Fatal(err)
	}
	buf := bufio.NewReader(f)

	var cmds []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		cmds = append(cmds, line)
		if err != nil {
			if err == io.EOF {
				goto Gen
			}
			t.Fatal(err)
		}
	}

Gen:

	var codeBuf bytes.Buffer
	tpl, _ := template.New("").Parse(goModTpl)
	_ = tpl.Execute(&codeBuf, cmds)

	err = ioutil.WriteFile("gomod.go", codeBuf.Bytes(), 0644)
	if err != nil {
		t.Fatal(err)
	}
}
