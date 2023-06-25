package config

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/zcubbs/zlogger/pkg/logger"
)

var Conn *sql.DB

func GetDbConnection(config DatabaseConfig) (*sql.DB, error) {
	if config.Postgres.Enabled {
		conn, err := connectToPostgres(config.Postgres)
		if err != nil {
			return nil, err
		}
		Conn = conn
		return conn, nil
	}

	return nil, errors.New("no database profile enabled")
}

func connectToPostgres(config PostgresConfig) (*sql.DB, error) {
	logger.L().Infof("Connecting to Postgres [host=%s, port=%d, dbname=%s]",
		config.Host,
		config.Port,
		config.DbName,
	)

	var sslMode string
	if config.SslMode {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}
	// Open a PostgresSQL database.
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DbName,
		sslMode,
	)
	postgresDb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	logger.L().Info("connected to Postgres database")

	return postgresDb, nil
}
