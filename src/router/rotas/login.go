package rotas

import (
	"GoTest/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
