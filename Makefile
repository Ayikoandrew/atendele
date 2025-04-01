.PHONY: gen-pb

run: build
	@./bin/atendele

build:
	@go build -o bin/atendele

test:
	@go test -v ./...


proto:
	@echo Starting generate pb
	@protoc --go_out=./types --go_opt=paths=source_relative \
		--go-grpc_out=./types --go-grpc_opt=paths=source_relative \
		core/block.proto
	@echo Successfully generated proto
	@echo Starting inject tags
	@protoc-go-inject-tag -input="./*.pb.go"
	@echo Successfully injected tags