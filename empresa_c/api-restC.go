// Empresa C - API REST para gerenciamento de reservas
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

var reservasC = make(map[string]Reserva)
var muC sync.Mutex

func criarReservaC(w http.ResponseWriter, r *http.Request) {
	var nova Reserva
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	muC.Lock()
	defer muC.Unlock()

	id := uuid.New().String()
	nova.ID = id
	nova.Status = "confirmada"
	reservasC[id] = nova

	json.NewEncoder(w).Encode(nova)
}

func statusReservaC(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/status/"):] // extrai o ID da URL
	muC.Lock()
	reserva, ok := reservasC[id]
	muC.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(reserva)
}

func main() {
	http.HandleFunc("/reserva", criarReservaC)
	http.HandleFunc("/status/", statusReservaC)
	fmt.Println("Empresa C rodando em :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
