package config

import "github.com/zcubbs/zlogger/pkg/logger"

type Config struct {
	HttpServer HttpServerConfig `mapstructure:"http_server" json:"http_server"`
	Grpc       GrpcConfig       `mapstructure:"grpc" json:"grpc"`
	Database   DatabaseConfig   `mapstructure:"database" json:"database"`
	Debug      bool             `mapstructure:"debug" json:"debug"`
	LoggerType logger.Type      `mapstructure:"logger_type" json:"logger_type"`
}

type HttpServerConfig struct {
	Port    int           `mapstructure:"port" json:"port"`
	Swagger SwaggerConfig `mapstructure:"swagger" json:"swagger"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled" json:"enabled"`
	Host    string `mapstructure:"host" json:"host"`
}

type GrpcConfig struct {
	Server  GrpcServer   `mapstructure:"server" json:"server"`
	Clients []GrpcClient `mapstructure:"clients" json:"clients"`
}

type GrpcServer struct {
	Port       int    `mapstructure:"port" json:"port"`
	TlsEnabled bool   `mapstructure:"tls_enabled" json:"tls_enabled"`
	TlsCertPem string `mapstructure:"tls_cert" json:"tls_cert_pem"`
	TlsCertKey string `mapstructure:"tls_key" json:"tls_cert_key"`
}

type GrpcClient struct {
	Name       string `mapstructure:"name" json:"name"`
	Port       int    `mapstructure:"port" json:"port"`
	TlsEnabled bool   `mapstructure:"tls_enabled" json:"tls_enabled"`
	TlsCertPem string `mapstructure:"tls_cert" json:"tls_cert_pem"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `mapstructure:"postgres" json:"postgres"`
}

type PostgresConfig struct {
	Enabled  bool   `mapstructure:"enabled" json:"enabled"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DbName   string `mapstructure:"db_name" json:"db_name"`
	SslMode  bool   `mapstructure:"ssl_mode" json:"ssl_mode"`
	Verbose  bool   `mapstructure:"verbose" json:"verbose"`
	CertPem  string `mapstructure:"cert_pem"`
	CertKey  string `mapstructure:"cert_key"`
}

type Auth0 struct {
	Domain   string   `mapstructure:"domain" json:"domain"`
	Audience []string `mapstructure:"audience" json:"audience"`
}
