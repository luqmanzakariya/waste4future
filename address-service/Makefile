.PHONY: proto

proto:
	protoc --go_out=. --go-grpc_out=. pb/*.proto

proto_protobuf_only:
	protoc --go_out=. pb/*.proto