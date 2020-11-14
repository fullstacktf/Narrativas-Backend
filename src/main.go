package main

import (
	config "github.com/fullstacktf/Narrativas-Backend/config"
	router "github.com/fullstacktf/Narrativas-Backend/router"
)

func main() {
	dbdata := config.DBConfig{
		Host:     "172.16.238.3",
		Port:     3306,
		User:     "rollify",
		DBName:   "rollify",
		Password: "password",
	}

	config.DbConnect(dbdata)
	r := router.InitRouter()
	r.Run(":9090")
}
