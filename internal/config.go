package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/pelletier/go-toml"
)

var (
	lazily_loaded_config *ConfigSchema
)

type ConfigSchema struct {
	Main struct {
		Db_prefix, Environ string
	}

	Database struct {
		Drv_name  string
		Conn_urls ConfigSchemaConnUrls
	}
}

type ConfigSchemaConnUrl struct {
	ConnMaxIdleTime time.Duration
	Partitions      uint
	MaxOpenConns    uint
	Url             string
}

type ConfigSchemaConnUrls map[string]ConfigSchemaConnUrl

func (c *ConfigSchema) IsDev() (env bool) {
	cenv := c.Main.Environ

	if cenv == "local" || cenv == "test" {
		env = true
	}

	return
}

func (c *ConfigSchema) Setenv(env string) {
	c.Main.Environ = env
}

func RootDir() string {
	return os.Getenv("ROOT_DIR")
}

func GetSchemaRootDir() string {
	return RootDir() + "/schemas"
}

func GetConfigPath() string {
	return fmt.Sprintf("%s/config.toml", RootDir())
}

func LoadConfig() (schema *ConfigSchema, err error) {
	if lazily_loaded_config != nil {
		return lazily_loaded_config, nil
	}

	schema = &ConfigSchema{
		Main: struct{ Db_prefix, Environ string }{},
		Database: struct {
			Drv_name  string
			Conn_urls ConfigSchemaConnUrls
		}{
			Conn_urls: ConfigSchemaConnUrls{},
		},
	}

	conf_bs := make([]byte, 0)

	if conf_bs, err = os.ReadFile(GetConfigPath()); err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(conf_bs, schema); err != nil {
		return nil, err
	}

	lazily_loaded_config = schema
	return
}
