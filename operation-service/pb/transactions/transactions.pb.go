// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: pb/transactions.proto

package transactions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentStatus int32

const (
	PaymentStatus_PENDING   PaymentStatus = 0
	PaymentStatus_COMPLETED PaymentStatus = 1
	PaymentStatus_FAILED    PaymentStatus = 2
	PaymentStatus_REFUNDED  PaymentStatus = 3
	PaymentStatus_CANCELLED PaymentStatus = 4
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PENDING",
		1: "COMPLETED",
		2: "FAILED",
		3: "REFUNDED",
		4: "CANCELLED",
	}
	PaymentStatus_value = map[string]int32{
		"PENDING":   0,
		"COMPLETED": 1,
		"FAILED":    2,
		"REFUNDED":  3,
		"CANCELLED": 4,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_transactions_proto_enumTypes[0].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_pb_transactions_proto_enumTypes[0]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{0}
}

type Transaction struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId         string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentMethod   string                 `protobuf:"bytes,3,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	GrandTotal      float64                `protobuf:"fixed64,4,opt,name=grand_total,json=grandTotal,proto3" json:"grand_total,omitempty"`
	PaymentStatus   string                 `protobuf:"bytes,5,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	TransactionDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_pb_transactions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Transaction) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *Transaction) GetGrandTotal() float64 {
	if x != nil {
		return x.GrandTotal
	}
	return 0
}

func (x *Transaction) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *Transaction) GetTransactionDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionDate
	}
	return nil
}

func (x *Transaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Transaction) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateTransactionRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	OrderId         string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentMethod   string                 `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	GrandTotal      float64                `protobuf:"fixed64,3,opt,name=grand_total,json=grandTotal,proto3" json:"grand_total,omitempty"`
	PaymentStatus   *string                `protobuf:"bytes,4,opt,name=payment_status,json=paymentStatus,proto3,oneof" json:"payment_status,omitempty"`
	TransactionDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=transaction_date,json=transactionDate,proto3,oneof" json:"transaction_date,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	mi := &file_pb_transactions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransactionRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateTransactionRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *CreateTransactionRequest) GetGrandTotal() float64 {
	if x != nil {
		return x.GrandTotal
	}
	return 0
}

func (x *CreateTransactionRequest) GetPaymentStatus() string {
	if x != nil && x.PaymentStatus != nil {
		return *x.PaymentStatus
	}
	return ""
}

func (x *CreateTransactionRequest) GetTransactionDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionDate
	}
	return nil
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	mi := &file_pb_transactions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTransactionsByOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionsByOrderRequest) Reset() {
	*x = GetTransactionsByOrderRequest{}
	mi := &file_pb_transactions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionsByOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsByOrderRequest) ProtoMessage() {}

func (x *GetTransactionsByOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsByOrderRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsByOrderRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionsByOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type UpdateTransactionStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentStatus string                 `protobuf:"bytes,2,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTransactionStatusRequest) Reset() {
	*x = UpdateTransactionStatusRequest{}
	mi := &file_pb_transactions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTransactionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionStatusRequest) ProtoMessage() {}

func (x *UpdateTransactionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateTransactionStatusRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTransactionStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTransactionStatusRequest) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type DeleteTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTransactionRequest) Reset() {
	*x = DeleteTransactionRequest{}
	mi := &file_pb_transactions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransactionRequest) ProtoMessage() {}

func (x *DeleteTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransactionRequest.ProtoReflect.Descriptor instead.
func (*DeleteTransactionRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTransactionResponse) Reset() {
	*x = DeleteTransactionResponse{}
	mi := &file_pb_transactions_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransactionResponse) ProtoMessage() {}

func (x *DeleteTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransactionResponse.ProtoReflect.Descriptor instead.
func (*DeleteTransactionResponse) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListTransactionsByDateRangeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartDate     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTransactionsByDateRangeRequest) Reset() {
	*x = ListTransactionsByDateRangeRequest{}
	mi := &file_pb_transactions_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTransactionsByDateRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsByDateRangeRequest) ProtoMessage() {}

func (x *ListTransactionsByDateRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsByDateRangeRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionsByDateRangeRequest) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{7}
}

func (x *ListTransactionsByDateRangeRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ListTransactionsByDateRangeRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type TransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transaction   *Transaction           `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	mi := &file_pb_transactions_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type TransactionListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transactions  []*Transaction         `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionListResponse) Reset() {
	*x = TransactionListResponse{}
	mi := &file_pb_transactions_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionListResponse) ProtoMessage() {}

func (x *TransactionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transactions_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionListResponse.ProtoReflect.Descriptor instead.
func (*TransactionListResponse) Descriptor() ([]byte, []int) {
	return file_pb_transactions_proto_rawDescGZIP(), []int{9}
}

func (x *TransactionListResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *TransactionListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_pb_transactions_proto protoreflect.FileDescriptor

var file_pb_transactions_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x02, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x9d, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2a,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x96, 0x01, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7c, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a,
	0x54, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x46, 0x55,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xaa, 0x05, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70,
	0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70,
	0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x72, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x70,
	0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70,
	0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e,
	0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pb_transactions_proto_rawDescOnce sync.Once
	file_pb_transactions_proto_rawDescData []byte
)

func file_pb_transactions_proto_rawDescGZIP() []byte {
	file_pb_transactions_proto_rawDescOnce.Do(func() {
		file_pb_transactions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_transactions_proto_rawDesc), len(file_pb_transactions_proto_rawDesc)))
	})
	return file_pb_transactions_proto_rawDescData
}

var file_pb_transactions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_transactions_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_transactions_proto_goTypes = []any{
	(PaymentStatus)(0),                         // 0: pb.transactions.PaymentStatus
	(*Transaction)(nil),                        // 1: pb.transactions.Transaction
	(*CreateTransactionRequest)(nil),           // 2: pb.transactions.CreateTransactionRequest
	(*GetTransactionRequest)(nil),              // 3: pb.transactions.GetTransactionRequest
	(*GetTransactionsByOrderRequest)(nil),      // 4: pb.transactions.GetTransactionsByOrderRequest
	(*UpdateTransactionStatusRequest)(nil),     // 5: pb.transactions.UpdateTransactionStatusRequest
	(*DeleteTransactionRequest)(nil),           // 6: pb.transactions.DeleteTransactionRequest
	(*DeleteTransactionResponse)(nil),          // 7: pb.transactions.DeleteTransactionResponse
	(*ListTransactionsByDateRangeRequest)(nil), // 8: pb.transactions.ListTransactionsByDateRangeRequest
	(*TransactionResponse)(nil),                // 9: pb.transactions.TransactionResponse
	(*TransactionListResponse)(nil),            // 10: pb.transactions.TransactionListResponse
	(*timestamppb.Timestamp)(nil),              // 11: google.protobuf.Timestamp
}
var file_pb_transactions_proto_depIdxs = []int32{
	11, // 0: pb.transactions.Transaction.transaction_date:type_name -> google.protobuf.Timestamp
	11, // 1: pb.transactions.Transaction.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: pb.transactions.Transaction.updated_at:type_name -> google.protobuf.Timestamp
	11, // 3: pb.transactions.CreateTransactionRequest.transaction_date:type_name -> google.protobuf.Timestamp
	11, // 4: pb.transactions.ListTransactionsByDateRangeRequest.start_date:type_name -> google.protobuf.Timestamp
	11, // 5: pb.transactions.ListTransactionsByDateRangeRequest.end_date:type_name -> google.protobuf.Timestamp
	1,  // 6: pb.transactions.TransactionResponse.transaction:type_name -> pb.transactions.Transaction
	1,  // 7: pb.transactions.TransactionListResponse.transactions:type_name -> pb.transactions.Transaction
	2,  // 8: pb.transactions.TransactionService.CreateTransaction:input_type -> pb.transactions.CreateTransactionRequest
	3,  // 9: pb.transactions.TransactionService.GetTransaction:input_type -> pb.transactions.GetTransactionRequest
	4,  // 10: pb.transactions.TransactionService.GetTransactionsByOrder:input_type -> pb.transactions.GetTransactionsByOrderRequest
	5,  // 11: pb.transactions.TransactionService.UpdateTransactionStatus:input_type -> pb.transactions.UpdateTransactionStatusRequest
	6,  // 12: pb.transactions.TransactionService.DeleteTransaction:input_type -> pb.transactions.DeleteTransactionRequest
	8,  // 13: pb.transactions.TransactionService.ListTransactionsByDateRange:input_type -> pb.transactions.ListTransactionsByDateRangeRequest
	9,  // 14: pb.transactions.TransactionService.CreateTransaction:output_type -> pb.transactions.TransactionResponse
	9,  // 15: pb.transactions.TransactionService.GetTransaction:output_type -> pb.transactions.TransactionResponse
	10, // 16: pb.transactions.TransactionService.GetTransactionsByOrder:output_type -> pb.transactions.TransactionListResponse
	9,  // 17: pb.transactions.TransactionService.UpdateTransactionStatus:output_type -> pb.transactions.TransactionResponse
	7,  // 18: pb.transactions.TransactionService.DeleteTransaction:output_type -> pb.transactions.DeleteTransactionResponse
	10, // 19: pb.transactions.TransactionService.ListTransactionsByDateRange:output_type -> pb.transactions.TransactionListResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pb_transactions_proto_init() }
func file_pb_transactions_proto_init() {
	if File_pb_transactions_proto != nil {
		return
	}
	file_pb_transactions_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_transactions_proto_rawDesc), len(file_pb_transactions_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_transactions_proto_goTypes,
		DependencyIndexes: file_pb_transactions_proto_depIdxs,
		EnumInfos:         file_pb_transactions_proto_enumTypes,
		MessageInfos:      file_pb_transactions_proto_msgTypes,
	}.Build()
	File_pb_transactions_proto = out.File
	file_pb_transactions_proto_goTypes = nil
	file_pb_transactions_proto_depIdxs = nil
}
