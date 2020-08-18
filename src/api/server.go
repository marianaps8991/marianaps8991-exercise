package api

import (
	"fmt"
	"log"
	"net/http"

	"api/auto"
	"api/router"
	"config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\t Listening http://localhost:%d \n", config.PORT)
	listen(9000)

}

func listen(port int) {
	r := router.New()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
