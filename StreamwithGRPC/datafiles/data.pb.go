// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: datafiles/data.proto

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

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datafiles_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datafiles_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_datafiles_data_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransactionRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransactionRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Step        int32  `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datafiles_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datafiles_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_datafiles_data_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransactionResponse) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *TransactionResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_datafiles_data_proto protoreflect.FileDescriptor

var file_datafiles_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x61, 0x74, 0x61, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x22, 0x50, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x63, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x68, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0f,
	0x4d, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datafiles_data_proto_rawDescOnce sync.Once
	file_datafiles_data_proto_rawDescData = file_datafiles_data_proto_rawDesc
)

func file_datafiles_data_proto_rawDescGZIP() []byte {
	file_datafiles_data_proto_rawDescOnce.Do(func() {
		file_datafiles_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_datafiles_data_proto_rawDescData)
	})
	return file_datafiles_data_proto_rawDescData
}

var file_datafiles_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datafiles_data_proto_goTypes = []interface{}{
	(*TransactionRequest)(nil),  // 0: datafiles.TransactionRequest
	(*TransactionResponse)(nil), // 1: datafiles.TransactionResponse
}
var file_datafiles_data_proto_depIdxs = []int32{
	0, // 0: datafiles.MoneyTransaction.MakeTransaction:input_type -> datafiles.TransactionRequest
	1, // 1: datafiles.MoneyTransaction.MakeTransaction:output_type -> datafiles.TransactionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datafiles_data_proto_init() }
func file_datafiles_data_proto_init() {
	if File_datafiles_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datafiles_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_datafiles_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_datafiles_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datafiles_data_proto_goTypes,
		DependencyIndexes: file_datafiles_data_proto_depIdxs,
		MessageInfos:      file_datafiles_data_proto_msgTypes,
	}.Build()
	File_datafiles_data_proto = out.File
	file_datafiles_data_proto_rawDesc = nil
	file_datafiles_data_proto_goTypes = nil
	file_datafiles_data_proto_depIdxs = nil
}
