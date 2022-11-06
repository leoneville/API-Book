package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StrConn string = ""
	APIPort int    = 0
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Não foi possível carregar o .Env", err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 9000
	}

	StrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}
