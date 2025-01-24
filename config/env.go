package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	SecretKey  string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load(".env")
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "administratorPruebas"),
		DBPassword: getEnv("DB_PASSWORD", "Z&gCeSYp4!Le8aev"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "3.137.48.211"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "mrtire"),
		SecretKey:  getEnv("SECRET_KEY", "marque"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
