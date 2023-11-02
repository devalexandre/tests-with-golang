// main.go
package main

import (
    "errors"
)

func Subtrair(a, b int) (int, error) {
    resultado := a - b
    if resultado < 0 {
        return 0, errors.New("erro: o resultado da subtração não pode ser negativo")
    }
    return resultado, nil
}
