package order

import (
	context "context"

	pb "GO_CLEAN_ARQ/proto" // ajuste de caminho para seu prototipo gRPC

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

// NewOrderService cria um novo serviço de ordens
func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

// ListOrders implementa o método para listar ordens via gRPC
func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	var orders []Order
	s.db.Find(&orders)

	response := &pb.ListOrdersResponse{}
	for _, order := range orders {
		response.Orders = append(response.Orders, &pb.Order{
			Id:       uint64(order.ID),
			Item:     order.Item,
			Quantity: int32(order.Quantity),
		})
	}
	return response, nil
}

// RegisterOrderService registra o serviço gRPC
func RegisterOrderService(s *grpc.Server, service *OrderService) {
	pb.RegisterOrderServiceServer(s, service)
}
