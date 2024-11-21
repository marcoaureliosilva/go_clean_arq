package order

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type OrderResolver struct {
	db *gorm.DB
}

// NewOrderResolver cria um novo resolver para GraphQL
func NewOrderResolver(db *gorm.DB) *OrderResolver {
	return &OrderResolver{db: db}
}

// ListOrders retorna todas as ordens
func (r *OrderResolver) ListOrders() ([]Order, error) {
	var orders []Order
	r.db.Find(&orders)
	return orders, nil
}

// CreateOrderType define o tipo de dados Order para GraphQL
func CreateOrderType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Uint},
			"item":     &graphql.Field{Type: graphql.String},
			"quantity": &graphql.Field{Type: graphql.Int},
		},
	})
}
