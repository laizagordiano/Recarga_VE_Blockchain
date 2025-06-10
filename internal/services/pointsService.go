package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/config"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/database"
	"github.com/laizagordiano/Recarga_VE_Blockchain/internal/models"
)

// Lógica de reserva/cancelamento de pontos de recarga
// Aqui ocorre a verificação de disponibilidade, reserva atômica e rollback
// Comunicação entre servidores (empresas) ocorre via HTTP
// Se a reserva envolver múltiplas empresas, é feita uma coordenação distribuída
// Se qualquer empresa recusar, ocorre rollback para garantir consistência

func GetAllPoints() []models.PontoRecarga {
	return database.Pontos
}

func isPontoDaEmpresa(ponto models.PontoRecarga) bool {
	idEsperado := config.EmpresaNomeParaID[config.NomeEmpresa]
	// Compara nome ou ID numérico
	match := ponto.EmpresaID == config.NomeEmpresa || ponto.EmpresaID == idEsperado
	return match
}

// Atualiza o campo Disponivel nas rotas sempre que um ponto é reservado ou liberado
func AtualizarDisponibilidadeNasRotas(pontoID string, disponivel bool) {
	for r := range database.Rotas {
		for p := range database.Rotas[r].Pontos {
			if database.Rotas[r].Pontos[p].ID == pontoID {
				database.Rotas[r].Pontos[p].Disponivel = disponivel
			}
		}
	}
}

// Função auxiliar para buscar o nome da empresa a partir do ID
func getEmpresaNomeDoID(id string) string {
	for nome, num := range config.EmpresaNomeParaID {
		if num == id {
			return nome
		}
	}
	return id // fallback
}

func ReservePoints(ids []string) (reservados []string, indisponiveis []string, err error) {
	type pontoStatus struct {
		id         string
		local      bool
		disponivel bool
	}
	var statusList []pontoStatus

	// 1. Verifica disponibilidade de todos os pontos (NÃO reserva ainda)
	for _, id := range ids {
		found := false
		for _, ponto := range database.Pontos {
			if ponto.ID == id {
				found = true
				if isPontoDaEmpresa(ponto) {
					if ponto.Disponivel {
						statusList = append(statusList, pontoStatus{id, true, true})
					} else {
						statusList = append(statusList, pontoStatus{id, true, false})
					}
				} else {
					empresaURL := getEmpresaURLDoPonto(ponto.EmpresaID)
					if empresaURL == "" {
						statusList = append(statusList, pontoStatus{id, false, false})
						continue
					}
					resp, err := http.Get(fmt.Sprintf("%s/points", empresaURL))
					if err != nil {
						statusList = append(statusList, pontoStatus{id, false, false})
						continue
					}
					body, _ := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					var pontosRemotos []models.PontoRecarga
					json.Unmarshal(body, &pontosRemotos)
					ok := false
					for _, pr := range pontosRemotos {
						if pr.ID == id && pr.Disponivel {
							ok = true
							break
						}
					}
					statusList = append(statusList, pontoStatus{id, false, ok})
				}
				break
			}
		}
		if !found {
			statusList = append(statusList, pontoStatus{id, false, false})
		}
	}

	// 2. Se algum indisponível, retorna erro
	for _, st := range statusList {
		if !st.disponivel {
			for _, s := range statusList {
				if !s.disponivel {
					indisponiveis = append(indisponiveis, s.id)
				}
			}
			err = errors.New("nenhum ponto foi reservado")
			return nil, indisponiveis, err
		}
	}

	// 3. Reserva todos (agora é garantido que todos estão disponíveis)
	for _, st := range statusList {
		if st.local {
			for i, ponto := range database.Pontos {
				if ponto.ID == st.id {
					database.Pontos[i].Disponivel = false
					AtualizarDisponibilidadeNasRotas(st.id, false)
					reservados = append(reservados, st.id)
					break
				}
			}
		} else {
			empresaURL := getEmpresaURLDoPonto(getEmpresaIDDoPonto(st.id))
			if empresaURL == "" {
				indisponiveis = append(indisponiveis, st.id)
				err = errors.New("nenhum ponto foi reservado")
				return nil, indisponiveis, err
			}
			urlReserva := fmt.Sprintf("%s/reserve-points/%s", empresaURL, st.id)
			req, _ := http.NewRequest("POST", urlReserva, nil)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				indisponiveis = append(indisponiveis, st.id)
				err = errors.New("nenhum ponto foi reservado")
				return nil, indisponiveis, err
			}
			defer resp.Body.Close()
			var resultado map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&resultado)
			if _, ok := resultado["reservados"]; ok {
				AtualizarDisponibilidadeNasRotas(st.id, false)
				reservados = append(reservados, st.id)
			} else {
				indisponiveis = append(indisponiveis, st.id)
				err = errors.New("nenhum ponto foi reservado")
				return nil, indisponiveis, err
			}
		}
	}
	return reservados, nil, nil
}

// Função auxiliar para buscar o EmpresaID de um ponto
func getEmpresaIDDoPonto(pontoID string) string {
	for _, ponto := range database.Pontos {
		if ponto.ID == pontoID {
			return ponto.EmpresaID
		}
	}
	return ""
}

// Função auxiliar para buscar a URL da empresa a partir do ID
func getEmpresaURLDoPonto(empresaID string) string {
	// Busca nas variáveis de ambiente
	if empresaID == "1" || empresaID == "empresa_a" {
		return os.Getenv("EMPRESA_A_URL")
	} else if empresaID == "2" || empresaID == "empresa_b" {
		return os.Getenv("EMPRESA_B_URL")
	} else if empresaID == "3" || empresaID == "empresa_c" {
		return os.Getenv("EMPRESA_C_URL")
	}
	return ""
}

// Função para cancelar reservas de múltiplos pontos (local e remoto)
func CancelPoints(ids []string) (cancelados []string, naoCancelados []string, err error) {
	type pontoStatus struct {
		id    string
		local bool
		ok    bool
	}
	var statusList []pontoStatus

	// 1. Verifica se todos os pontos existem
	for _, id := range ids {
		found := false
		for _, ponto := range database.Pontos {
			if ponto.ID == id {
				found = true
				if isPontoDaEmpresa(ponto) {
					if !ponto.Disponivel {
						statusList = append(statusList, pontoStatus{id, true, true})
					} else {
						statusList = append(statusList, pontoStatus{id, true, false})
					}
				} else {
					empresaURL := getEmpresaURLDoPonto(ponto.EmpresaID)
					if empresaURL == "" {
						statusList = append(statusList, pontoStatus{id, false, false})
						continue
					}
					// Consulta status remoto
					resp, err := http.Get(fmt.Sprintf("%s/points", empresaURL))
					if err != nil {
						statusList = append(statusList, pontoStatus{id, false, false})
						continue
					}
					body, _ := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					var pontosRemotos []models.PontoRecarga
					json.Unmarshal(body, &pontosRemotos)
					ok := false
					for _, pr := range pontosRemotos {
						if pr.ID == id && !pr.Disponivel {
							ok = true
							break
						}
					}
					statusList = append(statusList, pontoStatus{id, false, ok})
				}
				break
			}
		}
		if !found {
			statusList = append(statusList, pontoStatus{id, false, false})
		}
	}

	// 2. Se algum não está reservado, retorna erro
	for _, st := range statusList {
		if !st.ok {
			for _, s := range statusList {
				if !s.ok {
					naoCancelados = append(naoCancelados, s.id)
				}
			}
			err = errors.New("nenhum ponto foi cancelado")
			return nil, naoCancelados, err
		}
	}

	// 3. Cancela todos (agora é garantido que todos estão reservados)
	for _, st := range statusList {
		if st.local {
			for i, ponto := range database.Pontos {
				if ponto.ID == st.id {
					database.Pontos[i].Disponivel = true
					AtualizarDisponibilidadeNasRotas(st.id, true)
					cancelados = append(cancelados, st.id)
					break
				}
			}
		} else {
			empresaURL := getEmpresaURLDoPonto(getEmpresaIDDoPonto(st.id))
			if empresaURL == "" {
				naoCancelados = append(naoCancelados, st.id)
				err = errors.New("nenhum ponto foi cancelado")
				return nil, naoCancelados, err
			}
			urlCancel := fmt.Sprintf("%s/cancel-reservation/%s", empresaURL, st.id)
			req, _ := http.NewRequest("POST", urlCancel, nil)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				naoCancelados = append(naoCancelados, st.id)
				err = errors.New("nenhum ponto foi cancelado")
				return nil, naoCancelados, err
			}
			defer resp.Body.Close()
			var resultado map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&resultado)
			if _, ok := resultado["cancelados"]; ok {
				AtualizarDisponibilidadeNasRotas(st.id, true)
				cancelados = append(cancelados, st.id)
			} else {
				naoCancelados = append(naoCancelados, st.id)
				err = errors.New("nenhum ponto foi cancelado")
				return nil, naoCancelados, err
			}
		}
	}
	return cancelados, nil, nil
}

// Toda a lógica de reserva/cancelamento local foi removida.
// Agora, utilize apenas os endpoints blockchain para reservas e pagamentos.

// OBS: Este arquivo foi descontinuado. Toda a lógica de negócio está na blockchain.
