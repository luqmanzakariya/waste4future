package config

import (
	"fmt"
	"log"
	"os"
	// "user-service/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var dsn string
	var dialector gorm.Dialector

	dsn = fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	dialector = postgres.Open(dsn)

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Postgres Database connected successfully")

	// err = gormDB.AutoMigrate(&domain.User{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }

	return gormDB
}
