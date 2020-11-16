package main

import (
	"strconv"

	"github.com/fullstacktf/Narrativas-Backend/common"
	"github.com/fullstacktf/Narrativas-Backend/constants"
	"github.com/fullstacktf/Narrativas-Backend/router"
)

func main() {
	dbdata := common.BuildDBConfig()
	common.DbInit(dbdata)

	r := router.InitRouter()
	r.Run(":" + strconv.Itoa(constants.ServerPort))
}
