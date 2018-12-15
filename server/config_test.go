package server

import "testing"

func TestNewConfigFromFile(t *testing.T) {
	confPath := "../test/.config"
	conf, err := NewConfigFromFile(confPath)
	if err != nil {
		t.Fatalf("did not expect an err: %v", err)
	}

	if conf.Bind != "0.0.0.0:80" {
		t.Errorf("New Config Bind got %s expected %s", conf.Bind, "0.0.0.0:80")
	}

	if conf.CacheMaxAge != 0 {
		t.Errorf("New Config Cache Max age got %d expected %d", conf.CacheMaxAge, 0)
	}
}
