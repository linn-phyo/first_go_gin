package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/linn-phyo/go_gin_clean_architecture/src/config"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// db.AutoMigrate(&domain.Users{})

	return db, dbErr
}
