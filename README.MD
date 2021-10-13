# Prime number decomposition Grpc
#1
- Install protoc
    - pamac install protobuf || sudo pacman -S protobuf
    - more information about : https://grpc.io/docs/languages/go/quickstart/
    - protoc src/prime_proto/prime.proto --go_out=. --go-grpc_out=.
    - chmod +x gen.sh