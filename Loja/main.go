package main

import (
	"fmt"
	"net/http"
	"projeto/src/config"
	"projeto/src/routes"

	_ "github.com/lib/pq"
)

func main() {

	config.Carregando()
	routes.CarregarRotas()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	http.ListenAndServe(":8000", nil)
}
