# Shirabe 

## Architecture 

<img src="architecture.png" /> <br />

### Packages 
```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generate code

```
    protoc -I protos/ --go_out=. --go-grpc_out=. protos/currency.proto
```
