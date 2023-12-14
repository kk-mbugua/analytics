package config

import (
	"log"
	"os"
	"path/filepath"
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
		config.Port, _ = strconv.Atoi(port)
		config.JWTSecretKey = jwtSecretKey
		config.AuthSvcUrl = authSvcUrl
		config.PostgresHost = postgresHost
		config.PostgresPort = postgresPort
		config.PostgresDB = postgresDB
		config.PostgresUser = postgresUser
		config.PostgresPass = postgresPass

		return config, nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	configPath := filepath.Join(cwd, "envs/")

	viper.AddConfigPath(configPath)
	viper.SetConfigName("dev.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

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
