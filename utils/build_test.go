package utils

import "testing"

func TestBuild(t *testing.T) {
	err := Build("./caddy")
	if err != nil {
		t.Fatal(err)
	}
}
