Módulo 1: Introdução a testes de software

    O que são testes de software?
    Quais são os diferentes tipos de testes de software?
    O que são testes unitários?
    O que são testes de integração?
    O que são testes de aceitação?


### Run tests
```bash
go test -v ./...
```

### Run tests with coverage
```bash
go test -coverprofile=coverage.out ./...
```

### Run tests with coverage and generate html
```bash
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```