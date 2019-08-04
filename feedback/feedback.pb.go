// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedback.proto

package feedback

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type GetFeedbackReq struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	BookingCode          string   `protobuf:"bytes,2,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackReq) Reset()         { *m = GetFeedbackReq{} }
func (m *GetFeedbackReq) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackReq) ProtoMessage()    {}
func (*GetFeedbackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{1}
}

func (m *GetFeedbackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackReq.Unmarshal(m, b)
}
func (m *GetFeedbackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackReq.Marshal(b, m, deterministic)
}
func (m *GetFeedbackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackReq.Merge(m, src)
}
func (m *GetFeedbackReq) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackReq.Size(m)
}
func (m *GetFeedbackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackReq proto.InternalMessageInfo

func (m *GetFeedbackReq) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *GetFeedbackReq) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type GetFeedbackRes struct {
	PsgFeedback          *PassengerFeedback `protobuf:"bytes,1,opt,name=psg_feedback,json=psgFeedback,proto3" json:"psg_feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFeedbackRes) Reset()         { *m = GetFeedbackRes{} }
func (m *GetFeedbackRes) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackRes) ProtoMessage()    {}
func (*GetFeedbackRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{2}
}

func (m *GetFeedbackRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackRes.Unmarshal(m, b)
}
func (m *GetFeedbackRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackRes.Marshal(b, m, deterministic)
}
func (m *GetFeedbackRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackRes.Merge(m, src)
}
func (m *GetFeedbackRes) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackRes.Size(m)
}
func (m *GetFeedbackRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackRes proto.InternalMessageInfo

func (m *GetFeedbackRes) GetPsgFeedback() *PassengerFeedback {
	if m != nil {
		return m.PsgFeedback
	}
	return nil
}

type AddFeedbackReq struct {
	PsgFeedback          *PassengerFeedback `protobuf:"bytes,1,opt,name=psg_feedback,json=psgFeedback,proto3" json:"psg_feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddFeedbackReq) Reset()         { *m = AddFeedbackReq{} }
func (m *AddFeedbackReq) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackReq) ProtoMessage()    {}
func (*AddFeedbackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{3}
}

func (m *AddFeedbackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackReq.Unmarshal(m, b)
}
func (m *AddFeedbackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackReq.Marshal(b, m, deterministic)
}
func (m *AddFeedbackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackReq.Merge(m, src)
}
func (m *AddFeedbackReq) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackReq.Size(m)
}
func (m *AddFeedbackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackReq proto.InternalMessageInfo

func (m *AddFeedbackReq) GetPsgFeedback() *PassengerFeedback {
	if m != nil {
		return m.PsgFeedback
	}
	return nil
}

type AddFeedbackRes struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackRes) Reset()         { *m = AddFeedbackRes{} }
func (m *AddFeedbackRes) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackRes) ProtoMessage()    {}
func (*AddFeedbackRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{4}
}

func (m *AddFeedbackRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackRes.Unmarshal(m, b)
}
func (m *AddFeedbackRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackRes.Marshal(b, m, deterministic)
}
func (m *AddFeedbackRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackRes.Merge(m, src)
}
func (m *AddFeedbackRes) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackRes.Size(m)
}
func (m *AddFeedbackRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackRes proto.InternalMessageInfo

func (m *AddFeedbackRes) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddFeedbackRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RemoveFeedbackReq struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFeedbackReq) Reset()         { *m = RemoveFeedbackReq{} }
func (m *RemoveFeedbackReq) String() string { return proto.CompactTextString(m) }
func (*RemoveFeedbackReq) ProtoMessage()    {}
func (*RemoveFeedbackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{5}
}

func (m *RemoveFeedbackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFeedbackReq.Unmarshal(m, b)
}
func (m *RemoveFeedbackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFeedbackReq.Marshal(b, m, deterministic)
}
func (m *RemoveFeedbackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFeedbackReq.Merge(m, src)
}
func (m *RemoveFeedbackReq) XXX_Size() int {
	return xxx_messageInfo_RemoveFeedbackReq.Size(m)
}
func (m *RemoveFeedbackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFeedbackReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFeedbackReq proto.InternalMessageInfo

func (m *RemoveFeedbackReq) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type RemoveFeedbackRes struct {
	DeleteFeedback       *AddFeedbackRes `protobuf:"bytes,1,opt,name=delete_feedback,json=deleteFeedback,proto3" json:"delete_feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RemoveFeedbackRes) Reset()         { *m = RemoveFeedbackRes{} }
func (m *RemoveFeedbackRes) String() string { return proto.CompactTextString(m) }
func (*RemoveFeedbackRes) ProtoMessage()    {}
func (*RemoveFeedbackRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{6}
}

func (m *RemoveFeedbackRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFeedbackRes.Unmarshal(m, b)
}
func (m *RemoveFeedbackRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFeedbackRes.Marshal(b, m, deterministic)
}
func (m *RemoveFeedbackRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFeedbackRes.Merge(m, src)
}
func (m *RemoveFeedbackRes) XXX_Size() int {
	return xxx_messageInfo_RemoveFeedbackRes.Size(m)
}
func (m *RemoveFeedbackRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFeedbackRes.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFeedbackRes proto.InternalMessageInfo

func (m *RemoveFeedbackRes) GetDeleteFeedback() *AddFeedbackRes {
	if m != nil {
		return m.DeleteFeedback
	}
	return nil
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "feedback.PassengerFeedback")
	proto.RegisterType((*GetFeedbackReq)(nil), "feedback.GetFeedbackReq")
	proto.RegisterType((*GetFeedbackRes)(nil), "feedback.GetFeedbackRes")
	proto.RegisterType((*AddFeedbackReq)(nil), "feedback.AddFeedbackReq")
	proto.RegisterType((*AddFeedbackRes)(nil), "feedback.AddFeedbackRes")
	proto.RegisterType((*RemoveFeedbackReq)(nil), "feedback.RemoveFeedbackReq")
	proto.RegisterType((*RemoveFeedbackRes)(nil), "feedback.RemoveFeedbackRes")
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor_7b189e8c8330c03e) }

var fileDescriptor_7b189e8c8330c03e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x8a, 0xb5, 0x4e, 0x24, 0xd2, 0x45, 0x24, 0xc4, 0x4b, 0xc8, 0xa9, 0xa7, 0x22,
	0x15, 0xaf, 0x42, 0xac, 0x28, 0xf6, 0x54, 0x16, 0xf1, 0x2a, 0x89, 0x3b, 0x06, 0xa9, 0xed, 0xa6,
	0x9d, 0xe8, 0x2b, 0xfb, 0x1a, 0x92, 0x9a, 0x4d, 0x37, 0xbb, 0x44, 0x90, 0x1e, 0x67, 0x36, 0xf3,
	0xcd, 0xff, 0xcf, 0x4c, 0xc0, 0x7f, 0x43, 0x94, 0x59, 0xfa, 0xba, 0x18, 0x17, 0x1b, 0x55, 0x2a,
	0x3e, 0xd0, 0x71, 0x4c, 0x30, 0x9c, 0xa7, 0x44, 0xb8, 0xca, 0x71, 0x73, 0x5f, 0x27, 0x79, 0x04,
	0x5e, 0xa6, 0xd4, 0xe2, 0x7d, 0x95, 0x4f, 0x95, 0xc4, 0x80, 0x45, 0x6c, 0x74, 0x2c, 0xcc, 0x54,
	0xf5, 0x45, 0xa1, 0xcb, 0x1e, 0xef, 0x82, 0x5e, 0xc4, 0x46, 0x87, 0xc2, 0x4c, 0xf1, 0x10, 0x9a,
	0x26, 0xc1, 0xc1, 0x16, 0xb0, 0x6b, 0xfa, 0x04, 0xfe, 0x03, 0x96, 0xba, 0x9d, 0xc0, 0xb5, 0xcd,
	0x63, 0x2e, 0xcf, 0xd2, 0xd4, 0x73, 0x34, 0xc5, 0x73, 0x8b, 0x4a, 0xfc, 0x06, 0x4e, 0x0a, 0xca,
	0x5f, 0x1a, 0x1d, 0x15, 0xd6, 0x9b, 0x5c, 0x8c, 0x9b, 0x69, 0x38, 0xd6, 0x85, 0x57, 0x50, 0xae,
	0x83, 0x8a, 0x98, 0x48, 0x69, 0xea, 0xdc, 0x97, 0x78, 0x6b, 0x11, 0x89, 0x9f, 0x43, 0x9f, 0xca,
	0xb4, 0xfc, 0xa4, 0xda, 0x74, 0x1d, 0xf1, 0x00, 0x8e, 0x96, 0x48, 0x94, 0xe6, 0xda, 0xab, 0x0e,
	0xe3, 0x6b, 0x18, 0x0a, 0x5c, 0xaa, 0x2f, 0xfc, 0xd7, 0x00, 0xe3, 0x67, 0xb7, 0x8c, 0x78, 0x02,
	0xa7, 0x12, 0x3f, 0xb0, 0x44, 0xdb, 0x52, 0xb0, 0xb3, 0xd4, 0x16, 0x2c, 0xfc, 0xdf, 0x02, 0x9d,
	0x9a, 0x7c, 0x33, 0x18, 0x34, 0x97, 0x33, 0x83, 0xb3, 0x44, 0x4a, 0xf7, 0xa2, 0xba, 0x70, 0xeb,
	0xb0, 0xb3, 0x11, 0x9f, 0x82, 0x67, 0xec, 0xd3, 0x44, 0xb4, 0x8f, 0x27, 0xec, 0x7a, 0xa1, 0x4b,
	0xc6, 0x67, 0xe0, 0xb7, 0x5d, 0x73, 0x63, 0x59, 0xce, 0x18, 0xc3, 0x3f, 0x1e, 0x29, 0xeb, 0x6f,
	0x7f, 0x9e, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x39, 0x00, 0xa3, 0x4e, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedbackClient is the client API for Feedback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedbackClient interface {
	AddPassengerFeedback(ctx context.Context, in *AddFeedbackReq, opts ...grpc.CallOption) (*AddFeedbackRes, error)
	GetFeedback(ctx context.Context, in *GetFeedbackReq, opts ...grpc.CallOption) (Feedback_GetFeedbackClient, error)
	RemoveFeedback(ctx context.Context, in *RemoveFeedbackReq, opts ...grpc.CallOption) (*RemoveFeedbackRes, error)
}

type feedbackClient struct {
	cc *grpc.ClientConn
}

func NewFeedbackClient(cc *grpc.ClientConn) FeedbackClient {
	return &feedbackClient{cc}
}

func (c *feedbackClient) AddPassengerFeedback(ctx context.Context, in *AddFeedbackReq, opts ...grpc.CallOption) (*AddFeedbackRes, error) {
	out := new(AddFeedbackRes)
	err := c.cc.Invoke(ctx, "/feedback.Feedback/AddPassengerFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackClient) GetFeedback(ctx context.Context, in *GetFeedbackReq, opts ...grpc.CallOption) (Feedback_GetFeedbackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Feedback_serviceDesc.Streams[0], "/feedback.Feedback/GetFeedback", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackGetFeedbackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Feedback_GetFeedbackClient interface {
	Recv() (*GetFeedbackRes, error)
	grpc.ClientStream
}

type feedbackGetFeedbackClient struct {
	grpc.ClientStream
}

func (x *feedbackGetFeedbackClient) Recv() (*GetFeedbackRes, error) {
	m := new(GetFeedbackRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackClient) RemoveFeedback(ctx context.Context, in *RemoveFeedbackReq, opts ...grpc.CallOption) (*RemoveFeedbackRes, error) {
	out := new(RemoveFeedbackRes)
	err := c.cc.Invoke(ctx, "/feedback.Feedback/RemoveFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// feedbackServer is the server API for Feedback service.
type feedbackServer interface {
	AddPassengerFeedback(context.Context, *AddFeedbackReq) (*AddFeedbackRes, error)
	GetFeedback(*GetFeedbackReq, Feedback_GetFeedbackServer) error
	RemoveFeedback(context.Context, *RemoveFeedbackReq) (*RemoveFeedbackRes, error)
}

// UnimplementedFeedbackServer can be embedded to have forward compatible implementations.
type UnimplementedFeedbackServer struct {
}

func (*UnimplementedFeedbackServer) AddPassengerFeedback(ctx context.Context, req *AddFeedbackReq) (*AddFeedbackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassengerFeedback not implemented")
}
func (*UnimplementedFeedbackServer) GetFeedback(req *GetFeedbackReq, srv Feedback_GetFeedbackServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFeedback not implemented")
}
func (*UnimplementedFeedbackServer) RemoveFeedback(ctx context.Context, req *RemoveFeedbackReq) (*RemoveFeedbackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFeedback not implemented")
}

func RegisterFeedbackServer(s *grpc.Server, srv feedbackServer) {
	s.RegisterService(&_Feedback_serviceDesc, srv)
}

func _Feedback_AddPassengerFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(feedbackServer).AddPassengerFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.Feedback/AddPassengerFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(feedbackServer).AddPassengerFeedback(ctx, req.(*AddFeedbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feedback_GetFeedback_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFeedbackReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(feedbackServer).GetFeedback(m, &feedbackGetFeedbackServer{stream})
}

type Feedback_GetFeedbackServer interface {
	Send(*GetFeedbackRes) error
	grpc.ServerStream
}

type feedbackGetFeedbackServer struct {
	grpc.ServerStream
}

func (x *feedbackGetFeedbackServer) Send(m *GetFeedbackRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Feedback_RemoveFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFeedbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(feedbackServer).RemoveFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.Feedback/RemoveFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(feedbackServer).RemoveFeedback(ctx, req.(*RemoveFeedbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Feedback_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedback.Feedback",
	HandlerType: (*feedbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPassengerFeedback",
			Handler:    _Feedback_AddPassengerFeedback_Handler,
		},
		{
			MethodName: "RemoveFeedback",
			Handler:    _Feedback_RemoveFeedback_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFeedback",
			Handler:       _Feedback_GetFeedback_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "feedback.proto",
}
