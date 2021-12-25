package main

import (
	"log"
	"net/http"

	"tripit/routes"
	"tripit/utils"
)

func main() {
	utils.ConnectDatabase()

	r := routes.Router()

	log.Fatal(http.ListenAndServe(":8000", r))
}
