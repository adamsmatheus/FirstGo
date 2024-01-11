package rotas

import "net/http"

var rotaLogin = Rota{
	Uri:    "/login",
	Metodo: http.MethodPost,
	Funcao: func(writer http.ResponseWriter, request *http.Request) {

	},
	RequerAutenticacao: false,
}
