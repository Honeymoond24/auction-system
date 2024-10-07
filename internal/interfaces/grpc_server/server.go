package grpc_server

import (
	"context"
	"github.com/Honeymoond24/auction-system/internal/application"
	"github.com/Honeymoond24/auction-system/internal/interfaces/grpc_server/pb"
)

type AuctionServer struct {
	service *application.AuctionService
	pb.UnimplementedAuctionServiceServer
}

func NewGRPCAuctionServer(service *application.AuctionService) *AuctionServer {
	return &AuctionServer{service: service}
}

func (s *AuctionServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id, err := s.service.CreateUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Id: id}, nil
}

func (s *AuctionServer) CreateLot(ctx context.Context, req *pb.CreateLotRequest) (*pb.CreateLotResponse, error) {
	id, err := s.service.CreateLot(ctx, req.SellerId, req.Title, req.Description, req.StartingPrice, req.MinBidStep, req.AuctionStartTime, req.AuctionEndTime)
	if err != nil {
		return nil, err
	}

	return &pb.CreateLotResponse{Id: id}, nil
}

func (s *AuctionServer) PlaceBid(ctx context.Context, req *pb.PlaceBidRequest) (*pb.PlaceBidResponse, error) {
	err := s.service.PlaceBid(ctx, req.AuctionId, req.BidderId, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.PlaceBidResponse{Success: true}, nil
}

func (s *AuctionServer) CloseAuction(ctx context.Context, req *pb.CloseAuctionRequest) (*pb.CloseAuctionResponse, error) {
	winnerID, finalPrice, err := s.service.CloseAuction(ctx, req.AuctionId)
	if err != nil {
		return nil, err
	}

	return &pb.CloseAuctionResponse{WinnerId: winnerID, FinalPrice: finalPrice}, nil
}

func (s *AuctionServer) ReplenishBalance(ctx context.Context, req *pb.ReplenishBalanceRequest) (*pb.ReplenishBalanceResponse, error) {
	err := s.service.ReplenishBalance(ctx, req.UserId, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.ReplenishBalanceResponse{Success: true}, nil
}
