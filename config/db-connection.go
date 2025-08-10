package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dev-khalid/gin-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func initDB() {
	var err error

	user := GetEnv("DB_USER")
	pass := GetEnv("DB_PASS")
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	name := GetEnv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                              // Slow SQL threshold
			LogLevel:                  logger.Error | logger.Warn | logger.Info, // Log level
			IgnoreRecordNotFoundError: true,                                     // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                                    // Don't include params in the SQL log
			Colorful:                  true,                                     // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database connection established ✅")
	}

	migrateDB()
}

func migrateDB() {
	err := DB.AutoMigrate(&models.EventPricing{}, &models.Event{})

	if err != nil {
		panic("failed to migrate database")
	}
}
