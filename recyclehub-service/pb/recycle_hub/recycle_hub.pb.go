// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: pb/recycle_hub.proto

package recycle_hub

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CreateRecycleHubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressId     string                 `protobuf:"bytes,3,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	WasteTypeId   string                 `protobuf:"bytes,4,opt,name=waste_type_id,json=wasteTypeId,proto3" json:"waste_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRecycleHubRequest) Reset() {
	*x = CreateRecycleHubRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecycleHubRequest) ProtoMessage() {}

func (x *CreateRecycleHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*CreateRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRecycleHubRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRecycleHubRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateRecycleHubRequest) GetAddressId() string {
	if x != nil {
		return x.AddressId
	}
	return ""
}

func (x *CreateRecycleHubRequest) GetWasteTypeId() string {
	if x != nil {
		return x.WasteTypeId
	}
	return ""
}

type GetRecycleHubByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecycleHubByIDRequest) Reset() {
	*x = GetRecycleHubByIDRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecycleHubByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecycleHubByIDRequest) ProtoMessage() {}

func (x *GetRecycleHubByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecycleHubByIDRequest.ProtoReflect.Descriptor instead.
func (*GetRecycleHubByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecycleHubByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRecycleHubsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecycleHubsRequest) Reset() {
	*x = GetRecycleHubsRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecycleHubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecycleHubsRequest) ProtoMessage() {}

func (x *GetRecycleHubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecycleHubsRequest.ProtoReflect.Descriptor instead.
func (*GetRecycleHubsRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{2}
}

type GetRecycleHubsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecycleHubs   []*RecycleHubResponse  `protobuf:"bytes,1,rep,name=recycle_hubs,json=recycleHubs,proto3" json:"recycle_hubs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecycleHubsResponse) Reset() {
	*x = GetRecycleHubsResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecycleHubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecycleHubsResponse) ProtoMessage() {}

func (x *GetRecycleHubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecycleHubsResponse.ProtoReflect.Descriptor instead.
func (*GetRecycleHubsResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecycleHubsResponse) GetRecycleHubs() []*RecycleHubResponse {
	if x != nil {
		return x.RecycleHubs
	}
	return nil
}

type UpdateRecycleHubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressId     string                 `protobuf:"bytes,4,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	WasteTypeId   string                 `protobuf:"bytes,5,opt,name=waste_type_id,json=wasteTypeId,proto3" json:"waste_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRecycleHubRequest) Reset() {
	*x = UpdateRecycleHubRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecycleHubRequest) ProtoMessage() {}

func (x *UpdateRecycleHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*UpdateRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRecycleHubRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRecycleHubRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRecycleHubRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateRecycleHubRequest) GetAddressId() string {
	if x != nil {
		return x.AddressId
	}
	return ""
}

func (x *UpdateRecycleHubRequest) GetWasteTypeId() string {
	if x != nil {
		return x.WasteTypeId
	}
	return ""
}

type DeleteRecycleHubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRecycleHubRequest) Reset() {
	*x = DeleteRecycleHubRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecycleHubRequest) ProtoMessage() {}

func (x *DeleteRecycleHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*DeleteRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRecycleHubRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RecycleHubResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressId     string                 `protobuf:"bytes,4,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	WasteTypeId   string                 `protobuf:"bytes,5,opt,name=waste_type_id,json=wasteTypeId,proto3" json:"waste_type_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecycleHubResponse) Reset() {
	*x = RecycleHubResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecycleHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecycleHubResponse) ProtoMessage() {}

func (x *RecycleHubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecycleHubResponse.ProtoReflect.Descriptor instead.
func (*RecycleHubResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{6}
}

func (x *RecycleHubResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RecycleHubResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecycleHubResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RecycleHubResponse) GetAddressId() string {
	if x != nil {
		return x.AddressId
	}
	return ""
}

func (x *RecycleHubResponse) GetWasteTypeId() string {
	if x != nil {
		return x.WasteTypeId
	}
	return ""
}

func (x *RecycleHubResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RecycleHubResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteRecycleHubResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRecycleHubResponse) Reset() {
	*x = DeleteRecycleHubResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecycleHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecycleHubResponse) ProtoMessage() {}

func (x *DeleteRecycleHubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecycleHubResponse.ProtoReflect.Descriptor instead.
func (*DeleteRecycleHubResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRecycleHubResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteRecycleHubResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetWasteTypeByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWasteTypeByIDRequest) Reset() {
	*x = GetWasteTypeByIDRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWasteTypeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWasteTypeByIDRequest) ProtoMessage() {}

func (x *GetWasteTypeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWasteTypeByIDRequest.ProtoReflect.Descriptor instead.
func (*GetWasteTypeByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{8}
}

func (x *GetWasteTypeByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWasteTypesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWasteTypesRequest) Reset() {
	*x = GetWasteTypesRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWasteTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWasteTypesRequest) ProtoMessage() {}

func (x *GetWasteTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWasteTypesRequest.ProtoReflect.Descriptor instead.
func (*GetWasteTypesRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{9}
}

type GetWasteTypesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WasteTypes    []*WasteTypeResponse   `protobuf:"bytes,1,rep,name=waste_types,json=wasteTypes,proto3" json:"waste_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWasteTypesResponse) Reset() {
	*x = GetWasteTypesResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWasteTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWasteTypesResponse) ProtoMessage() {}

func (x *GetWasteTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWasteTypesResponse.ProtoReflect.Descriptor instead.
func (*GetWasteTypesResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{10}
}

func (x *GetWasteTypesResponse) GetWasteTypes() []*WasteTypeResponse {
	if x != nil {
		return x.WasteTypes
	}
	return nil
}

type WasteTypeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WasteTypeResponse) Reset() {
	*x = WasteTypeResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WasteTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasteTypeResponse) ProtoMessage() {}

func (x *WasteTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_recycle_hub_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasteTypeResponse.ProtoReflect.Descriptor instead.
func (*WasteTypeResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{11}
}

func (x *WasteTypeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WasteTypeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WasteTypeResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *WasteTypeResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WasteTypeResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_pb_recycle_hub_proto protoreflect.FileDescriptor

var file_pb_recycle_hub_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f,
	0x68, 0x75, 0x62, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48,
	0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x73, 0x22,
	0x96, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x73,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48,
	0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x73, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57,
	0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x5f, 0x68, 0x75, 0x62, 0x2e, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x77, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x32, 0xef, 0x03, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x73, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x48, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x5f, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x48, 0x75, 0x62, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0xcb, 0x01, 0x0a, 0x10, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x61,
	0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x72, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e,
	0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x61, 0x73,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68,
	0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pb_recycle_hub_proto_rawDescOnce sync.Once
	file_pb_recycle_hub_proto_rawDescData []byte
)

func file_pb_recycle_hub_proto_rawDescGZIP() []byte {
	file_pb_recycle_hub_proto_rawDescOnce.Do(func() {
		file_pb_recycle_hub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_recycle_hub_proto_rawDesc), len(file_pb_recycle_hub_proto_rawDesc)))
	})
	return file_pb_recycle_hub_proto_rawDescData
}

var file_pb_recycle_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pb_recycle_hub_proto_goTypes = []any{
	(*CreateRecycleHubRequest)(nil),  // 0: recycle_hub.CreateRecycleHubRequest
	(*GetRecycleHubByIDRequest)(nil), // 1: recycle_hub.GetRecycleHubByIDRequest
	(*GetRecycleHubsRequest)(nil),    // 2: recycle_hub.GetRecycleHubsRequest
	(*GetRecycleHubsResponse)(nil),   // 3: recycle_hub.GetRecycleHubsResponse
	(*UpdateRecycleHubRequest)(nil),  // 4: recycle_hub.UpdateRecycleHubRequest
	(*DeleteRecycleHubRequest)(nil),  // 5: recycle_hub.DeleteRecycleHubRequest
	(*RecycleHubResponse)(nil),       // 6: recycle_hub.RecycleHubResponse
	(*DeleteRecycleHubResponse)(nil), // 7: recycle_hub.DeleteRecycleHubResponse
	(*GetWasteTypeByIDRequest)(nil),  // 8: recycle_hub.GetWasteTypeByIDRequest
	(*GetWasteTypesRequest)(nil),     // 9: recycle_hub.GetWasteTypesRequest
	(*GetWasteTypesResponse)(nil),    // 10: recycle_hub.GetWasteTypesResponse
	(*WasteTypeResponse)(nil),        // 11: recycle_hub.WasteTypeResponse
}
var file_pb_recycle_hub_proto_depIdxs = []int32{
	6,  // 0: recycle_hub.GetRecycleHubsResponse.recycle_hubs:type_name -> recycle_hub.RecycleHubResponse
	11, // 1: recycle_hub.GetWasteTypesResponse.waste_types:type_name -> recycle_hub.WasteTypeResponse
	0,  // 2: recycle_hub.RecycleHubService.CreateRecycleHub:input_type -> recycle_hub.CreateRecycleHubRequest
	1,  // 3: recycle_hub.RecycleHubService.GetRecycleHubByID:input_type -> recycle_hub.GetRecycleHubByIDRequest
	2,  // 4: recycle_hub.RecycleHubService.GetAllRecycleHubs:input_type -> recycle_hub.GetRecycleHubsRequest
	4,  // 5: recycle_hub.RecycleHubService.UpdateRecycleHub:input_type -> recycle_hub.UpdateRecycleHubRequest
	5,  // 6: recycle_hub.RecycleHubService.DeleteRecycleHub:input_type -> recycle_hub.DeleteRecycleHubRequest
	8,  // 7: recycle_hub.WasteTypeService.GetWasteTypeByID:input_type -> recycle_hub.GetWasteTypeByIDRequest
	9,  // 8: recycle_hub.WasteTypeService.GetAllWasteTypes:input_type -> recycle_hub.GetWasteTypesRequest
	6,  // 9: recycle_hub.RecycleHubService.CreateRecycleHub:output_type -> recycle_hub.RecycleHubResponse
	6,  // 10: recycle_hub.RecycleHubService.GetRecycleHubByID:output_type -> recycle_hub.RecycleHubResponse
	3,  // 11: recycle_hub.RecycleHubService.GetAllRecycleHubs:output_type -> recycle_hub.GetRecycleHubsResponse
	6,  // 12: recycle_hub.RecycleHubService.UpdateRecycleHub:output_type -> recycle_hub.RecycleHubResponse
	7,  // 13: recycle_hub.RecycleHubService.DeleteRecycleHub:output_type -> recycle_hub.DeleteRecycleHubResponse
	11, // 14: recycle_hub.WasteTypeService.GetWasteTypeByID:output_type -> recycle_hub.WasteTypeResponse
	10, // 15: recycle_hub.WasteTypeService.GetAllWasteTypes:output_type -> recycle_hub.GetWasteTypesResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pb_recycle_hub_proto_init() }
func file_pb_recycle_hub_proto_init() {
	if File_pb_recycle_hub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_recycle_hub_proto_rawDesc), len(file_pb_recycle_hub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pb_recycle_hub_proto_goTypes,
		DependencyIndexes: file_pb_recycle_hub_proto_depIdxs,
		MessageInfos:      file_pb_recycle_hub_proto_msgTypes,
	}.Build()
	File_pb_recycle_hub_proto = out.File
	file_pb_recycle_hub_proto_goTypes = nil
	file_pb_recycle_hub_proto_depIdxs = nil
}
