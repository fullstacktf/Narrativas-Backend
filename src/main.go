package main

import (
	common "github.com/fullstacktf/Narrativas-Backend/common"
	router "github.com/fullstacktf/Narrativas-Backend/router"
)

func main() {
	dbdata := common.DBConfig{
		Host:     "172.16.238.3",
		Port:     3306,
		User:     "rollify",
		DBName:   "rollify",
		Password: "password",
	}

	common.DbInit(dbdata)

	r := router.InitRouter()
	r.Run(":10101")
}
