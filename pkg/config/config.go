package config

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST" validate:"required,gte=0,lte=80"`
	PostgresPort     uint32 `mapstructure:"POSTGRES_PORT" validate:"required,gte=0,lte=65535"`
	PostgresDBName   string `mapstructure:"POSTGRES_DB"  validate:"required,gte=0,lte=20"`
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"  validate:"required,gte=0,lte=80"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"  validate:"required,gte=0,lte=80"`

	AWSRegion      string `mapstructure:"AWS_REGION"  validate:"required,gte=0,lte=20"`
	AWSBucket      string `mapstructure:"AWS_BUCKET"  validate:"required,gte=0,lte=20"`
	AWSAccessKeyID string `mapstructure:"AWS_ACCESS_KEY_ID"  validate:"required,gte=0,lte=20"`
	AWSSecretKey   string `mapstructure:"AWS_SECRET_KEY"  validate:"required,gte=0,lte=80"`
}

func GetConfig(envFilePath string) (config Config, err error) {
	log.WithFields(log.Fields{"filepath": envFilePath}).Info("Loading env file")
	viper.SetConfigFile(envFilePath)
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.AutomaticEnv()
	viper.SetDefault("POSTGRES_PORT", 5432)

	println(config.PostgresHost)
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	err = validate(&config)
	return
}

func validate(config *Config) error {
	validate := validator.New()
	err := validate.Struct(config)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)

	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
