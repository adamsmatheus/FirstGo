package rotas

import (
	"GoTest/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rotas da api
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configura todas as rotas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.Uri,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao))).
				Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}
	return r
}
