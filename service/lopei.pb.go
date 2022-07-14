// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: model/lopei.proto

package service

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

// number1 + number2
type CheckBalanceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LopeiId int32 `protobuf:"varint,1,opt,name=lopeiId,proto3" json:"lopeiId,omitempty"`
}

func (x *CheckBalanceMessage) Reset() {
	*x = CheckBalanceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_lopei_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBalanceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBalanceMessage) ProtoMessage() {}

func (x *CheckBalanceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_model_lopei_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBalanceMessage.ProtoReflect.Descriptor instead.
func (*CheckBalanceMessage) Descriptor() ([]byte, []int) {
	return file_model_lopei_proto_rawDescGZIP(), []int{0}
}

func (x *CheckBalanceMessage) GetLopeiId() int32 {
	if x != nil {
		return x.LopeiId
	}
	return 0
}

type PaymentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LopeiId int32   `protobuf:"varint,1,opt,name=lopeiId,proto3" json:"lopeiId,omitempty"`
	Amount  float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentMessage) Reset() {
	*x = PaymentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_lopei_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMessage) ProtoMessage() {}

func (x *PaymentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_model_lopei_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMessage.ProtoReflect.Descriptor instead.
func (*PaymentMessage) Descriptor() ([]byte, []int) {
	return file_model_lopei_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMessage) GetLopeiId() int32 {
	if x != nil {
		return x.LopeiId
	}
	return 0
}

func (x *PaymentMessage) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_lopei_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_model_lopei_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_model_lopei_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResultMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultMessage) Reset() {
	*x = ResultMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_lopei_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultMessage) ProtoMessage() {}

func (x *ResultMessage) ProtoReflect() protoreflect.Message {
	mi := &file_model_lopei_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultMessage.ProtoReflect.Descriptor instead.
func (*ResultMessage) Descriptor() ([]byte, []int) {
	return file_model_lopei_proto_rawDescGZIP(), []int{3}
}

func (x *ResultMessage) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ResultMessage) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_model_lopei_proto protoreflect.FileDescriptor

var file_model_lopei_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6c, 0x6f, 0x70, 0x65, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f,
	0x70, 0x65, 0x69, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x6f, 0x70,
	0x65, 0x69, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x70, 0x65, 0x69, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x6f, 0x70, 0x65, 0x69, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x45, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x76, 0x0a, 0x0c, 0x4c, 0x6f, 0x70, 0x65, 0x69, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x09, 0x44, 0x6f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_lopei_proto_rawDescOnce sync.Once
	file_model_lopei_proto_rawDescData = file_model_lopei_proto_rawDesc
)

func file_model_lopei_proto_rawDescGZIP() []byte {
	file_model_lopei_proto_rawDescOnce.Do(func() {
		file_model_lopei_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_lopei_proto_rawDescData)
	})
	return file_model_lopei_proto_rawDescData
}

var file_model_lopei_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_lopei_proto_goTypes = []interface{}{
	(*CheckBalanceMessage)(nil), // 0: CheckBalanceMessage
	(*PaymentMessage)(nil),      // 1: PaymentMessage
	(*Error)(nil),               // 2: Error
	(*ResultMessage)(nil),       // 3: ResultMessage
}
var file_model_lopei_proto_depIdxs = []int32{
	2, // 0: ResultMessage.error:type_name -> Error
	0, // 1: LopeiPayment.CheckBalance:input_type -> CheckBalanceMessage
	1, // 2: LopeiPayment.DoPayment:input_type -> PaymentMessage
	3, // 3: LopeiPayment.CheckBalance:output_type -> ResultMessage
	3, // 4: LopeiPayment.DoPayment:output_type -> ResultMessage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_lopei_proto_init() }
func file_model_lopei_proto_init() {
	if File_model_lopei_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_lopei_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckBalanceMessage); i {
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
		file_model_lopei_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMessage); i {
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
		file_model_lopei_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_model_lopei_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultMessage); i {
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
			RawDescriptor: file_model_lopei_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_lopei_proto_goTypes,
		DependencyIndexes: file_model_lopei_proto_depIdxs,
		MessageInfos:      file_model_lopei_proto_msgTypes,
	}.Build()
	File_model_lopei_proto = out.File
	file_model_lopei_proto_rawDesc = nil
	file_model_lopei_proto_goTypes = nil
	file_model_lopei_proto_depIdxs = nil
}
