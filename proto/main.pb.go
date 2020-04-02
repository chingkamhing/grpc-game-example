// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/main.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MoveRequest_Direction int32

const (
	MoveRequest_UP    MoveRequest_Direction = 0
	MoveRequest_DOWN  MoveRequest_Direction = 1
	MoveRequest_LEFT  MoveRequest_Direction = 2
	MoveRequest_RIGHT MoveRequest_Direction = 3
	MoveRequest_STOP  MoveRequest_Direction = 4
)

var MoveRequest_Direction_name = map[int32]string{
	0: "UP",
	1: "DOWN",
	2: "LEFT",
	3: "RIGHT",
	4: "STOP",
}

var MoveRequest_Direction_value = map[string]int32{
	"UP":    0,
	"DOWN":  1,
	"LEFT":  2,
	"RIGHT": 3,
	"STOP":  4,
}

func (x MoveRequest_Direction) String() string {
	return proto.EnumName(MoveRequest_Direction_name, int32(x))
}

func (MoveRequest_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{3, 0}
}

type Coordinate struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{0}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinate) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ConnectRequest struct {
	Player               string   `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type ConnectResponse struct {
	Success              bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	StartPosition        *Coordinate `protobuf:"bytes,3,opt,name=startPosition,proto3" json:"startPosition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{2}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ConnectResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ConnectResponse) GetStartPosition() *Coordinate {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

type MoveRequest struct {
	Direction            MoveRequest_Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=proto.MoveRequest_Direction" json:"direction,omitempty"`
	Time                 *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MoveRequest) Reset()         { *m = MoveRequest{} }
func (m *MoveRequest) String() string { return proto.CompactTextString(m) }
func (*MoveRequest) ProtoMessage()    {}
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{3}
}

func (m *MoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveRequest.Unmarshal(m, b)
}
func (m *MoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveRequest.Marshal(b, m, deterministic)
}
func (m *MoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveRequest.Merge(m, src)
}
func (m *MoveRequest) XXX_Size() int {
	return xxx_messageInfo_MoveRequest.Size(m)
}
func (m *MoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveRequest proto.InternalMessageInfo

func (m *MoveRequest) GetDirection() MoveRequest_Direction {
	if m != nil {
		return m.Direction
	}
	return MoveRequest_UP
}

func (m *MoveRequest) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type ChangePositionResponse struct {
	Position             *Coordinate          `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Player               string               `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChangePositionResponse) Reset()         { *m = ChangePositionResponse{} }
func (m *ChangePositionResponse) String() string { return proto.CompactTextString(m) }
func (*ChangePositionResponse) ProtoMessage()    {}
func (*ChangePositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{4}
}

func (m *ChangePositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePositionResponse.Unmarshal(m, b)
}
func (m *ChangePositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePositionResponse.Marshal(b, m, deterministic)
}
func (m *ChangePositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePositionResponse.Merge(m, src)
}
func (m *ChangePositionResponse) XXX_Size() int {
	return xxx_messageInfo_ChangePositionResponse.Size(m)
}
func (m *ChangePositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePositionResponse proto.InternalMessageInfo

func (m *ChangePositionResponse) GetPosition() *Coordinate {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *ChangePositionResponse) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *ChangePositionResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type Request struct {
	// Types that are valid to be assigned to Message:
	//	*Request_Connect
	//	*Request_Move
	Message              isRequest_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{5}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type isRequest_Message interface {
	isRequest_Message()
}

type Request_Connect struct {
	Connect *ConnectRequest `protobuf:"bytes,2,opt,name=connect,proto3,oneof"`
}

type Request_Move struct {
	Move *MoveRequest `protobuf:"bytes,3,opt,name=move,proto3,oneof"`
}

func (*Request_Connect) isRequest_Message() {}

func (*Request_Move) isRequest_Message() {}

func (m *Request) GetMessage() isRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Request) GetConnect() *ConnectRequest {
	if x, ok := m.GetMessage().(*Request_Connect); ok {
		return x.Connect
	}
	return nil
}

func (m *Request) GetMove() *MoveRequest {
	if x, ok := m.GetMessage().(*Request_Move); ok {
		return x.Move
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_Connect)(nil),
		(*Request_Move)(nil),
	}
}

type Response struct {
	// Types that are valid to be assigned to Message:
	//	*Response_Connect
	//	*Response_ChangePosition
	Message              isResponse_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type isResponse_Message interface {
	isResponse_Message()
}

type Response_Connect struct {
	Connect *ConnectResponse `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type Response_ChangePosition struct {
	ChangePosition *ChangePositionResponse `protobuf:"bytes,2,opt,name=change_position,json=changePosition,proto3,oneof"`
}

func (*Response_Connect) isResponse_Message() {}

func (*Response_ChangePosition) isResponse_Message() {}

func (m *Response) GetMessage() isResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Response) GetConnect() *ConnectResponse {
	if x, ok := m.GetMessage().(*Response_Connect); ok {
		return x.Connect
	}
	return nil
}

func (m *Response) GetChangePosition() *ChangePositionResponse {
	if x, ok := m.GetMessage().(*Response_ChangePosition); ok {
		return x.ChangePosition
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Connect)(nil),
		(*Response_ChangePosition)(nil),
	}
}

func init() {
	proto.RegisterEnum("proto.MoveRequest_Direction", MoveRequest_Direction_name, MoveRequest_Direction_value)
	proto.RegisterType((*Coordinate)(nil), "proto.Coordinate")
	proto.RegisterType((*ConnectRequest)(nil), "proto.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "proto.ConnectResponse")
	proto.RegisterType((*MoveRequest)(nil), "proto.MoveRequest")
	proto.RegisterType((*ChangePositionResponse)(nil), "proto.ChangePositionResponse")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() {
	proto.RegisterFile("proto/main.proto", fileDescriptor_098391ad7281b52b)
}

var fileDescriptor_098391ad7281b52b = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xe3, 0x36, 0xfd, 0xc9, 0x2d, 0xb4, 0xc1, 0x82, 0xaa, 0xaa, 0x40, 0x8c, 0xb2, 0xca,
	0x86, 0x14, 0xc2, 0x62, 0x24, 0xc4, 0x8a, 0x0e, 0x4c, 0x90, 0x80, 0xa9, 0x3c, 0x45, 0x2c, 0x51,
	0x26, 0x63, 0x4a, 0xa4, 0x49, 0x9c, 0xb1, 0xdd, 0xd1, 0x94, 0x97, 0x60, 0xc1, 0xc3, 0xf0, 0x7a,
	0x28, 0x76, 0xec, 0xa6, 0xfc, 0x48, 0xac, 0xea, 0xe3, 0x1e, 0xdf, 0x7b, 0xee, 0x77, 0x03, 0x7e,
	0xc5, 0x99, 0x64, 0x8b, 0x22, 0xcd, 0xcb, 0x48, 0x1d, 0x71, 0x4f, 0xfd, 0xcc, 0x1f, 0x6f, 0x18,
	0xdb, 0x5c, 0xd1, 0x85, 0x52, 0x17, 0xdb, 0x2f, 0x0b, 0x99, 0x17, 0x54, 0xc8, 0xb4, 0xa8, 0xb4,
	0x2f, 0x08, 0x01, 0x96, 0x8c, 0xf1, 0xcb, 0xbc, 0x4c, 0x25, 0xc5, 0x77, 0x00, 0xdd, 0xce, 0xd0,
	0x11, 0x0a, 0x7b, 0x04, 0xdd, 0xd6, 0x6a, 0x37, 0xeb, 0x68, 0xb5, 0x0b, 0x42, 0x18, 0x2f, 0x59,
	0x59, 0xd2, 0x4c, 0x12, 0x7a, 0xbd, 0xa5, 0x42, 0xe2, 0x29, 0xf4, 0xab, 0xab, 0x74, 0x47, 0xb9,
	0x7a, 0xe2, 0x91, 0x46, 0x05, 0xdf, 0x60, 0x62, 0x9d, 0xa2, 0x62, 0xa5, 0xa0, 0x78, 0x06, 0x03,
	0xb1, 0xcd, 0x32, 0x2a, 0x84, 0xf2, 0x0e, 0x89, 0x91, 0xf8, 0x3e, 0xf4, 0x28, 0xe7, 0x8c, 0xab,
	0x46, 0x1e, 0xd1, 0x02, 0x1f, 0xc3, 0x5d, 0x21, 0x53, 0x2e, 0x57, 0x4c, 0xe4, 0x32, 0x67, 0xe5,
	0xac, 0x7b, 0x84, 0xc2, 0x51, 0x7c, 0x4f, 0xa7, 0x8e, 0xf6, 0x91, 0xc9, 0xa1, 0x2f, 0xf8, 0x89,
	0x60, 0xf4, 0x9e, 0xdd, 0x50, 0x93, 0xf1, 0x05, 0x78, 0x97, 0x39, 0xa7, 0x99, 0x2a, 0x52, 0xb7,
	0x1e, 0xc7, 0x0f, 0x9b, 0x22, 0x2d, 0x5b, 0x74, 0x62, 0x3c, 0x64, 0x6f, 0xc7, 0x11, 0xb8, 0x35,
	0x2e, 0x95, 0x6c, 0x14, 0xcf, 0x23, 0xcd, 0x32, 0x32, 0x2c, 0xa3, 0xb5, 0x61, 0x49, 0x94, 0x2f,
	0x78, 0x09, 0x9e, 0xad, 0x83, 0xfb, 0xd0, 0xf9, 0xb8, 0xf2, 0x1d, 0x3c, 0x04, 0xf7, 0xe4, 0xec,
	0xd3, 0x07, 0x1f, 0xd5, 0xa7, 0x77, 0xaf, 0xdf, 0xac, 0xfd, 0x0e, 0xf6, 0xa0, 0x47, 0xde, 0x9e,
	0x26, 0x6b, 0xbf, 0x5b, 0x5f, 0x9e, 0xaf, 0xcf, 0x56, 0xbe, 0x1b, 0x7c, 0x47, 0x30, 0x5d, 0x7e,
	0x4d, 0xcb, 0x0d, 0x35, 0xc3, 0x58, 0x7a, 0x4f, 0x60, 0x58, 0x19, 0x10, 0xe8, 0x5f, 0x20, 0xac,
	0xa5, 0xb5, 0x97, 0x4e, 0x7b, 0x2f, 0x76, 0x9e, 0xee, 0x7f, 0xce, 0x73, 0x0d, 0x03, 0x83, 0xf1,
	0x19, 0x0c, 0x32, 0xbd, 0xd2, 0x86, 0xc6, 0x03, 0x1b, 0xa0, 0xfd, 0x49, 0x24, 0x0e, 0x31, 0x3e,
	0x1c, 0x82, 0x5b, 0xb0, 0x1b, 0xd3, 0x0d, 0xff, 0x09, 0x3d, 0x71, 0x88, 0x72, 0xbc, 0xf2, 0x60,
	0x50, 0x50, 0x21, 0xd2, 0x0d, 0x0d, 0x7e, 0x20, 0x18, 0xda, 0xb1, 0xe3, 0x7d, 0x53, 0x3d, 0xf5,
	0xf4, 0xf7, 0xa6, 0xda, 0xd8, 0xee, 0x9a, 0xc0, 0x24, 0x53, 0x10, 0x3f, 0x5b, 0x62, 0x3a, 0xf0,
	0x23, 0xf3, 0xf6, 0xaf, 0x88, 0x13, 0x87, 0x8c, 0xb3, 0x83, 0x7f, 0x5a, 0xa9, 0xe2, 0x63, 0x70,
	0x4f, 0xd3, 0x82, 0xe2, 0x05, 0xf4, 0xcf, 0x25, 0xa7, 0x69, 0x81, 0xc7, 0x4d, 0xb5, 0x66, 0x94,
	0xf9, 0xc4, 0x6a, 0x5d, 0x2f, 0x70, 0x42, 0xf4, 0x14, 0x5d, 0xf4, 0xd5, 0xed, 0xf3, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0xa7, 0x0a, 0x32, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Game_serviceDesc.Streams[0], "/proto.Game/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameStreamClient{stream}
	return x, nil
}

type Game_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type gameStreamClient struct {
	grpc.ClientStream
}

func (x *gameStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	Stream(Game_StreamServer) error
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) Stream(srv Game_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Stream(&gameStreamServer{stream})
}

type Game_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type gameStreamServer struct {
	grpc.ServerStream
}

func (x *gameStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Game_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/main.proto",
}
