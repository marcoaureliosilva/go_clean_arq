package main

import (
	"log"
	"net"

	"myapp/internal/order"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// ajuste de caminho para seu prototipo gRPC
)

func main() {
	// Configurar banco de dados
	db, err := gorm.Open(postgres.Open("host=db user=user password=password dbname=orders port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	order.Migrate(db)

	// Configurar servidor REST
	r := gin.Default()
	orderHandler := order.NewHandler(db)
	r.GET("/order", orderHandler.ListOrders)

	// Configurar gRPC
	grpcServer := grpc.NewServer()
	orderService := order.NewOrderService(db)
	order.RegisterOrderService(grpcServer, orderService)

	// Iniciar servidores
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("failed to run REST server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server listening on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
