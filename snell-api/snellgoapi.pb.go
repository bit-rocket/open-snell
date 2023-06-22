// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: snellgoapi.proto

package snellapi

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

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"` //reserved for server api
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snellgoapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snellgoapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_snellgoapi_proto_rawDescGZIP(), []int{0}
}

func (x *StatsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type StatsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadTraffic   uint64 `protobuf:"varint,1,opt,name=upload_traffic,json=uploadTraffic,proto3" json:"upload_traffic,omitempty"`
	DownloadTraffic uint64 `protobuf:"varint,2,opt,name=download_traffic,json=downloadTraffic,proto3" json:"download_traffic,omitempty"`
	UploadSpeed     uint64 `protobuf:"varint,3,opt,name=upload_speed,json=uploadSpeed,proto3" json:"upload_speed,omitempty"`
	DownloadSpeed   uint64 `protobuf:"varint,4,opt,name=download_speed,json=downloadSpeed,proto3" json:"download_speed,omitempty"`
}

func (x *StatsReply) Reset() {
	*x = StatsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snellgoapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsReply) ProtoMessage() {}

func (x *StatsReply) ProtoReflect() protoreflect.Message {
	mi := &file_snellgoapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsReply.ProtoReflect.Descriptor instead.
func (*StatsReply) Descriptor() ([]byte, []int) {
	return file_snellgoapi_proto_rawDescGZIP(), []int{1}
}

func (x *StatsReply) GetUploadTraffic() uint64 {
	if x != nil {
		return x.UploadTraffic
	}
	return 0
}

func (x *StatsReply) GetDownloadTraffic() uint64 {
	if x != nil {
		return x.DownloadTraffic
	}
	return 0
}

func (x *StatsReply) GetUploadSpeed() uint64 {
	if x != nil {
		return x.UploadSpeed
	}
	return 0
}

func (x *StatsReply) GetDownloadSpeed() uint64 {
	if x != nil {
		return x.DownloadSpeed
	}
	return 0
}

var File_snellgoapi_proto protoreflect.FileDescriptor

var file_snellgoapi_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6e, 0x65, 0x6c, 0x6c, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x6e, 0x65, 0x6c, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12,
	0x29, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x32, 0x4e, 0x0a, 0x0c, 0x53, 0x6e, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x6e, 0x65, 0x6c, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x6e,
	0x65, 0x6c, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x6e, 0x65, 0x6c, 0x6c, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snellgoapi_proto_rawDescOnce sync.Once
	file_snellgoapi_proto_rawDescData = file_snellgoapi_proto_rawDesc
)

func file_snellgoapi_proto_rawDescGZIP() []byte {
	file_snellgoapi_proto_rawDescOnce.Do(func() {
		file_snellgoapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_snellgoapi_proto_rawDescData)
	})
	return file_snellgoapi_proto_rawDescData
}

var file_snellgoapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_snellgoapi_proto_goTypes = []interface{}{
	(*StatsRequest)(nil), // 0: snell.api.StatsRequest
	(*StatsReply)(nil),   // 1: snell.api.StatsReply
}
var file_snellgoapi_proto_depIdxs = []int32{
	0, // 0: snell.api.SnellService.QueryStats:input_type -> snell.api.StatsRequest
	1, // 1: snell.api.SnellService.QueryStats:output_type -> snell.api.StatsReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_snellgoapi_proto_init() }
func file_snellgoapi_proto_init() {
	if File_snellgoapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snellgoapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
		file_snellgoapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsReply); i {
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
			RawDescriptor: file_snellgoapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snellgoapi_proto_goTypes,
		DependencyIndexes: file_snellgoapi_proto_depIdxs,
		MessageInfos:      file_snellgoapi_proto_msgTypes,
	}.Build()
	File_snellgoapi_proto = out.File
	file_snellgoapi_proto_rawDesc = nil
	file_snellgoapi_proto_goTypes = nil
	file_snellgoapi_proto_depIdxs = nil
}
