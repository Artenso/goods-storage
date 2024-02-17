// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: goods_storage.proto

package goods_storage

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AddProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddProductRequest) Reset() {
	*x = AddProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductRequest) ProtoMessage() {}

func (x *AddProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductRequest.ProtoReflect.Descriptor instead.
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return file_goods_storage_proto_rawDescGZIP(), []int{0}
}

func (x *AddProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddProductResponse) Reset() {
	*x = AddProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductResponse) ProtoMessage() {}

func (x *AddProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductResponse.ProtoReflect.Descriptor instead.
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return file_goods_storage_proto_rawDescGZIP(), []int{1}
}

func (x *AddProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_goods_storage_proto protoreflect.FileDescriptor

var file_goods_storage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x24, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xac, 0x01, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_goods_storage_proto_rawDescOnce sync.Once
	file_goods_storage_proto_rawDescData = file_goods_storage_proto_rawDesc
)

func file_goods_storage_proto_rawDescGZIP() []byte {
	file_goods_storage_proto_rawDescOnce.Do(func() {
		file_goods_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_storage_proto_rawDescData)
	})
	return file_goods_storage_proto_rawDescData
}

var file_goods_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goods_storage_proto_goTypes = []interface{}{
	(*AddProductRequest)(nil),  // 0: github.com.Artenso.goods_storage.api.goods_storage.AddProductRequest
	(*AddProductResponse)(nil), // 1: github.com.Artenso.goods_storage.api.goods_storage.AddProductResponse
}
var file_goods_storage_proto_depIdxs = []int32{
	0, // 0: github.com.Artenso.goods_storage.api.goods_storage.GoodsStorage.AddProduct:input_type -> github.com.Artenso.goods_storage.api.goods_storage.AddProductRequest
	1, // 1: github.com.Artenso.goods_storage.api.goods_storage.GoodsStorage.AddProduct:output_type -> github.com.Artenso.goods_storage.api.goods_storage.AddProductResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goods_storage_proto_init() }
func file_goods_storage_proto_init() {
	if File_goods_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductRequest); i {
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
		file_goods_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductResponse); i {
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
			RawDescriptor: file_goods_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_storage_proto_goTypes,
		DependencyIndexes: file_goods_storage_proto_depIdxs,
		MessageInfos:      file_goods_storage_proto_msgTypes,
	}.Build()
	File_goods_storage_proto = out.File
	file_goods_storage_proto_rawDesc = nil
	file_goods_storage_proto_goTypes = nil
	file_goods_storage_proto_depIdxs = nil
}
