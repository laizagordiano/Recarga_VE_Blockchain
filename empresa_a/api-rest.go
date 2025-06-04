// Empresa A - API REST para gerenciamento de reservas
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/google/uuid"
)

type Reserva struct {
	ID      string `json:"id"`
	Veiculo string `json:"veiculo"`
	Hora    string `json:"hora"`
	Status  string `json:"status"`
}

var reservasA = make(map[string]Reserva)
var muA sync.Mutex

func criarReservaA(w http.ResponseWriter, r *http.Request) {
	var nova Reserva
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	muA.Lock()
	defer muA.Unlock()

	id := uuid.New().String()
	nova.ID = id
	nova.Status = "confirmada"
	reservasA[id] = nova

	json.NewEncoder(w).Encode(nova)
}

func statusReservaA(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/status/"):] // extrai o ID da URL
	muA.Lock()
	reserva, ok := reservasA[id]
	muA.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(reserva)
}

func main() {
	http.HandleFunc("/reserva", criarReservaA)
	http.HandleFunc("/status/", statusReservaA)
	fmt.Println("Empresa A rodando em :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
