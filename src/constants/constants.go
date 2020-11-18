package constants

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseHost     string
	DatabaseUser     string
	DatabaseName     string
	DatabasePassword string
	DatabasePort     int
	ServerPort       string
)

func init() {
	godotenv.Load()
	godotenv.Load("../env/.env")
	DatabaseHost = os.Getenv("IPV4_DATABASE_ADDRESS")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	ServerPort = os.Getenv("SERVER_PORT")
}
