package config

import "testing"

func TestGetDbConfig(t *testing.T) {
	config, err := getDbConfig(".env")

	if err != nil {
		t.Errorf("Unable to Env File: %v", err)
	}

	if config.DB_HOST != "localhost" {
		t.Errorf("Database Host is not Correct, Expected: %s, Actual: %s", "localhost", config.DB_HOST)
	}
}
