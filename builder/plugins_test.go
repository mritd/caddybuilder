package builder

import "testing"

func TestFind(t *testing.T) {
	ps := Find("geoip", "authz", "ipfilter")
	if len(ps) != 3 {
		t.Fatal("builder missing")
	} else {
		t.Log(ps)
	}
}

func TestGenerateCode(t *testing.T) {
	err := GenerateCode("geoip", "authz", "ipfilter")
	if err != nil {
		t.Fatal(err)
	}
}
