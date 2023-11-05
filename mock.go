package pokemon

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func mockHTTPGet(url string) (*http.Response, error) {
	// Crie uma resposta mock
	mockResp := `{"name":"pikachu"}`
	r := io.NopCloser(bytes.NewReader([]byte(mockResp))) // Cria um io.ReadCloser
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

// case error status
func mockHTTPGetError(url string) (*http.Response, error) {
	// Crie uma resposta mock
	mockResp := `{"name":"pikachu"}`
	r := io.NopCloser(bytes.NewReader([]byte(mockResp))) // Cria um io.ReadCloser
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       r,
	}, nil
}

// error request
func mockHTTPGetErrorRequest(url string) (*http.Response, error) {
	// Crie uma resposta mock
	return nil, fmt.Errorf("error request")
}
