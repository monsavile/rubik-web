// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: resolver.proto

package resolver_v1

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

type ResolveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cube          string                 `protobuf:"bytes,1,opt,name=cube,proto3" json:"cube,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	mi := &file_resolver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_resolver_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveRequest) GetCube() string {
	if x != nil {
		return x.Cube
	}
	return ""
}

type ResolveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cube          string                 `protobuf:"bytes,1,opt,name=cube,proto3" json:"cube,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	mi := &file_resolver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_resolver_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveResponse) GetCube() string {
	if x != nil {
		return x.Cube
	}
	return ""
}

var File_resolver_proto protoreflect.FileDescriptor

var file_resolver_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x22, 0x24, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x75, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x75, 0x62, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x75, 0x62, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x75, 0x62, 0x65, 0x32, 0x52, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x56, 0x31, 0x12, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e,
	0x73, 0x61, 0x76, 0x69, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x62, 0x69, 0x6b, 0x2d, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_resolver_proto_rawDescOnce sync.Once
	file_resolver_proto_rawDescData []byte
)

func file_resolver_proto_rawDescGZIP() []byte {
	file_resolver_proto_rawDescOnce.Do(func() {
		file_resolver_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resolver_proto_rawDesc), len(file_resolver_proto_rawDesc)))
	})
	return file_resolver_proto_rawDescData
}

var file_resolver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resolver_proto_goTypes = []any{
	(*ResolveRequest)(nil),  // 0: resolver_v1.ResolveRequest
	(*ResolveResponse)(nil), // 1: resolver_v1.ResolveResponse
}
var file_resolver_proto_depIdxs = []int32{
	0, // 0: resolver_v1.ResolverV1.Resolve:input_type -> resolver_v1.ResolveRequest
	1, // 1: resolver_v1.ResolverV1.Resolve:output_type -> resolver_v1.ResolveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resolver_proto_init() }
func file_resolver_proto_init() {
	if File_resolver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resolver_proto_rawDesc), len(file_resolver_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resolver_proto_goTypes,
		DependencyIndexes: file_resolver_proto_depIdxs,
		MessageInfos:      file_resolver_proto_msgTypes,
	}.Build()
	File_resolver_proto = out.File
	file_resolver_proto_goTypes = nil
	file_resolver_proto_depIdxs = nil
}
