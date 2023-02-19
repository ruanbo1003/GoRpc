
### install go plugins for protocol compiler
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### build and run
* protoc build: ./gen_proto.sh
* build execute binary: make
* run server: ./bin/server
* run client: ./bin/client

