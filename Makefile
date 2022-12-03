
all: proto
	go build -o client cmd/client/main.go
	go build -o server cmd/server/main.go

proto:
	@protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. proto/dialog.proto

.PHONY: proto
