package rotas

import (
	"net/http"

	"github.com/ojaoferreira/devops-utils/src/controllers"
)

var rotasRandomBase64 = []Rota{
	{
		URI:    "/random-base64",
		Metodo: http.MethodPost,
		Funcao: controllers.RandomBase64,
	},
}
