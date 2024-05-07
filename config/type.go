package config

import (
	"github.com/golang-jwt/jwt"
)

type (
	ApiConfig struct {
		Url string
	}

	DbConfig struct {
		DataSourceName string
	}

	RedisConfig struct {
		Address  string
		Password string
		Database int
	}

	TokenConfig struct {
		ApplicationName string
		// JwtSigningMethod *jwt.SigningMethodHMAC
		JwtSigningMethod *jwt.SigningMethodHMAC

		JwtSignatureKey []byte
	}

	MailConfig struct {
		CONFIG_SMTP_HOST     string
		CONFIG_SMTP_PORT     int
		CONFIG_SENDER_NAME   string
		CONFIG_AUTH_EMAIL    string
		CONFIG_AUTH_PASSWORD string
	}

	NotifConfig struct {
		NotifUrl string
	}

	Config struct {
		ApiConfig
		DbConfig
		RedisConfig
		TokenConfig
		MailConfig
		NotifConfig
	}
)
