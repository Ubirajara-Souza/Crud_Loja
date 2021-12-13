package db

import (
	"database/sql"
	"projeto/src/config"

	_ "github.com/lib/pq" // Driver
)

//ConectaComBancoDeDados se conectar com o banco de dados
func ConectaComBancoDeDados() *sql.DB {

	db, err := sql.Open("postgres", config.StringConexaoBanco)
	if err != nil {
		panic(err.Error())
	}
	return db
}
