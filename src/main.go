package main

import (
	server "github.com/fullstacktf/Narrativas-Backend/api/server"

	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func main() {
	config := &DBConfig{
		Host:     "172.26.0.3",
		Port:     3306,
		User:     "rollify",
		DBName:   "rollify",
		Password: "password",
	}

	_, err := gorm.Open(mysql.Open(DbURL(config)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error BBDD")
	} else {
		server.Serve()
	}
}
