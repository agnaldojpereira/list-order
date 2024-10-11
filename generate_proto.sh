#!/bin/bash

# Certifique-se de que o diretório proto existe
mkdir -p proto

# Gere os arquivos Go a partir do .proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/order.proto

echo "Arquivos Go gerados com sucesso a partir do proto!"