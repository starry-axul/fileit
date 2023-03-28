package bootstrap

import (
	"fmt"
	"os"

	"github.com/digitalhouse-dev/dh-kit/logger"
	"github.com/ncostamagna/axul_notifications/internal/notification"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectLocal func
func ConnectLocal(l logger.Logger) (*gorm.DB, string, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, "", l.CatchError(err)
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		// Migrate the schema
		err := db.AutoMigrate(&notification.Notification{})
		_ = l.CatchError(err)
	}

	return db, dsn, nil
}

// InitLogger -
func InitLogger() logger.Logger {
	return logger.New(logger.LogOption{})
}
