package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error

	user := GetEnv("DB_USER")
	pass := GetEnv("DB_PASS")
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	name := GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database connection established ✅")
	}

	// Migrate the schema
	migrateDB()
}

func migrateDB() {
	// Run the database migrations
	err := DB.AutoMigrate()

	if err != nil {
		panic("failed to migrate database")
	}
}
