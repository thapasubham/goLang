
build:
	@if not exist bin mkdir bin
	@go build -o bin/eco cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/eco
