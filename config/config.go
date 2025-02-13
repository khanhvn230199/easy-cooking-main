package config

import (
	"easy-cooking/internal/models/dto"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config dto.AppConfig

func LoadConfig() {
	viper.AutomaticEnv()

	requiredEnvs := []string{
		"DATABASE_HOST",
		"DATABASE_PORT",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
		"SERVER_PORT",
	}

	for _, env := range requiredEnvs {
		if !viper.IsSet(env) {
			log.Fatalf("Required environment variable %s is not set", env)
		}
	}

	Config = dto.AppConfig{
		DatabaseDSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			strings.TrimSpace(strings.Trim(viper.GetString("DATABASE_HOST"), "\"")),
			strings.TrimSpace(strings.Trim(viper.GetString("DATABASE_PORT"), "\"")),
			strings.TrimSpace(strings.Trim(viper.GetString("DATABASE_USER"), "\"")),
			strings.TrimSpace(strings.Trim(viper.GetString("DATABASE_PASSWORD"), "\"")),
			strings.TrimSpace(strings.Trim(viper.GetString("DATABASE_NAME"), "\"")),
		),
		ServerPort: viper.GetString("SERVER_PORT"),
		//JWTConfig: dto.JWTConfig{
		//	SecretKey:        viper.GetString("JWT_SECRET_KEY"),
		//	AccessExpiration: viper.GetDuration("JWT_ACCESS_EXPIRATION"),
		//},
	}

	log.Printf("DatabaseDSN: %s\n", Config.DatabaseDSN)
	log.Printf("ServerPort: %s\n", Config.ServerPort)
}
