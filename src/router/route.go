package router

import (
	"github.com/gorilla/mux"
	"github.com/ojaoferreira/devops-utils/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
