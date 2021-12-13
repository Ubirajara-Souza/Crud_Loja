package models

// Produto representa os itens do cadastro de produtos
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}
