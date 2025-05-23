// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: pb/address.proto

package address

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

type CreateAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Latitude      string                 `protobuf:"bytes,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAddressRequest) Reset() {
	*x = CreateAddressRequest{}
	mi := &file_pb_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAddressRequest) ProtoMessage() {}

func (x *CreateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAddressRequest.ProtoReflect.Descriptor instead.
func (*CreateAddressRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAddressRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAddressRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *CreateAddressRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type GetAddressByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAddressByIDRequest) Reset() {
	*x = GetAddressByIDRequest{}
	mi := &file_pb_address_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAddressByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressByIDRequest) ProtoMessage() {}

func (x *GetAddressByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressByIDRequest.ProtoReflect.Descriptor instead.
func (*GetAddressByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{1}
}

func (x *GetAddressByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAddressesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Latitude      string                 `protobuf:"bytes,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAddressesRequest) Reset() {
	*x = GetAddressesRequest{}
	mi := &file_pb_address_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAddressesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressesRequest) ProtoMessage() {}

func (x *GetAddressesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressesRequest.ProtoReflect.Descriptor instead.
func (*GetAddressesRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{2}
}

func (x *GetAddressesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAddressesRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *GetAddressesRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type GetAddressesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []*AddressResponse     `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAddressesResponse) Reset() {
	*x = GetAddressesResponse{}
	mi := &file_pb_address_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAddressesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressesResponse) ProtoMessage() {}

func (x *GetAddressesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressesResponse.ProtoReflect.Descriptor instead.
func (*GetAddressesResponse) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{3}
}

func (x *GetAddressesResponse) GetAddresses() []*AddressResponse {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type UpdateAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Latitude      string                 `protobuf:"bytes,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAddressRequest) Reset() {
	*x = UpdateAddressRequest{}
	mi := &file_pb_address_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddressRequest) ProtoMessage() {}

func (x *UpdateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddressRequest.ProtoReflect.Descriptor instead.
func (*UpdateAddressRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAddressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAddressRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAddressRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *UpdateAddressRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type DeleteAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAddressRequest) Reset() {
	*x = DeleteAddressRequest{}
	mi := &file_pb_address_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressRequest) ProtoMessage() {}

func (x *DeleteAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressRequest.ProtoReflect.Descriptor instead.
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAddressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Latitude      string                 `protobuf:"bytes,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressResponse) Reset() {
	*x = AddressResponse{}
	mi := &file_pb_address_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResponse) ProtoMessage() {}

func (x *AddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResponse.ProtoReflect.Descriptor instead.
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{6}
}

func (x *AddressResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddressResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddressResponse) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *AddressResponse) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *AddressResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AddressResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAddressResponse) Reset() {
	*x = DeleteAddressResponse{}
	mi := &file_pb_address_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressResponse) ProtoMessage() {}

func (x *DeleteAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressResponse.ProtoReflect.Descriptor instead.
func (*DeleteAddressResponse) Descriptor() ([]byte, []int) {
	return file_pb_address_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAddressResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteAddressResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_address_proto protoreflect.FileDescriptor

const file_pb_address_proto_rawDesc = "" +
	"\n" +
	"\x10pb/address.proto\x12\aaddress\"d\n" +
	"\x14CreateAddressRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x03 \x01(\tR\tlongitude\"'\n" +
	"\x15GetAddressByIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"c\n" +
	"\x13GetAddressesRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x03 \x01(\tR\tlongitude\"N\n" +
	"\x14GetAddressesResponse\x126\n" +
	"\taddresses\x18\x01 \x03(\v2\x18.address.AddressResponseR\taddresses\"t\n" +
	"\x14UpdateAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\blatitude\x18\x03 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x04 \x01(\tR\tlongitude\"&\n" +
	"\x14DeleteAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\xad\x01\n" +
	"\x0fAddressResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\blatitude\x18\x03 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x04 \x01(\tR\tlongitude\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\"K\n" +
	"\x15DeleteAddressResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\x9a\x03\n" +
	"\x0eAddressService\x12J\n" +
	"\rCreateAddress\x12\x1d.address.CreateAddressRequest\x1a\x18.address.AddressResponse\"\x00\x12L\n" +
	"\x0eGetAddressByID\x12\x1e.address.GetAddressByIDRequest\x1a\x18.address.AddressResponse\"\x00\x12P\n" +
	"\x0fGetAllAddresses\x12\x1c.address.GetAddressesRequest\x1a\x1d.address.GetAddressesResponse\"\x00\x12J\n" +
	"\rUpdateAddress\x12\x1d.address.UpdateAddressRequest\x1a\x18.address.AddressResponse\"\x00\x12P\n" +
	"\rDeleteAddress\x12\x1d.address.DeleteAddressRequest\x1a\x1e.address.DeleteAddressResponse\"\x00B\fZ\n" +
	"pb/addressb\x06proto3"

var (
	file_pb_address_proto_rawDescOnce sync.Once
	file_pb_address_proto_rawDescData []byte
)

func file_pb_address_proto_rawDescGZIP() []byte {
	file_pb_address_proto_rawDescOnce.Do(func() {
		file_pb_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_address_proto_rawDesc), len(file_pb_address_proto_rawDesc)))
	})
	return file_pb_address_proto_rawDescData
}

var file_pb_address_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_address_proto_goTypes = []any{
	(*CreateAddressRequest)(nil),  // 0: address.CreateAddressRequest
	(*GetAddressByIDRequest)(nil), // 1: address.GetAddressByIDRequest
	(*GetAddressesRequest)(nil),   // 2: address.GetAddressesRequest
	(*GetAddressesResponse)(nil),  // 3: address.GetAddressesResponse
	(*UpdateAddressRequest)(nil),  // 4: address.UpdateAddressRequest
	(*DeleteAddressRequest)(nil),  // 5: address.DeleteAddressRequest
	(*AddressResponse)(nil),       // 6: address.AddressResponse
	(*DeleteAddressResponse)(nil), // 7: address.DeleteAddressResponse
}
var file_pb_address_proto_depIdxs = []int32{
	6, // 0: address.GetAddressesResponse.addresses:type_name -> address.AddressResponse
	0, // 1: address.AddressService.CreateAddress:input_type -> address.CreateAddressRequest
	1, // 2: address.AddressService.GetAddressByID:input_type -> address.GetAddressByIDRequest
	2, // 3: address.AddressService.GetAllAddresses:input_type -> address.GetAddressesRequest
	4, // 4: address.AddressService.UpdateAddress:input_type -> address.UpdateAddressRequest
	5, // 5: address.AddressService.DeleteAddress:input_type -> address.DeleteAddressRequest
	6, // 6: address.AddressService.CreateAddress:output_type -> address.AddressResponse
	6, // 7: address.AddressService.GetAddressByID:output_type -> address.AddressResponse
	3, // 8: address.AddressService.GetAllAddresses:output_type -> address.GetAddressesResponse
	6, // 9: address.AddressService.UpdateAddress:output_type -> address.AddressResponse
	7, // 10: address.AddressService.DeleteAddress:output_type -> address.DeleteAddressResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_address_proto_init() }
func file_pb_address_proto_init() {
	if File_pb_address_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_address_proto_rawDesc), len(file_pb_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_address_proto_goTypes,
		DependencyIndexes: file_pb_address_proto_depIdxs,
		MessageInfos:      file_pb_address_proto_msgTypes,
	}.Build()
	File_pb_address_proto = out.File
	file_pb_address_proto_goTypes = nil
	file_pb_address_proto_depIdxs = nil
}
