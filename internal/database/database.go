package database

import (
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/models"
)

var Empresas = []models.Empresa{
	{
		ID:   "1",
		Nome: "HenryVolt",
		Pontos: []models.PontoRecarga{
			{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
			{ID: "2", Localizacao: "Salvador", Disponivel: true, EmpresaID: "1"},
		},
	},
	{
		ID:   "2",
		Nome: "LaizaCharge",
		Pontos: []models.PontoRecarga{
			{ID: "3", Localizacao: "Valente", Disponivel: true, EmpresaID: "2"},
			{ID: "4", Localizacao: "Santaluz", Disponivel: true, EmpresaID: "2"},
		},
	},
	{
		ID:   "3",
		Nome: "MarioPower",
		Pontos: []models.PontoRecarga{
			{ID: "5", Localizacao: "Amargosa", Disponivel: true, EmpresaID: "3"},
		},
	},
}

var Pontos = []models.PontoRecarga{
	{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
	{ID: "2", Localizacao: "Salvador", Disponivel: true, EmpresaID: "1"},
	{ID: "3", Localizacao: "Valente", Disponivel: true, EmpresaID: "2"},
	{ID: "4", Localizacao: "Santaluz", Disponivel: true, EmpresaID: "2"},
	{ID: "5", Localizacao: "Amargosa", Disponivel: true, EmpresaID: "3"},
	{ID: "6", Localizacao: "Amélia Rodrigues", Disponivel: true, EmpresaID: "3"},
	{ID: "7", Localizacao: "Camaçari", Disponivel: true, EmpresaID: "3"},
	{ID: "8", Localizacao: "Candeias", Disponivel: true, EmpresaID: "2"},
	{ID: "9", Localizacao: "Cruz das Almas", Disponivel: true, EmpresaID: "1"},
	{ID: "10", Localizacao: "Santo Antonio de Jesus", Disponivel: true, EmpresaID: "1"},
	{ID: "11", Localizacao: "Itabuna", Disponivel: true, EmpresaID: "2"},
	{ID: "12", Localizacao: "Ilheus", Disponivel: true, EmpresaID: "2"},
	{ID: "13", Localizacao: "Porto Seguro", Disponivel: true, EmpresaID: "3"},
	{ID: "14", Localizacao: "Pereira", Disponivel: true, EmpresaID: "3"},
}

var Rotas = []models.Rota{
	{
		ID:   "1",
		Nome: "Rota Feira de Santana - Salvador",
		Pontos: []models.PontoRecarga{
			{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
			{ID: "6", Localizacao: "Amélia Rodrigues", Disponivel: true, EmpresaID: "3"},
			{ID: "7", Localizacao: "Camaçari", Disponivel: true, EmpresaID: "3"},
			{ID: "8", Localizacao: "Candeias", Disponivel: true, EmpresaID: "2"},
			{ID: "2", Localizacao: "Salvador", Disponivel: true, EmpresaID: "1"},
		},
	},
	{
		ID:   "2",
		Nome: "Rota Salvador - Porto Seguro",
		Pontos: []models.PontoRecarga{
			{ID: "2", Localizacao: "Salvador", Disponivel: true, EmpresaID: "1"},
			{ID: "9", Localizacao: "Cruz das Almas", Disponivel: true, EmpresaID: "1"},
			{ID: "10", Localizacao: "Santo Antonio de Jesus", Disponivel: true, EmpresaID: "1"},
			{ID: "5", Localizacao: "Amargosa", Disponivel: true, EmpresaID: "3"},
			{ID: "11", Localizacao: "Itabuna", Disponivel: true, EmpresaID: "2"},
			{ID: "12", Localizacao: "Ilheus", Disponivel: true, EmpresaID: "2"},
			{ID: "13", Localizacao: "Porto Seguro", Disponivel: true, EmpresaID: "3"},
		},
	},
	{
		ID:   "3",
		Nome: "Rota Feira de Santana  - Pereira",
		Pontos: []models.PontoRecarga{
			{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
			{ID: "6", Localizacao: "Amélia Rodrigues", Disponivel: true, EmpresaID: "3"},
			{ID: "3", Localizacao: "Valente", Disponivel: true, EmpresaID: "2"},
			{ID: "4", Localizacao: "Santaluz", Disponivel: true, EmpresaID: "2"},
			{ID: "14", Localizacao: "Pereira", Disponivel: true, EmpresaID: "3"},
		},
	},
	{
		ID:   "4",
		Nome: "Rota Salvador - Feira de Santana",
		Pontos: []models.PontoRecarga{
			{ID: "2", Localizacao: "Salvador", Disponivel: true, EmpresaID: "1"},
			{ID: "8", Localizacao: "Candeias", Disponivel: true, EmpresaID: "2"},
			{ID: "7", Localizacao: "Camaçari", Disponivel: true, EmpresaID: "3"},
			{ID: "6", Localizacao: "Amélia Rodrigues", Disponivel: true, EmpresaID: "3"},
			{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
		},
	},
	{
		ID:   "5",
		Nome: "Rota Valente - Amargosa",
		Pontos: []models.PontoRecarga{
			{ID: "3", Localizacao: "Valente", Disponivel: true, EmpresaID: "2"},
			{ID: "4", Localizacao: "Santaluz", Disponivel: true, EmpresaID: "2"},
			{ID: "1", Localizacao: "Feira de Santana", Disponivel: true, EmpresaID: "1"},
			{ID: "9", Localizacao: "Cruz das Almas", Disponivel: true, EmpresaID: "1"},
			{ID: "10", Localizacao: "Santo Antonio de Jesus", Disponivel: true, EmpresaID: "1"},
			{ID: "5", Localizacao: "Amargosa", Disponivel: true, EmpresaID: "3"},
		},
	},
}
