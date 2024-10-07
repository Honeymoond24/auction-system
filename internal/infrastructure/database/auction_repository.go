package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Honeymoond24/auction-system/internal/application"
	"github.com/Honeymoond24/auction-system/internal/domain"
	"github.com/jackc/pgx/v5"
	"log"
)

type AuctionRepository struct {
	db *DBPool
}

func NewAuctionRepository(db *DBPool) application.AuctionRepository {
	return &AuctionRepository{db: db}
}

func (r *AuctionRepository) CreateAuction(ctx context.Context, auction *domain.Auction) (int64, error) {
	if auction == nil || auction.Lot == nil {
		return 0, errors.New("invalid auction data")
	}

	tx, err := r.db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			if errors.Is(err, pgx.ErrTxClosed) {
				return
			}
			log.Println(err)
		}
	}(tx, ctx)

	queryLot := `
		INSERT INTO lots (seller_id, title, description, starting_price, min_bid_step)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var lotID int64
	err = tx.QueryRow(ctx, queryLot,
		auction.Lot.SellerID,
		auction.Lot.Title,
		auction.Lot.Description,
		auction.Lot.StartingPrice,
		auction.Lot.MinBidStep,
	).Scan(&lotID)
	if err != nil {
		return 0, err
	}

	queryAuction := `
		INSERT INTO auctions (lot_id, start_time, end_time, is_closed)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var auctionID int64
	err = tx.QueryRow(ctx, queryAuction,
		lotID,
		auction.StartTime,
		auction.EndTime,
		auction.IsClosed,
	).Scan(&auctionID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return auctionID, nil
}

func (r *AuctionRepository) GetAuctionByID(ctx context.Context, auctionID int64) (*domain.Auction, error) {
	if auctionID <= 0 {
		return nil, errors.New("invalid auction ID")
	}

	query := `
		SELECT a.id, a.start_time, a.end_time, a.winner_id, a.is_closed, a.created_at, a.updated_at,
		       l.id, l.seller_id, l.title, l.description, l.starting_price, l.min_bid_step
		FROM auctions a
		JOIN lots l ON a.lot_id = l.id
		WHERE a.id = $1
	`

	var auction domain.Auction
	var lot domain.Lot
	err := r.db.conn.QueryRow(ctx, query, auctionID).Scan(
		&auction.ID,
		&auction.StartTime,
		&auction.EndTime,
		&auction.WinnerID,
		&auction.IsClosed,
		&auction.CreatedAt,
		&auction.UpdatedAt,
		&lot.ID,
		&lot.SellerID,
		&lot.Title,
		&lot.Description,
		&lot.StartingPrice,
		&lot.MinBidStep,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	auction.Lot = &lot
	return &auction, nil
}

func (r *AuctionRepository) PlaceBid(ctx context.Context, bid *domain.Bid) error {
	if bid == nil {
		return errors.New("invalid bid data")
	}

	auction, err := r.GetAuctionByID(ctx, bid.AuctionID)
	if err != nil {
		return err
	}
	if auction == nil {
		return errors.New("auction not found")
	}

	if auction.IsClosed {
		return errors.New("auction is closed")
	}

	query := `
		INSERT INTO bids (auction_id, bidder_id, amount)
		VALUES ($1, $2, $3)
	`
	_, err = r.db.conn.Exec(ctx, query, bid.AuctionID, bid.BidderID, bid.Amount)

	if err != nil {
		return err
	}
	return nil
}

func (r *AuctionRepository) GetHighestBid(ctx context.Context, auctionID int64) (*domain.Bid, error) {
	if auctionID <= 0 {
		return nil, errors.New("invalid auction ID")
	}

	query := `
		SELECT id, auction_id, bidder_id, amount, created_at
		FROM bids
		WHERE auction_id = $1
		ORDER BY amount DESC
		LIMIT 1
	`

	var bid domain.Bid
	err := r.db.conn.QueryRow(ctx, query, auctionID).Scan(
		&bid.ID,
		&bid.AuctionID,
		&bid.BidderID,
		&bid.Amount,
		&bid.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &bid, nil
}
