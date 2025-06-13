package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	NomeEmpresa       string
	Porta             string
	URLsEmpresas      map[string]string
	EmpresaNomeParaID map[string]string
)

func CarregarVariaveis() {
	// Tente carregar o .env no diretório atual, mas NÃO sobrescreva variáveis já definidas no ambiente
	err := godotenv.Load(".env")
	if err != nil {
		// Se não encontrar, tente carregar o .env no caminho do Docker
		err = godotenv.Load("/app/.env")
		if err != nil {
			log.Println("⚠️  Arquivo .env não encontrado, usando variáveis do sistema")
		}
	}

	NomeEmpresa = os.Getenv("NOME_EMPRESA")
	Porta = os.Getenv("PORTA")

	URLsEmpresas = map[string]string{
		"empresa_a": os.Getenv("EMPRESA_A_URL"),
		"empresa_b": os.Getenv("EMPRESA_B_URL"),
		"empresa_c": os.Getenv("EMPRESA_C_URL"),
	}

	EmpresaNomeParaID = map[string]string{
		"empresa_a": "1",
		"empresa_b": "2",
		"empresa_c": "3",
	}
}
