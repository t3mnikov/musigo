build:
	go build -v -o bin/ ./cmd/musigo

.DEFAULT_GOAL := build
