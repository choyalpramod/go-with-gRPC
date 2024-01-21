Run this command to generate the greet.pb.go and greet_grpc.pb.go files:

`protoc --go_out=. --go-grpc_out=. proto/greet.proto`