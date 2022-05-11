// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: protocol/transaction.proto

package protocol

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_protocol_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    *Address `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To        *Address `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount    uint64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Nonce     uint64   `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Input     []byte   `protobuf:"bytes,5,opt,name=input,proto3" json:"input,omitempty"`
	Sign      []byte   `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	TimeStamp uint64   `protobuf:"varint,7,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_protocol_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Tx) GetSender() *Address {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Tx) GetTo() *Address {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Tx) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Tx) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Tx) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Tx) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *Tx) GetTimeStamp() uint64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

var File_protocol_transaction_proto protoreflect.FileDescriptor

var file_protocol_transaction_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x23, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x02,
	0x74, 0x78, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x29, 0x5a, 0x27, 0x68, 0x65, 0x79, 0x75, 0x61,
	0x6e, 0x6c, 0x6f, 0x6e, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x73, 0x74, 0x65, 0x70,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_transaction_proto_rawDescOnce sync.Once
	file_protocol_transaction_proto_rawDescData = file_protocol_transaction_proto_rawDesc
)

func file_protocol_transaction_proto_rawDescGZIP() []byte {
	file_protocol_transaction_proto_rawDescOnce.Do(func() {
		file_protocol_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_transaction_proto_rawDescData)
	})
	return file_protocol_transaction_proto_rawDescData
}

var file_protocol_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_transaction_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: protocol.address
	(*Tx)(nil),      // 1: protocol.tx
}
var file_protocol_transaction_proto_depIdxs = []int32{
	0, // 0: protocol.tx.sender:type_name -> protocol.address
	0, // 1: protocol.tx.to:type_name -> protocol.address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_transaction_proto_init() }
func file_protocol_transaction_proto_init() {
	if File_protocol_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_protocol_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
			RawDescriptor: file_protocol_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_transaction_proto_goTypes,
		DependencyIndexes: file_protocol_transaction_proto_depIdxs,
		MessageInfos:      file_protocol_transaction_proto_msgTypes,
	}.Build()
	File_protocol_transaction_proto = out.File
	file_protocol_transaction_proto_rawDesc = nil
	file_protocol_transaction_proto_goTypes = nil
	file_protocol_transaction_proto_depIdxs = nil
}
