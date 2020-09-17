// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: function.proto

package handler

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fuction string   `protobuf:"bytes,1,opt,name=Fuction,proto3" json:"Fuction,omitempty"`
	Au      *AppUser `protobuf:"bytes,2,opt,name=Au,proto3" json:"Au,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetFuction() string {
	if x != nil {
		return x.Fuction
	}
	return ""
}

func (x *ClientRequest) GetAu() *AppUser {
	if x != nil {
		return x.Au
	}
	return nil
}

var File_function_proto protoreflect.FileDescriptor

var file_function_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x02, 0x41, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x02,
	0x41, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_function_proto_rawDescOnce sync.Once
	file_function_proto_rawDescData = file_function_proto_rawDesc
)

func file_function_proto_rawDescGZIP() []byte {
	file_function_proto_rawDescOnce.Do(func() {
		file_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_function_proto_rawDescData)
	})
	return file_function_proto_rawDescData
}

var file_function_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_function_proto_goTypes = []interface{}{
	(*ClientRequest)(nil), // 0: handler.ClientRequest
	(*AppUser)(nil),       // 1: handler.AppUser
}
var file_function_proto_depIdxs = []int32{
	1, // 0: handler.ClientRequest.Au:type_name -> handler.AppUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_function_proto_init() }
func file_function_proto_init() {
	if File_function_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
			RawDescriptor: file_function_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_function_proto_goTypes,
		DependencyIndexes: file_function_proto_depIdxs,
		MessageInfos:      file_function_proto_msgTypes,
	}.Build()
	File_function_proto = out.File
	file_function_proto_rawDesc = nil
	file_function_proto_goTypes = nil
	file_function_proto_depIdxs = nil
}
