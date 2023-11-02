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

func GetPokemon(name string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("erro ao buscar Pokémon: status %d", resp.StatusCode)
	}

	body, err := ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
