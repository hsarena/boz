package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`

	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOSTNAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_DBNAME"`
	MigrationPath string `mapstructure:"MIGRATION_PATH"`
	DBUrl         string
}

func LoadConfig(name string, path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config: %v", err)
		return
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("config: %v", err)
		return
	}
	return
}
