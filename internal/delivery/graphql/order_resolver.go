package graphql

import (
	"context"
	"net/http"

	"github.com/agnaldojpereira/list-order/internal/usecase"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Resolver struct {
	listOrdersUseCase usecase.ListOrdersUseCase
}

func NewResolver(listOrdersUseCase usecase.ListOrdersUseCase) *Resolver {
	return &Resolver{listOrdersUseCase: listOrdersUseCase}
}

type orderResolver struct {
	o *usecase.Order
}

func (r *orderResolver) ID() graphql.ID {
	return graphql.ID(r.o.ID)
}

func (r *orderResolver) UserID() int32 {
	return int32(r.o.UserID)
}

func (r *orderResolver) Amount() float64 {
	return r.o.Amount
}

func (r *orderResolver) Status() string {
	return r.o.Status
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*orderResolver, error) {
	orders, err := r.listOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var resolvers []*orderResolver
	for _, o := range orders {
		resolvers = append(resolvers, &orderResolver{o: o})
	}
	return resolvers, nil
}

func NewHandler(listOrdersUseCase usecase.ListOrdersUseCase) http.Handler {
	schema := graphql.MustParseSchema(Schema, NewResolver(listOrdersUseCase))
	return &relay.Handler{Schema: schema}
}
