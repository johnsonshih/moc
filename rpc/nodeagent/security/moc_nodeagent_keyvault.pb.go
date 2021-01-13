// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_keyvault.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type KeyVaultRequest struct {
	KeyVaults            []*KeyVault      `protobuf:"bytes,1,rep,name=KeyVaults,proto3" json:"KeyVaults,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KeyVaultRequest) Reset()         { *m = KeyVaultRequest{} }
func (m *KeyVaultRequest) String() string { return proto.CompactTextString(m) }
func (*KeyVaultRequest) ProtoMessage()    {}
func (*KeyVaultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_316bd727a7e0a804, []int{0}
}

func (m *KeyVaultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVaultRequest.Unmarshal(m, b)
}
func (m *KeyVaultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVaultRequest.Marshal(b, m, deterministic)
}
func (m *KeyVaultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVaultRequest.Merge(m, src)
}
func (m *KeyVaultRequest) XXX_Size() int {
	return xxx_messageInfo_KeyVaultRequest.Size(m)
}
func (m *KeyVaultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVaultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVaultRequest proto.InternalMessageInfo

func (m *KeyVaultRequest) GetKeyVaults() []*KeyVault {
	if m != nil {
		return m.KeyVaults
	}
	return nil
}

func (m *KeyVaultRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type KeyVaultResponse struct {
	KeyVaults            []*KeyVault         `protobuf:"bytes,1,rep,name=KeyVaults,proto3" json:"KeyVaults,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KeyVaultResponse) Reset()         { *m = KeyVaultResponse{} }
func (m *KeyVaultResponse) String() string { return proto.CompactTextString(m) }
func (*KeyVaultResponse) ProtoMessage()    {}
func (*KeyVaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_316bd727a7e0a804, []int{1}
}

func (m *KeyVaultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVaultResponse.Unmarshal(m, b)
}
func (m *KeyVaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVaultResponse.Marshal(b, m, deterministic)
}
func (m *KeyVaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVaultResponse.Merge(m, src)
}
func (m *KeyVaultResponse) XXX_Size() int {
	return xxx_messageInfo_KeyVaultResponse.Size(m)
}
func (m *KeyVaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVaultResponse proto.InternalMessageInfo

func (m *KeyVaultResponse) GetKeyVaults() []*KeyVault {
	if m != nil {
		return m.KeyVaults
	}
	return nil
}

func (m *KeyVaultResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *KeyVaultResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type KeyVault struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Secrets              []*Secret      `protobuf:"bytes,3,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	GroupName            string         `protobuf:"bytes,4,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,6,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KeyVault) Reset()         { *m = KeyVault{} }
func (m *KeyVault) String() string { return proto.CompactTextString(m) }
func (*KeyVault) ProtoMessage()    {}
func (*KeyVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_316bd727a7e0a804, []int{2}
}

func (m *KeyVault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVault.Unmarshal(m, b)
}
func (m *KeyVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVault.Marshal(b, m, deterministic)
}
func (m *KeyVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVault.Merge(m, src)
}
func (m *KeyVault) XXX_Size() int {
	return xxx_messageInfo_KeyVault.Size(m)
}
func (m *KeyVault) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVault.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVault proto.InternalMessageInfo

func (m *KeyVault) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyVault) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KeyVault) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *KeyVault) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *KeyVault) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *KeyVault) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *KeyVault) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyVaultRequest)(nil), "moc.nodeagent.security.KeyVaultRequest")
	proto.RegisterType((*KeyVaultResponse)(nil), "moc.nodeagent.security.KeyVaultResponse")
	proto.RegisterType((*KeyVault)(nil), "moc.nodeagent.security.KeyVault")
}

func init() { proto.RegisterFile("moc_nodeagent_keyvault.proto", fileDescriptor_316bd727a7e0a804) }

var fileDescriptor_316bd727a7e0a804 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xcd, 0xee, 0x36, 0x35, 0xb3, 0x74, 0x95, 0x41, 0x34, 0x84, 0x5a, 0x96, 0xf5, 0xc2,
	0xbd, 0x9a, 0x60, 0xf4, 0xc2, 0x2b, 0xc1, 0x42, 0x2f, 0x44, 0x50, 0x98, 0x96, 0x5e, 0x08, 0x52,
	0x66, 0x67, 0x4f, 0x63, 0x68, 0x32, 0x27, 0xce, 0x9f, 0x4a, 0x9e, 0xc0, 0x97, 0xf0, 0x1d, 0x7d,
	0x05, 0xc9, 0xe4, 0xcf, 0x5a, 0x50, 0xf6, 0xa2, 0x57, 0x49, 0xce, 0xf9, 0xcd, 0xf7, 0x7d, 0x73,
	0x4e, 0xc8, 0x71, 0x85, 0xf2, 0x4a, 0xe1, 0x16, 0x44, 0x0e, 0xca, 0x5e, 0xdd, 0x40, 0x73, 0x2b,
	0x5c, 0x69, 0x59, 0xad, 0xd1, 0x22, 0x7d, 0x5a, 0xa1, 0x64, 0x63, 0x97, 0x19, 0x90, 0x4e, 0x17,
	0xb6, 0x49, 0x4e, 0x72, 0xc4, 0xbc, 0x84, 0xd4, 0x53, 0x1b, 0x77, 0x9d, 0xfe, 0xd0, 0xa2, 0xae,
	0x41, 0x9b, 0xee, 0x5c, 0xf2, 0xac, 0x55, 0x95, 0x58, 0x55, 0xa8, 0xfa, 0x47, 0xdf, 0x48, 0xee,
	0xda, 0x19, 0x90, 0x1a, 0x7a, 0xb3, 0xd5, 0xcf, 0x80, 0x3c, 0xfa, 0x08, 0xcd, 0x65, 0xeb, 0xcf,
	0xe1, 0xbb, 0x03, 0x63, 0xe9, 0x3b, 0x12, 0x0d, 0x25, 0x13, 0x07, 0xcb, 0xe9, 0x7a, 0x9e, 0x2d,
	0xd9, 0xbf, 0x43, 0xb1, 0xf1, 0xec, 0xee, 0x08, 0x7d, 0x43, 0x8e, 0x3e, 0xd7, 0xa0, 0x85, 0x2d,
	0x50, 0x5d, 0x34, 0x35, 0xc4, 0x93, 0x65, 0xb0, 0x5e, 0x64, 0x0b, 0xaf, 0x31, 0x76, 0xf8, 0x5d,
	0x68, 0xf5, 0x2b, 0x20, 0x8f, 0x77, 0x49, 0x4c, 0x8d, 0xca, 0xc0, 0xbd, 0xa3, 0x64, 0x24, 0xe4,
	0x60, 0x5c, 0x69, 0x7d, 0x86, 0x79, 0x96, 0xb0, 0x6e, 0x88, 0x6c, 0x18, 0x22, 0x3b, 0x45, 0x2c,
	0x2f, 0x45, 0xe9, 0x80, 0xf7, 0x24, 0x7d, 0x42, 0x0e, 0xce, 0xb4, 0x46, 0x1d, 0x4f, 0x97, 0xc1,
	0x3a, 0xe2, 0xdd, 0xc7, 0xea, 0x77, 0x40, 0x1e, 0x0e, 0xba, 0x94, 0x92, 0x99, 0x12, 0x15, 0xc4,
	0x81, 0x27, 0xfc, 0x3b, 0x5d, 0x90, 0x49, 0xb1, 0xf5, 0x36, 0x11, 0x9f, 0x14, 0x5b, 0xfa, 0x96,
	0x1c, 0x9e, 0xfb, 0x49, 0x9b, 0x78, 0xea, 0x83, 0x9f, 0xfc, 0x2f, 0x78, 0x87, 0xf1, 0x01, 0xa7,
	0xc7, 0x24, 0xca, 0x35, 0xba, 0xfa, 0x53, 0x6b, 0x31, 0xf3, 0x82, 0xbb, 0x02, 0x7d, 0x41, 0x42,
	0x63, 0x85, 0x75, 0x26, 0x3e, 0xf0, 0x57, 0x9a, 0x7b, 0xd9, 0x73, 0x5f, 0xe2, 0x7d, 0xab, 0x85,
	0x40, 0xd9, 0xc2, 0x36, 0x71, 0xf8, 0x17, 0x74, 0xe6, 0x4b, 0xbc, 0x6f, 0xd1, 0xe7, 0x64, 0x66,
	0x45, 0x6e, 0xe2, 0x43, 0x8f, 0x44, 0x1e, 0xb9, 0x10, 0xb9, 0xe1, 0xbe, 0x9c, 0x29, 0x72, 0x34,
	0x5c, 0xf8, 0x7d, 0x1b, 0x98, 0x7e, 0x25, 0xe1, 0x07, 0x75, 0x8b, 0x37, 0x40, 0x5f, 0xee, 0xdd,
	0x41, 0xf7, 0x2b, 0x25, 0xeb, 0xfd, 0x60, 0xb7, 0xe9, 0xd5, 0x83, 0xd3, 0x57, 0x5f, 0xd2, 0xbc,
	0xb0, 0xdf, 0xdc, 0x86, 0x49, 0xac, 0xd2, 0xaa, 0x90, 0x1a, 0x0d, 0x5e, 0xdb, 0xb4, 0x42, 0x99,
	0xea, 0x5a, 0xa6, 0xa3, 0x4a, 0x3a, 0xa8, 0x6c, 0x42, 0xbf, 0xc6, 0xd7, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x6a, 0xe1, 0xb4, 0xda, 0x51, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyVaultAgentClient is the client API for KeyVaultAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyVaultAgentClient interface {
	Invoke(ctx context.Context, in *KeyVaultRequest, opts ...grpc.CallOption) (*KeyVaultResponse, error)
}

type keyVaultAgentClient struct {
	cc *grpc.ClientConn
}

func NewKeyVaultAgentClient(cc *grpc.ClientConn) KeyVaultAgentClient {
	return &keyVaultAgentClient{cc}
}

func (c *keyVaultAgentClient) Invoke(ctx context.Context, in *KeyVaultRequest, opts ...grpc.CallOption) (*KeyVaultResponse, error) {
	out := new(KeyVaultResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.KeyVaultAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyVaultAgentServer is the server API for KeyVaultAgent service.
type KeyVaultAgentServer interface {
	Invoke(context.Context, *KeyVaultRequest) (*KeyVaultResponse, error)
}

// UnimplementedKeyVaultAgentServer can be embedded to have forward compatible implementations.
type UnimplementedKeyVaultAgentServer struct {
}

func (*UnimplementedKeyVaultAgentServer) Invoke(ctx context.Context, req *KeyVaultRequest) (*KeyVaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterKeyVaultAgentServer(s *grpc.Server, srv KeyVaultAgentServer) {
	s.RegisterService(&_KeyVaultAgent_serviceDesc, srv)
}

func _KeyVaultAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyVaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyVaultAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.KeyVaultAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyVaultAgentServer).Invoke(ctx, req.(*KeyVaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyVaultAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.KeyVaultAgent",
	HandlerType: (*KeyVaultAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _KeyVaultAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_keyvault.proto",
}