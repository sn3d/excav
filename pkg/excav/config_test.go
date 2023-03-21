package excav

import (
	"testing"
)

func TestReadYamlConfig(t *testing.T) {
	// when we read testdata/config.yaml
	config, err := ReadYamlConfig("testdata/config.yaml")
	if err != nil {
		t.FailNow()
	}

	// then config must contain correct workspace
	if config.WorkspaceDir != "/tmp/excav/workspace" {
		t.FailNow()
	}
}

func TestCurrentBulk(t *testing.T) {
	cfg, _ := TestConfiguration()

	err := cfg.SetCurrentBulk("/path/to/bulk")
	if err != nil {
		t.FailNow()
	}

	bulkDir := cfg.GetCurrentBulk()
	if bulkDir != "/path/to/bulk" {
		t.FailNow()
	}
}
