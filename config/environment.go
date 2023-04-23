package config

import (
	"log"

	"github.com/spf13/viper"
)

type Environments struct {
	APIPort  string
	AppName  string
	Env      string
	Database Database
	Token    Token
}

const (
	defaultEnv     = "dev"
	defaultAPIPort = "8080"
	defaultAppName = "gobank"
)

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("ENV", defaultEnv)
	viper.SetDefault("API_PORT", defaultAPIPort)
	viper.SetDefault("APP_NAME", defaultAppName)
	viper.SetDefault("DB_HOST", defaultDBHost)
	viper.SetDefault("DB_USER", defaultDBUser)
	viper.SetDefault("DB_PASSWORD", defaultDBPassword)
	viper.SetDefault("DB_NAME", defaultDBName)
	viper.SetDefault("DB_PORT", defaultDBPort)
	viper.SetDefault("DB_SSL_MODE", defaultDBSSLMode)
	viper.SetDefault("DB_DRIVER", defaultDBDriver)
	viper.SetDefault("TOKEN_SYMMETRIC_KEY", defaultTokenSymmetricKey)
	viper.SetDefault("ACCESS_TOKEN_DURATION", defaultAccessTokenDuration)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("unable to find or read config file %w", err)
	}

	env := &Environments{
		Env:     viper.GetString("ENV"),
		APIPort: viper.GetString("API_PORT"),
		AppName: viper.GetString("APP_NAME"),
		Database: Database{
			Host:     viper.GetString("DB_HOST"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			Port:     viper.GetString("DB_PORT"),
			SSLMode:  viper.GetString("DB_SSL_MODE"),
			Driver:   viper.GetString("DB_DRIVER"),
		},
		Token: Token{
			TokenSynmmetricKey:  viper.GetString("TOKEN_SYMMETRIC_KEY"),
			AccessTokenDuration: viper.GetDuration("ACCESS_TOKEN_DURATION"),
		},
	}

	return env
}
