package routes

import (
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(router *gin.Engine) {
	router.GET("/points", controllers.GetAllPoints)                        // lista todos os pontos de recarga
	router.POST("/reserve-points/:ids", controllers.PostPoints)            // reserva múltiplos pontos (por IDs)
	router.GET("/routes", controllers.GetAllRoutes)                        // lista todas as rotas cadastradas
	router.GET("/routes/:id", controllers.GetRouteByID)                    // obtém uma rota específica por ID
	router.POST("/plan-trip", controllers.PlanTrip)                        // retorna pontos disponíveis entre origem e destino
	router.POST("/reserve-sequence", controllers.ReserveSequence)          // reserva uma sequência de pontos de recarga
	router.POST("/cancel-reservation", controllers.CancelReservation)      // cancela reservas por ID
	router.POST("/cancel-reservation/:ids", controllers.CancelPointsByIDs) // cancela reservas por lista de IDs na URL
	router.GET("/health", controllers.HealthCheck)                         // status do servidor

	// --- ENDPOINTS RECOMENDADOS PARA PRODUÇÃO ---
	// Blockchain: máxima transparência, segurança e auditabilidade
	router.POST("/blockchain/reservar", controllers.BlockchainReservar)   // inicia sessão (reserva) na blockchain
	router.POST("/blockchain/finalizar", controllers.BlockchainFinalizar) // finaliza sessão/pagamento na blockchain
	router.GET("/blockchain/historico", controllers.BlockchainHistorico)  // histórico de sessões/pagamentos
	router.GET("/blockchain/sessao/:id", controllers.BlockchainSessao)    // detalhes de uma sessão

}
