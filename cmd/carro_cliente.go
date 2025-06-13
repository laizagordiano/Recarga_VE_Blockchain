package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Estrutura que representa um ponto de recarga retornado pelo backend
// Comunicação cliente-servidor ocorre via HTTP
// O cliente pode buscar rotas, reservar e cancelar pontos
// Comunicação HTTP: GET /routes, POST /reserve, POST /cancel

type PontoRecarga struct {
	ID          string `json:"id"`
	Localizacao string `json:"localizacao"`
	Disponivel  bool   `json:"disponivel"`
	EmpresaID   string `json:"empresa_id"`
}

type Rota struct {
	ID     string         `json:"id"`
	Nome   string         `json:"nome"`
	Pontos []PontoRecarga `json:"pontos"`
}

// Função auxiliar para converter um slice em JSON
func toJSON(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	servidor := os.Getenv("SERVER_URL")
	if servidor == "" {
		servidor = "http://localhost:8085"
	}

	// 1. Buscar rotas disponíveis
	// Aqui o cliente consulta o backend para obter todas as rotas possíveis
	// O backend pode consultar outros servidores para montar rotas distribuídas
	resp, err := http.Get(servidor + "/routes")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var rotas []Rota
	if err := json.NewDecoder(resp.Body).Decode(&rotas); err != nil {
		panic(err)
	}
	fmt.Println("Rotas disponíveis:")
	for i, r := range rotas {
		fmt.Printf("%d - %s\n", i+1, r.Nome)
	}

	// 2. Escolhe uma rota
	rota := rotas[rand.Intn(len(rotas))]
	fmt.Printf("\nRota escolhida: %s\n", rota.Nome)
	fmt.Println("Pontos dessa rota:")
	for i, p := range rota.Pontos {
		fmt.Printf("%d - %s (%s) [Disponível: %v]\n", i+1, p.ID, p.Localizacao, p.Disponivel)
	}

	// 3. Escolher até 3 pontos disponíveis aleatórios
	var pontosDisponiveis []PontoRecarga
	for _, p := range rota.Pontos {
		if p.Disponivel {
			pontosDisponiveis = append(pontosDisponiveis, p)
		}
	}
	if len(pontosDisponiveis) == 0 {
		fmt.Println("Nenhum ponto disponível para reservar nesta rota!")
		return
	}

	// Selecionar até 3 pontos aleatórios
	numReservas := 3
	if len(pontosDisponiveis) < numReservas {
		numReservas = len(pontosDisponiveis)
	}
	pontosEscolhidos := pontosDisponiveis[:numReservas]
	fmt.Println("\nPontos escolhidos:")
	for _, ponto := range pontosEscolhidos {
		fmt.Printf("- %s (%s)\n", ponto.ID, ponto.Localizacao)
	}

	// 4. Reservar os pontos na blockchain (um idSessao para cada ponto)
	sessoes := make([]string, 0, len(pontosEscolhidos))
	for _, ponto := range pontosEscolhidos {
		resp, err := http.Post(servidor+"/blockchain/reservar", "application/json", nil)
		if err != nil {
			fmt.Printf("Erro ao reservar ponto %s: %v\n", ponto.ID, err)
			continue
		}
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		resp.Body.Close()
		if id, ok := res["idSessao"]; ok {
			fmt.Printf("Sessão blockchain criada para ponto %s: idSessao=%v\n", ponto.ID, id)
			sessoes = append(sessoes, fmt.Sprintf("%v", id))
		} else {
			fmt.Printf("Falha ao reservar ponto %s na blockchain: %v\n", ponto.ID, res["erro"])
		}
	}
	if len(sessoes) == 0 {
		fmt.Println("Nenhuma sessão foi criada na blockchain. Encerrando fluxo.")
		return
	}

	// 5. Simular viagem em partes e finalizar/pagar cada sessão na blockchain
	bateria := 100    // porcentagem inicial da bateria
	consumoPorKm := 2 // consumo de bateria por km em %
	fmt.Println("\nIniciando a viagem...")
	for i := 0; i < len(pontosEscolhidos); i++ {
		var origem, destino string
		if i == 0 {
			origem = "Início da rota"
		} else {
			origem = pontosEscolhidos[i-1].Localizacao
		}
		destino = pontosEscolhidos[i].Localizacao

		distancia := rand.Intn(16) + 5 // distância aleatória entre 5 e 20 km
		consumoTrecho := distancia * consumoPorKm

		fmt.Printf("Bateria antes do trecho: %d%%\n", bateria)
		fmt.Printf("Trecho de %s para %s: %d km, consumo estimado: %d%%\n", origem, destino, distancia, consumoTrecho)
		if bateria <= 0 {
			fmt.Println("Bateria esgotada! Viagem interrompida.")
			break
		}

		duracao := rand.Intn(3) + 7 // Tempo entre 3 e 7 segundos para cada trecho
		fmt.Printf("Viajando de %s para %s...\n", origem, destino)
		time.Sleep(time.Duration(duracao) * time.Second)
		bateria -= consumoTrecho
		if bateria < 0 {
			bateria = 0
		}

		// Finalizar/pagar sessão na blockchain
		idSessao := sessoes[i]
		energiaConsumida := consumoTrecho // pode ser ajustado conforme lógica real
		body := fmt.Sprintf(`{"idSessao":"%s","energiaConsumida":"%d"}`, idSessao, energiaConsumida)
		resp, err := http.Post(servidor+"/blockchain/finalizar", "application/json", strings.NewReader(body))
		if err != nil {
			fmt.Printf("Erro ao finalizar sessão %s: %v\n", idSessao, err)
			continue
		}
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		resp.Body.Close()
		if _, ok := res["mensagem"]; ok {
			fmt.Printf("Pagamento efetivado na blockchain! Chegou em %s! Bateria restante: %d%%\n", destino, bateria)
		} else {
			fmt.Printf("Falha ao finalizar sessão %s: %v\n", idSessao, res["erro"])
		}
		if bateria == 0 {
			fmt.Println("Bateria esgotada! Viagem encerrada.")
			break
		}
	}
	if bateria > 0 {
		fmt.Println("Viagem concluída!")
	}
}
