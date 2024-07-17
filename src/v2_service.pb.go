//
//Akeyless API
//
//The purpose of this application is to provide access to Akeyless API.
//
//The version of the OpenAPI document: 2.0
//
//Contact: support@akeyless.io
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: services/v2_service.proto

package akeyless_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *Auth `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetBody() *Auth {
	if x != nil {
		return x.Body
	}
	return nil
}

type CreateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *CreateSecret `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateSecretRequest) Reset() {
	*x = CreateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRequest) ProtoMessage() {}

func (x *CreateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSecretRequest) GetBody() *CreateSecret {
	if x != nil {
		return x.Body
	}
	return nil
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *DeleteItem `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteItemRequest) GetBody() *DeleteItem {
	if x != nil {
		return x.Body
	}
	return nil
}

type DescribeItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *DescribeItem `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DescribeItemRequest) Reset() {
	*x = DescribeItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeItemRequest) ProtoMessage() {}

func (x *DescribeItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeItemRequest.ProtoReflect.Descriptor instead.
func (*DescribeItemRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeItemRequest) GetBody() *DescribeItem {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetRotatedSecretValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *GetRotatedSecretValue `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetRotatedSecretValueRequest) Reset() {
	*x = GetRotatedSecretValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRotatedSecretValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRotatedSecretValueRequest) ProtoMessage() {}

func (x *GetRotatedSecretValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRotatedSecretValueRequest.ProtoReflect.Descriptor instead.
func (*GetRotatedSecretValueRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRotatedSecretValueRequest) GetBody() *GetRotatedSecretValue {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetRotatedSecretValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetRotatedSecretValueResponse) Reset() {
	*x = GetRotatedSecretValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRotatedSecretValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRotatedSecretValueResponse) ProtoMessage() {}

func (x *GetRotatedSecretValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRotatedSecretValueResponse.ProtoReflect.Descriptor instead.
func (*GetRotatedSecretValueResponse) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRotatedSecretValueResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSecretValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *GetSecretValue `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetSecretValueRequest) Reset() {
	*x = GetSecretValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretValueRequest) ProtoMessage() {}

func (x *GetSecretValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretValueRequest.ProtoReflect.Descriptor instead.
func (*GetSecretValueRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSecretValueRequest) GetBody() *GetSecretValue {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetSecretValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSecretValueResponse) Reset() {
	*x = GetSecretValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretValueResponse) ProtoMessage() {}

func (x *GetSecretValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretValueResponse.ProtoReflect.Descriptor instead.
func (*GetSecretValueResponse) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetSecretValueResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *ListItems `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListItemsRequest) GetBody() *ListItems {
	if x != nil {
		return x.Body
	}
	return nil
}

type UpdateSecretValRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *UpdateSecretVal `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateSecretValRequest) Reset() {
	*x = UpdateSecretValRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_v2_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSecretValRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretValRequest) ProtoMessage() {}

func (x *UpdateSecretValRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_v2_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretValRequest.ProtoReflect.Descriptor instead.
func (*UpdateSecretValRequest) Descriptor() ([]byte, []int) {
	return file_services_v2_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateSecretValRequest) GetBody() *UpdateSecretVal {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_services_v2_service_proto protoreflect.FileDescriptor

var file_services_v2_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x61, 0x6b, 0x65,
	0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69,
	0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x5f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x46, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x42, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6b,
	0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x46, 0x0a, 0x13,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x58, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x4c,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x45, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x4c, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6b, 0x65, 0x79,
	0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32,
	0x8a, 0x07, 0x0a, 0x11, 0x41, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x56, 0x32, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2d, 0x2e,
	0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x68, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x35, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65,
	0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x62, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x33, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x5a, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61,
	0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x2e, 0x61, 0x6b,
	0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x61, 0x6b,
	0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x83, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x37, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c,
	0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x65, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x32, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x71, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x38, 0x2e, 0x61,
	0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73,
	0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x51, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0xaa, 0x02, 0x0d, 0x41, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0xca, 0x02, 0x0d, 0x41, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5c,
	0x67, 0x52, 0x50, 0x43, 0xe2, 0x02, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x00, 0x50, 0x01, 0x50, 0x02, 0x50, 0x03, 0x50, 0x04, 0x50, 0x05, 0x50, 0x06, 0x50, 0x07, 0x50,
	0x08, 0x50, 0x09, 0x50, 0x0a, 0x50, 0x0b, 0x50, 0x0c, 0x50, 0x0d, 0x50, 0x0e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_v2_service_proto_rawDescOnce sync.Once
	file_services_v2_service_proto_rawDescData = file_services_v2_service_proto_rawDesc
)

func file_services_v2_service_proto_rawDescGZIP() []byte {
	file_services_v2_service_proto_rawDescOnce.Do(func() {
		file_services_v2_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_v2_service_proto_rawDescData)
	})
	return file_services_v2_service_proto_rawDescData
}

var file_services_v2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_v2_service_proto_goTypes = []any{
	(*AuthRequest)(nil),                   // 0: akeyless_grpc.services.v2service.AuthRequest
	(*CreateSecretRequest)(nil),           // 1: akeyless_grpc.services.v2service.CreateSecretRequest
	(*DeleteItemRequest)(nil),             // 2: akeyless_grpc.services.v2service.DeleteItemRequest
	(*DescribeItemRequest)(nil),           // 3: akeyless_grpc.services.v2service.DescribeItemRequest
	(*GetRotatedSecretValueRequest)(nil),  // 4: akeyless_grpc.services.v2service.GetRotatedSecretValueRequest
	(*GetRotatedSecretValueResponse)(nil), // 5: akeyless_grpc.services.v2service.GetRotatedSecretValueResponse
	(*GetSecretValueRequest)(nil),         // 6: akeyless_grpc.services.v2service.GetSecretValueRequest
	(*GetSecretValueResponse)(nil),        // 7: akeyless_grpc.services.v2service.GetSecretValueResponse
	(*ListItemsRequest)(nil),              // 8: akeyless_grpc.services.v2service.ListItemsRequest
	(*UpdateSecretValRequest)(nil),        // 9: akeyless_grpc.services.v2service.UpdateSecretValRequest
	(*Auth)(nil),                          // 10: akeyless_grpc.Auth
	(*CreateSecret)(nil),                  // 11: akeyless_grpc.CreateSecret
	(*DeleteItem)(nil),                    // 12: akeyless_grpc.DeleteItem
	(*DescribeItem)(nil),                  // 13: akeyless_grpc.DescribeItem
	(*GetRotatedSecretValue)(nil),         // 14: akeyless_grpc.GetRotatedSecretValue
	(*structpb.Struct)(nil),               // 15: google.protobuf.Struct
	(*GetSecretValue)(nil),                // 16: akeyless_grpc.GetSecretValue
	(*ListItems)(nil),                     // 17: akeyless_grpc.ListItems
	(*UpdateSecretVal)(nil),               // 18: akeyless_grpc.UpdateSecretVal
	(*AuthOutput)(nil),                    // 19: akeyless_grpc.AuthOutput
	(*CreateSecretOutput)(nil),            // 20: akeyless_grpc.CreateSecretOutput
	(*DeleteItemOutput)(nil),              // 21: akeyless_grpc.DeleteItemOutput
	(*Item)(nil),                          // 22: akeyless_grpc.Item
	(*ListItemsInPathOutput)(nil),         // 23: akeyless_grpc.ListItemsInPathOutput
	(*UpdateSecretValOutput)(nil),         // 24: akeyless_grpc.UpdateSecretValOutput
}
var file_services_v2_service_proto_depIdxs = []int32{
	10, // 0: akeyless_grpc.services.v2service.AuthRequest.body:type_name -> akeyless_grpc.Auth
	11, // 1: akeyless_grpc.services.v2service.CreateSecretRequest.body:type_name -> akeyless_grpc.CreateSecret
	12, // 2: akeyless_grpc.services.v2service.DeleteItemRequest.body:type_name -> akeyless_grpc.DeleteItem
	13, // 3: akeyless_grpc.services.v2service.DescribeItemRequest.body:type_name -> akeyless_grpc.DescribeItem
	14, // 4: akeyless_grpc.services.v2service.GetRotatedSecretValueRequest.body:type_name -> akeyless_grpc.GetRotatedSecretValue
	15, // 5: akeyless_grpc.services.v2service.GetRotatedSecretValueResponse.data:type_name -> google.protobuf.Struct
	16, // 6: akeyless_grpc.services.v2service.GetSecretValueRequest.body:type_name -> akeyless_grpc.GetSecretValue
	15, // 7: akeyless_grpc.services.v2service.GetSecretValueResponse.data:type_name -> google.protobuf.Struct
	17, // 8: akeyless_grpc.services.v2service.ListItemsRequest.body:type_name -> akeyless_grpc.ListItems
	18, // 9: akeyless_grpc.services.v2service.UpdateSecretValRequest.body:type_name -> akeyless_grpc.UpdateSecretVal
	0,  // 10: akeyless_grpc.services.v2service.AkeylessV2Service.Auth:input_type -> akeyless_grpc.services.v2service.AuthRequest
	1,  // 11: akeyless_grpc.services.v2service.AkeylessV2Service.CreateSecret:input_type -> akeyless_grpc.services.v2service.CreateSecretRequest
	2,  // 12: akeyless_grpc.services.v2service.AkeylessV2Service.DeleteItem:input_type -> akeyless_grpc.services.v2service.DeleteItemRequest
	3,  // 13: akeyless_grpc.services.v2service.AkeylessV2Service.DescribeItem:input_type -> akeyless_grpc.services.v2service.DescribeItemRequest
	4,  // 14: akeyless_grpc.services.v2service.AkeylessV2Service.GetRotatedSecretValue:input_type -> akeyless_grpc.services.v2service.GetRotatedSecretValueRequest
	6,  // 15: akeyless_grpc.services.v2service.AkeylessV2Service.GetSecretValue:input_type -> akeyless_grpc.services.v2service.GetSecretValueRequest
	8,  // 16: akeyless_grpc.services.v2service.AkeylessV2Service.ListItems:input_type -> akeyless_grpc.services.v2service.ListItemsRequest
	9,  // 17: akeyless_grpc.services.v2service.AkeylessV2Service.UpdateSecretVal:input_type -> akeyless_grpc.services.v2service.UpdateSecretValRequest
	19, // 18: akeyless_grpc.services.v2service.AkeylessV2Service.Auth:output_type -> akeyless_grpc.AuthOutput
	20, // 19: akeyless_grpc.services.v2service.AkeylessV2Service.CreateSecret:output_type -> akeyless_grpc.CreateSecretOutput
	21, // 20: akeyless_grpc.services.v2service.AkeylessV2Service.DeleteItem:output_type -> akeyless_grpc.DeleteItemOutput
	22, // 21: akeyless_grpc.services.v2service.AkeylessV2Service.DescribeItem:output_type -> akeyless_grpc.Item
	5,  // 22: akeyless_grpc.services.v2service.AkeylessV2Service.GetRotatedSecretValue:output_type -> akeyless_grpc.services.v2service.GetRotatedSecretValueResponse
	7,  // 23: akeyless_grpc.services.v2service.AkeylessV2Service.GetSecretValue:output_type -> akeyless_grpc.services.v2service.GetSecretValueResponse
	23, // 24: akeyless_grpc.services.v2service.AkeylessV2Service.ListItems:output_type -> akeyless_grpc.ListItemsInPathOutput
	24, // 25: akeyless_grpc.services.v2service.AkeylessV2Service.UpdateSecretVal:output_type -> akeyless_grpc.UpdateSecretValOutput
	18, // [18:26] is the sub-list for method output_type
	10, // [10:18] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_v2_service_proto_init() }
func file_services_v2_service_proto_init() {
	if File_services_v2_service_proto != nil {
		return
	}
	file_models_auth_proto_init()
	file_models_auth_output_proto_init()
	file_models_create_secret_proto_init()
	file_models_create_secret_output_proto_init()
	file_models_delete_item_proto_init()
	file_models_delete_item_output_proto_init()
	file_models_describe_item_proto_init()
	file_models_get_rotated_secret_value_proto_init()
	file_models_get_secret_value_proto_init()
	file_models_item_proto_init()
	file_models_json_error_proto_init()
	file_models_list_items_proto_init()
	file_models_list_items_in_path_output_proto_init()
	file_models_update_secret_val_proto_init()
	file_models_update_secret_val_output_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_v2_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuthRequest); i {
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
		file_services_v2_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSecretRequest); i {
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
		file_services_v2_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteItemRequest); i {
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
		file_services_v2_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DescribeItemRequest); i {
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
		file_services_v2_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetRotatedSecretValueRequest); i {
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
		file_services_v2_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetRotatedSecretValueResponse); i {
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
		file_services_v2_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetSecretValueRequest); i {
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
		file_services_v2_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetSecretValueResponse); i {
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
		file_services_v2_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListItemsRequest); i {
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
		file_services_v2_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSecretValRequest); i {
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
			RawDescriptor: file_services_v2_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_v2_service_proto_goTypes,
		DependencyIndexes: file_services_v2_service_proto_depIdxs,
		MessageInfos:      file_services_v2_service_proto_msgTypes,
	}.Build()
	File_services_v2_service_proto = out.File
	file_services_v2_service_proto_rawDesc = nil
	file_services_v2_service_proto_goTypes = nil
	file_services_v2_service_proto_depIdxs = nil
}
