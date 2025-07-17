package config

import (
	"os"
	"strconv"
)

type Config struct {
	App   AppConfig
	MySQL MySQLConfig
	Redis RedisConfig
	JWT   JWTConfig
}

type AppConfig struct {
	Port string
}

type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

var GlobalConfig Config

func init() {

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	accessTokenExpiry, _ := strconv.Atoi(getEnv("JWT_ACCESS_TOKEN_EXPIRY", "36000"))
	refreshTokenExpiry, _ := strconv.Atoi(getEnv("JWT_REFRESH_TOKEN_EXPIRY", "604800"))

	GlobalConfig = Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),
		},
		MySQL: MySQLConfig{
			Host:     getEnv("MYSQL_HOST", "localhost"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			Username: getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", "password"),
			Database: getEnv("MYSQL_DATABASE", "heritage_tour"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
		JWT: JWTConfig{
			Secret:             getEnv("JWT_SECRET", "heritage-tour-secret-key"),
			AccessTokenExpiry:  accessTokenExpiry,
			RefreshTokenExpiry: refreshTokenExpiry,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
