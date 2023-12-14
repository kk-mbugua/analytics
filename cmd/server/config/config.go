package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Port         int    `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	AuthSvcUrl   string `mapstructure:"AUTH_SVC_URL"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
	PostgresDB   string `mapstructure:"POSTGRES_DB"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
}

func LoadConfig() (config Config, err error) {

	port, jwtSecretKey, authSvcUrl, postgresHost, postgresPort, postgresDB, postgresUser, postgresPass :=
		os.Getenv("PORT"), os.Getenv("JWT_SECRET_KEY"), os.Getenv("AUTH_SVC_URL"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS")

	if port != "" && jwtSecretKey != "" && authSvcUrl != "" {
		log.Println("Port is: ", port)
		config.Port, err = strconv.Atoi(port)
		if err != nil {
			log.Fatalf("Error converting PORT to int: %v", err)
		}

		config.JWTSecretKey = jwtSecretKey
		config.AuthSvcUrl = authSvcUrl
		config.PostgresHost = postgresHost
		config.PostgresPort = postgresPort
		config.PostgresDB = postgresDB
		config.PostgresUser = postgresUser
		config.PostgresPass = postgresPass

		return config, nil
	}

	viper.AutomaticEnv()
	config.Port = viper.GetInt("PORT")
	config.PostgresHost = viper.GetString("POSTGRES_HOST")
	config.PostgresPort = viper.GetString("POSTGRES_PORT")
	config.PostgresDB = viper.GetString("POSTGRES_DB")
	config.PostgresUser = viper.GetString("POSTGRES_USER")
	config.PostgresPass = viper.GetString("POSTGRES_PASSWORD")

	if config.Port == 0 || config.PostgresHost == "" || config.PostgresPort == "" || config.PostgresDB == "" || config.PostgresUser == "" || config.PostgresPass == "" {
		log.Fatalf("Some required configuration is missing: Missing values: %v", config)
	}

	return config, nil
}
