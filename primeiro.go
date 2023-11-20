package main

import (
	"GoTest/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Olá Mundo!")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
