package domain

import (
	"errors"
	"time"
)

type Auction struct {
	ID        int64
	Lot       *Lot
	Bids      []*Bid
	WinnerID  int64
	IsClosed  bool
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Lot struct {
	ID            int64
	SellerID      int64
	Title         string
	Description   string
	StartingPrice float64
	MinBidStep    float64
	CreatedAt     time.Time
}

type Bid struct {
	ID        int64
	AuctionID int64
	BidderID  int64
	Amount    float64
	CreatedAt time.Time
}

func (a *Auction) PlaceBid(bid *Bid) error {
	if a.IsClosed {
		return errors.New("auction is closed")
	}

	if bid.Amount < a.GetHighestBidAmount()+a.Lot.MinBidStep {
		return errors.New("bid amount is too low")
	}

	a.Bids = append(a.Bids, bid)
	return nil
}

func (a *Auction) Close() error {
	if time.Now().Before(a.EndTime) {
		return errors.New("auction is still active")
	}

	a.IsClosed = true
	if len(a.Bids) > 0 {
		highestBid := a.getHighestBid()
		a.WinnerID = highestBid.BidderID
	}

	return nil
}

func (a *Auction) getHighestBid() *Bid {
	var highestBid *Bid
	for _, bid := range a.Bids {
		if highestBid == nil || bid.Amount > highestBid.Amount {
			highestBid = bid
		}
	}
	return highestBid
}

func (a *Auction) GetHighestBidAmount() float64 {
	highestBid := a.getHighestBid()
	if highestBid != nil {
		return highestBid.Amount
	}
	return a.Lot.StartingPrice
}
