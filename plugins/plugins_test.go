package plugins

import "testing"

func TestGeneratePluginsCode(t *testing.T) {
	err := GeneratePluginsCode([]string{"geoip", "authz", "ipfilter"})
	if err != nil {
		t.Fatal(err)
	}
}
