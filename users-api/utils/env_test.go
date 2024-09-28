package utils

import (
	"fmt"
	"testing"
)

func TestGetDatabaseURLSet(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgresql://test:test@localhost:5432/test")
	_ = GetDatabaseURL()
}

func TestGetDatabaseURLUnset(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetDatabaseURL() did not panic")
		}
	}()

	_ = GetDatabaseURL()
}

func TestGetServerPortDefault(t *testing.T) {
	server_port := GetServerPort()
	if server_port != DefaultServerPort {
		t.Errorf("GetServerPort() = %v, want = %v", server_port, DefaultServerPort)
	}
}

func TestGetServerPortCustom(t *testing.T) {
	expected := 6969
	t.Setenv("SERVER_PORT", fmt.Sprint(expected))
	server_port := GetServerPort()
	if server_port == DefaultServerPort {
		t.Errorf("GetServerPort() = %v, want = %v", DefaultServerPort, expected)
	}
}
