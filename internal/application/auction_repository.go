package application

import (
	"context"
	"github.com/Honeymoond24/auction-system/internal/domain"
)

type AuctionRepository interface {
	CreateAuction(ctx context.Context, auction *domain.Auction) (int64, error)
	GetAuctionByID(ctx context.Context, auctionID int64) (*domain.Auction, error)
	PlaceBid(ctx context.Context, bid *domain.Bid) error
	GetHighestBid(ctx context.Context, auctionID int64) (*domain.Bid, error)
}
