package app

import (
	"backend-inventory-app/migrations"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// SET ENV VARIABLE
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := "" + dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	// MIGRATE DB
	db.AutoMigrate(&migrations.Role{})
	db.AutoMigrate(&migrations.User{})
	db.AutoMigrate(&migrations.Category{})
	db.AutoMigrate(&migrations.Product{})
	db.AutoMigrate(&migrations.Transaction{})
	db.AutoMigrate(&migrations.TransactionDetail{})
	db.AutoMigrate(&migrations.UserRole{})
	db.AutoMigrate(&migrations.HakAkses{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
