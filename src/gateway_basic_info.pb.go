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
// source: models/gateway_basic_info.proto

package akeyless_grpc

import (
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

type GatewayBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterDisplayName string `protobuf:"bytes,339533302,opt,name=cluster_display_name,proto3" json:"cluster_display_name,omitempty"`
	ClusterId          int64  `protobuf:"varint,240280960,opt,name=cluster_id,proto3" json:"cluster_id,omitempty"`
	ClusterName        string `protobuf:"bytes,481210961,opt,name=cluster_name,proto3" json:"cluster_name,omitempty"`
	ClusterUrl         string `protobuf:"bytes,67470936,opt,name=cluster_url,proto3" json:"cluster_url,omitempty"`
}

func (x *GatewayBasicInfo) Reset() {
	*x = GatewayBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_gateway_basic_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayBasicInfo) ProtoMessage() {}

func (x *GatewayBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_gateway_basic_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayBasicInfo.ProtoReflect.Descriptor instead.
func (*GatewayBasicInfo) Descriptor() ([]byte, []int) {
	return file_models_gateway_basic_info_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayBasicInfo) GetClusterDisplayName() string {
	if x != nil {
		return x.ClusterDisplayName
	}
	return ""
}

func (x *GatewayBasicInfo) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *GatewayBasicInfo) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GatewayBasicInfo) GetClusterUrl() string {
	if x != nil {
		return x.ClusterUrl
	}
	return ""
}

var File_models_gateway_basic_info_proto protoreflect.FileDescriptor

var file_models_gateway_basic_info_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a,
	0x10, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x36, 0x0a, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xf6, 0xbb, 0xf3, 0xa1, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x80, 0xcb, 0xc9, 0x72, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xd1, 0xe4, 0xba,
	0xe5, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0xd8, 0x8c, 0x96, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x51, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01,
	0x5a, 0x0f, 0x2e, 0x3b, 0x61, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0xaa, 0x02, 0x0d, 0x41, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0xca, 0x02, 0x0d, 0x41, 0x6b, 0x65, 0x79, 0x6c, 0x65, 0x73, 0x73, 0x5c, 0x67, 0x52, 0x50,
	0x43, 0xe2, 0x02, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_gateway_basic_info_proto_rawDescOnce sync.Once
	file_models_gateway_basic_info_proto_rawDescData = file_models_gateway_basic_info_proto_rawDesc
)

func file_models_gateway_basic_info_proto_rawDescGZIP() []byte {
	file_models_gateway_basic_info_proto_rawDescOnce.Do(func() {
		file_models_gateway_basic_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_gateway_basic_info_proto_rawDescData)
	})
	return file_models_gateway_basic_info_proto_rawDescData
}

var file_models_gateway_basic_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_gateway_basic_info_proto_goTypes = []any{
	(*GatewayBasicInfo)(nil), // 0: akeyless_grpc.GatewayBasicInfo
}
var file_models_gateway_basic_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_gateway_basic_info_proto_init() }
func file_models_gateway_basic_info_proto_init() {
	if File_models_gateway_basic_info_proto != nil {
		return
	}
	file_models_special_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_gateway_basic_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GatewayBasicInfo); i {
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
			RawDescriptor: file_models_gateway_basic_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_gateway_basic_info_proto_goTypes,
		DependencyIndexes: file_models_gateway_basic_info_proto_depIdxs,
		MessageInfos:      file_models_gateway_basic_info_proto_msgTypes,
	}.Build()
	File_models_gateway_basic_info_proto = out.File
	file_models_gateway_basic_info_proto_rawDesc = nil
	file_models_gateway_basic_info_proto_goTypes = nil
	file_models_gateway_basic_info_proto_depIdxs = nil
}
