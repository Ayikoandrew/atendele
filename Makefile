run: build
	@./bin/atendele

build:
	@go build -o bin/atendele

test:
	@go test -v ./...