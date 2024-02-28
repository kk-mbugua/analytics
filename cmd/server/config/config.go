package config

/*
This file is used to define the configuration for the service.

*/

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	AuthSvcUrl         string `mapstructure:"AUTH_SVC_URL"`
	OrganisationSvcUrl string `mapstructure:"ORGANISATION_SVC_URL"`
	ContactSvcUrl      string `mapstructure:"CONTACT_SVC_URL"`
	JwtSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	MailgunAPIKey      string `mapstructure:"MAILGUN_API_KEY"`
	MailgunDomain      string `mapstructure:"MAILGUN_DOMAIN"`
	MailgunSender      string `mapstructure:"MAILGUN_SENDER"`
	PostgresHost       string `mapstructure:"POSTGRES_HOST"`
	PostgresPort       string `mapstructure:"POSTGRES_PORT"`
	PostgresDB         string `mapstructure:"POSTGRES_DB"`
	PostgresUser       string `mapstructure:"POSTGRES_USER"`
	PostgresPass       string `mapstructure:"POSTGRES_PASS"`
	Environment        string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig() (config Config, err error) {
	port, jwtSecretKey, authSvcUrl, postgresHost, postgresPort, postgresDB, postgresUser, postgresPass, environment :=
		os.Getenv("PORT"), os.Getenv("JWT_SECRET_KEY"), os.Getenv("AUTH_SVC_URL"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"), os.Getenv("ENVIRONMENT")

	if port != "" && jwtSecretKey != "" && authSvcUrl != "" {
		log.Println("Port is: ", port)
		config.Port = port
		config.JwtSecretKey = jwtSecretKey
		config.AuthSvcUrl = authSvcUrl
		config.PostgresHost = postgresHost
		config.PostgresPort = postgresPort
		config.PostgresDB = postgresDB
		config.PostgresUser = postgresUser
		config.PostgresPass = postgresPass
		config.Environment = environment

		return config, nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	configPath := filepath.Join(cwd, "envs")

	viper.AddConfigPath(configPath)
	viper.SetConfigName("dev.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config, nil
}
