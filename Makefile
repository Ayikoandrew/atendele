run: build
	@./bin/atendele

build:
	@go build -o bin/atendele

test:
	@go test -v ./...

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		core/block.proto