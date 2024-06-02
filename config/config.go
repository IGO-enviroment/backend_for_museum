package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Log       `yaml:"logger"`
		PG        `yaml:"postgres"`
		ENV       `yaml:"env"`
		StorageS3 `yaml:"s3"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		HOST           string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port           string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		JwtSecretKey   string `env-required:"true" yaml:"jwt_secret_key" env:"JWT_SECRET_KEY"`
		JwtSeparateKey string `env-required:"true" yaml:"jwt_separate_key" env:"JWT_SEPARATE_KEY"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" env:"PG_URL"`
	}

	Redis struct {
		URL string `env-required:"true" env:"REDIS_URL"`
	}

	ENV struct {
		Level string `env-required:"true" yaml:"level" env:"ENV_LEVEL"`
	}

	StorageS3 struct {
		EndPoint        string `env-required:"true" yaml:"end_point" env:"S3_ENDPOINT"`
		AccessKey       string `env-required:"true" yaml:"access_key" env:"S3_ACCESS_KEY"`
		SecretAccessKey string `env-required:"true" yaml:"secret_access_key" env:"S3_SECRET_ACCESS_KEY"`
		BucketName      string `env-required:"true" yaml:"bucket_name" env:"S3_BUCKET_NAME"`
		Region          string `env-required:"true" yaml:"region" env:"S3_REGION"`
	}
)

// Загрузка настроек приложения.
func NewConfig() (*Config, error) {
	var err error

	cfg := &Config{}

	err = cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return cfg, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

// Получение ранее ранее установленых настроек
func GetConf() (*Config, error) {
	return NewConfig()
}

func (c *Config) Development() bool {
	return c.ENV.Level == "development"
}
