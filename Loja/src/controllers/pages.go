package controllers

import (
	"html/template"
	"net/http"
	"projeto/src/repositories"
)

var temp = template.Must(template.ParseGlob("src/templates/*.html"))

// index direcionar para a pagina de listar de produtos cadastrados na loja.
func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := repositories.BuscarTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

// New direcionar para a pagina para inserir um novo produto na loja.
func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}
