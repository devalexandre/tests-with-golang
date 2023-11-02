// pokemon/pokemon_test.go
package pokemon

import (
	"fmt"
	"io"
	"testing"
)

func restore() {
	ReadAll = io.ReadAll
}

func TestGetPokemon(t *testing.T) {
	t.Run("Pokemon encontrado", func(t *testing.T) {

		pokemon, err := GetPokemon("ditto")
		if err != nil {
			t.Fatalf("não deveria ter erro, mas obteve %v", err)
		}

		if pokemon.Name != "ditto" {
			t.Errorf("esperava nome 'ditto', mas obteve '%s'", pokemon.Name)
		}
	})

	t.Run("Erro de status", func(t *testing.T) {

		_, err := GetPokemon("unknown")
		if err == nil {
			t.Fatal("deveria ter erro, mas não obteve")
		}
	})

	t.Run("Erro de deserialização do JSON", func(t *testing.T) {

		defer restore()

		//simula retorno de JSON inválido
		ReadAll = func(r io.Reader) ([]byte, error) {
			return []byte(`{"Name": "ditto"`), nil
		}

		_, err := GetPokemon("ditto")
		if err == nil {
			t.Fatal("Deveria ter erro, mas não obteve")
		}
	})

	t.Run("Erro de deserialização do JSON", func(t *testing.T) {

		defer restore()

		//simula retorno de JSON inválido
		ReadAll = func(r io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("erro ao ler body")
		}

		_, err := GetPokemon("ditto")
		if err == nil {
			t.Fatal("Deveria ter erro, mas não obteve")
		}
	})

}
