package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/config"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/routes"
)

func main() {
	config.CarregarVariaveis()

	// Inicializa o backend apenas com HTTP REST
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.ConfigurarRotas(router)
	router.Run(":" + config.Porta)
}
