package application

import (
	"context"
	"testing"
	"time"

	"github.com/Honeymoond24/auction-system/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MockAuctionRepository is a mock implementation of the AuctionRepository interface
type MockAuctionRepository struct {
	mock.Mock
}

func (m *MockAuctionRepository) PlaceBid(ctx context.Context, bid *domain.Bid) error {
	args := m.Called(ctx, bid)
	return args.Error(0)
}

func (m *MockAuctionRepository) CreateAuction(ctx context.Context, auction *domain.Auction) (int64, error) {
	args := m.Called(ctx, auction)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAuctionRepository) GetAuctionByID(ctx context.Context, auctionID int64) (*domain.Auction, error) {
	args := m.Called(ctx, auctionID)
	return args.Get(0).(*domain.Auction), args.Error(1)
}

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) (int64, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) ReplenishBalanceTx(ctx context.Context, userID int64, amount float64) error {
	args := m.Called(ctx, userID, amount)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewAuctionService(nil, mockUserRepo)

	mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(int64(1), nil)

	userID, err := service.CreateUser(context.Background(), "john_doe")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)
	mockUserRepo.AssertExpectations(t)
}

func TestCreateLot(t *testing.T) {
	mockAuctionRepo := new(MockAuctionRepository)
	service := NewAuctionService(mockAuctionRepo, nil)

	startTime := timestamppb.New(time.Now())
	endTime := timestamppb.New(time.Now().Add(1 * time.Hour))

	mockAuctionRepo.On("CreateAuction", mock.Anything, mock.AnythingOfType("*domain.Auction")).Return(int64(1), nil)

	lotID, err := service.CreateLot(context.Background(), 1, "title", "description", 100.0, 10.0, startTime, endTime)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), lotID)
	mockAuctionRepo.AssertExpectations(t)
}

func TestPlaceBid(t *testing.T) {
	mockAuctionRepo := new(MockAuctionRepository)
	service := NewAuctionService(mockAuctionRepo, nil)

	auction := &domain.Auction{}
	mockAuctionRepo.On("GetAuctionByID", mock.Anything, int64(1)).Return(auction, nil)
	mockAuctionRepo.On("PlaceBid", mock.Anything, mock.AnythingOfType("*domain.Bid")).Return(nil)

	err := service.PlaceBid(context.Background(), 1, 1, 100.0)
	assert.NoError(t, err)
	mockAuctionRepo.AssertExpectations(t)
}

func TestReplenishBalance(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewAuctionService(nil, mockUserRepo)

	mockUserRepo.On("ReplenishBalanceTx", mock.Anything, int64(1), 100.0).Return(nil)

	err := service.ReplenishBalance(context.Background(), 1, 100.0)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestCloseAuction(t *testing.T) {
	mockAuctionRepo := new(MockAuctionRepository)
	service := NewAuctionService(mockAuctionRepo, nil)

	auction := &domain.Auction{WinnerID: 1}
	mockAuctionRepo.On("GetAuctionByID", mock.Anything, int64(1)).Return(auction, nil)
	mockAuctionRepo.On("Close", mock.Anything).Return(nil)

	winnerID, highestBidAmount, err := service.CloseAuction(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), winnerID)
	assert.Equal(t, 100.0, highestBidAmount)
	mockAuctionRepo.AssertExpectations(t)
}
