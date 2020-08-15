package config

import (
	"fmt"
	"log"
	"strconv"
)

var (
	PORT      = 0
	DB_URL    = ""
	DB_DRIVER = ""
	SECRET    []byte
)

func Load() {
	var err error

	PORT, err = strconv.Atoi("9000")
	if err != nil {
		fmt.Println("bitch")
		log.Println(err)
		PORT = 8000
	}
	DB_DRIVER = "mysql"
	DB_URL = fmt.Sprintf("root:marianasantos8991@tcp(127.0.0.1:3306)/teste")

	SECRET = []byte("lknavbchghqhvhhojhc3583729fhwbjxiyjgfhagcu1")
}
