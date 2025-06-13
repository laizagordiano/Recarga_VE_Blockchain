package controllers

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/services"
)

var blockchainService *services.BlockchainService
var blockchainInitErr error

func init() {
	blockchainService, blockchainInitErr = services.NewBlockchainService()
}

// Blockchain endpoints para reserva, finalização e histórico

// POST /blockchain/reservar
func BlockchainReservar(c *gin.Context) {
	if blockchainInitErr != nil || blockchainService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "BlockchainService não inicializado: " + blockchainInitErr.Error()})
		return
	}
	id, err := blockchainService.IniciarSessao()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"idSessao": id.String()})
}

// POST /blockchain/finalizar
func BlockchainFinalizar(c *gin.Context) {
	if blockchainInitErr != nil || blockchainService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "BlockchainService não inicializado: " + blockchainInitErr.Error()})
		return
	}
	var req struct {
		IDSessao         string `json:"idSessao"`
		EnergiaConsumida string `json:"energiaConsumida"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato inválido"})
		return
	}
	id := new(big.Int)
	id.SetString(req.IDSessao, 10)
	energia := new(big.Int)
	energia.SetString(req.EnergiaConsumida, 10)
	err := blockchainService.FinalizarSessao(id, energia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Sessão finalizada e pagamento registrado na blockchain"})
}

// GET /blockchain/historico
func BlockchainHistorico(c *gin.Context) {
	if blockchainInitErr != nil || blockchainService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "BlockchainService não inicializado: " + blockchainInitErr.Error()})
		return
	}
	historico, err := blockchainService.GetHistorico()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, historico)
}

// GET /blockchain/sessao/:id
func BlockchainSessao(c *gin.Context) {
	if blockchainInitErr != nil || blockchainService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "BlockchainService não inicializado: " + blockchainInitErr.Error()})
		return
	}
	idStr := c.Param("id")
	id := new(big.Int)
	id.SetString(idStr, 10)
	usuario, energia, valor, finalizada, err := blockchainService.GetSessao(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"usuario":          usuario.Hex(),
		"energiaConsumida": energia.String(),
		"valorPago":        valor.String(),
		"finalizada":       finalizada,
	})
}
