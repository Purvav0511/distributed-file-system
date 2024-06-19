// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/chunkserver.proto

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

type ReadChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
}

func (x *ReadChunkRequest) Reset() {
	*x = ReadChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunkserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadChunkRequest) ProtoMessage() {}

func (x *ReadChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunkserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadChunkRequest.ProtoReflect.Descriptor instead.
func (*ReadChunkRequest) Descriptor() ([]byte, []int) {
	return file_proto_chunkserver_proto_rawDescGZIP(), []int{0}
}

func (x *ReadChunkRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

type ReadChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadChunkResponse) Reset() {
	*x = ReadChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunkserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadChunkResponse) ProtoMessage() {}

func (x *ReadChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunkserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadChunkResponse.ProtoReflect.Descriptor instead.
func (*ReadChunkResponse) Descriptor() ([]byte, []int) {
	return file_proto_chunkserver_proto_rawDescGZIP(), []int{1}
}

func (x *ReadChunkResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WriteChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WriteChunkRequest) Reset() {
	*x = WriteChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunkserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteChunkRequest) ProtoMessage() {}

func (x *WriteChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunkserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteChunkRequest.ProtoReflect.Descriptor instead.
func (*WriteChunkRequest) Descriptor() ([]byte, []int) {
	return file_proto_chunkserver_proto_rawDescGZIP(), []int{2}
}

func (x *WriteChunkRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *WriteChunkRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WriteChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteChunkResponse) Reset() {
	*x = WriteChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunkserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteChunkResponse) ProtoMessage() {}

func (x *WriteChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunkserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteChunkResponse.ProtoReflect.Descriptor instead.
func (*WriteChunkResponse) Descriptor() ([]byte, []int) {
	return file_proto_chunkserver_proto_rawDescGZIP(), []int{3}
}

var File_proto_chunkserver_proto protoreflect.FileDescriptor

var file_proto_chunkserver_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x22,
	0x27, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x90, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x75, 0x72, 0x76, 0x61, 0x76, 0x30, 0x35, 0x31, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chunkserver_proto_rawDescOnce sync.Once
	file_proto_chunkserver_proto_rawDescData = file_proto_chunkserver_proto_rawDesc
)

func file_proto_chunkserver_proto_rawDescGZIP() []byte {
	file_proto_chunkserver_proto_rawDescOnce.Do(func() {
		file_proto_chunkserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chunkserver_proto_rawDescData)
	})
	return file_proto_chunkserver_proto_rawDescData
}

var file_proto_chunkserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_chunkserver_proto_goTypes = []any{
	(*ReadChunkRequest)(nil),   // 0: proto.ReadChunkRequest
	(*ReadChunkResponse)(nil),  // 1: proto.ReadChunkResponse
	(*WriteChunkRequest)(nil),  // 2: proto.WriteChunkRequest
	(*WriteChunkResponse)(nil), // 3: proto.WriteChunkResponse
}
var file_proto_chunkserver_proto_depIdxs = []int32{
	0, // 0: proto.ChunkServer.ReadChunk:input_type -> proto.ReadChunkRequest
	2, // 1: proto.ChunkServer.WriteChunk:input_type -> proto.WriteChunkRequest
	1, // 2: proto.ChunkServer.ReadChunk:output_type -> proto.ReadChunkResponse
	3, // 3: proto.ChunkServer.WriteChunk:output_type -> proto.WriteChunkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chunkserver_proto_init() }
func file_proto_chunkserver_proto_init() {
	if File_proto_chunkserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chunkserver_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReadChunkRequest); i {
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
		file_proto_chunkserver_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReadChunkResponse); i {
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
		file_proto_chunkserver_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WriteChunkRequest); i {
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
		file_proto_chunkserver_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WriteChunkResponse); i {
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
			RawDescriptor: file_proto_chunkserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chunkserver_proto_goTypes,
		DependencyIndexes: file_proto_chunkserver_proto_depIdxs,
		MessageInfos:      file_proto_chunkserver_proto_msgTypes,
	}.Build()
	File_proto_chunkserver_proto = out.File
	file_proto_chunkserver_proto_rawDesc = nil
	file_proto_chunkserver_proto_goTypes = nil
	file_proto_chunkserver_proto_depIdxs = nil
}