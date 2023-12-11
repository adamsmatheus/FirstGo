package controllers

import "net/http"

// Funcao que cria usuarios no sistema
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario especifico"))
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando os usuarios"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando os usuarios"))
}
