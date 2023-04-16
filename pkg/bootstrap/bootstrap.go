package bootstrap

import (
	"fmt"
	"github.com/starry-axul/fileit/internal/domain"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectLocal func
func ConnectLocal() (*gorm.DB, string, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, "", err
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&domain.Client{}); err != nil {
			return nil, "", err
		}
	}

	return db, dsn, nil
}
