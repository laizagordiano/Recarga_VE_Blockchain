package services

import (
	"errors"

	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/database"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/models"
)

// Serviço responsável pelo cálculo de rotas
// O backend pode consultar outros servidores para montar rotas distribuídas
// O resultado pode incluir pontos de recarga de múltiplas empresas
// O sistema pode mostrar todas as rotas possíveis considerando todos os pontos disponíveis

func GetAllRoutes() []models.Rota {
	return database.Rotas
}

func GetRouteByID(id string) (models.Rota, error) {
	for _, rota := range database.Rotas {
		if rota.ID == id {
			return rota, nil
		}
	}
	return models.Rota{}, errors.New("rota não encontrada")
}
