.PHONY: all build clean

MODULE_NAME := lcp

all: build

build:
	cd $(MODULE_NAME)/proto && protoc --go_out=. --go-grpc_out=. services.proto || echo "protoc ya existe"
	cd $(MODULE_NAME) && go mod init $(MODULE_NAME) || echo "go.mod ya existe"
	cd $(MODULE_NAME) && go mod tidy || echo "go.sum ya existe"

clean:
	cd $(MODULE_NAME)/proto/grpc-server/proto && del services_grpc.pb.go
	cd $(MODULE_NAME)/proto/grpc-server/proto && del services.pb.go
	cd $(MODULE_NAME)/proto/grpc-server && rmdir /q proto
	cd $(MODULE_NAME)/proto && rmdir /q grpc-server
	cd $(MODULE_NAME) && del go.mod
	cd $(MODULE_NAME) && del go.sum
