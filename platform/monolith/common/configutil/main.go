package configutil

import (
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/spf13/viper"
)

// LoadConfig loads and unmarshalls configuration into the given config address. configPath is an optional path to the
// yaml configuration file, if none is provided a file named config.yml at the current or root directory is expected.
func LoadConfig(configPath string, logger log.Logger, config interface{}) error {
	// Configure viper to read a YML configuration file from various paths
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}

	// Configure viper to override values from env variables by replacing the YML delimiters with env var delimiters
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		level.Error(logger).Log("msg", "Error reading config file", "err", err)
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		level.Error(logger).Log("msg", "Unable to decode into struct", "err", err)
		return err
	}

	return nil
}
