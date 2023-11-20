package app

import "github.com/urfave/cli"

// Teste
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação do comando"
	app.Usage = "Busca IP"

	return app
}
