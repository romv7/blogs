package internal_test

import (
	"os"
	"testing"

	"github.com/romv7/blogs/internal"
)

var (
	configLoc = internal.GetConfigPath()
)

const (
	origConfigPath = "/tmp/main_config.toml"
	mockConfigFile = `[main]

[database]
drv_name = "mongodb"
conn_urls = {
	local = {
		conn_max_idle_time = 3000,
		max_open_conns = 10,
		url = "rommms:07261999@unix(/var/run/mysqld/mysqld.sock)/rommms.nation"
	}
}
`
)

func Test_mustParseAndLoadTheConfigFile(t *testing.T) {
	attachMockConfig()

	config, err := internal.LoadConfig()
	if err != nil {
		t.Errorf("[fail] %s", err.Error())
	}

	failMsg := "[fail] did not properly load the expected config file."

	if config.Database.Drv_name != "mongodb" {
		t.Errorf(failMsg)
	}

	if _, exists := config.Database.Conn_urls["local"]; !exists {
		t.Errorf(failMsg)
	}

	detachMockConfig()
}

func attachMockConfig() {
	b, _ := os.ReadFile(configLoc)
	os.WriteFile(origConfigPath, b, 0644)
	os.Remove(configLoc)
	os.WriteFile(configLoc, []byte(mockConfigFile), 0755)
}

func detachMockConfig() {
	b, _ := os.ReadFile(origConfigPath)
	os.Remove(configLoc)
	os.Remove(origConfigPath)
	os.WriteFile(configLoc, b, 0644)
}
