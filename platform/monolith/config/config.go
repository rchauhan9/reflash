package config

import (
	"flag"
	"github.com/rchauhan9/reflash/monolith/common/configutil"
	"github.com/rchauhan9/reflash/monolith/logging"
	"sync"
)

type Config struct {
	Server       ServerConfig
	Clerk        ClerkConfig        `mapstructure:"clerk"`
	StudyService StudyServiceConfig `mapstructure:"study"`
}

type ServerConfig struct {
	HTTPAddress string `mapstructure:"http-address"`
}

type StudyServiceConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	MigrationPath string `mapstructure:"migrationPath"`
	URL           string
}

type ClerkConfig struct {
	ApiKey string `mapstructure:"api-key"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		configPath := flag.String("config-dir", "./config", "Directory containing config.yml")
		flag.Parse()
		logger := logging.GetLogger()
		err := configutil.LoadConfig(*configPath, logger, &config)
		if err != nil {
			panic(err)
		}
	})
	return config
}
