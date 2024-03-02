package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBDriver            string `json:"DB_DRIVER" binding:"required"`
	DBSource            string `json:"DB_SOURCE" binding:"required"`
	CDN_URL             string `json:"CDN_URL"`
	CDN_BUCKET          string `json:"CDN_BUCKET"`
	CDN_MUG             string `json:"CDN_MUG"`
	PORT                string `json:"PORT"`
	HOSTPATH            string `json:"HOST_PATH"`
	APIURL              string `json:"API_URL"`
	ADMINURL            string `json:"ADMIN_URL"`
	BASEPATH            string `json:"BASE_PATH"`
	TokenType           string `json:"TOKEN_TYPE" binding:"required"`
	TokenSymmetricKey   string `json:"TOKEN_SYMMETRIC_KEY" binding:"required"`
	AccessTokenDuration string `json:"ACCESS_TOKEN_DURATION"`
	SmptHost            string `json:"SMTP_HOST"`
	SmtpPort            string `json:"SMTP_PORT"`
	Smtp_Encryption     string `json:"SMTP_ENCRYPTION"`
	SmtpUsername        string `json:"SMTP_USERNAME"`
	SmtpPassword        string `json:"SMTP_PASSWORD"`
	SmtpFromAddress     string `json:"SMTP_FROM_ADDRESS"`
	Debug               string `json:"DEBUG"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	kvs := os.Environ()
	mp := map[string]string{}
	for _, kv := range kvs {
		sp := strings.SplitN(kv, "=", 2)
		mp[sp[0]] = sp[1]
	}
	// mp, err := godotenv.Read()
	// if err != nil {
	// 	return
	// }
	dataBytes, err := json.Marshal(mp)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataBytes, &config)
	if err != nil {
		return
	}
	logLevel := logrus.InfoLevel
	if config.Debug == "true" {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
	logrus.Info("Successfully loaded configuration.")
	return
}
