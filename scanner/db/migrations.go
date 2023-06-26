package sql

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/mattes/migrate/source/file"
	"github.com/zcubbs/zlogger/pkg/logger"
	"net/http"
	"zel/sbom-prototype/scanner/internal/config"
)

//go:embed migrations
var migrations embed.FS

func MigrateDB(dbCfg config.DatabaseConfig) {
	driver, err := postgres.WithInstance(config.DbConn, &postgres.Config{
		DatabaseName:          dbCfg.Postgres.DbName,
		SchemaName:            "public",
		MultiStatementEnabled: true,
	})
	if err != nil {
		logger.L().Fatal(err, "failed to create postgres driver")
	}

	source, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		logger.L().Fatal(err, "failed to create migration source")
	}

	m, err := migrate.NewWithInstance(
		"FS",
		source,
		dbCfg.Postgres.DbName,
		driver,
	)
	if err != nil {
		logger.L().Fatal(err, "failed to create migration")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.L().Fatal(err, "failed to apply migration")
	}

	logger.L().Info("database migration completed")
}
