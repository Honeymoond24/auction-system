// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: internal/interfaces/grpc_server/pb/auction.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateLotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId         int64                  `protobuf:"varint,1,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
	Title            string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description      string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartingPrice    float64                `protobuf:"fixed64,4,opt,name=starting_price,json=startingPrice,proto3" json:"starting_price,omitempty"`
	MinBidStep       float64                `protobuf:"fixed64,5,opt,name=min_bid_step,json=minBidStep,proto3" json:"min_bid_step,omitempty"`
	AuctionStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=auction_start_time,json=auctionStartTime,proto3" json:"auction_start_time,omitempty"`
	AuctionEndTime   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=auction_end_time,json=auctionEndTime,proto3" json:"auction_end_time,omitempty"`
}

func (x *CreateLotRequest) Reset() {
	*x = CreateLotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLotRequest) ProtoMessage() {}

func (x *CreateLotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLotRequest.ProtoReflect.Descriptor instead.
func (*CreateLotRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLotRequest) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *CreateLotRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateLotRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLotRequest) GetStartingPrice() float64 {
	if x != nil {
		return x.StartingPrice
	}
	return 0
}

func (x *CreateLotRequest) GetMinBidStep() float64 {
	if x != nil {
		return x.MinBidStep
	}
	return 0
}

func (x *CreateLotRequest) GetAuctionStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AuctionStartTime
	}
	return nil
}

func (x *CreateLotRequest) GetAuctionEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AuctionEndTime
	}
	return nil
}

type CreateLotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLotResponse) Reset() {
	*x = CreateLotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLotResponse) ProtoMessage() {}

func (x *CreateLotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLotResponse.ProtoReflect.Descriptor instead.
func (*CreateLotResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLotResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PlaceBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId int64   `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	BidderId  int64   `protobuf:"varint,2,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
	Amount    float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PlaceBidRequest) Reset() {
	*x = PlaceBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceBidRequest) ProtoMessage() {}

func (x *PlaceBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceBidRequest.ProtoReflect.Descriptor instead.
func (*PlaceBidRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceBidRequest) GetAuctionId() int64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

func (x *PlaceBidRequest) GetBidderId() int64 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

func (x *PlaceBidRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PlaceBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PlaceBidResponse) Reset() {
	*x = PlaceBidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceBidResponse) ProtoMessage() {}

func (x *PlaceBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceBidResponse.ProtoReflect.Descriptor instead.
func (*PlaceBidResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{5}
}

func (x *PlaceBidResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CloseAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId int64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
}

func (x *CloseAuctionRequest) Reset() {
	*x = CloseAuctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseAuctionRequest) ProtoMessage() {}

func (x *CloseAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseAuctionRequest.ProtoReflect.Descriptor instead.
func (*CloseAuctionRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{6}
}

func (x *CloseAuctionRequest) GetAuctionId() int64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

type CloseAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WinnerId   int64   `protobuf:"varint,1,opt,name=winner_id,json=winnerId,proto3" json:"winner_id,omitempty"`
	FinalPrice float64 `protobuf:"fixed64,2,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *CloseAuctionResponse) Reset() {
	*x = CloseAuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseAuctionResponse) ProtoMessage() {}

func (x *CloseAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseAuctionResponse.ProtoReflect.Descriptor instead.
func (*CloseAuctionResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{7}
}

func (x *CloseAuctionResponse) GetWinnerId() int64 {
	if x != nil {
		return x.WinnerId
	}
	return 0
}

func (x *CloseAuctionResponse) GetFinalPrice() float64 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type GetAuctionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId int64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
}

func (x *GetAuctionStatusRequest) Reset() {
	*x = GetAuctionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionStatusRequest) ProtoMessage() {}

func (x *GetAuctionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAuctionStatusRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{8}
}

func (x *GetAuctionStatusRequest) GetAuctionId() int64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

type GetAuctionStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	WinnerId   int64   `protobuf:"varint,2,opt,name=winner_id,json=winnerId,proto3" json:"winner_id,omitempty"`
	HighestBid float64 `protobuf:"fixed64,3,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
}

func (x *GetAuctionStatusResponse) Reset() {
	*x = GetAuctionStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuctionStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionStatusResponse) ProtoMessage() {}

func (x *GetAuctionStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionStatusResponse.ProtoReflect.Descriptor instead.
func (*GetAuctionStatusResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{9}
}

func (x *GetAuctionStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetAuctionStatusResponse) GetWinnerId() int64 {
	if x != nil {
		return x.WinnerId
	}
	return 0
}

func (x *GetAuctionStatusResponse) GetHighestBid() float64 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type ReplenishBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ReplenishBalanceRequest) Reset() {
	*x = ReplenishBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplenishBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplenishBalanceRequest) ProtoMessage() {}

func (x *ReplenishBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplenishBalanceRequest.ProtoReflect.Descriptor instead.
func (*ReplenishBalanceRequest) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{10}
}

func (x *ReplenishBalanceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReplenishBalanceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ReplenishBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ReplenishBalanceResponse) Reset() {
	*x = ReplenishBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplenishBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplenishBalanceResponse) ProtoMessage() {}

func (x *ReplenishBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplenishBalanceResponse.ProtoReflect.Descriptor instead.
func (*ReplenishBalanceResponse) Descriptor() ([]byte, []int) {
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP(), []int{11}
}

func (x *ReplenishBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_internal_interfaces_grpc_server_pb_auction_proto protoreflect.FileDescriptor

var file_internal_interfaces_grpc_server_pb_auction_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xc0, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x62, 0x69, 0x64, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x42,
	0x69, 0x64, 0x53, 0x74, 0x65, 0x70, 0x12, 0x48, 0x0a, 0x12, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x44, 0x0a, 0x10, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0f, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x34, 0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x77,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x65,
	0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x68, 0x69,
	0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c,
	0x65, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73,
	0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdb, 0x03, 0x0a, 0x0e, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x74, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x42, 0x69, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x6e, 0x65, 0x79, 0x6d, 0x6f, 0x6f, 0x6e,
	0x64, 0x32, 0x34, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_interfaces_grpc_server_pb_auction_proto_rawDescOnce sync.Once
	file_internal_interfaces_grpc_server_pb_auction_proto_rawDescData = file_internal_interfaces_grpc_server_pb_auction_proto_rawDesc
)

func file_internal_interfaces_grpc_server_pb_auction_proto_rawDescGZIP() []byte {
	file_internal_interfaces_grpc_server_pb_auction_proto_rawDescOnce.Do(func() {
		file_internal_interfaces_grpc_server_pb_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_interfaces_grpc_server_pb_auction_proto_rawDescData)
	})
	return file_internal_interfaces_grpc_server_pb_auction_proto_rawDescData
}

var file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_interfaces_grpc_server_pb_auction_proto_goTypes = []any{
	(*CreateUserRequest)(nil),        // 0: auction.CreateUserRequest
	(*CreateUserResponse)(nil),       // 1: auction.CreateUserResponse
	(*CreateLotRequest)(nil),         // 2: auction.CreateLotRequest
	(*CreateLotResponse)(nil),        // 3: auction.CreateLotResponse
	(*PlaceBidRequest)(nil),          // 4: auction.PlaceBidRequest
	(*PlaceBidResponse)(nil),         // 5: auction.PlaceBidResponse
	(*CloseAuctionRequest)(nil),      // 6: auction.CloseAuctionRequest
	(*CloseAuctionResponse)(nil),     // 7: auction.CloseAuctionResponse
	(*GetAuctionStatusRequest)(nil),  // 8: auction.GetAuctionStatusRequest
	(*GetAuctionStatusResponse)(nil), // 9: auction.GetAuctionStatusResponse
	(*ReplenishBalanceRequest)(nil),  // 10: auction.ReplenishBalanceRequest
	(*ReplenishBalanceResponse)(nil), // 11: auction.ReplenishBalanceResponse
	(*timestamppb.Timestamp)(nil),    // 12: google.protobuf.Timestamp
}
var file_internal_interfaces_grpc_server_pb_auction_proto_depIdxs = []int32{
	12, // 0: auction.CreateLotRequest.auction_start_time:type_name -> google.protobuf.Timestamp
	12, // 1: auction.CreateLotRequest.auction_end_time:type_name -> google.protobuf.Timestamp
	0,  // 2: auction.AuctionService.CreateUser:input_type -> auction.CreateUserRequest
	2,  // 3: auction.AuctionService.CreateLot:input_type -> auction.CreateLotRequest
	4,  // 4: auction.AuctionService.PlaceBid:input_type -> auction.PlaceBidRequest
	6,  // 5: auction.AuctionService.CloseAuction:input_type -> auction.CloseAuctionRequest
	8,  // 6: auction.AuctionService.GetAuctionStatus:input_type -> auction.GetAuctionStatusRequest
	10, // 7: auction.AuctionService.ReplenishBalance:input_type -> auction.ReplenishBalanceRequest
	1,  // 8: auction.AuctionService.CreateUser:output_type -> auction.CreateUserResponse
	3,  // 9: auction.AuctionService.CreateLot:output_type -> auction.CreateLotResponse
	5,  // 10: auction.AuctionService.PlaceBid:output_type -> auction.PlaceBidResponse
	7,  // 11: auction.AuctionService.CloseAuction:output_type -> auction.CloseAuctionResponse
	9,  // 12: auction.AuctionService.GetAuctionStatus:output_type -> auction.GetAuctionStatusResponse
	11, // 13: auction.AuctionService.ReplenishBalance:output_type -> auction.ReplenishBalanceResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_internal_interfaces_grpc_server_pb_auction_proto_init() }
func file_internal_interfaces_grpc_server_pb_auction_proto_init() {
	if File_internal_interfaces_grpc_server_pb_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PlaceBidRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PlaceBidResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CloseAuctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CloseAuctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuctionStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuctionStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ReplenishBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ReplenishBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_interfaces_grpc_server_pb_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_interfaces_grpc_server_pb_auction_proto_goTypes,
		DependencyIndexes: file_internal_interfaces_grpc_server_pb_auction_proto_depIdxs,
		MessageInfos:      file_internal_interfaces_grpc_server_pb_auction_proto_msgTypes,
	}.Build()
	File_internal_interfaces_grpc_server_pb_auction_proto = out.File
	file_internal_interfaces_grpc_server_pb_auction_proto_rawDesc = nil
	file_internal_interfaces_grpc_server_pb_auction_proto_goTypes = nil
	file_internal_interfaces_grpc_server_pb_auction_proto_depIdxs = nil
}
