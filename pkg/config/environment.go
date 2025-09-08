package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig DBConfig
	DB_MYSQL MySQLConfig
	Port     string
}

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
}

type DBConfig struct {
	Driver string
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Database string
	SSLMode  string
	CertPath string
	Port     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	cfg := &Config{
		Port: os.Getenv("PORT"),
		DbConfig: DBConfig{
			Driver: os.Getenv("DB_DRIVER"),
		},
		DB_MYSQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("DB_DATABASE"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			CertPath: os.Getenv("DB_CERT_PATH"),
		},
	}

	println("Variables:")
	println("CERT_PATH:", cfg.DB_MYSQL.CertPath)
	println("HOST:", cfg.DB_MYSQL.Host)
	println("USERNAME:", cfg.DB_MYSQL.Username)
	println("PASSWORD:", cfg.DB_MYSQL.Password)
	println("DBNAME:", cfg.DB_MYSQL.Database)
	println("PORT:", cfg.DB_MYSQL.Port)
	println("SSL_MODE:", cfg.DB_MYSQL.SSLMode)

	return cfg, nil
}

var AppConfig *Config

func init() {
	AppConfig, _ = LoadConfig()
}
