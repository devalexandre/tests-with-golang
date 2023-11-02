// main_test.go
package main

import "testing"

func TestSubtrair(t *testing.T) {
    casosDeTeste := []struct {
        a, b      int
        esperado  int
        esperadoErro bool
    }{
        {5, 3, 2, false},
        {3, 5, 0, true},
        {7, 2, 5, false},
    }

    for _, caso := range casosDeTeste {
        resultado, err := Subtrair(caso.a, caso.b)
        if caso.esperadoErro && err == nil {
            t.Errorf("Subtrair(%d, %d) esperava um erro, mas não obteve", caso.a, caso.b)
        } else if !caso.esperadoErro && resultado != caso.esperado {
            t.Errorf("Subtrair(%d, %d) esperava %d, mas obteve %d", caso.a, caso.b, caso.esperado, resultado)
        }
    }
}

