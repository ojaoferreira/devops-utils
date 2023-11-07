package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ojaoferreira/devops-utils/src/utils"
)

func RandomBase64(w http.ResponseWriter, r *http.Request) {
	tamanhoStr := r.URL.Query().Get("tamanho")
	if tamanhoStr == "" {
		tamanhoStr = "64"
	}

	tamanho, erro := strconv.Atoi(tamanhoStr)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(utils.RandomBase64(tamanho)))
}
