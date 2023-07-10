package gorm_db

import (
	"github.com/abdullayev13/gorm_vs_bun/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(db *config.DB) *gorm.DB {
	dsn := db.Dsn()
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return gormDb
}
