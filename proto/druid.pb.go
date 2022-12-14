// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/druid.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AcknowlegeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AcknowlegeResponse) Reset() {
	*x = AcknowlegeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_druid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcknowlegeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcknowlegeResponse) ProtoMessage() {}

func (x *AcknowlegeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_druid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcknowlegeResponse.ProtoReflect.Descriptor instead.
func (*AcknowlegeResponse) Descriptor() ([]byte, []int) {
	return file_proto_druid_proto_rawDescGZIP(), []int{0}
}

func (x *AcknowlegeResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type DruidData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Query    string `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *DruidData) Reset() {
	*x = DruidData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_druid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DruidData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DruidData) ProtoMessage() {}

func (x *DruidData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_druid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DruidData.ProtoReflect.Descriptor instead.
func (*DruidData) Descriptor() ([]byte, []int) {
	return file_proto_druid_proto_rawDescGZIP(), []int{1}
}

func (x *DruidData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DruidData) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *DruidData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DruidData) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

var File_proto_druid_proto protoreflect.FileDescriptor

var file_proto_druid_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x12, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x09,
	0x44, 0x72, 0x75, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x32, 0x52, 0x0a, 0x0f, 0x44, 0x72, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x72, 0x75, 0x69,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x72, 0x75, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_druid_proto_rawDescOnce sync.Once
	file_proto_druid_proto_rawDescData = file_proto_druid_proto_rawDesc
)

func file_proto_druid_proto_rawDescGZIP() []byte {
	file_proto_druid_proto_rawDescOnce.Do(func() {
		file_proto_druid_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_druid_proto_rawDescData)
	})
	return file_proto_druid_proto_rawDescData
}

var file_proto_druid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_druid_proto_goTypes = []interface{}{
	(*AcknowlegeResponse)(nil), // 0: proto.AcknowlegeResponse
	(*DruidData)(nil),          // 1: proto.DruidData
}
var file_proto_druid_proto_depIdxs = []int32{
	1, // 0: proto.DruidManagement.SendDruidQuery:input_type -> proto.DruidData
	0, // 1: proto.DruidManagement.SendDruidQuery:output_type -> proto.AcknowlegeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_druid_proto_init() }
func file_proto_druid_proto_init() {
	if File_proto_druid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_druid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcknowlegeResponse); i {
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
		file_proto_druid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DruidData); i {
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
			RawDescriptor: file_proto_druid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_druid_proto_goTypes,
		DependencyIndexes: file_proto_druid_proto_depIdxs,
		MessageInfos:      file_proto_druid_proto_msgTypes,
	}.Build()
	File_proto_druid_proto = out.File
	file_proto_druid_proto_rawDesc = nil
	file_proto_druid_proto_goTypes = nil
	file_proto_druid_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DruidManagementClient is the client API for DruidManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DruidManagementClient interface {
	SendDruidQuery(ctx context.Context, in *DruidData, opts ...grpc.CallOption) (*AcknowlegeResponse, error)
}

type druidManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewDruidManagementClient(cc grpc.ClientConnInterface) DruidManagementClient {
	return &druidManagementClient{cc}
}

func (c *druidManagementClient) SendDruidQuery(ctx context.Context, in *DruidData, opts ...grpc.CallOption) (*AcknowlegeResponse, error) {
	out := new(AcknowlegeResponse)
	err := c.cc.Invoke(ctx, "/proto.DruidManagement/SendDruidQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DruidManagementServer is the server API for DruidManagement service.
type DruidManagementServer interface {
	SendDruidQuery(context.Context, *DruidData) (*AcknowlegeResponse, error)
}

// UnimplementedDruidManagementServer can be embedded to have forward compatible implementations.
type UnimplementedDruidManagementServer struct {
}

func (*UnimplementedDruidManagementServer) SendDruidQuery(context.Context, *DruidData) (*AcknowlegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDruidQuery not implemented")
}

func RegisterDruidManagementServer(s *grpc.Server, srv DruidManagementServer) {
	s.RegisterService(&_DruidManagement_serviceDesc, srv)
}

func _DruidManagement_SendDruidQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DruidData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DruidManagementServer).SendDruidQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DruidManagement/SendDruidQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DruidManagementServer).SendDruidQuery(ctx, req.(*DruidData))
	}
	return interceptor(ctx, in, info, handler)
}

var _DruidManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DruidManagement",
	HandlerType: (*DruidManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDruidQuery",
			Handler:    _DruidManagement_SendDruidQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/druid.proto",
}
