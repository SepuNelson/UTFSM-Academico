.PHONY: all build clean

NAME := player1

all: build

build:
	cd proto && protoc --go_out=. --go-grpc_out=. services.proto
	go mod init $(NAME)
	go mod tidy

clean:
	cd proto/grpc-server/proto && del services_grpc.pb.go
	cd proto/grpc-server/proto && del services.pb.go
	cd proto/grpc-server && rmdir /q proto
	cd proto && rmdir /q grpc-server
	del go.mod
	del go.sum