package utils

import (
	"errors"
	"testing"
)

func TestCheckAndExit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
		} else {
			t.Fatal("error check failed")
		}
	}()
	CheckAndExit(errors.New("this is a test error"))
}

func TestShortenString(t *testing.T) {
	if ShortenString("12345678", 7) != "1234567" {
		t.Fatal("test short string failed")
	}
}

func TestCheckGoCommand(t *testing.T) {
	t.Log(CheckCommand("go", "version"))
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
