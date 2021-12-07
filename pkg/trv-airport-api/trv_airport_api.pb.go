// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ozonmp/trv_airport_api/v1/trv_airport_api.proto

package trv_airport_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Airport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Airport) Reset() {
	*x = Airport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Airport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Airport) ProtoMessage() {}

func (x *Airport) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Airport.ProtoReflect.Descriptor instead.
func (*Airport) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{0}
}

func (x *Airport) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Airport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Airport) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type CreateAirportV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateAirportV1Request) Reset() {
	*x = CreateAirportV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAirportV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAirportV1Request) ProtoMessage() {}

func (x *CreateAirportV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAirportV1Request.ProtoReflect.Descriptor instead.
func (*CreateAirportV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAirportV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAirportV1Request) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type CreateAirportV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateAirportV1Response) Reset() {
	*x = CreateAirportV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAirportV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAirportV1Response) ProtoMessage() {}

func (x *CreateAirportV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAirportV1Response.ProtoReflect.Descriptor instead.
func (*CreateAirportV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAirportV1Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeAirportV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirportId uint64 `protobuf:"varint,1,opt,name=airport_id,json=airportId,proto3" json:"airport_id,omitempty"`
}

func (x *DescribeAirportV1Request) Reset() {
	*x = DescribeAirportV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeAirportV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAirportV1Request) ProtoMessage() {}

func (x *DescribeAirportV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAirportV1Request.ProtoReflect.Descriptor instead.
func (*DescribeAirportV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeAirportV1Request) GetAirportId() uint64 {
	if x != nil {
		return x.AirportId
	}
	return 0
}

type DescribeAirportV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Airport `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribeAirportV1Response) Reset() {
	*x = DescribeAirportV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeAirportV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAirportV1Response) ProtoMessage() {}

func (x *DescribeAirportV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAirportV1Response.ProtoReflect.Descriptor instead.
func (*DescribeAirportV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeAirportV1Response) GetValue() *Airport {
	if x != nil {
		return x.Value
	}
	return nil
}

type ListAirportsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor uint64 `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListAirportsV1Request) Reset() {
	*x = ListAirportsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAirportsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAirportsV1Request) ProtoMessage() {}

func (x *ListAirportsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAirportsV1Request.ProtoReflect.Descriptor instead.
func (*ListAirportsV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListAirportsV1Request) GetCursor() uint64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *ListAirportsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListAirportsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airports []*Airport `protobuf:"bytes,1,rep,name=airports,proto3" json:"airports,omitempty"`
}

func (x *ListAirportsV1Response) Reset() {
	*x = ListAirportsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAirportsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAirportsV1Response) ProtoMessage() {}

func (x *ListAirportsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAirportsV1Response.ProtoReflect.Descriptor instead.
func (*ListAirportsV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListAirportsV1Response) GetAirports() []*Airport {
	if x != nil {
		return x.Airports
	}
	return nil
}

type DeleteAirportV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirportId uint64 `protobuf:"varint,1,opt,name=airport_id,json=airportId,proto3" json:"airport_id,omitempty"`
}

func (x *DeleteAirportV1Request) Reset() {
	*x = DeleteAirportV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAirportV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAirportV1Request) ProtoMessage() {}

func (x *DeleteAirportV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAirportV1Request.ProtoReflect.Descriptor instead.
func (*DeleteAirportV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAirportV1Request) GetAirportId() uint64 {
	if x != nil {
		return x.AirportId
	}
	return 0
}

type DeleteAirportV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirportId uint64 `protobuf:"varint,1,opt,name=airport_id,json=airportId,proto3" json:"airport_id,omitempty"`
}

func (x *DeleteAirportV1Response) Reset() {
	*x = DeleteAirportV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAirportV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAirportV1Response) ProtoMessage() {}

func (x *DeleteAirportV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAirportV1Response.ProtoReflect.Descriptor instead.
func (*DeleteAirportV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAirportV1Response) GetAirportId() uint64 {
	if x != nil {
		return x.AirportId
	}
	return 0
}

var File_ozonmp_trv_airport_api_v1_trv_airport_api_proto protoreflect.FileDescriptor

var file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x76, 0x5f,
	0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x07, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x03,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x02, 0x18, 0xff, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x18, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x09, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a,
	0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0a, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x61, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x32, 0xed, 0x04, 0x0a, 0x14, 0x54, 0x72, 0x76, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x12, 0x31,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa1,
	0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x56, 0x31, 0x12, 0x33, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74, 0x72,
	0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x56, 0x31, 0x12, 0x30, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74,
	0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x8e, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x56, 0x31, 0x12, 0x31, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x74, 0x72,
	0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x74, 0x72, 0x76, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x74, 0x72, 0x76, 0x2d, 0x61, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x76, 0x2d, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x74, 0x72, 0x76, 0x5f, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescOnce sync.Once
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescData = file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDesc
)

func file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescGZIP() []byte {
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescOnce.Do(func() {
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescData)
	})
	return file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDescData
}

var file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_goTypes = []interface{}{
	(*Airport)(nil),                   // 0: ozonmp.trv_airport_api.v1.Airport
	(*CreateAirportV1Request)(nil),    // 1: ozonmp.trv_airport_api.v1.CreateAirportV1Request
	(*CreateAirportV1Response)(nil),   // 2: ozonmp.trv_airport_api.v1.CreateAirportV1Response
	(*DescribeAirportV1Request)(nil),  // 3: ozonmp.trv_airport_api.v1.DescribeAirportV1Request
	(*DescribeAirportV1Response)(nil), // 4: ozonmp.trv_airport_api.v1.DescribeAirportV1Response
	(*ListAirportsV1Request)(nil),     // 5: ozonmp.trv_airport_api.v1.ListAirportsV1Request
	(*ListAirportsV1Response)(nil),    // 6: ozonmp.trv_airport_api.v1.ListAirportsV1Response
	(*DeleteAirportV1Request)(nil),    // 7: ozonmp.trv_airport_api.v1.DeleteAirportV1Request
	(*DeleteAirportV1Response)(nil),   // 8: ozonmp.trv_airport_api.v1.DeleteAirportV1Response
}
var file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_depIdxs = []int32{
	0, // 0: ozonmp.trv_airport_api.v1.DescribeAirportV1Response.value:type_name -> ozonmp.trv_airport_api.v1.Airport
	0, // 1: ozonmp.trv_airport_api.v1.ListAirportsV1Response.airports:type_name -> ozonmp.trv_airport_api.v1.Airport
	1, // 2: ozonmp.trv_airport_api.v1.TrvAirportApiService.CreateAirportV1:input_type -> ozonmp.trv_airport_api.v1.CreateAirportV1Request
	3, // 3: ozonmp.trv_airport_api.v1.TrvAirportApiService.DescribeAirportV1:input_type -> ozonmp.trv_airport_api.v1.DescribeAirportV1Request
	5, // 4: ozonmp.trv_airport_api.v1.TrvAirportApiService.ListAirportsV1:input_type -> ozonmp.trv_airport_api.v1.ListAirportsV1Request
	7, // 5: ozonmp.trv_airport_api.v1.TrvAirportApiService.DeleteAirportV1:input_type -> ozonmp.trv_airport_api.v1.DeleteAirportV1Request
	2, // 6: ozonmp.trv_airport_api.v1.TrvAirportApiService.CreateAirportV1:output_type -> ozonmp.trv_airport_api.v1.CreateAirportV1Response
	4, // 7: ozonmp.trv_airport_api.v1.TrvAirportApiService.DescribeAirportV1:output_type -> ozonmp.trv_airport_api.v1.DescribeAirportV1Response
	6, // 8: ozonmp.trv_airport_api.v1.TrvAirportApiService.ListAirportsV1:output_type -> ozonmp.trv_airport_api.v1.ListAirportsV1Response
	8, // 9: ozonmp.trv_airport_api.v1.TrvAirportApiService.DeleteAirportV1:output_type -> ozonmp.trv_airport_api.v1.DeleteAirportV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_init() }
func file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_init() {
	if File_ozonmp_trv_airport_api_v1_trv_airport_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Airport); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAirportV1Request); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAirportV1Response); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeAirportV1Request); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeAirportV1Response); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAirportsV1Request); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAirportsV1Response); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAirportV1Request); i {
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
		file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAirportV1Response); i {
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
			RawDescriptor: file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_goTypes,
		DependencyIndexes: file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_depIdxs,
		MessageInfos:      file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_msgTypes,
	}.Build()
	File_ozonmp_trv_airport_api_v1_trv_airport_api_proto = out.File
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_rawDesc = nil
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_goTypes = nil
	file_ozonmp_trv_airport_api_v1_trv_airport_api_proto_depIdxs = nil
}
