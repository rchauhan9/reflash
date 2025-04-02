package config

type Config struct {
	Server       ServerConfig
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
