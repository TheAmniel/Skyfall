package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"skyfall/services/config"
	"skyfall/types"
	"skyfall/utils"
)

type Database = gorm.DB

func New(cfg *config.DatabaseConfig) *Database {
	dir, _, err := utils.Executable()
	if err != nil {
		panic(err)
	}
	if cfg.Type != "local" && cfg.Type != "memory" {
		panic("[Database]: Invalid database type")
	}

	connection := buildConnection(cfg.Type, cfg.Name, dir)
	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{
		SkipDefaultTransaction: cfg.SkipDefaultTransaction,
		PrepareStmt:            cfg.PrepareStmt,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.File{}, &types.Short{}, &types.Ban{}, &types.Traffic{}, &types.Visitor{})
	return db
}

func buildConnection(t, n, d string) string {
	if t == "memory" {
		return ":memory:?mode=rw&cache=shared"
	}
	return fmt.Sprintf("%s%s.db?mode=rw", d, n)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
