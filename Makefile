
all: proto
	go build cmd/client.go
	go build cmd/server.go

proto:
	@protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. proto/dialog.proto

.PHONY: proto
