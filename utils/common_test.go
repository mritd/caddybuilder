package utils

import "testing"

func TestCheckGoCommand(t *testing.T) {
	t.Log(CheckGoCommand())
}

func TestGetGoPath(t *testing.T) {
	goPath := GetGoPath()
	if goPath == "" {
		t.Fatal("get GOPATH failed")
	} else {
		t.Log(goPath)
	}
}

func TestGetCaddyRepoPath(t *testing.T) {
	caddyRepoPath := GetCaddyRepoPath()
	if caddyRepoPath == "" {
		t.Fatal("get caddy repo path failed")
	} else {
		t.Log(caddyRepoPath)
	}
}
