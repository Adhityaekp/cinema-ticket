package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	AppBaseURL string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	JWTSecret              string
	JWTExpireHours         int
	JWTAccessExpireMinutes int
	JWTRefreshExpireDays   int

	CookieSecure bool
	CookieDomain string

	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string
}

func Load() *Config {
	_ = godotenv.Load()

	jwtExpireHours, _ := strconv.Atoi(
		os.Getenv("JWT_EXPIRE_HOURS"),
	)

	jwtAccessExpireMinutes, _ := strconv.Atoi(
		os.Getenv("JWT_ACCESS_EXPIRE_MINUTES"),
	)

	jwtRefreshExpireDays, _ := strconv.Atoi(
		os.Getenv("JWT_REFRESH_EXPIRE_DAYS"),
	)

	cookieSecure := os.Getenv("COOKIE_SECURE") == "true"

	return &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppPort:    os.Getenv("APP_PORT"),
		AppBaseURL: os.Getenv("APP_BASE_URL"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       os.Getenv("REDIS_DB"),

		JWTSecret:              os.Getenv("JWT_SECRET"),
		JWTExpireHours:         jwtExpireHours,
		JWTAccessExpireMinutes: jwtAccessExpireMinutes,
		JWTRefreshExpireDays:   jwtRefreshExpireDays,

		CookieSecure: cookieSecure,
		CookieDomain: os.Getenv("COOKIE_DOMAIN"),

		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		SMTPFrom:     os.Getenv("SMTP_FROM"),
	}
}
