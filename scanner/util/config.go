package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zcubbs/zlogger/pkg/logger"
	"strings"
	"sync"
)

const (
	ViperConfigName      = "config"
	ViperConfigType      = "yaml"
	ViperConfigEnvPrefix = "SBOM"
	DefaultDbName        = "scanner"
)

var (
	viperConfigPaths = [...]string{"./config"}
	config           Config
	onceConfig       sync.Once
	defaults         = map[string]interface{}{
		"debug":                       false,
		"logger_type":                 logger.ZapLogger,
		"http_server.port":            8000,
		"http_server.swagger.enabled": false,
		"http_server.swagger.host":    "localhost:8000",
		"grpc.server.port":            50051,
		"grpc.server.tls_enabled":     false,
		"grpc.server.tls_cert":        "./scripts/server.crt",
		"grpc.server.tls_key":         "./scripts/server.key",
		"database.postgres.enabled":   true,
		"database.postgres.db_name":   DefaultDbName,
		"database.postgres.host":      "127.0.0.1",
		"database.postgres.port":      5432,
		"database.postgres.username":  "postgres",
		"database.postgres.password":  "postgres",
		"database.postgres.database":  "postgres",
		"database.postgres.ssl_mode":  false,
		"database.postgres.verbose":   false,
		"nats.enabled":                false,
		"nats.host":                   "localhost",
		"nats.port":                   4222,
	}
)

var allowedEnvVarKeys = []string{
	"debug",
	"logger_type",
	"grpc.server.port",
	"grpc.server.tls_enabled",
	"grpc.server.tls_cert",
	"grpc.server.tls_key",
	"http_server.port",
	"http_server.swagger.enabled",
	"http_server.swagger.host",
	"database.postgres.enabled",
	"database.postgres.host",
	"database.postgres.port",
	"database.postgres.username",
	"database.postgres.password",
	"database.postgres.database",
	"database.postgres.ssl_mode",
	"database.postgres.verbose",
	"nats.enabled",
	"nats.host",
	"nats.port",
}

func Bootstrap() Config {
	onceConfig.Do(loadConfig)
	return config
}

func loadConfig() {

	err := godotenv.Load(".env")

	if err != nil {
		if viper.GetString("debug") == "true" {
			fmt.Println("loading .env file")
		}
	}

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	for _, p := range viperConfigPaths {
		viper.AddConfigPath(p)
	}

	viper.SetConfigType(ViperConfigType)
	viper.SetConfigName(ViperConfigName)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("warn: %s\n", err)
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(ViperConfigEnvPrefix)

	for _, key := range allowedEnvVarKeys {
		err := viper.BindEnv(key)
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("warn: could not decode config into struct: %v", err)
	}
}
