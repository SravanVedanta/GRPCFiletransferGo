// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfer.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{0}
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	TotalSize            string   `protobuf:"bytes,2,opt,name=totalSize,proto3" json:"totalSize,omitempty"`
	Received             string   `protobuf:"bytes,3,opt,name=received,proto3" json:"received,omitempty"`
	Filename             string   `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{0}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chunk) GetTotalSize() string {
	if m != nil {
		return m.TotalSize
	}
	return ""
}

func (m *Chunk) GetReceived() string {
	if m != nil {
		return m.Received
	}
	return ""
}

func (m *Chunk) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=pb.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{1}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("pb.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*Chunk)(nil), "pb.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "pb.UploadStatus")
}

func init() {
	proto.RegisterFile("transfer.proto", fileDescriptor_96c3e6bcafb460d3)
}

var fileDescriptor_96c3e6bcafb460d3 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xd7, 0x3a, 0x3b, 0xfb, 0x1c, 0x25, 0x3c, 0x3c, 0x84, 0xe1, 0x61, 0xf4, 0x54, 0x3d,
	0xf4, 0xb0, 0x1d, 0x3d, 0x16, 0xf4, 0x24, 0x42, 0xc6, 0x7e, 0x40, 0xba, 0xbe, 0x69, 0x68, 0x4d,
	0x4a, 0x9a, 0x6e, 0xe0, 0xaf, 0x97, 0xa6, 0x54, 0xc7, 0x8e, 0x5f, 0xbe, 0xbc, 0xf7, 0xc1, 0x83,
	0xc4, 0x59, 0xa9, 0xbb, 0x23, 0xd9, 0xbc, 0xb5, 0xc6, 0x19, 0x0c, 0xdb, 0x32, 0x3d, 0xc3, 0x6d,
	0xf1, 0xd5, 0xeb, 0x1a, 0x39, 0x2c, 0x0a, 0xa3, 0x1d, 0x69, 0xc7, 0x83, 0x75, 0x90, 0x2d, 0xc5,
	0x84, 0xf8, 0x08, 0xb1, 0x33, 0x4e, 0x36, 0x3b, 0xf5, 0x43, 0x3c, 0x5c, 0x07, 0x59, 0x2c, 0xfe,
	0x1f, 0x70, 0x05, 0x77, 0x96, 0x0e, 0xa4, 0x4e, 0x54, 0xf1, 0x1b, 0x2f, 0xff, 0x78, 0x70, 0x47,
	0xd5, 0x90, 0x96, 0xdf, 0xc4, 0xe7, 0xa3, 0x9b, 0x38, 0x15, 0xb0, 0xdc, 0xb7, 0x8d, 0x91, 0xd5,
	0xce, 0x49, 0xd7, 0x77, 0x43, 0xff, 0x9d, 0xba, 0x4e, 0x7e, 0x92, 0xef, 0xc7, 0x62, 0x42, 0xcc,
	0x60, 0x5e, 0x98, 0x6a, 0x4c, 0x27, 0x9b, 0x87, 0xbc, 0x2d, 0xf3, 0xcb, 0xc9, 0xc1, 0x09, 0xff,
	0xe3, 0x79, 0x0b, 0xec, 0xda, 0xe0, 0x3d, 0x2c, 0xf6, 0xba, 0xd6, 0xe6, 0xac, 0xd9, 0x0c, 0x23,
	0x08, 0x3f, 0x6a, 0x16, 0x20, 0x40, 0xf4, 0x2a, 0x55, 0x43, 0x15, 0x0b, 0x37, 0x2f, 0x90, 0xbc,
	0xf5, 0xe3, 0x14, 0xd9, 0x93, 0x3a, 0x10, 0x3e, 0x41, 0x34, 0xae, 0xc1, 0x78, 0x88, 0xf9, 0xfb,
	0xac, 0xd8, 0x75, 0x37, 0x9d, 0x65, 0x41, 0x19, 0xf9, 0x4b, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x55, 0xf5, 0x6e, 0x8e, 0x5b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GuploadServiceClient is the client API for GuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error)
}

type guploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuploadServiceClient(cc grpc.ClientConnInterface) GuploadServiceClient {
	return &guploadServiceClient{cc}
}

func (c *guploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GuploadService_serviceDesc.Streams[0], "/pb.GuploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceUploadClient{stream}
	return x, nil
}

type GuploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type guploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuploadServiceServer is the server API for GuploadService service.
type GuploadServiceServer interface {
	Upload(GuploadService_UploadServer) error
}

// UnimplementedGuploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuploadServiceServer struct {
}

func (*UnimplementedGuploadServiceServer) Upload(srv GuploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterGuploadServiceServer(s *grpc.Server, srv GuploadServiceServer) {
	s.RegisterService(&_GuploadService_serviceDesc, srv)
}

func _GuploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuploadServiceServer).Upload(&guploadServiceUploadServer{stream})
}

type GuploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type guploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GuploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GuploadService",
	HandlerType: (*GuploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GuploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transfer.proto",
}
