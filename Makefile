
all:
	docker-compose build

server:
	docker-compose up server

dev:
	docker-compose up devserver

test:
	docker-compose up test

grpc:
	protoc --include_imports --include_source_info -I${GAPIS} protobuf/move.proto  --proto_path protobuf/  --descriptor_set_out deployment/endpoint/move.pb
	protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GAPIS} --go_out=plugins=grpc:moveio protobuf/move.proto
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:moveio protobuf/move.proto
