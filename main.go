package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ojaoferreira/devops-utils/src/config"
	"github.com/ojaoferreira/devops-utils/src/router"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	log.Printf("[INFO] API executando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
