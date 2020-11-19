package main

import (
	"github.com/fullstacktf/Narrativas-Backend/common"
	"github.com/fullstacktf/Narrativas-Backend/constants"
	"github.com/fullstacktf/Narrativas-Backend/router"
)

func main() {

	dbdata := common.BuildDBConfig()
	common.DbInit(dbdata)

	r := router.InitRouter()
	r.Run(":" + constants.ServerPort)
}
