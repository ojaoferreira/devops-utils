package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Porta int
)

func Carregar() {
	var erro error

	if _, erro := os.Stat(".env"); erro == nil {
		if erro = godotenv.Load(); erro != nil {
			log.Fatal(erro)
		}
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000
	}
}
