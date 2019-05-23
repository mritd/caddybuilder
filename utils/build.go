package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func InitDep(packages ...string) {
	if len(packages) == 0 {
		logrus.Warn("no package specified, skip init dependencies!")
		return
	}

	for _, p := range packages {
		cmd := exec.Command("go", "get", p)
		cmd.Dir = GetCaddyRepoPath()
		cmd.Env = append(cmd.Env, os.Environ()...)
		cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+GetGoPath())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		CheckAndExit(err)
	}
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
	cmd.Dir = filepath.Join(GetCaddyRepoPath(), "caddy")
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPATH="+GetGoPath())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
