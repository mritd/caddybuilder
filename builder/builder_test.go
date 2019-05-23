package builder

import "testing"

func TestInitDep(t *testing.T) {
	err := InitDep("geoip", "authz", "ipfilter")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuild(t *testing.T) {
	err := Build("./caddy")
	if err != nil {
		t.Fatal(err)
	}
}
