package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/agnaldojpereira/list-order/internal/delivery/grpc"
	"github.com/agnaldojpereira/list-order/internal/delivery/http"
	"github.com/agnaldojpereira/list-order/internal/repository"
	"github.com/agnaldojpereira/list-order/internal/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Configuração da conexão com o banco de dados
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Conecta ao banco de dados
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Verifica a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	// Inicializa o repositório
	repo := repository.NewOrderRepository(db)

	// Inicializa o caso de uso
	listOrdersUseCase := usecase.NewListOrdersUseCase(repo)

	// Inicializa o manipulador HTTP
	httpHandler := http.NewOrderHandler(listOrdersUseCase)

	// Inicializa o servidor gRPC
	grpcServer := grpc.NewOrderServer(listOrdersUseCase)

	// Inicializa o resolvedor GraphQL
	graphqlHandler := graphql.NewHandler(listOrdersUseCase)

	// Inicia o servidor HTTP
	go func() {
		log.Println("Iniciando servidor HTTP na porta 8080")
		http.Handle("/order", httpHandler)
		http.Handle("/graphql", graphqlHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Inicia o servidor gRPC
	go func() {
		log.Println("Iniciando servidor gRPC na porta 50051")
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
		}
		log.Fatal(grpcServer.Serve(listener))
	}()

	// Mantém a aplicação em execução
	select {}
}
