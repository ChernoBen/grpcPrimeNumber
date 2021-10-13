# Prime number decomposition Grpc
Breve estudo sobre funcionamento do server streaming
## Iniciar aplicação
Execute : ./server.sh para iniciar o servidor
Execute : ./client.sh para iniciar a funcionalidade
#   Configurando ambiente
- 
    - go mod download
    - pamac install protobuf || sudo pacman -S protobuf
    - more information about : https://grpc.io/docs/languages/go/quickstart/
    - protoc src/prime_proto/prime.proto --go_out=. --go-grpc_out=.
    - chmod +x arquivos.sh