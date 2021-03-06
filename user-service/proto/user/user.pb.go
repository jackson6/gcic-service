// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type Id struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MemberId             string   `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Initial              string   `protobuf:"bytes,6,opt,name=initial,proto3" json:"initial,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Trn                  string   `protobuf:"bytes,8,opt,name=trn,proto3" json:"trn,omitempty"`
	WorkPhone            string   `protobuf:"bytes,9,opt,name=work_phone,json=workPhone,proto3" json:"work_phone,omitempty"`
	HomePhone            string   `protobuf:"bytes,10,opt,name=home_phone,json=homePhone,proto3" json:"home_phone,omitempty"`
	CellPhone            string   `protobuf:"bytes,11,opt,name=cell_phone,json=cellPhone,proto3" json:"cell_phone,omitempty"`
	SponsorId            string   `protobuf:"bytes,12,opt,name=sponsor_id,json=sponsorId,proto3" json:"sponsor_id,omitempty"`
	Address              string   `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	Address2             string   `protobuf:"bytes,14,opt,name=address2,proto3" json:"address2,omitempty"`
	Parish               string   `protobuf:"bytes,15,opt,name=parish,proto3" json:"parish,omitempty"`
	Country              string   `protobuf:"bytes,16,opt,name=country,proto3" json:"country,omitempty"`
	Question             string   `protobuf:"bytes,17,opt,name=question,proto3" json:"question,omitempty"`
	Answer               string   `protobuf:"bytes,18,opt,name=answer,proto3" json:"answer,omitempty"`
	Dob                  int64    `protobuf:"varint,19,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender               string   `protobuf:"bytes,20,opt,name=gender,proto3" json:"gender,omitempty"`
	StripeId             string   `protobuf:"bytes,21,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	PlanId               string   `protobuf:"bytes,22,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	ReferralCode         string   `protobuf:"bytes,23,opt,name=referral_code,json=referralCode,proto3" json:"referral_code,omitempty"`
	MembershipExp        int64    `protobuf:"varint,24,opt,name=membership_exp,json=membershipExp,proto3" json:"membership_exp,omitempty"`
	ProfilePic           string   `protobuf:"bytes,25,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
	Count                int32    `protobuf:"varint,26,opt,name=count,proto3" json:"count,omitempty"`
	Level                int32    `protobuf:"varint,27,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
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

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetInitial() string {
	if m != nil {
		return m.Initial
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetTrn() string {
	if m != nil {
		return m.Trn
	}
	return ""
}

func (m *User) GetWorkPhone() string {
	if m != nil {
		return m.WorkPhone
	}
	return ""
}

func (m *User) GetHomePhone() string {
	if m != nil {
		return m.HomePhone
	}
	return ""
}

func (m *User) GetCellPhone() string {
	if m != nil {
		return m.CellPhone
	}
	return ""
}

func (m *User) GetSponsorId() string {
	if m != nil {
		return m.SponsorId
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetAddress2() string {
	if m != nil {
		return m.Address2
	}
	return ""
}

func (m *User) GetParish() string {
	if m != nil {
		return m.Parish
	}
	return ""
}

func (m *User) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *User) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *User) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *User) GetDob() int64 {
	if m != nil {
		return m.Dob
	}
	return 0
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *User) GetReferralCode() string {
	if m != nil {
		return m.ReferralCode
	}
	return ""
}

func (m *User) GetMembershipExp() int64 {
	if m != nil {
		return m.MembershipExp
	}
	return 0
}

func (m *User) GetProfilePic() string {
	if m != nil {
		return m.ProfilePic
	}
	return ""
}

func (m *User) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *User) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Request struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	SaveCard             bool     `protobuf:"varint,3,opt,name=save_card,json=saveCard,proto3" json:"save_card,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
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

func (m *Request) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Request) GetSaveCard() bool {
	if m != nil {
		return m.SaveCard
	}
	return false
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
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

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
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
	return fileDescriptor_9b283a848145d6b7, []int{5}
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

func init() {
	proto.RegisterType((*Id)(nil), "user.Id")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Error)(nil), "user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	GetUsers(ctx context.Context, in *Id, opts ...client.CallOption) (*Response, error)
	GetByEmail(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetByMemberId(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetUserReferral(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetReferred(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *Id, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUsers", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByEmail(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetByEmail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByMemberId(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetByMemberId", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserReferral(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUserReferral", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetReferred(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetReferred", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *Request, *Response) error
	Update(context.Context, *User, *Response) error
	Delete(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	All(context.Context, *Request, *Response) error
	ValidateToken(context.Context, *Token, *Token) error
	GetUsers(context.Context, *Id, *Response) error
	GetByEmail(context.Context, *User, *Response) error
	GetByMemberId(context.Context, *User, *Response) error
	GetUserReferral(context.Context, *User, *Response) error
	GetReferred(context.Context, *User, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Update(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *UserService) Delete(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) All(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.All(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func (h *UserService) GetUsers(ctx context.Context, in *Id, out *Response) error {
	return h.UserServiceHandler.GetUsers(ctx, in, out)
}

func (h *UserService) GetByEmail(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetByEmail(ctx, in, out)
}

func (h *UserService) GetByMemberId(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetByMemberId(ctx, in, out)
}

func (h *UserService) GetUserReferral(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetUserReferral(ctx, in, out)
}

func (h *UserService) GetReferred(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetReferred(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x6f, 0xe3, 0x44,
	0x10, 0x26, 0x71, 0x9c, 0x38, 0xe3, 0x4b, 0xee, 0x58, 0x7a, 0x77, 0x4b, 0x4e, 0x40, 0x48, 0x05,
	0x0a, 0x9c, 0xb8, 0x93, 0xc2, 0x33, 0x0f, 0x50, 0xaa, 0x28, 0x0f, 0xa0, 0x93, 0xa1, 0x88, 0x07,
	0xa4, 0x68, 0xeb, 0x9d, 0x92, 0x55, 0x1d, 0xdb, 0xac, 0xdd, 0xb4, 0xfd, 0x3f, 0xf8, 0xef, 0x78,
	0xe4, 0x1f, 0x41, 0x33, 0x63, 0xb7, 0x11, 0xea, 0x35, 0x2f, 0xd5, 0x7e, 0xdf, 0x37, 0xfb, 0xcd,
	0x8f, 0xcc, 0xba, 0xf0, 0xbc, 0xf4, 0x45, 0x5d, 0xbc, 0xbd, 0xaa, 0xd0, 0xf3, 0x9f, 0x37, 0x8c,
	0x55, 0x8f, 0xce, 0xb3, 0x23, 0xe8, 0xae, 0xac, 0x1a, 0x43, 0xd7, 0x59, 0xdd, 0x99, 0x06, 0xf3,
	0x61, 0xd2, 0x75, 0x76, 0xf6, 0x4f, 0x08, 0xbd, 0xb3, 0x0a, 0xfd, 0x9d, 0xd0, 0x11, 0x41, 0xbd,
	0x84, 0x01, 0x5d, 0x5b, 0x3b, 0xab, 0xbb, 0x4c, 0xf6, 0x09, 0xae, 0xac, 0x7a, 0x05, 0xc3, 0x2d,
	0x6e, 0xcf, 0x45, 0x0a, 0x58, 0x8a, 0x84, 0x58, 0x59, 0xf5, 0x09, 0xc0, 0x85, 0xf3, 0x55, 0xbd,
	0xce, 0xcd, 0x16, 0x75, 0x8f, 0xd5, 0x21, 0x33, 0x3f, 0x9b, 0x2d, 0xd2, 0xdd, 0xcc, 0xb4, 0x6a,
	0x28, 0x77, 0x89, 0x60, 0x51, 0xc3, 0xc0, 0xe5, 0xae, 0x76, 0x26, 0xd3, 0x7d, 0x96, 0x5a, 0xa8,
	0x8e, 0x20, 0xc4, 0xad, 0x71, 0x99, 0x1e, 0x30, 0x2f, 0x40, 0x3d, 0x83, 0xa0, 0xf6, 0xb9, 0x8e,
	0x98, 0xa3, 0x23, 0x65, 0xbf, 0x2e, 0xfc, 0xe5, 0xba, 0xdc, 0x14, 0x39, 0xea, 0xa1, 0x64, 0x27,
	0xe6, 0x1d, 0x11, 0x24, 0x6f, 0x8a, 0x2d, 0x36, 0x32, 0x88, 0x4c, 0xcc, 0x9d, 0x9c, 0x62, 0x96,
	0x35, 0x72, 0x2c, 0x32, 0x31, 0x77, 0x72, 0x55, 0x16, 0x79, 0x55, 0x70, 0xe3, 0x4f, 0x44, 0x6e,
	0x98, 0x95, 0xa5, 0xea, 0x8d, 0xb5, 0x1e, 0xab, 0x4a, 0x8f, 0xa4, 0xfa, 0x06, 0xaa, 0x09, 0x44,
	0xcd, 0x71, 0xa1, 0xc7, 0xd2, 0x73, 0x8b, 0xd5, 0x0b, 0xe8, 0x97, 0xc6, 0xbb, 0x6a, 0xa3, 0x9f,
	0xca, 0x90, 0x05, 0x91, 0x5b, 0x5a, 0x5c, 0xe5, 0xb5, 0xbf, 0xd5, 0xcf, 0xc4, 0xad, 0x81, 0xe4,
	0xf6, 0xd7, 0x15, 0x56, 0xb5, 0x2b, 0x72, 0xfd, 0xa1, 0xb8, 0xb5, 0x98, 0xdc, 0x4c, 0x5e, 0x5d,
	0xa3, 0xd7, 0x4a, 0xdc, 0x04, 0xd1, 0xa4, 0x6c, 0x71, 0xae, 0x3f, 0x9a, 0x76, 0xe6, 0x41, 0x42,
	0x47, 0x8a, 0xfc, 0x13, 0x73, 0x8b, 0x5e, 0x1f, 0x49, 0xa4, 0x20, 0xfa, 0x81, 0xaa, 0xda, 0xbb,
	0x12, 0xa9, 0xc7, 0xe7, 0x62, 0x2f, 0xc4, 0x8a, 0x57, 0xa2, 0xcc, 0x4c, 0x4e, 0xd2, 0x8b, 0xa6,
	0xda, 0xcc, 0xe4, 0x2b, 0xab, 0x8e, 0x61, 0xe4, 0xf1, 0x02, 0xbd, 0x37, 0xd9, 0x3a, 0x2d, 0x2c,
	0xea, 0x97, 0x2c, 0x3f, 0x69, 0xc9, 0x93, 0xc2, 0xa2, 0xfa, 0x02, 0xc6, 0xb2, 0x26, 0xd5, 0xc6,
	0x95, 0x6b, 0xbc, 0x29, 0xb5, 0xe6, 0x7a, 0x46, 0xf7, 0xec, 0xe9, 0x4d, 0xa9, 0x3e, 0x83, 0xb8,
	0xf4, 0xc5, 0x85, 0xcb, 0x70, 0x5d, 0xba, 0x54, 0x7f, 0xcc, 0x4e, 0xd0, 0x50, 0xef, 0x5c, 0x4a,
	0xcb, 0xc0, 0xb3, 0xd0, 0x93, 0x69, 0x67, 0x1e, 0x26, 0x02, 0x88, 0xcd, 0x70, 0x87, 0x99, 0x7e,
	0x25, 0x2c, 0x83, 0xd9, 0xef, 0x10, 0xfe, 0x5a, 0x5c, 0x62, 0x4e, 0x72, 0x4d, 0x87, 0x66, 0xc1,
	0x05, 0x10, 0xbb, 0x33, 0x59, 0xb3, 0xe1, 0x51, 0x22, 0x40, 0x1d, 0x43, 0x1f, 0xbd, 0x2f, 0x7c,
	0xa5, 0x83, 0x69, 0x30, 0x8f, 0x17, 0xf1, 0x1b, 0x7e, 0x4b, 0xa7, 0xc4, 0x25, 0x8d, 0x34, 0xfb,
	0x03, 0x06, 0x09, 0xf2, 0xe0, 0xd5, 0xa7, 0xc0, 0x0f, 0x8c, 0xad, 0xe3, 0x05, 0x48, 0x34, 0xbd,
	0xa9, 0x84, 0xf9, 0xfb, 0xdc, 0xdd, 0xfd, 0xdc, 0x34, 0x69, 0xb3, 0xc3, 0x75, 0x6a, 0xbc, 0x3c,
	0xa3, 0x28, 0x89, 0x88, 0x38, 0x31, 0xde, 0xce, 0xfe, 0xee, 0x40, 0x94, 0x20, 0x2f, 0x17, 0x1e,
	0xf4, 0xff, 0x1c, 0x42, 0x2e, 0x8a, 0xfd, 0xff, 0x57, 0xae, 0x28, 0x6a, 0x0a, 0x21, 0x91, 0x6d,
	0x47, 0xfb, 0x1e, 0x22, 0xec, 0x35, 0xdd, 0x7b, 0x7f, 0xd3, 0xdf, 0x41, 0xc8, 0x84, 0x52, 0xd0,
	0xe3, 0xdf, 0xb9, 0xc3, 0xc3, 0xe6, 0xb3, 0x9a, 0x42, 0x6c, 0xb1, 0x4a, 0xbd, 0x2b, 0x79, 0x37,
	0xa5, 0xd9, 0x7d, 0x6a, 0xf1, 0x6f, 0x00, 0x31, 0xe5, 0xfc, 0x05, 0xfd, 0xce, 0xa5, 0xa8, 0xbe,
	0x82, 0xfe, 0x89, 0x47, 0x53, 0xa3, 0x1a, 0x49, 0xb6, 0x66, 0xa2, 0x93, 0x71, 0x0b, 0x65, 0x02,
	0xb3, 0x0f, 0xd4, 0x97, 0xd0, 0x3f, 0x2b, 0x2d, 0x85, 0xee, 0xd5, 0xfe, 0x70, 0xdc, 0x8f, 0x98,
	0xe1, 0xc1, 0xb8, 0x63, 0x08, 0x96, 0x58, 0x1f, 0x34, 0x0b, 0xbe, 0xcf, 0xb2, 0xc3, 0xc5, 0xbd,
	0x86, 0xd1, 0x6f, 0xb4, 0x39, 0xa6, 0x46, 0xd9, 0xb6, 0x66, 0x78, 0x0c, 0x26, 0xfb, 0x80, 0x4d,
	0xa3, 0x25, 0xd6, 0x67, 0x3c, 0xf4, 0x48, 0xa4, 0x95, 0x7d, 0xc0, 0xf4, 0x6b, 0x80, 0x25, 0xd6,
	0x3f, 0xdc, 0x9e, 0xf2, 0xb7, 0xee, 0xf1, 0x42, 0xbf, 0x81, 0x11, 0xc7, 0xfe, 0xd4, 0x7e, 0x86,
	0x1f, 0x0f, 0x7f, 0x0b, 0x4f, 0x9b, 0x12, 0x92, 0xe6, 0x81, 0x1e, 0xb8, 0xf0, 0x1a, 0xe2, 0x25,
	0xd6, 0x12, 0x8c, 0x07, 0xdc, 0xcf, 0xfb, 0xfc, 0x4f, 0xe7, 0xdb, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0x24, 0x85, 0xb3, 0x8d, 0x06, 0x00, 0x00,
}
