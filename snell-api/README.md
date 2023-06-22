### Generate grpc and pb files

> reference: https://grpc.io/docs/languages/go/quickstart/
Install protoc-gen-go
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

Install protoc-gen-go-grpc
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Generate grpc & pb files
```
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative snellgoapi.proto
```

