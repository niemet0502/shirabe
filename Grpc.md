## Grpc commands

1. At the top of the proto file you need:
`option go_package="/currency";`
2. The command for generation can now be something like this:
`protoc -I protos/ --go_out=. --go-grpc_out=. protos/currency.proto`
3. grpcul --plaintext localhost:9002 list
4. grpcul --plaintext localhost:9002 describe service