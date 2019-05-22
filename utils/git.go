package utils

import (
	"os"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"gopkg.in/src-d/go-git.v4"
)

const CaddyRepoAddr = "https://github.com/mholt/caddy.git"

func InitCaddyRepo(tag string) error {

	repoPath := GetCaddyRepoPath()

	// always delete caddy local repo dir
	err := os.RemoveAll(repoPath)
	if err != nil {
		return err
	}

	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:           CaddyRepoAddr,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewTagReferenceName(tag),
	})
	if err != nil {
		return err
	}
	return nil
}
