package grpc

import (
	"context"

	"github.com/agnaldojpereira/list-order/internal/usecase"
	pb "github.com/agnaldojpereira/list-order/proto"
	"google.golang.org/grpc"
)

// OrderServer implementa o servidor gRPC para serviços relacionados a pedidos
type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	listOrdersUseCase usecase.ListOrdersUseCase
}

// NewOrderServer cria uma nova instância do servidor gRPC
func NewOrderServer(listOrdersUseCase usecase.ListOrdersUseCase) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &OrderServer{listOrdersUseCase: listOrdersUseCase})
	return server
}

// ListOrders implementa o método RPC para listar pedidos
func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.listOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:     order.ID,
			UserId: order.UserID,
			Amount: order.Amount,
			Status: order.Status,
		})
	}

	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
