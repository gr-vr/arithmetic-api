package main

import (
	"fmt"
	"github.com/satanCoffee/arithmetic-api/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.Router()
	log.Printf("Starting server on the port %s...\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
