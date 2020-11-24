package common

import (
	"fmt"
	"log"

	"github.com/fullstacktf/Narrativas-Backend/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     constants.DatabaseHost,
		Port:     constants.DatabasePort,
		User:     constants.DatabaseUser,
		Password: constants.DatabasePassword,
		DBName:   constants.DatabaseName,
	}
	return &dbConfig
}

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

type Tabler interface {
	TableName() string
}

func DbInit(config *DBConfig) {

	db, err := gorm.Open(mysql.Open(dbURL(config)), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connection to BBDD")
	}

	DB = db
}
