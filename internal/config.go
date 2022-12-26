package internal

import (
	"fmt"
	"go/build"
	"os"
	"time"

	"github.com/pelletier/go-toml"
)

var (
	lazily_loaded_config *ConfigSchema
)

type ConfigSchema struct {
	Main struct {
		Db_prefix, Environ, Project_root string
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
	return build.Default.GOPATH + "/src"
}

func GetSchemaRootDir() string {
	config, _ := LoadConfig()
	return RootDir() + "/" + config.Main.Project_root + "/schemas"
}

func GetConfigPath() string {
	return fmt.Sprintf("%s/src/blogs/config.toml", build.Default.GOPATH)
}

func LoadConfig() (schema *ConfigSchema, err error) {
	if lazily_loaded_config != nil {
		return lazily_loaded_config, nil
	}

	schema = &ConfigSchema{
		Main: struct{ Db_prefix, Environ, Project_root string }{},
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
