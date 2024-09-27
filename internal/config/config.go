package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port int    `json:"port" mapstructure:"port"`
	Host string `json:"host" mapstructure:"host"`
}

type JwtConfig struct {
	Secret            string `json:"secret" mapstructure:"secret"`
	Expiration        int    `json:"expiration" mapstructure:"expiration"`
	RefreshExpiration int    `json:"refresh_expiration" mapstructure:"refresh_expiration"`
}

type DbConfig struct {
	Model  string        `json:"model" mapstructure:"model"`
	Sqlite *SqliteConfig `json:"sqlite" mapstructure:"sqlite"`
	Mysql  *MysqlConfig  `json:"mysql" mapstructure:"mysql"`
}

type SqliteConfig struct {
	Path string `json:"path" mapstructure:"path"`
}

type MysqlConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

// get config by viper
func GetAppConfig() *AppConfig {
	app := &AppConfig{
		Port: 8081,
		Host: "localhost",
	}
	viper.Unmarshal(&app)

	return app
}

func GetJwtConfig() *JwtConfig {
	jwt := &JwtConfig{
		Secret:            "secret",
		Expiration:        3600,
		RefreshExpiration: 86400,
	}
	viper.UnmarshalKey("jwt", &jwt)

	return jwt
}

func GetDbConfig() *DbConfig {
	// get project path
	projectPath, _ := os.Getwd()

	psPath := filepath.Base(projectPath)

	homePath := os.Getenv("HOME")
	db := &DbConfig{
		Model: "sqlite",
		Sqlite: &SqliteConfig{
			Path: filepath.Join(homePath, ".cache", "coderx", psPath, "db.sqlite"),
		},
	}
	viper.UnmarshalKey("db", &db)

	return db
}
