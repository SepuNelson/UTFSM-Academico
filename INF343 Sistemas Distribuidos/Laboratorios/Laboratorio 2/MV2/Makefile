.PHONY: all build clean

MODULE_NAME_CDP := cdp
MODULE_NAME_ENTRENADOR := entrenador

all: build

build:
	cd $(MODULE_NAME_ENTRENADOR)/proto && protoc --go_out=. --go-grpc_out=. services.proto || echo "protoc ya existe"
	cd $(MODULE_NAME_ENTRENADOR) && go mod init $(MODULE_NAME_ENTRENADOR) || echo "go.mod ya existe"
	cd $(MODULE_NAME_ENTRENADOR) && go mod tidy || echo "go.sum ya existe"

	cd $(MODULE_NAME_CDP)/proto && protoc --go_out=. --go-grpc_out=. services.proto || echo "protoc ya existe"
	cd $(MODULE_NAME_CDP) && go mod init $(MODULE_NAME_CDP) || echo "go.mod ya existe"
	cd $(MODULE_NAME_CDP) && go mod tidy || echo "go.sum ya existe"

clean:
	cd $(MODULE_NAME_ENTRENADOR)/proto/grpc-server/proto && del services_grpc.pb.go
	cd $(MODULE_NAME_ENTRENADOR)/proto/grpc-server/proto && del services.pb.go
	cd $(MODULE_NAME_ENTRENADOR)/proto/grpc-server && rmdir /q proto
	cd $(MODULE_NAME_ENTRENADOR)/proto && rmdir /q grpc-server
	cd $(MODULE_NAME_ENTRENADOR) && del go.mod
	cd $(MODULE_NAME_ENTRENADOR) && del go.sum

	cd $(MODULE_NAME_CDP)/proto/grpc-server/proto && del services_grpc.pb.go
	cd $(MODULE_NAME_CDP)/proto/grpc-server/proto && del services.pb.go
	cd $(MODULE_NAME_CDP)/proto/grpc-server && rmdir /q proto
	cd $(MODULE_NAME_CDP)/proto && rmdir /q grpc-server
	cd $(MODULE_NAME_CDP) && del go.mod
	cd $(MODULE_NAME_CDP) && del go.sum
