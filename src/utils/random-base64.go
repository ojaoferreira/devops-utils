package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func RandomBase64(tamanho int) string {
	chave := make([]byte, tamanho)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	chaveBase64 := base64.StdEncoding.EncodeToString(chave)

	return chaveBase64
}
