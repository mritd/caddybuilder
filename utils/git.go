package utils

import (
	"os"
	"os/exec"
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
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "checkout", "tags/"+tag, "-b", tag)
	cmd.Dir = repoPath
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
