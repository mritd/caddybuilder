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

func TestSort(t *testing.T) {
	ps := Sort()
	t.Log(ps)
}

func TestMerge(t *testing.T) {
	err := Merge("../resources/other_plugins.json")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateCode(t *testing.T) {
	err := GenerateCode("geoip", "authz", "ipfilter")
	if err != nil {
		t.Fatal(err)
	}
}
