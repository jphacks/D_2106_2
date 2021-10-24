package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_NAME string
	DB_PORT string
}

type AwsConfig struct {
	ACCESS_KEY string
	SECRET_KEY string
	REGION     string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		// return nil, err
		fmt.Println(err)
	}

	config := &Config{
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PORT: os.Getenv("DB_PORT"),
	}

	return config, nil
}

func GetAwsConfig() (*AwsConfig, error) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}

	config := &AwsConfig{
		ACCESS_KEY: os.Getenv("AWS_ACCESS_KEY"),
		SECRET_KEY: os.Getenv("AWS_SECRET_KEY"),
		REGION:     os.Getenv("AWS_REGION"),
	}

	return config, nil
}
