package utils

import "testing"

func TestInitCaddyRepo(t *testing.T) {
	err := InitCaddyRepo("v1.0.0")
	if err != nil {
		t.Fatal(err)
	}
}
