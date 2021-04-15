.PHONY: protos

protos:
	protoc -I protos/ protos/inventory.proto --go-grpc_out=protos/inventory --go_out=protos/inventory --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	protoc -I protos/ protos/user.proto --go-grpc_out=protos/user --go_out=protos/user --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	protoc -I protos/ protos/cost.proto --go-grpc_out=protos/cost --go_out=protos/cost --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	protoc -I protos/ protos/order.proto --go-grpc_out=protos/order --go_out=protos/order --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
	#protoc -I protos/ protos/item.proto --go-grpc_out=protos/item --go_out=protos/item --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false
