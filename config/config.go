package config

import (
	"github.com/joho/godotenv"
	"log"
)

var (
	StringConexao = ""
	Porta         = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
}
