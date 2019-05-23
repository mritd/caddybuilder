package utils

import (
	"os"
	"os/exec"

	"github.com/mritd/caddybuilder/config"
)

const CaddyRepoAddr = "https://github.com/mholt/caddy.git"

func InitCaddyRepo(tag string) error {

	repoPath := GetCaddyRepoPath()

	// always delete caddy local repo dir
	err := os.RemoveAll(repoPath)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", CaddyRepoAddr, repoPath)
	if config.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "checkout", "tags/"+tag, "-b", tag)
	cmd.Dir = repoPath
	if config.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
