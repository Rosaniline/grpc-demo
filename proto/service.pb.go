// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetPersonRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPersonRequest) Reset()         { *m = GetPersonRequest{} }
func (m *GetPersonRequest) String() string { return proto.CompactTextString(m) }
func (*GetPersonRequest) ProtoMessage()    {}
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetPersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPersonRequest.Unmarshal(m, b)
}
func (m *GetPersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPersonRequest.Marshal(b, m, deterministic)
}
func (m *GetPersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPersonRequest.Merge(m, src)
}
func (m *GetPersonRequest) XXX_Size() int {
	return xxx_messageInfo_GetPersonRequest.Size(m)
}
func (m *GetPersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPersonRequest proto.InternalMessageInfo

func (m *GetPersonRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPersonRequest)(nil), "getPersonRequest")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x29, 0x48, 0x2d, 0x2a, 0xce,
	0xcf, 0x83, 0xf2, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34,
	0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0xa4, 0x92, 0x12, 0x97, 0x40, 0x7a, 0x6a, 0x49, 0x00,
	0x58, 0x7d, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0xd1, 0x34, 0x46, 0x2e, 0x56, 0x8f, 0xd2,
	0x44, 0xb7, 0x52, 0x21, 0x4d, 0x2e, 0x4e, 0x77, 0x98, 0x6a, 0x21, 0x41, 0x3d, 0x74, 0x9d, 0x52,
	0xec, 0x7a, 0x10, 0xbe, 0x12, 0x83, 0x90, 0x3e, 0x17, 0x77, 0x60, 0x69, 0x6a, 0x51, 0x25, 0x31,
	0x8a, 0x35, 0x18, 0x0d, 0x18, 0x85, 0x0c, 0xb9, 0xb8, 0x7c, 0x32, 0x8b, 0x4b, 0x02, 0x52, 0xf3,
	0x0b, 0x72, 0x52, 0x85, 0xc4, 0xf4, 0x20, 0xae, 0xd6, 0x83, 0xb9, 0x5a, 0xcf, 0x15, 0xe4, 0x6a,
	0x24, 0x4d, 0x06, 0x8c, 0x4e, 0xec, 0x51, 0xac, 0x10, 0x69, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xc5, 0x20, 0x37, 0x43, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HuaFuClient is the client API for HuaFu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HuaFuClient interface {
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*Person, error)
	QueryPerson(ctx context.Context, opts ...grpc.CallOption) (HuaFu_QueryPersonClient, error)
	ListPeople(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (HuaFu_ListPeopleClient, error)
}

type huaFuClient struct {
	cc *grpc.ClientConn
}

func NewHuaFuClient(cc *grpc.ClientConn) HuaFuClient {
	return &huaFuClient{cc}
}

func (c *huaFuClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/HuaFu/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huaFuClient) QueryPerson(ctx context.Context, opts ...grpc.CallOption) (HuaFu_QueryPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HuaFu_serviceDesc.Streams[0], "/HuaFu/QueryPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &huaFuQueryPersonClient{stream}
	return x, nil
}

type HuaFu_QueryPersonClient interface {
	Send(*GetPersonRequest) error
	Recv() (*Person, error)
	grpc.ClientStream
}

type huaFuQueryPersonClient struct {
	grpc.ClientStream
}

func (x *huaFuQueryPersonClient) Send(m *GetPersonRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *huaFuQueryPersonClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *huaFuClient) ListPeople(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (HuaFu_ListPeopleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HuaFu_serviceDesc.Streams[1], "/HuaFu/ListPeople", opts...)
	if err != nil {
		return nil, err
	}
	x := &huaFuListPeopleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HuaFu_ListPeopleClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type huaFuListPeopleClient struct {
	grpc.ClientStream
}

func (x *huaFuListPeopleClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HuaFuServer is the server API for HuaFu service.
type HuaFuServer interface {
	GetPerson(context.Context, *GetPersonRequest) (*Person, error)
	QueryPerson(HuaFu_QueryPersonServer) error
	ListPeople(*empty.Empty, HuaFu_ListPeopleServer) error
}

// UnimplementedHuaFuServer can be embedded to have forward compatible implementations.
type UnimplementedHuaFuServer struct {
}

func (*UnimplementedHuaFuServer) GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedHuaFuServer) QueryPerson(srv HuaFu_QueryPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryPerson not implemented")
}
func (*UnimplementedHuaFuServer) ListPeople(req *empty.Empty, srv HuaFu_ListPeopleServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPeople not implemented")
}

func RegisterHuaFuServer(s *grpc.Server, srv HuaFuServer) {
	s.RegisterService(&_HuaFu_serviceDesc, srv)
}

func _HuaFu_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuaFuServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HuaFu/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuaFuServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuaFu_QueryPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HuaFuServer).QueryPerson(&huaFuQueryPersonServer{stream})
}

type HuaFu_QueryPersonServer interface {
	Send(*Person) error
	Recv() (*GetPersonRequest, error)
	grpc.ServerStream
}

type huaFuQueryPersonServer struct {
	grpc.ServerStream
}

func (x *huaFuQueryPersonServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

func (x *huaFuQueryPersonServer) Recv() (*GetPersonRequest, error) {
	m := new(GetPersonRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HuaFu_ListPeople_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HuaFuServer).ListPeople(m, &huaFuListPeopleServer{stream})
}

type HuaFu_ListPeopleServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type huaFuListPeopleServer struct {
	grpc.ServerStream
}

func (x *huaFuListPeopleServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

var _HuaFu_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HuaFu",
	HandlerType: (*HuaFuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _HuaFu_GetPerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryPerson",
			Handler:       _HuaFu_QueryPerson_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListPeople",
			Handler:       _HuaFu_ListPeople_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
