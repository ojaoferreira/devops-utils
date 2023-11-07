package rotas

import "net/http"

var rotasHealthcheck = Rota{
	URI:    "/health",
	Metodo: http.MethodGet,
	Funcao: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	},
	RequerAutenticacao: false,
}
