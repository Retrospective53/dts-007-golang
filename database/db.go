package database

import (
	"fmt"
	"log"

	"challenge-10/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "flanerie"
	dbPort   = "5432"
	dbname   = "productuser"
	db       *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("successfully connect to database")
	db.AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}