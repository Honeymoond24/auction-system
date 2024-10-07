package application

import (
	"context"
	"errors"
	"github.com/Honeymoond24/auction-system/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuctionService struct {
	repo     AuctionRepository
	userRepo UserRepository
}

func NewAuctionService(repo AuctionRepository, userRepo UserRepository) *AuctionService {
	return &AuctionService{repo: repo, userRepo: userRepo}
}

func (s *AuctionService) CreateUser(ctx context.Context, username string) (int64, error) {
	user := &domain.User{
		Username: username,
	}
	return s.userRepo.CreateUser(ctx, user)
}

func (s *AuctionService) CreateLot(
	ctx context.Context, sellerID int64, title, description string,
	startingPrice, minBidStep float64, startTime *timestamppb.Timestamp, endTime *timestamppb.Timestamp,
) (int64, error) {
	lot := &domain.Lot{
		SellerID:      sellerID,
		Title:         title,
		Description:   description,
		StartingPrice: startingPrice,
		MinBidStep:    minBidStep,
	}
	if startTime == nil || endTime == nil || startTime.AsTime().After(endTime.AsTime()) {
		return 0, errors.New("invalid auction time")
	}
	auction := &domain.Auction{
		Lot:       lot,
		StartTime: startTime.AsTime(),
		EndTime:   endTime.AsTime(),
	}
	return s.repo.CreateAuction(ctx, auction)
}

func (s *AuctionService) PlaceBid(ctx context.Context, auctionID, bidderID int64, amount float64) error {
	auction, err := s.repo.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	bid := &domain.Bid{
		AuctionID: auctionID,
		BidderID:  bidderID,
		Amount:    amount,
	}

	return auction.PlaceBid(bid)
}

func (s *AuctionService) ReplenishBalance(ctx context.Context, userID int64, amount float64) error {
	return s.userRepo.ReplenishBalanceTx(ctx, userID, amount)
}

func (s *AuctionService) CloseAuction(ctx context.Context, auctionID int64) (int64, float64, error) {
	auction, err := s.repo.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return 0, 0, err
	}

	err = auction.Close()
	if err != nil {
		return 0, 0, err
	}
	highestBid, err := s.repo.GetHighestBid(ctx, auctionID)
	if err != nil {
		return 0, 0, err
	}
	return auction.WinnerID, highestBid.Amount, nil
}
