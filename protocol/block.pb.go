// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: protocol/block.proto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash  string `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Hash        string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	BlockNum    uint64 `protobuf:"varint,3,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	ReceiptHash string `protobuf:"bytes,4,opt,name=receipt_hash,json=receiptHash,proto3" json:"receipt_hash,omitempty"`
	Txs         []*Tx  `protobuf:"bytes,5,rep,name=txs,proto3" json:"txs,omitempty"`
	TxsRoot     []byte `protobuf:"bytes,6,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
	Difficulty  string `protobuf:"bytes,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Nonce       uint64 `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TimeStamp   uint64 `protobuf:"varint,9,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_protocol_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *Block) GetReceiptHash() string {
	if x != nil {
		return x.ReceiptHash
	}
	return ""
}

func (x *Block) GetTxs() []*Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *Block) GetTxsRoot() []byte {
	if x != nil {
		return x.TxsRoot
	}
	return nil
}

func (x *Block) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Block) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetTimeStamp() uint64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

var File_protocol_block_proto protoreflect.FileDescriptor

var file_protocol_block_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x03, 0x74,
	0x78, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x78, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74,
	0x78, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x2e, 0x5a, 0x2c, 0x68,
	0x65, 0x79, 0x75, 0x61, 0x6e, 0x6c, 0x6f, 0x6e, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2d, 0x73, 0x74, 0x65, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protocol_block_proto_rawDescOnce sync.Once
	file_protocol_block_proto_rawDescData = file_protocol_block_proto_rawDesc
)

func file_protocol_block_proto_rawDescGZIP() []byte {
	file_protocol_block_proto_rawDescOnce.Do(func() {
		file_protocol_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_block_proto_rawDescData)
	})
	return file_protocol_block_proto_rawDescData
}

var file_protocol_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protocol_block_proto_goTypes = []interface{}{
	(*Block)(nil), // 0: protocol.block
	(*Tx)(nil),    // 1: protocol.tx
}
var file_protocol_block_proto_depIdxs = []int32{
	1, // 0: protocol.block.txs:type_name -> protocol.tx
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_block_proto_init() }
func file_protocol_block_proto_init() {
	if File_protocol_block_proto != nil {
		return
	}
	file_protocol_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_protocol_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_block_proto_goTypes,
		DependencyIndexes: file_protocol_block_proto_depIdxs,
		MessageInfos:      file_protocol_block_proto_msgTypes,
	}.Build()
	File_protocol_block_proto = out.File
	file_protocol_block_proto_rawDesc = nil
	file_protocol_block_proto_goTypes = nil
	file_protocol_block_proto_depIdxs = nil
}
