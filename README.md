# Serviço de Pedidos

Este projeto implementa um serviço de listagem de pedidos usando Go e Arquitetura Limpa. Ele fornece endpoints REST, gRPC e GraphQL para listar pedidos.

## Pré-requisitos

- Docker
- Docker Compose

## Começando

1. Clone o repositório:
   ```
   git clone https://github.com/agnaldojpereira/list-order.git
   cd order-service
   ```

2. Inicie a aplicação:
   ```
   docker-compose up
   ```

3. Os serviços estarão disponíveis em:
   - API REST: http://localhost:8080
   - gRPC: localhost:50051
   - GraphQL: http://localhost:8080/graphql (TODO: Implementar)

## Endpoints da API

### REST

- Listar Pedidos: GET /order

### gRPC

- ListOrders: Implementado no OrderServer

### GraphQL

- Query: listOrders (TODO: Implementar)

## Testando

Use o arquivo `api.http` para testar os endpoints REST. Você pode usar ferramentas como [Postman](https://www.postman.com/) ou [curl](https://curl.se/) para testar os endpoints gRPC e GraphQL.

## Estrutura do Projeto

O projeto segue os princípios da Arquitetura Limpa:

```
.
├── cmd
│   └── main.go                 # Ponto de entrada da aplicação
├── internal
│   ├── domain
│   │   └── order.go            # Define a entidade Order
│   ├── usecase
│   │   └── list_orders.go      # Implementa o caso de uso de listagem de pedidos
│   ├── repository
│   │   └── order_repository.go # Lida com a persistência de dados
│   └── delivery
│       ├── http
│       │   └── order_handler.go # Manipulador HTTP
│       ├── grpc
│       │   └── order_service.go # Servidor gRPC
│       └── graphql
│           └── order_resolver.go # Resolvedor GraphQL
├── migrations
│   └── 001_create_orders_table.up.sql # Script de migração do banco de dados
├── api.http                    # Arquivo de teste para requisições HTTP
├── Dockerfile                  # Configuração do Docker
├── docker-compose.yaml         # Configuração do Docker Compose
└── README.md                   # Este arquivo
```

## TODO

- Implementar o resolvedor e servidor GraphQL
- Adicionar tratamento de erros mais abrangente
- Implementar a integração com o banco de dados
- Adicionar testes unitários e de integração

## Contribuindo

Contribuições são bem-vindas! Por favor, sinta-se à vontade para enviar um Pull Request.

# list-order
