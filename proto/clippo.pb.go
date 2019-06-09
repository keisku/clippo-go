// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/clippo.proto

package clippopb

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

type ArticleURLRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleURLRequest) Reset()         { *m = ArticleURLRequest{} }
func (m *ArticleURLRequest) String() string { return proto.CompactTextString(m) }
func (*ArticleURLRequest) ProtoMessage()    {}
func (*ArticleURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef53a253686b3d8, []int{0}
}

func (m *ArticleURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleURLRequest.Unmarshal(m, b)
}
func (m *ArticleURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleURLRequest.Marshal(b, m, deterministic)
}
func (m *ArticleURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleURLRequest.Merge(m, src)
}
func (m *ArticleURLRequest) XXX_Size() int {
	return xxx_messageInfo_ArticleURLRequest.Size(m)
}
func (m *ArticleURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleURLRequest proto.InternalMessageInfo

func (m *ArticleURLRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ArticleTitleDescriptionImgResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleTitleDescriptionImgResponse) Reset()         { *m = ArticleTitleDescriptionImgResponse{} }
func (m *ArticleTitleDescriptionImgResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleTitleDescriptionImgResponse) ProtoMessage()    {}
func (*ArticleTitleDescriptionImgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef53a253686b3d8, []int{1}
}

func (m *ArticleTitleDescriptionImgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleTitleDescriptionImgResponse.Unmarshal(m, b)
}
func (m *ArticleTitleDescriptionImgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleTitleDescriptionImgResponse.Marshal(b, m, deterministic)
}
func (m *ArticleTitleDescriptionImgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleTitleDescriptionImgResponse.Merge(m, src)
}
func (m *ArticleTitleDescriptionImgResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleTitleDescriptionImgResponse.Size(m)
}
func (m *ArticleTitleDescriptionImgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleTitleDescriptionImgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleTitleDescriptionImgResponse proto.InternalMessageInfo

func (m *ArticleTitleDescriptionImgResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleTitleDescriptionImgResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArticleTitleDescriptionImgResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*ArticleURLRequest)(nil), "clippo.ArticleURLRequest")
	proto.RegisterType((*ArticleTitleDescriptionImgResponse)(nil), "clippo.ArticleTitleDescriptionImgResponse")
}

func init() { proto.RegisterFile("proto/clippo.proto", fileDescriptor_8ef53a253686b3d8) }

var fileDescriptor_8ef53a253686b3d8 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xc9, 0x2c, 0x28, 0xc8, 0xd7, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x3c, 0x25,
	0x55, 0x2e, 0x41, 0xc7, 0xa2, 0x92, 0xcc, 0xe4, 0x9c, 0xd4, 0xd0, 0x20, 0x9f, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xa2, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x10, 0x53, 0xa9, 0x88, 0x4b, 0x09, 0xaa, 0x2c, 0x24, 0xb3, 0x24, 0x27, 0xd5, 0x25,
	0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x33, 0x3f, 0xcf, 0x33, 0x37, 0x3d, 0x28, 0xb5, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0xb5, 0x04, 0x24, 0x0d, 0xd5, 0x09, 0xe1, 0x08,
	0x29, 0x70, 0x71, 0xa7, 0x20, 0xd4, 0x4b, 0x30, 0x81, 0xe5, 0x90, 0x85, 0x40, 0xfa, 0x32, 0x73,
	0x13, 0xd3, 0x53, 0x25, 0x98, 0x21, 0xfa, 0xc0, 0x1c, 0xa3, 0x2a, 0x2e, 0x3e, 0xa8, 0x9d, 0xc1,
	0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x19, 0x5c, 0xb2, 0xee, 0xa9, 0x25, 0xb8, 0x1d, 0x22,
	0x24, 0xa9, 0x07, 0xf5, 0x24, 0x86, 0x9f, 0xa4, 0xb4, 0xd0, 0xa4, 0xf0, 0xf8, 0x43, 0x89, 0xc1,
	0x89, 0x2b, 0x8a, 0x03, 0xa2, 0xbc, 0x20, 0x29, 0x89, 0x0d, 0x1c, 0x62, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xac, 0x46, 0x8a, 0xa3, 0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	GetArticleTitleDescriptionImg(ctx context.Context, in *ArticleURLRequest, opts ...grpc.CallOption) (*ArticleTitleDescriptionImgResponse, error)
}

type articleServiceClient struct {
	cc *grpc.ClientConn
}

func NewArticleServiceClient(cc *grpc.ClientConn) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) GetArticleTitleDescriptionImg(ctx context.Context, in *ArticleURLRequest, opts ...grpc.CallOption) (*ArticleTitleDescriptionImgResponse, error) {
	out := new(ArticleTitleDescriptionImgResponse)
	err := c.cc.Invoke(ctx, "/clippo.ArticleService/GetArticleTitleDescriptionImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	GetArticleTitleDescriptionImg(context.Context, *ArticleURLRequest) (*ArticleTitleDescriptionImgResponse, error)
}

// UnimplementedArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (*UnimplementedArticleServiceServer) GetArticleTitleDescriptionImg(ctx context.Context, req *ArticleURLRequest) (*ArticleTitleDescriptionImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleTitleDescriptionImg not implemented")
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_GetArticleTitleDescriptionImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleTitleDescriptionImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clippo.ArticleService/GetArticleTitleDescriptionImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleTitleDescriptionImg(ctx, req.(*ArticleURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clippo.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticleTitleDescriptionImg",
			Handler:    _ArticleService_GetArticleTitleDescriptionImg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/clippo.proto",
}
