.PHONY: all build clean

NAME := server1

all: build

build:
	cd $(NAME)/proto && protoc --go_out=. --go-grpc_out=. services.proto
	cd $(NAME) && go mod init $(NAME)
	cd $(NAME) && go mod tidy

clean:
	cd $(NAME)/proto/grpc-server/proto && del services_grpc.pb.go
	cd $(NAME)/proto/grpc-server/proto && del services.pb.go
	cd $(NAME)/proto/grpc-server && rmdir /q proto
	cd $(NAME)/proto && rmdir /q grpc-server
	cd $(NAME) && del go.mod
	cd $(NAME) && del go.sum