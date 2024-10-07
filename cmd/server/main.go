package main

import (
	"context"
	"github.com/Honeymoond24/auction-system/internal/application"
	"github.com/Honeymoond24/auction-system/internal/config"
	"github.com/Honeymoond24/auction-system/internal/infrastructure/database"
	"github.com/Honeymoond24/auction-system/internal/interfaces/grpc_server"
	"github.com/Honeymoond24/auction-system/internal/interfaces/grpc_server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := config.NewConfig()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := database.NewDB(cfg.DatabaseURI)

	auctionRepo := database.NewAuctionRepository(db)
	userRepo := database.NewUserRepository(db)
	auctionService := application.NewAuctionService(auctionRepo, userRepo)
	grpcServer := grpc.NewServer()

	pb.RegisterAuctionServiceServer(grpcServer, grpc_server.NewGRPCAuctionServer(auctionService))
	reflection.Register(grpcServer)

	go func() {
		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	log.Println("Starting gRPC REST Gateway server on port 8080...")
	if err := grpc_server.RunRESTGateway(ctx, "localhost:50051"); err != nil {
		log.Fatalf("failed to start REST gateway: %v", err)
	}
}
