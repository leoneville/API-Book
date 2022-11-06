package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leoneville/api-book/src/config"
	"github.com/leoneville/api-book/src/router"
)

func init() {
	config.Load()
}

func main() {

	r := router.GenerateRoutes()

	fmt.Printf("Escutando na porta %d...", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r))
}
