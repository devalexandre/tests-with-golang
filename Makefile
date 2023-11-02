.PHONY: test test-coverage test-coverage-html

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...

test-coverage-html:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

run: test test-coverage test-coverage-html
