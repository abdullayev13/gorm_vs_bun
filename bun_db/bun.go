package bun_db

import (
	"database/sql"
	"github.com/abdullayev13/gorm_vs_bun/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func New(db *config.DB) *bun.DB {
	dsn := db.Dsn()
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	if db.DebugMode {
		bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return bunDB
}
