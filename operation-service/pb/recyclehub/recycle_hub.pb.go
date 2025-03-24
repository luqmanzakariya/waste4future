// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: pb/recycle_hub.proto

package recyclehub

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

// Empty is an empty message for requests/responses that don't need parameters.
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_pb_recycle_hub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{0}
}

// WasteType represents a single waste type.
type WasteType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WasteType) Reset() {
	*x = WasteType{}
	mi := &file_pb_recycle_hub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WasteType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasteType) ProtoMessage() {}

func (x *WasteType) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use WasteType.ProtoReflect.Descriptor instead.
func (*WasteType) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{1}
}

func (x *WasteType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WasteType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WasteType) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// WasteTypeList represents a list of waste types.
type WasteTypeList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WasteTypes    []*WasteType           `protobuf:"bytes,1,rep,name=waste_types,json=wasteTypes,proto3" json:"waste_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WasteTypeList) Reset() {
	*x = WasteTypeList{}
	mi := &file_pb_recycle_hub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WasteTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasteTypeList) ProtoMessage() {}

func (x *WasteTypeList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use WasteTypeList.ProtoReflect.Descriptor instead.
func (*WasteTypeList) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{2}
}

func (x *WasteTypeList) GetWasteTypes() []*WasteType {
	if x != nil {
		return x.WasteTypes
	}
	return nil
}

// GetWasteTypeByIDRequest contains the ID of the waste type to retrieve.
type GetWasteTypeByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // The MongoDB _id field
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWasteTypeByIDRequest) Reset() {
	*x = GetWasteTypeByIDRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWasteTypeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWasteTypeByIDRequest) ProtoMessage() {}

func (x *GetWasteTypeByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetWasteTypeByIDRequest.ProtoReflect.Descriptor instead.
func (*GetWasteTypeByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{3}
}

func (x *GetWasteTypeByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// CreateRecycleHubRequest contains the data needed to create a new recycle hub.
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
	mi := &file_pb_recycle_hub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecycleHubRequest) ProtoMessage() {}

func (x *CreateRecycleHubRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*CreateRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{4}
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

// GetRecycleHubByIDRequest contains the ID of the recycle hub to retrieve.
type GetRecycleHubByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // The MongoDB _id field
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecycleHubByIDRequest) Reset() {
	*x = GetRecycleHubByIDRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecycleHubByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecycleHubByIDRequest) ProtoMessage() {}

func (x *GetRecycleHubByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetRecycleHubByIDRequest.ProtoReflect.Descriptor instead.
func (*GetRecycleHubByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{5}
}

func (x *GetRecycleHubByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// UpdateRecycleHubRequest contains the data needed to update an existing recycle hub.
type UpdateRecycleHubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // The MongoDB _id field
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressId     string                 `protobuf:"bytes,4,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	WasteTypeId   string                 `protobuf:"bytes,5,opt,name=waste_type_id,json=wasteTypeId,proto3" json:"waste_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRecycleHubRequest) Reset() {
	*x = UpdateRecycleHubRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecycleHubRequest) ProtoMessage() {}

func (x *UpdateRecycleHubRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*UpdateRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{6}
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

// DeleteRecycleHubRequest contains the ID of the recycle hub to delete.
type DeleteRecycleHubRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // The MongoDB _id field
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRecycleHubRequest) Reset() {
	*x = DeleteRecycleHubRequest{}
	mi := &file_pb_recycle_hub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecycleHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecycleHubRequest) ProtoMessage() {}

func (x *DeleteRecycleHubRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteRecycleHubRequest.ProtoReflect.Descriptor instead.
func (*DeleteRecycleHubRequest) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRecycleHubRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// RecycleHubResponse represents a single recycle hub.
type RecycleHubResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // The MongoDB _id field
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	AddressId     string                 `protobuf:"bytes,4,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	WasteTypeId   string                 `protobuf:"bytes,5,opt,name=waste_type_id,json=wasteTypeId,proto3" json:"waste_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecycleHubResponse) Reset() {
	*x = RecycleHubResponse{}
	mi := &file_pb_recycle_hub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecycleHubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecycleHubResponse) ProtoMessage() {}

func (x *RecycleHubResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RecycleHubResponse.ProtoReflect.Descriptor instead.
func (*RecycleHubResponse) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{8}
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

// RecycleHubList represents a list of recycle hubs.
type RecycleHubList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecycleHubs   []*RecycleHubResponse  `protobuf:"bytes,1,rep,name=recycle_hubs,json=recycleHubs,proto3" json:"recycle_hubs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecycleHubList) Reset() {
	*x = RecycleHubList{}
	mi := &file_pb_recycle_hub_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecycleHubList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecycleHubList) ProtoMessage() {}

func (x *RecycleHubList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RecycleHubList.ProtoReflect.Descriptor instead.
func (*RecycleHubList) Descriptor() ([]byte, []int) {
	return file_pb_recycle_hub_proto_rawDescGZIP(), []int{9}
}

func (x *RecycleHubList) GetRecycleHubs() []*RecycleHubResponse {
	if x != nil {
		return x.RecycleHubs
	}
	return nil
}

var File_pb_recycle_hub_proto protoreflect.FileDescriptor

var file_pb_recycle_hub_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x68, 0x75, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x45,
	0x0a, 0x09, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x0d, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x57, 0x61, 0x73, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x77, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x86, 0x01, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x73, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x96, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77,
	0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61,
	0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x52, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0c, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75,
	0x62, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62,
	0x73, 0x32, 0xb0, 0x01, 0x0a, 0x10, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62,
	0x2e, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x54,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68,
	0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x73, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x57, 0x61, 0x73, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x32, 0xce, 0x03, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x48, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x12, 0x26,
	0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x73, 0x12, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75,
	0x62, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62,
	0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x68, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pb_recycle_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_recycle_hub_proto_goTypes = []any{
	(*Empty)(nil),                    // 0: pb.recyclehub.Empty
	(*WasteType)(nil),                // 1: pb.recyclehub.WasteType
	(*WasteTypeList)(nil),            // 2: pb.recyclehub.WasteTypeList
	(*GetWasteTypeByIDRequest)(nil),  // 3: pb.recyclehub.GetWasteTypeByIDRequest
	(*CreateRecycleHubRequest)(nil),  // 4: pb.recyclehub.CreateRecycleHubRequest
	(*GetRecycleHubByIDRequest)(nil), // 5: pb.recyclehub.GetRecycleHubByIDRequest
	(*UpdateRecycleHubRequest)(nil),  // 6: pb.recyclehub.UpdateRecycleHubRequest
	(*DeleteRecycleHubRequest)(nil),  // 7: pb.recyclehub.DeleteRecycleHubRequest
	(*RecycleHubResponse)(nil),       // 8: pb.recyclehub.RecycleHubResponse
	(*RecycleHubList)(nil),           // 9: pb.recyclehub.RecycleHubList
}
var file_pb_recycle_hub_proto_depIdxs = []int32{
	1, // 0: pb.recyclehub.WasteTypeList.waste_types:type_name -> pb.recyclehub.WasteType
	8, // 1: pb.recyclehub.RecycleHubList.recycle_hubs:type_name -> pb.recyclehub.RecycleHubResponse
	0, // 2: pb.recyclehub.WasteTypeService.GetAllWasteTypes:input_type -> pb.recyclehub.Empty
	3, // 3: pb.recyclehub.WasteTypeService.GetWasteTypeByID:input_type -> pb.recyclehub.GetWasteTypeByIDRequest
	4, // 4: pb.recyclehub.RecycleHubService.CreateRecycleHub:input_type -> pb.recyclehub.CreateRecycleHubRequest
	0, // 5: pb.recyclehub.RecycleHubService.GetAllRecycleHubs:input_type -> pb.recyclehub.Empty
	5, // 6: pb.recyclehub.RecycleHubService.GetRecycleHubByID:input_type -> pb.recyclehub.GetRecycleHubByIDRequest
	6, // 7: pb.recyclehub.RecycleHubService.UpdateRecycleHub:input_type -> pb.recyclehub.UpdateRecycleHubRequest
	7, // 8: pb.recyclehub.RecycleHubService.DeleteRecycleHub:input_type -> pb.recyclehub.DeleteRecycleHubRequest
	2, // 9: pb.recyclehub.WasteTypeService.GetAllWasteTypes:output_type -> pb.recyclehub.WasteTypeList
	1, // 10: pb.recyclehub.WasteTypeService.GetWasteTypeByID:output_type -> pb.recyclehub.WasteType
	8, // 11: pb.recyclehub.RecycleHubService.CreateRecycleHub:output_type -> pb.recyclehub.RecycleHubResponse
	9, // 12: pb.recyclehub.RecycleHubService.GetAllRecycleHubs:output_type -> pb.recyclehub.RecycleHubList
	8, // 13: pb.recyclehub.RecycleHubService.GetRecycleHubByID:output_type -> pb.recyclehub.RecycleHubResponse
	8, // 14: pb.recyclehub.RecycleHubService.UpdateRecycleHub:output_type -> pb.recyclehub.RecycleHubResponse
	0, // 15: pb.recyclehub.RecycleHubService.DeleteRecycleHub:output_type -> pb.recyclehub.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			NumMessages:   10,
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
