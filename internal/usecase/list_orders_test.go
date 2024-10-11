package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/agnaldojpereira/list-order/internal/domain"
	"github.com/agnaldojpereira/list-order/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) ListOrders(ctx context.Context) ([]domain.Order, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Order), args.Error(1)
}

func TestListOrdersUseCase_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockOrderRepository)
		mockOrders := []domain.Order{
			{ID: 1, UserID: 1, Amount: 100.0, Status: "pending"},
			{ID: 2, UserID: 2, Amount: 200.0, Status: "completed"},
		}
		mockRepo.On("ListOrders", mock.Anything).Return(mockOrders, nil)

		useCase := usecase.NewListOrdersUseCase(mockRepo)
		orders, err := useCase.Execute(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, mockOrders, orders)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockOrderRepository)
		mockRepo.On("ListOrders", mock.Anything).Return([]domain.Order{}, errors.New("database error"))

		useCase := usecase.NewListOrdersUseCase(mockRepo)
		orders, err := useCase.Execute(context.Background())

		assert.Error(t, err)
		assert.Empty(t, orders)
		mockRepo.AssertExpectations(t)
	})
}
