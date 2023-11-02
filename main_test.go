// pokemon/pokemon_test.go
package pokemon

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func restore() {
	ReadAll = io.ReadAll
}

func TestGetPokemon(t *testing.T) {
	t.Run("Pokemon encontrado", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"name": "ditto"}`)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		pokemon, err := GetPokemon("ditto")
		if err != nil {
			t.Fatalf("não deveria ter erro, mas obteve %v", err)
		}

		if pokemon.Name != "ditto" {
			t.Errorf("esperava nome 'ditto', mas obteve '%s'", pokemon.Name)
		}
	})

	t.Run("Erro de status", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

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

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"name": "ditto"`) // JSON inválido para simular erro de deserialização
		})

		server := httptest.NewServer(handler)
		defer server.Close()

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

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"name": "ditto"`) // JSON inválido para simular erro de deserialização
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		_, err := GetPokemon("ditto")
		if err == nil {
			t.Fatal("Deveria ter erro, mas não obteve")
		}
	})

}
