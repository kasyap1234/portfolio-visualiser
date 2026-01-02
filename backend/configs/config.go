package configs

import "time"

type Config struct {
	DatabaseConfig *DatabaseConfig
	JWTConfig      *JWTConfig
	AuthConfig     *AuthConfig
	EmailConfig    *EmailConfig
	RedisConfig    *RedisConfig
	StorageConfig  *StorageConfig
}


func New() *Config {
	return &Config{
		DatabaseConfig: &DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "password",
			Database: "portfolio_visualiser",
		},
		JWTConfig: &JWTConfig{
			Issuer: "portfolio-visualiser",
			Expire: time.Hour * 24,
		},
		AuthConfig: &AuthConfig{
			JWTSecret: "secret",
		},
		EmailConfig: &EmailConfig{
			Host:     "smtp.gmail.com",
			Port:     587,
			Username: "username",
			Password: "password",
		},
		RedisConfig: &RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
		},
		StorageConfig: &StorageConfig{
			BucketName: "portfolio-visualiser",
			Region:     "us-east-1",
			AccessKey:  "access_key",
			SecretKey:  "secret_key",
			Endpoint:   "https://s3.amazonaws.com",
			SSL:        true,
		},
	}
}


type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type StorageConfig struct {
	BucketName string
	Region     string
	AccessKey  string
	SecretKey  string
	Endpoint   string
	SSL        bool
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type AuthConfig struct {
	JWTSecret string
}

type JWTConfig struct {
	Issuer string
	Expire time.Duration
}
