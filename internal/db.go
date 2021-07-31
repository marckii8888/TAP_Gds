package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// TODO: Change to read from config file
const (
	DB_USERNAME = "root"
	DB_PASSWORD = "root123"
	DB_NAME = "gds_assessment"
	DB_HOST = "127.0.0.1"
	DB_PORT = "3306"
)

func InitDB() *gorm.DB {
	db := connectDB()
	return db
}

func connectDB() *gorm.DB {
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database. Error = %+v", err)
		return nil
	}
	return db
}