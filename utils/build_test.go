package utils

import (
	"testing"
)

func TestInitDep(t *testing.T) {
	InitDep("github.com/kodnaplakal/caddy-geoip", "github.com/casbin/caddy-authz", "github.com/pyed/ipfilter")
}

func TestBuild(t *testing.T) {
	err := Build("./caddy")
	if err != nil {
		t.Fatal(err)
	}
}
