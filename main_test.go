// pokemon/pokemon_test.go
package pokemon

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func resetReadAll() {
	ReadAll = io.ReadAll
	HTTPGet = http.Get
}

func TestMain(m *testing.M) {
	resetReadAll()
	m.Run()
}

func TestGetPokemon(t *testing.T) {
	t.Run("Pokemon encontrado", func(t *testing.T) {
		ReadAll = func(r io.Reader) ([]byte, error) {
			return json.Marshal(Pokemon{Name: "pikachu"})
		}

		HTTPGet = mockHTTPGet

		// Crie o seu handler http
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r)
		})

		// Crie um servidor de teste
		server := httptest.NewServer(handler)
		defer server.Close()

		// Construa a URL com parâmetros de query
		params := url.Values{}
		params.Add("name", "pikachu")
		req, _ := http.NewRequest(http.MethodGet, server.URL+"?"+params.Encode(), nil)

		// Execute a requisição ao servidor de teste
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

		// Verifique o status code da resposta
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK; got %v", res.StatusCode)
		}

		// Decode the response and check if it's the expected pokemon
		var p Pokemon
		err = json.NewDecoder(res.Body).Decode(&p)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		if p.Name != "pikachu" {
			t.Errorf("Expected pokemon name to be 'pikachu', got '%v'", p.Name)
		}
	})

	t.Run("Error ReadAll", func(t *testing.T) {
		ReadAll = func(r io.Reader) ([]byte, error) {
			return nil, io.ErrUnexpectedEOF
		}

		HTTPGet = mockHTTPGet

		// Crie o seu handler http
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r)
		})

		// Crie um servidor de teste
		server := httptest.NewServer(handler)
		defer server.Close()

		// Construa a URL com parâmetros de query
		params := url.Values{}
		params.Add("name", "pikachu")
		req, _ := http.NewRequest(http.MethodGet, server.URL+"?"+params.Encode(), nil)

		// Execute a requisição ao servidor de teste
		res, err := http.DefaultClient.Do(req)
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status Internal Server Error; got %v", res.StatusCode)
		}

		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

	})

	t.Run("Error Unmarshal", func(t *testing.T) {
		ReadAll = func(r io.Reader) ([]byte, error) {
			return []byte(`{"Name":123}`), nil
		}

		HTTPGet = mockHTTPGet

		// Crie o seu handler http
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r)
		})

		// Crie um servidor de teste
		server := httptest.NewServer(handler)
		defer server.Close()

		// Construa a URL com parâmetros de query
		params := url.Values{}
		params.Add("name", "pikachu")
		req, _ := http.NewRequest(http.MethodGet, server.URL+"?"+params.Encode(), nil)

		// Execute a requisição ao servidor de teste
		res, err := http.DefaultClient.Do(req)
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status Internal Server Error; got %v", res.StatusCode)
		}

		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

	})

	t.Run("Error request", func(t *testing.T) {
		HTTPGet = mockHTTPGetError

		// Crie o seu handler http
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r)
		})

		// Crie um servidor de teste
		server := httptest.NewServer(handler)
		defer server.Close()

		// Construa a URL com parâmetros de query
		params := url.Values{}
		params.Add("name", "pikachu")
		req, _ := http.NewRequest(http.MethodGet, server.URL+"?"+params.Encode(), nil)

		// Execute a requisição ao servidor de teste
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

		// Verifique o status code da resposta
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status Internal Server Error; got %v", res.StatusCode)
		}
	})

	t.Run("Error request http", func(t *testing.T) {
		HTTPGet = mockHTTPGetErrorRequest

		// Crie o seu handler http
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r)
		})

		// Crie um servidor de teste
		server := httptest.NewServer(handler)
		defer server.Close()

		// Construa a URL com parâmetros de query
		params := url.Values{}
		params.Add("name", "pikachu")
		req, _ := http.NewRequest(http.MethodGet, server.URL+"?"+params.Encode(), nil)

		// Execute a requisição ao servidor de teste
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

		// Verifique o status code da resposta
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status Internal Server Error; got %v", res.StatusCode)
		}
	})
}
