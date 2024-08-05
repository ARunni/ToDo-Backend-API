run:
	go run cmd/main.go

swag: 
	swag init -g cmd/main.go -o ./cmd/docs

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/todo.proto

deps:
	go mod tidy