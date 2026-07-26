package config

import (
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	ServerPort         string   `mapstructure:"PORT"`
	MongoURI           string   `mapstructure:"MONGO_URI"`
	DBName             string   `mapstructure:"DB_NAME"`
	JWTSecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int      `mapstructure:"JWT_EXPIRATION_HOURS"`
	EnableCache        bool     `mapstructure:"ENABLE_CACHE"`
	RedisAddr          string   `mapstructure:"REDIS_HOST"`
	RedisPassword      string   `mapstructure:"REDIS_PASSWORD"`
	LogLevel           string   `mapstructure:"LOG_LEVEL"`
	LogFormat          string   `mapstructure:"LOG_FORMAT"`
	CookieDomains      []string `mapstructure:"COOKIE_DOMAINS"`
	SecureCookie       bool     `mapstructure:"SECURE_COOKIE"`
	AllowedOrigins     []string `mapstructure:"ALLOWED_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Bind environment variables
	viper.BindEnv("PORT")
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("JWT_SECRET_KEY")
	viper.BindEnv("JWT_EXPIRATION_HOURS")
	viper.BindEnv("ENABLE_CACHE")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("LOG_FORMAT")
	viper.BindEnv("COOKIE_DOMAINS")
	viper.BindEnv("SECURE_COOKIE")
	viper.BindEnv("ALLOWED_ORIGINS")

	// Default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENABLE_CACHE", false)
	viper.SetDefault("JWT_EXPIRATION_HOURS", 72)
	viper.SetDefault("COOKIE_DOMAINS", []string{"localhost"})
	viper.SetDefault("SECURE_COOKIE", false)
	viper.SetDefault("ALLOWED_ORIGINS", []string{"http://localhost:5173"})

	// Read .env if it exists
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	// Read environment variables first
	config.ServerPort = os.Getenv("PORT")
	if config.ServerPort == "" {
		config.ServerPort = viper.GetString("PORT")
	}

	config.MongoURI = os.Getenv("MONGO_URI")
	if config.MongoURI == "" {
		config.MongoURI = viper.GetString("MONGO_URI")
	}

	config.DBName = os.Getenv("DB_NAME")
	if config.DBName == "" {
		config.DBName = viper.GetString("DB_NAME")
	}

	config.RedisAddr = os.Getenv("REDIS_HOST")
	if config.RedisAddr == "" {
		config.RedisAddr = viper.GetString("REDIS_HOST")
	}

	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	if config.RedisPassword == "" {
		config.RedisPassword = viper.GetString("REDIS_PASSWORD")
	}

	config.JWTSecretKey = viper.GetString("JWT_SECRET_KEY")
	config.JWTExpirationHours = viper.GetInt("JWT_EXPIRATION_HOURS")
	config.EnableCache = viper.GetBool("ENABLE_CACHE")
	config.LogLevel = viper.GetString("LOG_LEVEL")
	config.LogFormat = viper.GetString("LOG_FORMAT")
	config.CookieDomains = viper.GetStringSlice("COOKIE_DOMAINS")
	config.SecureCookie = viper.GetBool("SECURE_COOKIE")
	config.AllowedOrigins = viper.GetStringSlice("ALLOWED_ORIGINS")

	return config, nil
}