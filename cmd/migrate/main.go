package main

import (
	"go-casbin/internal/config"
	"go-casbin/internal/model/gormModel"
	"go-casbin/internal/repository"

	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadAppConfig()
	db, _ := repository.NewPostgresConnection(*cfg)
	GormMigrate(db)
}

func GormMigrate(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(
		&gormModel.User{},
		&gormModel.Task{},
	)
}
