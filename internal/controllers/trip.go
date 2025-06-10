package controllers

import (
	"net/http"

	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/models"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/services"

	"github.com/gin-gonic/gin"
)

func PlanTrip(c *gin.Context) {
	c.JSON(http.StatusGone, gin.H{"erro": "Este endpoint está obsoleto. Use os endpoints /blockchain para consultar pontos e planejar viagens reais."})
}

func ReserveSequence(c *gin.Context) {
	var req models.SequenciaReserva
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato inválido"})
		return
	}

	ids := []string{}
	for _, ponto := range req.Pontos {
		ids = append(ids, ponto.ID)
	}

	reservados, indisponiveis, err := services.ReservePoints(ids)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"erro": err.Error(), "indisponiveis": indisponiveis})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservados":    reservados,
		"indisponiveis": indisponiveis,
	})
}

func CancelReservation(c *gin.Context) {
	c.JSON(http.StatusGone, gin.H{"erro": "Este endpoint está obsoleto. Use os endpoints /blockchain para cancelar reservas reais."})
}
