package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndValidateWorkflowStyleConfig(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	configPath := filepath.Join(dir, "config.yaml")
	content := `profile:
  first_name: "Test"
  last_name: "User"
  email: "test@example.com"
email:
  provider: smtp
  from: "test@example.com"
  smtp:
    host: smtp.gmail.com
    port: 465
    username: "smtp-user@example.com"
    password: "app-password"
    use_tls: true
options:
  template: generic
`

	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}
