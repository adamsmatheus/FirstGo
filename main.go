package main

import (
	"GoTest/src/router"
	"log"
	"net/http"
)

func main() {

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}
