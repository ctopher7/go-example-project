// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: part1/proto/ohlc.proto

package __

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

type GetSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockName string `protobuf:"bytes,1,opt,name=StockName,proto3" json:"StockName,omitempty"`
}

func (x *GetSummaryRequest) Reset() {
	*x = GetSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_part1_proto_ohlc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryRequest) ProtoMessage() {}

func (x *GetSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_part1_proto_ohlc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetSummaryRequest) Descriptor() ([]byte, []int) {
	return file_part1_proto_ohlc_proto_rawDescGZIP(), []int{0}
}

func (x *GetSummaryRequest) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

type GetSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousPrice int32 `protobuf:"varint,1,opt,name=PreviousPrice,proto3" json:"PreviousPrice,omitempty"`
	OpenPrice     int32 `protobuf:"varint,2,opt,name=OpenPrice,proto3" json:"OpenPrice,omitempty"`
	HighestPrice  int32 `protobuf:"varint,3,opt,name=HighestPrice,proto3" json:"HighestPrice,omitempty"`
	LowestPrice   int32 `protobuf:"varint,4,opt,name=LowestPrice,proto3" json:"LowestPrice,omitempty"`
	ClosePrice    int32 `protobuf:"varint,5,opt,name=ClosePrice,proto3" json:"ClosePrice,omitempty"`
	AveragePrice  int32 `protobuf:"varint,6,opt,name=AveragePrice,proto3" json:"AveragePrice,omitempty"`
	Volume        int32 `protobuf:"varint,7,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Value         int32 `protobuf:"varint,8,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetSummaryResponse) Reset() {
	*x = GetSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_part1_proto_ohlc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryResponse) ProtoMessage() {}

func (x *GetSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_part1_proto_ohlc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetSummaryResponse) Descriptor() ([]byte, []int) {
	return file_part1_proto_ohlc_proto_rawDescGZIP(), []int{1}
}

func (x *GetSummaryResponse) GetPreviousPrice() int32 {
	if x != nil {
		return x.PreviousPrice
	}
	return 0
}

func (x *GetSummaryResponse) GetOpenPrice() int32 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *GetSummaryResponse) GetHighestPrice() int32 {
	if x != nil {
		return x.HighestPrice
	}
	return 0
}

func (x *GetSummaryResponse) GetLowestPrice() int32 {
	if x != nil {
		return x.LowestPrice
	}
	return 0
}

func (x *GetSummaryResponse) GetClosePrice() int32 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *GetSummaryResponse) GetAveragePrice() int32 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

func (x *GetSummaryResponse) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *GetSummaryResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_part1_proto_ohlc_proto protoreflect.FileDescriptor

var file_part1_proto_ohlc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x68,
	0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x90, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x4b, 0x0a, 0x04, 0x4f, 0x68, 0x6c, 0x63, 0x12, 0x43, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_part1_proto_ohlc_proto_rawDescOnce sync.Once
	file_part1_proto_ohlc_proto_rawDescData = file_part1_proto_ohlc_proto_rawDesc
)

func file_part1_proto_ohlc_proto_rawDescGZIP() []byte {
	file_part1_proto_ohlc_proto_rawDescOnce.Do(func() {
		file_part1_proto_ohlc_proto_rawDescData = protoimpl.X.CompressGZIP(file_part1_proto_ohlc_proto_rawDescData)
	})
	return file_part1_proto_ohlc_proto_rawDescData
}

var file_part1_proto_ohlc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_part1_proto_ohlc_proto_goTypes = []interface{}{
	(*GetSummaryRequest)(nil),  // 0: proto.GetSummaryRequest
	(*GetSummaryResponse)(nil), // 1: proto.GetSummaryResponse
}
var file_part1_proto_ohlc_proto_depIdxs = []int32{
	0, // 0: proto.Ohlc.GetSummary:input_type -> proto.GetSummaryRequest
	1, // 1: proto.Ohlc.GetSummary:output_type -> proto.GetSummaryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_part1_proto_ohlc_proto_init() }
func file_part1_proto_ohlc_proto_init() {
	if File_part1_proto_ohlc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_part1_proto_ohlc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryRequest); i {
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
		file_part1_proto_ohlc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryResponse); i {
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
			RawDescriptor: file_part1_proto_ohlc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_part1_proto_ohlc_proto_goTypes,
		DependencyIndexes: file_part1_proto_ohlc_proto_depIdxs,
		MessageInfos:      file_part1_proto_ohlc_proto_msgTypes,
	}.Build()
	File_part1_proto_ohlc_proto = out.File
	file_part1_proto_ohlc_proto_rawDesc = nil
	file_part1_proto_ohlc_proto_goTypes = nil
	file_part1_proto_ohlc_proto_depIdxs = nil
}