package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port      string
	AdminsUrl string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.Server.AdminsUrl = os.Getenv("ADMINS_URL")
	c.Server.Port = os.Getenv("SERVER_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")

	return nil
}
