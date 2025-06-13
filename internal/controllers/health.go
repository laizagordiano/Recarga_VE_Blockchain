package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



 //HealthCheck é um endpoint simples que retorna o status do servidor.
 // Ele pode ser usado para verificar se o servidor está em execução e acessível.
 
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"empresa": "HenryVolt", // pode ser parametrizado
	})
}
