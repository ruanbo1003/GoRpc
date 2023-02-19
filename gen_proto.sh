
protoc --go_out=libs/rpc/src --go-grpc_out=require_unimplemented_servers=false:libs/rpc/src libs/rpc/proto/*.proto
