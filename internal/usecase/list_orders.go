package usecase

import (
	"context"

	"github.com/agnaldojpereira/list-order/internal/domain"
	"github.com/agnaldojpereira/list-order/internal/repository"
)

// ListOrdersUseCase define a interface para o caso de uso de listagem de pedidos
type ListOrdersUseCase interface {
	Execute(ctx context.Context) ([]domain.Order, error)
}

// listOrdersUseCase é a implementação concreta do caso de uso ListOrdersUseCase
type listOrdersUseCase struct {
	repo repository.OrderRepository
}

// NewListOrdersUseCase cria uma nova instância de ListOrdersUseCase
func NewListOrdersUseCase(repo repository.OrderRepository) ListOrdersUseCase {
	return &listOrdersUseCase{repo: repo}
}

// Execute realiza a operação de listagem de pedidos
func (uc *listOrdersUseCase) Execute(ctx context.Context) ([]domain.Order, error) {
	return uc.repo.ListOrders(ctx)
}
