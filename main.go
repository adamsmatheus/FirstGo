package main

import (
	"GoTest/src/config"
	"GoTest/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
func main() {
	config.Carregar()

	println(config.Porta)
	println(config.StringConexao)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
