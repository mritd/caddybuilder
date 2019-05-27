package builder

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mritd/caddybuilder/conf"

	"github.com/sirupsen/logrus"

	"github.com/mritd/caddybuilder/utils"
)

func InitDep(names ...string) error {
	if len(names) == 0 {
		logrus.Warn("no package specified, skip init dependencies!")
		return nil
	}

	for _, p := range Find(names...) {
		cmd := exec.Command("go", "get", p.Package)
		cmd.Dir = utils.GetCaddyRepoPath()
		cmd.Env = append(cmd.Env, os.Environ()...)
		cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+utils.GetGoPath())
		if conf.Debug {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func PatchDep() error {

	if conf.ModCmd != "" {
		f, err := os.Open(conf.ModCmd)
		if err != nil {
			return err
		}
		buf := bufio.NewReader(f)
		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					goto Patch
				}
				return err
			} else {
				conf.GoModCmds = append(conf.GoModCmds, strings.TrimSpace(line))
			}
		}
	}

Patch:

	for _, l := range conf.GoModCmds {
		cmds := strings.Fields(l)
		if len(cmds) == 0 {
			continue
		}

		cmd := exec.Command(cmds[0], cmds[1:]...)
		cmd.Dir = utils.GetCaddyRepoPath()
		cmd.Env = append(cmd.Env, os.Environ()...)
		cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+utils.GetGoPath())
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil

}

func Build(out string) error {

	if out == "" {
		if runtime.GOOS == "windows" {
			out = "./caddy.exe"
		} else {
			out = "./caddy"
		}
	}

	if !filepath.IsAbs(out) {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		out = filepath.Join(pwd, out)
	}

	cmd := exec.Command("go", "build", "-o", out)
	cmd.Dir = filepath.Join(utils.GetCaddyRepoPath(), "caddy")
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+utils.GetGoPath())
	if conf.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}
