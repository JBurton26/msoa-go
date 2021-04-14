.PHONY: protos

protos:
	protoc -I protos/ protos/inventory.proto --go-grpc_out=protos/inventory --go_out=protos/inventory --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	protoc -I protos/ protos/user.proto --go-grpc_out=protos/user --go_out=protos/user --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	protoc -I protos/ protos/cost.proto --go-grpc_out=protos/cost --go_out=protos/cost --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
