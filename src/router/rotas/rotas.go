package rotas

import (
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
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
