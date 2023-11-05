// pokemon/pokemon.go
package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Declaração de uma variável que pode ser sobreposta pelos testes.
var (
	ReadAll = io.ReadAll
	HTTPGet = http.Get
)

type Pokemon struct {
	Name string `json:"name"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := HTTPGet(url) // Aqui usamos a variável sobreposta.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, fmt.Sprintf("erro ao buscar Pokémon: status %d", resp.StatusCode), resp.StatusCode)
		return
	}

	body, err := ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&pokemon)
}
