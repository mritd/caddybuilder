package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

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
