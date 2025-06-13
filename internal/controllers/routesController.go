package controllers

import (
	"net/http"

	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/services"

	"github.com/gin-gonic/gin"
)

// Controller de rotas: expõe endpoints HTTP para o cliente buscar rotas
// Chama os serviços de rotas, que podem consultar outros servidores para montar rotas distribuídas
// Exemplo de uso: GET /routes (cliente busca todas as rotas possíveis)
func GetAllRoutes(c *gin.Context) {
	routes := services.GetAllRoutes()
	c.JSON(http.StatusOK, routes)
}

func GetRouteByID(c *gin.Context) {
	id := c.Param("id")
	route, err := services.GetRouteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, route)
}
