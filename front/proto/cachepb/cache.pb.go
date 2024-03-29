// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cachepb/cache.proto

package cachepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SetTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenRequest) Reset()         { *m = SetTokenRequest{} }
func (m *SetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*SetTokenRequest) ProtoMessage()    {}
func (*SetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{0}
}

func (m *SetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenRequest.Unmarshal(m, b)
}
func (m *SetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenRequest.Marshal(b, m, deterministic)
}
func (m *SetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenRequest.Merge(m, src)
}
func (m *SetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_SetTokenRequest.Size(m)
}
func (m *SetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenRequest proto.InternalMessageInfo

func (m *SetTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SetTokenRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SetTokenResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenResponse) Reset()         { *m = SetTokenResponse{} }
func (m *SetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*SetTokenResponse) ProtoMessage()    {}
func (*SetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{1}
}

func (m *SetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenResponse.Unmarshal(m, b)
}
func (m *SetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenResponse.Marshal(b, m, deterministic)
}
func (m *SetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenResponse.Merge(m, src)
}
func (m *SetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_SetTokenResponse.Size(m)
}
func (m *SetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenResponse proto.InternalMessageInfo

func (m *SetTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetTokenRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{2}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResponse) Reset()         { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()    {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{3}
}

func (m *GetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResponse.Unmarshal(m, b)
}
func (m *GetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResponse.Merge(m, src)
}
func (m *GetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenResponse.Size(m)
}
func (m *GetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResponse proto.InternalMessageInfo

func (m *GetTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteTokenRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenRequest) Reset()         { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()    {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{4}
}

func (m *DeleteTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenRequest.Unmarshal(m, b)
}
func (m *DeleteTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenRequest.Merge(m, src)
}
func (m *DeleteTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenRequest.Size(m)
}
func (m *DeleteTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenRequest proto.InternalMessageInfo

func (m *DeleteTokenRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteTokenResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResponse) Reset()         { *m = DeleteTokenResponse{} }
func (m *DeleteTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResponse) ProtoMessage()    {}
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{5}
}

func (m *DeleteTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResponse.Unmarshal(m, b)
}
func (m *DeleteTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResponse.Merge(m, src)
}
func (m *DeleteTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResponse.Size(m)
}
func (m *DeleteTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResponse proto.InternalMessageInfo

func (m *DeleteTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SetIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetIDRequest) Reset()         { *m = SetIDRequest{} }
func (m *SetIDRequest) String() string { return proto.CompactTextString(m) }
func (*SetIDRequest) ProtoMessage()    {}
func (*SetIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{6}
}

func (m *SetIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIDRequest.Unmarshal(m, b)
}
func (m *SetIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIDRequest.Marshal(b, m, deterministic)
}
func (m *SetIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIDRequest.Merge(m, src)
}
func (m *SetIDRequest) XXX_Size() int {
	return xxx_messageInfo_SetIDRequest.Size(m)
}
func (m *SetIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetIDRequest proto.InternalMessageInfo

func (m *SetIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetIDRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SetIDResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetIDResponse) Reset()         { *m = SetIDResponse{} }
func (m *SetIDResponse) String() string { return proto.CompactTextString(m) }
func (*SetIDResponse) ProtoMessage()    {}
func (*SetIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{7}
}

func (m *SetIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIDResponse.Unmarshal(m, b)
}
func (m *SetIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIDResponse.Marshal(b, m, deterministic)
}
func (m *SetIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIDResponse.Merge(m, src)
}
func (m *SetIDResponse) XXX_Size() int {
	return xxx_messageInfo_SetIDResponse.Size(m)
}
func (m *SetIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetIDResponse proto.InternalMessageInfo

func (m *SetIDResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetIDRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDRequest) Reset()         { *m = GetIDRequest{} }
func (m *GetIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetIDRequest) ProtoMessage()    {}
func (*GetIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{8}
}

func (m *GetIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDRequest.Unmarshal(m, b)
}
func (m *GetIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDRequest.Marshal(b, m, deterministic)
}
func (m *GetIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDRequest.Merge(m, src)
}
func (m *GetIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetIDRequest.Size(m)
}
func (m *GetIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDRequest proto.InternalMessageInfo

func (m *GetIDRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetIDResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDResponse) Reset()         { *m = GetIDResponse{} }
func (m *GetIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetIDResponse) ProtoMessage()    {}
func (*GetIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{9}
}

func (m *GetIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDResponse.Unmarshal(m, b)
}
func (m *GetIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDResponse.Marshal(b, m, deterministic)
}
func (m *GetIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDResponse.Merge(m, src)
}
func (m *GetIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetIDResponse.Size(m)
}
func (m *GetIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDResponse proto.InternalMessageInfo

func (m *GetIDResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteIDRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteIDRequest) Reset()         { *m = DeleteIDRequest{} }
func (m *DeleteIDRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteIDRequest) ProtoMessage()    {}
func (*DeleteIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{10}
}

func (m *DeleteIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteIDRequest.Unmarshal(m, b)
}
func (m *DeleteIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteIDRequest.Marshal(b, m, deterministic)
}
func (m *DeleteIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteIDRequest.Merge(m, src)
}
func (m *DeleteIDRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteIDRequest.Size(m)
}
func (m *DeleteIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteIDRequest proto.InternalMessageInfo

func (m *DeleteIDRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteIDResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteIDResponse) Reset()         { *m = DeleteIDResponse{} }
func (m *DeleteIDResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteIDResponse) ProtoMessage()    {}
func (*DeleteIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd5525426c464b4, []int{11}
}

func (m *DeleteIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteIDResponse.Unmarshal(m, b)
}
func (m *DeleteIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteIDResponse.Marshal(b, m, deterministic)
}
func (m *DeleteIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteIDResponse.Merge(m, src)
}
func (m *DeleteIDResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteIDResponse.Size(m)
}
func (m *DeleteIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteIDResponse proto.InternalMessageInfo

func (m *DeleteIDResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SetTokenRequest)(nil), "cachepb.SetTokenRequest")
	proto.RegisterType((*SetTokenResponse)(nil), "cachepb.SetTokenResponse")
	proto.RegisterType((*GetTokenRequest)(nil), "cachepb.GetTokenRequest")
	proto.RegisterType((*GetTokenResponse)(nil), "cachepb.GetTokenResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "cachepb.DeleteTokenRequest")
	proto.RegisterType((*DeleteTokenResponse)(nil), "cachepb.DeleteTokenResponse")
	proto.RegisterType((*SetIDRequest)(nil), "cachepb.SetIDRequest")
	proto.RegisterType((*SetIDResponse)(nil), "cachepb.SetIDResponse")
	proto.RegisterType((*GetIDRequest)(nil), "cachepb.GetIDRequest")
	proto.RegisterType((*GetIDResponse)(nil), "cachepb.GetIDResponse")
	proto.RegisterType((*DeleteIDRequest)(nil), "cachepb.DeleteIDRequest")
	proto.RegisterType((*DeleteIDResponse)(nil), "cachepb.DeleteIDResponse")
}

func init() { proto.RegisterFile("proto/cachepb/cache.proto", fileDescriptor_dbd5525426c464b4) }

var fileDescriptor_dbd5525426c464b4 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4d, 0x4f, 0xb3, 0x40,
	0x10, 0xc7, 0xfb, 0x92, 0x3e, 0xed, 0x33, 0x16, 0x21, 0xeb, 0x4b, 0x28, 0x9a, 0xd8, 0xac, 0x89,
	0xa9, 0x49, 0x43, 0x8d, 0x5e, 0xf4, 0x58, 0x6d, 0xb2, 0xd1, 0x63, 0xf1, 0xe4, 0xad, 0xa5, 0x13,
	0x25, 0xd5, 0x82, 0x65, 0x35, 0xf1, 0x93, 0xf8, 0x75, 0x0d, 0xb0, 0xc0, 0xb2, 0xa0, 0x9c, 0x60,
	0x66, 0xe7, 0xff, 0xdb, 0x61, 0xfe, 0x03, 0x0c, 0x82, 0xad, 0xcf, 0xfd, 0x89, 0xbb, 0x70, 0x5f,
	0x30, 0x58, 0x26, 0x4f, 0x3b, 0xce, 0x91, 0xae, 0x48, 0xd2, 0x1b, 0xd0, 0x1d, 0xe4, 0x8f, 0xfe,
	0x1a, 0x37, 0x73, 0x7c, 0xff, 0xc0, 0x90, 0x93, 0x7d, 0xe8, 0xf0, 0x28, 0x36, 0x9b, 0xc3, 0xe6,
	0xe8, 0xff, 0x3c, 0x09, 0x88, 0x01, 0xed, 0x35, 0x7e, 0x99, 0xad, 0x38, 0x17, 0xbd, 0xd2, 0x31,
	0x18, 0xb9, 0x34, 0x0c, 0xfc, 0x4d, 0x88, 0xc4, 0x84, 0xee, 0x1b, 0x86, 0xe1, 0xe2, 0x19, 0x85,
	0x3a, 0x0d, 0xe9, 0x29, 0xe8, 0x4c, 0xb9, 0x48, 0x20, 0x9b, 0x39, 0x72, 0x04, 0x06, 0x53, 0x91,
	0x95, 0xed, 0xd0, 0x33, 0x20, 0x33, 0x7c, 0x45, 0x8e, 0x35, 0xc4, 0x09, 0xec, 0x15, 0xea, 0x6a,
	0xfb, 0xbc, 0x80, 0xbe, 0x83, 0xfc, 0x7e, 0x96, 0x22, 0x77, 0xa1, 0xe5, 0xad, 0x44, 0x51, 0xcb,
	0x5b, 0x55, 0xcc, 0xe1, 0x1c, 0x34, 0xa1, 0xa8, 0x85, 0x0f, 0xa1, 0xcf, 0x64, 0x78, 0xb9, 0xdf,
	0x13, 0xd0, 0x58, 0x01, 0xa6, 0xdc, 0x1f, 0xcd, 0x31, 0xf9, 0xa0, 0xbf, 0x28, 0x63, 0x30, 0xf2,
	0xa2, 0xba, 0xae, 0x2e, 0xbf, 0xdb, 0xd0, 0xbf, 0x8b, 0xf6, 0xc1, 0xc1, 0xed, 0xa7, 0xe7, 0x22,
	0x99, 0x42, 0x2f, 0x75, 0x96, 0x98, 0xb6, 0x58, 0x15, 0x5b, 0xd9, 0x13, 0x6b, 0x50, 0x71, 0x92,
	0xdc, 0x45, 0x1b, 0x11, 0x82, 0x95, 0x11, 0xec, 0x57, 0x04, 0x2b, 0x23, 0x1e, 0x60, 0x47, 0xb2,
	0x8e, 0x1c, 0x65, 0xb5, 0x65, 0xe3, 0xad, 0xe3, 0xea, 0xc3, 0x8c, 0x75, 0x0d, 0x9d, 0xd8, 0x23,
	0x72, 0x20, 0x37, 0x9d, 0x8d, 0xd0, 0x3a, 0x54, 0xd3, 0xb2, 0x92, 0x29, 0x4a, 0x56, 0xad, 0x64,
	0x8a, 0x72, 0x0a, 0xbd, 0xd4, 0x04, 0x69, 0x04, 0x8a, 0x79, 0xd2, 0x08, 0x54, 0xc7, 0x68, 0xe3,
	0x56, 0x7f, 0xd2, 0x0a, 0xff, 0xf0, 0xf2, 0x5f, 0x1c, 0x5e, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xb0, 0x02, 0x2d, 0xdb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenResponse, error)
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
	SetID(ctx context.Context, in *SetIDRequest, opts ...grpc.CallOption) (*SetIDResponse, error)
	GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error)
	DeleteID(ctx context.Context, in *DeleteIDRequest, opts ...grpc.CallOption) (*DeleteIDResponse, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenResponse, error) {
	out := new(SetTokenResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/SetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) SetID(ctx context.Context, in *SetIDRequest, opts ...grpc.CallOption) (*SetIDResponse, error) {
	out := new(SetIDResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/SetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error) {
	out := new(GetIDResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteID(ctx context.Context, in *DeleteIDRequest, opts ...grpc.CallOption) (*DeleteIDResponse, error) {
	out := new(DeleteIDResponse)
	err := c.cc.Invoke(ctx, "/cachepb.CacheService/DeleteID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	SetToken(context.Context, *SetTokenRequest) (*SetTokenResponse, error)
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
	SetID(context.Context, *SetIDRequest) (*SetIDResponse, error)
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)
	DeleteID(context.Context, *DeleteIDRequest) (*DeleteIDResponse, error)
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_SetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).SetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/SetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).SetToken(ctx, req.(*SetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_SetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).SetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/SetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).SetID(ctx, req.(*SetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cachepb.CacheService/DeleteID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteID(ctx, req.(*DeleteIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cachepb.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetToken",
			Handler:    _CacheService_SetToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _CacheService_GetToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _CacheService_DeleteToken_Handler,
		},
		{
			MethodName: "SetID",
			Handler:    _CacheService_SetID_Handler,
		},
		{
			MethodName: "GetID",
			Handler:    _CacheService_GetID_Handler,
		},
		{
			MethodName: "DeleteID",
			Handler:    _CacheService_DeleteID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cachepb/cache.proto",
}
