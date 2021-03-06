// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat/chat.proto

package chat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReferralCode         string   `protobuf:"bytes,2,opt,name=referral_code,json=referralCode,proto3" json:"referral_code,omitempty"`
	SponsorId            string   `protobuf:"bytes,3,opt,name=sponsor_id,json=sponsorId,proto3" json:"sponsor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetReferralCode() string {
	if m != nil {
		return m.ReferralCode
	}
	return ""
}

func (m *User) GetSponsorId() string {
	if m != nil {
		return m.SponsorId
	}
	return ""
}

type MessageReq struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	NextId               string   `protobuf:"bytes,3,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	NextTime             int64    `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReq) Reset()         { *m = MessageReq{} }
func (m *MessageReq) String() string { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()    {}
func (*MessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{1}
}

func (m *MessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReq.Unmarshal(m, b)
}
func (m *MessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReq.Marshal(b, m, deterministic)
}
func (m *MessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReq.Merge(m, src)
}
func (m *MessageReq) XXX_Size() int {
	return xxx_messageInfo_MessageReq.Size(m)
}
func (m *MessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReq proto.InternalMessageInfo

func (m *MessageReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MessageReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MessageReq) GetNextId() string {
	if m != nil {
		return m.NextId
	}
	return ""
}

func (m *MessageReq) GetNextTime() int64 {
	if m != nil {
		return m.NextTime
	}
	return 0
}

type MessageResp struct {
	From                 string     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	NextTime             int64      `protobuf:"varint,2,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
	NextId               string     `protobuf:"bytes,3,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	Messages             []*Message `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageResp) Reset()         { *m = MessageResp{} }
func (m *MessageResp) String() string { return proto.CompactTextString(m) }
func (*MessageResp) ProtoMessage()    {}
func (*MessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{2}
}

func (m *MessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResp.Unmarshal(m, b)
}
func (m *MessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResp.Marshal(b, m, deterministic)
}
func (m *MessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResp.Merge(m, src)
}
func (m *MessageResp) XXX_Size() int {
	return xxx_messageInfo_MessageResp.Size(m)
}
func (m *MessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResp proto.InternalMessageInfo

func (m *MessageResp) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MessageResp) GetNextTime() int64 {
	if m != nil {
		return m.NextTime
	}
	return 0
}

func (m *MessageResp) GetNextId() string {
	if m != nil {
		return m.NextId
	}
	return ""
}

func (m *MessageResp) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Time                 int64    `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	Next                 string   `protobuf:"bytes,6,opt,name=next,proto3" json:"next,omitempty"`
	Seen                 bool     `protobuf:"varint,7,opt,name=seen,proto3" json:"seen,omitempty"`
	Received             bool     `protobuf:"varint,8,opt,name=received,proto3" json:"received,omitempty"`
	Event                string   `protobuf:"bytes,9,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *Message) GetSeen() bool {
	if m != nil {
		return m.Seen
	}
	return false
}

func (m *Message) GetReceived() bool {
	if m != nil {
		return m.Received
	}
	return false
}

func (m *Message) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

type ContactResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Unread               int64    `protobuf:"varint,2,opt,name=unread,proto3" json:"unread,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	LastMessage          *Message `protobuf:"bytes,5,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
	ProfilePic           string   `protobuf:"bytes,6,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactResp) Reset()         { *m = ContactResp{} }
func (m *ContactResp) String() string { return proto.CompactTextString(m) }
func (*ContactResp) ProtoMessage()    {}
func (*ContactResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{4}
}

func (m *ContactResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactResp.Unmarshal(m, b)
}
func (m *ContactResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactResp.Marshal(b, m, deterministic)
}
func (m *ContactResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactResp.Merge(m, src)
}
func (m *ContactResp) XXX_Size() int {
	return xxx_messageInfo_ContactResp.Size(m)
}
func (m *ContactResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactResp.DiscardUnknown(m)
}

var xxx_messageInfo_ContactResp proto.InternalMessageInfo

func (m *ContactResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactResp) GetUnread() int64 {
	if m != nil {
		return m.Unread
	}
	return 0
}

func (m *ContactResp) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ContactResp) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ContactResp) GetLastMessage() *Message {
	if m != nil {
		return m.LastMessage
	}
	return nil
}

func (m *ContactResp) GetProfilePic() string {
	if m != nil {
		return m.ProfilePic
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Response struct {
	Contacts             []*ContactResp `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	Messages             *MessageResp   `protobuf:"bytes,2,opt,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{6}
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

func (m *Response) GetContacts() []*ContactResp {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *Response) GetMessages() *MessageResp {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_411570473cc66b06, []int{7}
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

func init() {
	proto.RegisterType((*User)(nil), "chat.User")
	proto.RegisterType((*MessageReq)(nil), "chat.MessageReq")
	proto.RegisterType((*MessageResp)(nil), "chat.MessageResp")
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterType((*ContactResp)(nil), "chat.ContactResp")
	proto.RegisterType((*Error)(nil), "chat.Error")
	proto.RegisterType((*Response)(nil), "chat.Response")
	proto.RegisterType((*Request)(nil), "chat.Request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ChatService service

type ChatServiceClient interface {
	Messages(ctx context.Context, in *MessageReq, opts ...client.CallOption) (*Response, error)
	Contacts(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type chatServiceClient struct {
	c           client.Client
	serviceName string
}

func NewChatServiceClient(serviceName string, c client.Client) ChatServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chat"
	}
	return &chatServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *chatServiceClient) Messages(ctx context.Context, in *MessageReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.Messages", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Contacts(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.Contacts", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatService service

type ChatServiceHandler interface {
	Messages(context.Context, *MessageReq, *Response) error
	Contacts(context.Context, *User, *Response) error
}

func RegisterChatServiceHandler(s server.Server, hdlr ChatServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ChatService{hdlr}, opts...))
}

type ChatService struct {
	ChatServiceHandler
}

func (h *ChatService) Messages(ctx context.Context, in *MessageReq, out *Response) error {
	return h.ChatServiceHandler.Messages(ctx, in, out)
}

func (h *ChatService) Contacts(ctx context.Context, in *User, out *Response) error {
	return h.ChatServiceHandler.Contacts(ctx, in, out)
}

func init() { proto.RegisterFile("proto/chat/chat.proto", fileDescriptor_411570473cc66b06) }

var fileDescriptor_411570473cc66b06 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0xfd, 0x74, 0x26, 0xbb, 0x2b, 0xd6, 0xe2, 0xc3, 0x5a, 0x84, 0xa8, 0xc2, 0xa5, 0x1c,
	0x58, 0xd0, 0x72, 0xe6, 0x54, 0x71, 0xd8, 0x03, 0x08, 0x05, 0xb8, 0x70, 0xa9, 0x42, 0x3c, 0xdd,
	0x5a, 0x6a, 0xec, 0xd4, 0xf6, 0x56, 0x7b, 0xe6, 0x8f, 0x71, 0xe1, 0x87, 0x21, 0x4f, 0x9c, 0xd0,
	0x6e, 0xe1, 0x52, 0xcd, 0xbc, 0x99, 0xbe, 0x99, 0xf7, 0x32, 0x86, 0xc7, 0x8d, 0x35, 0xde, 0xbc,
	0xa9, 0xd6, 0xa5, 0xa7, 0x9f, 0x4b, 0xca, 0xf9, 0x28, 0xc4, 0xf9, 0x77, 0x18, 0x7d, 0x73, 0x68,
	0xf9, 0x19, 0x0c, 0x94, 0x14, 0xc9, 0x2c, 0x99, 0xa7, 0xc5, 0x40, 0x49, 0xfe, 0x12, 0x4e, 0x2d,
	0xae, 0xd0, 0xda, 0x72, 0xb3, 0xac, 0x8c, 0x44, 0x31, 0xa0, 0xd2, 0x49, 0x07, 0x2e, 0x8c, 0x44,
	0xfe, 0x1c, 0xc0, 0x35, 0x46, 0x3b, 0x63, 0x97, 0x4a, 0x8a, 0x21, 0x75, 0xa4, 0x11, 0xb9, 0x96,
	0xf9, 0x0a, 0xe0, 0x23, 0x3a, 0x57, 0xde, 0x60, 0x81, 0xdb, 0x30, 0xc1, 0x9b, 0x6e, 0x82, 0x37,
	0x9c, 0xc3, 0x68, 0x65, 0x4d, 0x1d, 0x89, 0x29, 0xe6, 0x4f, 0x61, 0xaa, 0xf1, 0xce, 0xff, 0x65,
	0x9b, 0x84, 0xf4, 0x5a, 0xf2, 0x67, 0x90, 0x52, 0xc1, 0xab, 0x1a, 0xc5, 0x68, 0x96, 0xcc, 0x87,
	0x05, 0x0b, 0xc0, 0x57, 0x55, 0x63, 0xfe, 0x33, 0x81, 0xac, 0x1f, 0xe4, 0x9a, 0x9e, 0x39, 0xd9,
	0x63, 0x3e, 0x20, 0x18, 0x1c, 0x12, 0xfc, 0x7f, 0xec, 0x2b, 0x60, 0x75, 0x4b, 0xec, 0xc4, 0x68,
	0x36, 0x9c, 0x67, 0x57, 0xa7, 0x97, 0x64, 0x61, 0x37, 0xae, 0x2f, 0xe7, 0xbf, 0x12, 0x98, 0x46,
	0xf4, 0xc8, 0xcc, 0x7f, 0x49, 0x6d, 0xed, 0x18, 0xee, 0xdb, 0xe1, 0xf1, 0xce, 0x93, 0xb8, 0xb4,
	0xa0, 0x98, 0xb0, 0xb0, 0xef, 0x98, 0xf6, 0xa5, 0x38, 0x60, 0x61, 0x39, 0x31, 0x69, 0xfb, 0x74,
	0xec, 0x73, 0x88, 0x5a, 0x4c, 0x67, 0xc9, 0x9c, 0x15, 0x14, 0xf3, 0x0b, 0x60, 0x16, 0x2b, 0x54,
	0x3b, 0x94, 0x82, 0x11, 0xde, 0xe7, 0xfc, 0x11, 0x8c, 0x71, 0x87, 0xda, 0x8b, 0x94, 0x48, 0xda,
	0x24, 0xff, 0x9d, 0x40, 0xb6, 0x30, 0xda, 0x97, 0x95, 0x27, 0x1b, 0xef, 0xab, 0x78, 0x02, 0x93,
	0x5b, 0x6d, 0xb1, 0x94, 0xd1, 0xbf, 0x98, 0x85, 0x2b, 0x58, 0x29, 0xeb, 0xfc, 0x52, 0x97, 0x35,
	0x76, 0x57, 0x40, 0xc8, 0xa7, 0xb2, 0xc6, 0xe0, 0xfc, 0xa6, 0xec, 0xaa, 0xad, 0x3a, 0x16, 0x00,
	0x2a, 0xbe, 0x85, 0x13, 0x2a, 0x46, 0x1b, 0x49, 0xe9, 0x91, 0xc9, 0x59, 0x68, 0xe9, 0xbc, 0x7d,
	0x01, 0x59, 0x63, 0xcd, 0x4a, 0x6d, 0x70, 0xd9, 0xa8, 0x2a, 0xda, 0x00, 0x11, 0xfa, 0xac, 0xaa,
	0xfc, 0x3d, 0x8c, 0x3f, 0x58, 0x6b, 0x6c, 0x70, 0x85, 0x2e, 0x37, 0x28, 0x18, 0x17, 0x14, 0xf3,
	0x19, 0x64, 0x12, 0x5d, 0x65, 0x55, 0xe3, 0x95, 0xd1, 0xf1, 0x83, 0xec, 0x43, 0xf9, 0x1a, 0x58,
	0x50, 0x6f, 0xb4, 0x43, 0xfe, 0x1a, 0x58, 0xd5, 0x1a, 0xe2, 0x44, 0x42, 0x9f, 0xff, 0xbc, 0xdd,
	0x6c, 0xcf, 0xa6, 0xa2, 0x6f, 0x09, 0xed, 0xfd, 0xb5, 0x0c, 0x48, 0xc8, 0xf9, 0xa1, 0x10, 0x6a,
	0xef, 0x2f, 0x26, 0x85, 0x69, 0x81, 0xdb, 0x5b, 0x74, 0xfe, 0xea, 0x06, 0xb2, 0xc5, 0xba, 0xf4,
	0x5f, 0xd0, 0xee, 0x54, 0x85, 0xfc, 0x12, 0x58, 0xfc, 0x8b, 0xe3, 0x0f, 0xef, 0x51, 0x6c, 0x2f,
	0xce, 0x5a, 0xa4, 0xdb, 0x32, 0x7f, 0xc0, 0xe7, 0xc0, 0x16, 0xdd, 0x12, 0xd0, 0x56, 0xc3, 0xa3,
	0x3e, 0xee, 0xfc, 0x31, 0xa1, 0xb7, 0xff, 0xee, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xef,
	0x12, 0xc5, 0x14, 0x04, 0x00, 0x00,
}
