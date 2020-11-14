package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig : Struct that configs database
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// dbURL : generates MySQL URL with the given config
func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// DbConnect : Connects to database
func DbConnect(config DBConfig) {

	db, err := gorm.Open(mysql.Open(dbURL(&config)), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connection to BBDD")
	}

	return db
}
