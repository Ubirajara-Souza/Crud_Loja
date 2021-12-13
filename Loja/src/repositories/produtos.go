package repositories

import (
	"projeto/src/db"
	"projeto/src/models"
)

// BuscarTodosOsProdutos exibir a lista de todos os produtos
func BuscarTodosOsProdutos() []models.Produto {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}
	produtos := []models.Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// CriarNovoProduto criar um novo produto na loja.
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare(
		"INSERT INTO produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

// DeletaProduto excluir um produto da loja.
func DeletaProduto(id string) {

	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

// EditaProduto editar um produto da loja.
func EditaProdut(id string) models.Produto {

	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := models.Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}

	defer db.Close()
	return produtoParaAtualizar

}

// AtualizarProduto atualizar um produto da loja.
func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	AtualizarProduto, err := db.Prepare(
		"UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5 ")
	if err != nil {
		panic(err.Error())
	}

	AtualizarProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
