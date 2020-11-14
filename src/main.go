package main

import (
	server "github.com/fullstacktf/Narrativas-Backend/api/server"
	config "github.com/fullstacktf/Narrativas-Backend/config"
)

func main() {
	dbdata := config.DBConfig{
		Host:     "172.16.238.3",
		Port:     3306,
		User:     "rollify",
		DBName:   "rollify",
		Password: "password",
	}

	db := config.DbConnect(dbdata)
	server.Serve()
}
