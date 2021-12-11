.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	# flags allow revive to pick up uncommented exported funcs
	golangci-lint run --exclude-use-default=false --enable=revive
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

pb:
	protoc --go_out=./pkg/hellopb \
	--go_opt=paths=source_relative \
	--go-grpc_out=./pkg/hellopb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=../proto \
	hello.proto

build: vet
	go build .
.PHONY:build