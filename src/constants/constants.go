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
	JWTSecret        string
	AccessSecret     string
)

func init() {
	godotenv.Load()
	godotenv.Load("../.env")
	DatabaseHost = os.Getenv("IPV4_DATABASE_ADDRESS")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	ServerPort = os.Getenv("SERVER_PORT")
	JWTSecret = os.Getenv("JWT_SECRET")
	AccessSecret = os.Getenv("ACCESS_SECRET")
}
