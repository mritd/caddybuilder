package builder

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

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

func Build(out string) error {

	if out == "" {
		if runtime.GOOS == "windows" {
			out = "./caddy.exe"
		} else {
			out = "./caddy"
		}
	}

	cmd := exec.Command("go", "build", "-o", out)
	cmd.Dir = filepath.Join(utils.GetCaddyRepoPath(), "caddy")
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+utils.GetGoPath())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
