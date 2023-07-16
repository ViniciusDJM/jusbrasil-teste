all: build

install-deps:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golang/mock/mockgen@latest

code-gen:
	go generate ./...
	swag init --parseInternal --parseDepth 0 -o ./api

build: code-gen
	go build -ldflags="-w -s" -o bin/jusbrasil_process-finder.exe .

clean:
	rm bin/*
