package config

import (
	"fmt"

	"github.com/pedromartinsb/gomerit/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root1234"
const DB_NAME = "merit"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var Db *gorm.DB

func InitializeMySQL() (*gorm.DB, error) {
	Db, err := connectDB()
	if err != nil {
		return nil, err
	}

	return Db, nil
}

func connectDB() (*gorm.DB, error) {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.Holding{})
	if err != nil {
		logger.Errorf("mysql automigrate error: %v", err)
		return nil, err
	}

	// Return the db
	return db, nil
}
