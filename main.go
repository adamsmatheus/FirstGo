package main

import (
	"GoTest/config"
	"GoTest/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	println(config.Porta)
	println(config.StringConexao)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
