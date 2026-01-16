package repo

import (
	"fff"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	goose.SetBaseFS(fff.EmbeddedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		return nil, err
	}

	return db, nil
}
