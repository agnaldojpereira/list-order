package repository

import (
	"context"
	"database/sql"

	"github.com/agnaldojpereira/list-order/internal/domain"
)

// OrderRepository define a interface para operações de repositório relacionadas a pedidos
type OrderRepository interface {
	ListOrders(ctx context.Context) ([]domain.Order, error)
}

// orderRepository é a implementação concreta de OrderRepository
type orderRepository struct {
	db *sql.DB
}

// NewOrderRepository cria uma nova instância de OrderRepository
func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

// ListOrders recupera todos os pedidos do repositório
func (r *orderRepository) ListOrders(ctx context.Context) ([]domain.Order, error) {
	query := `SELECT id, user_id, amount, status FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Amount, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
