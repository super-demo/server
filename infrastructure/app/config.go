package app

import (
	"log"

	"github.com/spf13/viper"
)

// Global Configuration variable
var Config *EnvConfigs

func InitEnvConfigs() {
	Config = LoadConfig()
}

type JwtConfigs struct {
	JwtSecretKey   string `mapstructure:"JWT_SECRET_KEY"`
	JwtTokenExpire int    `mapstructure:"JWT_TOKEN_EXPIRES_IN"`
}

type PostgresConfigs struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Database string `mapstructure:"POSTGRES_DATABASE"`
	SSLMode  string `mapstructure:"POSTGRES_SSLMODE"`
	Timezone string `mapstructure:"POSTGRES_TIMEZONE"`
}

type EnvConfigs struct {
	Environment   string `mapstructure:"APP_ENV"`
	Host          string `mapstructure:"APP_HOST"`
	Port          string `mapstructure:"APP_PORT"`
	AppSecret     string `mapstructure:"APP_SECRET"`
	ALLOW_ORIGINS string `mapstructure:"ALLOW_ORIGINS"`
	Postgres      PostgresConfigs
	Jwt           JwtConfigs
}

func SetDefaultConfig() {
	viper.SetDefault("APP_ENV", "local")
	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("POSTGRES_HOST", "postgres")
	viper.SetDefault("POSTGRES_PORT", 5432)
	viper.SetDefault("POSTGRES_USER", "user")
	viper.SetDefault("POSTGRES_PASSWORD", "password")
	viper.SetDefault("POSTGRES_DATABASE", "server_db_local")
	viper.SetDefault("POSTGRES_SSLMODE", "disable")
	viper.SetDefault("POSTGRES_TIMEZONE", "Asia/Bangkok")
	viper.SetDefault("APP_SECRET", "app-secret")
	viper.SetDefault("JWT_SECRET_KEY", "jwtsecretkey")
	viper.SetDefault("JWT_TOKEN_EXPIRES_IN", 24)
	viper.SetDefault("ALLOW_ORIGINS", "http://localhost:3000")
}

func LoadConfig() (config *EnvConfigs) {
	SetDefaultConfig()
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config.Postgres); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config.Jwt); err != nil {
		log.Fatal(err)
	}

	return
}
