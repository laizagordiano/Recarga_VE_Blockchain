package models

type RequisicaoVeiculo struct {
	VeiculoID string `json:"veiculo_id"`
	Bateria   int    `json:"bateria"`
	Local     string `json:"local"`
	Destino   string `json:"destino"`
}
