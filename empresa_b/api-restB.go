// Empresa B - API REST para gerenciamento de reservas
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

var reservasB = make(map[string]Reserva)
var muB sync.Mutex

func criarReservaB(w http.ResponseWriter, r *http.Request) {
	var nova Reserva
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	muB.Lock()
	defer muB.Unlock()

	id := uuid.New().String()
	nova.ID = id
	nova.Status = "confirmada"
	reservasB[id] = nova

	json.NewEncoder(w).Encode(nova)
}

func statusReservaB(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/status/"):] // extrai o ID da URL
	muB.Lock()
	reserva, ok := reservasB[id]
	muB.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(reserva)
}

func main() {
	http.HandleFunc("/reserva", criarReservaB)
	http.HandleFunc("/status/", statusReservaB)
	fmt.Println("Empresa B rodando em :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
