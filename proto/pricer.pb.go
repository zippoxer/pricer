// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/pricer.proto

package proto

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

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pricer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pricer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_pricer_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PriceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PriceReply) Reset() {
	*x = PriceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pricer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceReply) ProtoMessage() {}

func (x *PriceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pricer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceReply.ProtoReflect.Descriptor instead.
func (*PriceReply) Descriptor() ([]byte, []int) {
	return file_proto_pricer_proto_rawDescGZIP(), []int{1}
}

func (x *PriceReply) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_pricer_proto protoreflect.FileDescriptor

var file_proto_pricer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x3b, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x47, 0x0a, 0x14, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x70, 0x70, 0x6f, 0x78, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_pricer_proto_rawDescOnce sync.Once
	file_proto_pricer_proto_rawDescData = file_proto_pricer_proto_rawDesc
)

func file_proto_pricer_proto_rawDescGZIP() []byte {
	file_proto_pricer_proto_rawDescOnce.Do(func() {
		file_proto_pricer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pricer_proto_rawDescData)
	})
	return file_proto_pricer_proto_rawDescData
}

var file_proto_pricer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pricer_proto_goTypes = []interface{}{
	(*PriceRequest)(nil), // 0: proto.PriceRequest
	(*PriceReply)(nil),   // 1: proto.PriceReply
}
var file_proto_pricer_proto_depIdxs = []int32{
	0, // 0: proto.Pricer.Price:input_type -> proto.PriceRequest
	1, // 1: proto.Pricer.Price:output_type -> proto.PriceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pricer_proto_init() }
func file_proto_pricer_proto_init() {
	if File_proto_pricer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pricer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_proto_pricer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceReply); i {
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
			RawDescriptor: file_proto_pricer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pricer_proto_goTypes,
		DependencyIndexes: file_proto_pricer_proto_depIdxs,
		MessageInfos:      file_proto_pricer_proto_msgTypes,
	}.Build()
	File_proto_pricer_proto = out.File
	file_proto_pricer_proto_rawDesc = nil
	file_proto_pricer_proto_goTypes = nil
	file_proto_pricer_proto_depIdxs = nil
}
