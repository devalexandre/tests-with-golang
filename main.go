// pokemon/pokemon.go
package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var ReadAll = io.ReadAll

type Pokemon struct {
	Name string `json:"name"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)
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
