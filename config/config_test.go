package config

import "testing"

func TestNewAppConfig(t *testing.T) {
	cfg := NewAppConfig(".")
	if cfg.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", cfg.Host)
	}
	if cfg.Port != 3128 {
		t.Errorf("expected port to be 8087, got %d", cfg.Port)
	}
}
