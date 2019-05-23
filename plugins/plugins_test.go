package plugins

import "testing"

func TestFindPlugins(t *testing.T) {
	ps := FindPlugins("geoip", "authz", "ipfilter")
	if len(ps) != 3 {
		t.Fatal("plugins missing")
	} else {
		t.Log(ps)
	}
}

func TestGeneratePluginsCode(t *testing.T) {
	err := GeneratePluginsCode("geoip", "authz", "ipfilter")
	if err != nil {
		t.Fatal(err)
	}
}
