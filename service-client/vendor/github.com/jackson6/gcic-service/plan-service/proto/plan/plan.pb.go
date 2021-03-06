// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/plan/plan.proto

package plan

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

type Plan struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Includes             []string `protobuf:"bytes,5,rep,name=includes,proto3" json:"includes,omitempty"`
	Levels               []*Level `protobuf:"bytes,6,rep,name=levels,proto3" json:"levels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_91877109aacab46f, []int{0}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plan) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetIncludes() []string {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Plan) GetLevels() []*Level {
	if m != nil {
		return m.Levels
	}
	return nil
}

type Level struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Level) Reset()         { *m = Level{} }
func (m *Level) String() string { return proto.CompactTextString(m) }
func (*Level) ProtoMessage()    {}
func (*Level) Descriptor() ([]byte, []int) {
	return fileDescriptor_91877109aacab46f, []int{1}
}

func (m *Level) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Level.Unmarshal(m, b)
}
func (m *Level) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Level.Marshal(b, m, deterministic)
}
func (m *Level) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Level.Merge(m, src)
}
func (m *Level) XXX_Size() int {
	return xxx_messageInfo_Level.Size(m)
}
func (m *Level) XXX_DiscardUnknown() {
	xxx_messageInfo_Level.DiscardUnknown(m)
}

var xxx_messageInfo_Level proto.InternalMessageInfo

func (m *Level) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Level) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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
	return fileDescriptor_91877109aacab46f, []int{2}
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
	return fileDescriptor_91877109aacab46f, []int{3}
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
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Plan                 *Plan    `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Plans                []*Plan  `protobuf:"bytes,3,rep,name=plans,proto3" json:"plans,omitempty"`
	Errors               []*Error `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_91877109aacab46f, []int{4}
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

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

func (m *Response) GetPlans() []*Plan {
	if m != nil {
		return m.Plans
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*Plan)(nil), "plan.Plan")
	proto.RegisterType((*Level)(nil), "plan.Level")
	proto.RegisterType((*Request)(nil), "plan.Request")
	proto.RegisterType((*Error)(nil), "plan.Error")
	proto.RegisterType((*Response)(nil), "plan.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PlanService service

type PlanServiceClient interface {
	Create(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error)
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type planServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPlanServiceClient(serviceName string, c client.Client) PlanServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "plan"
	}
	return &planServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *planServiceClient) Create(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) Update(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) Delete(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) Get(ctx context.Context, in *Plan, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlanService service

type PlanServiceHandler interface {
	Create(context.Context, *Plan, *Response) error
	Update(context.Context, *Plan, *Response) error
	Delete(context.Context, *Plan, *Response) error
	Get(context.Context, *Plan, *Response) error
	All(context.Context, *Request, *Response) error
}

func RegisterPlanServiceHandler(s server.Server, hdlr PlanServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PlanService{hdlr}, opts...))
}

type PlanService struct {
	PlanServiceHandler
}

func (h *PlanService) Create(ctx context.Context, in *Plan, out *Response) error {
	return h.PlanServiceHandler.Create(ctx, in, out)
}

func (h *PlanService) Update(ctx context.Context, in *Plan, out *Response) error {
	return h.PlanServiceHandler.Update(ctx, in, out)
}

func (h *PlanService) Delete(ctx context.Context, in *Plan, out *Response) error {
	return h.PlanServiceHandler.Delete(ctx, in, out)
}

func (h *PlanService) Get(ctx context.Context, in *Plan, out *Response) error {
	return h.PlanServiceHandler.Get(ctx, in, out)
}

func (h *PlanService) All(ctx context.Context, in *Request, out *Response) error {
	return h.PlanServiceHandler.All(ctx, in, out)
}

func init() { proto.RegisterFile("proto/plan/plan.proto", fileDescriptor_91877109aacab46f) }

var fileDescriptor_91877109aacab46f = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x7d, 0xc9, 0x24, 0x79, 0xed, 0x0d, 0xaf, 0x8b, 0xe1, 0x29, 0x43, 0x17, 0x12, 0x52, 0x90,
	0xac, 0x2a, 0xd4, 0xb5, 0x88, 0xa8, 0xb8, 0x71, 0x21, 0x23, 0x7e, 0x40, 0x4c, 0xee, 0x22, 0x30,
	0x9d, 0x89, 0x33, 0x69, 0xff, 0xc0, 0x4f, 0xf1, 0x87, 0xfc, 0x22, 0x99, 0x9b, 0x54, 0x82, 0x15,
	0xda, 0x4d, 0x72, 0xef, 0x39, 0xe7, 0x1e, 0xce, 0xbd, 0x09, 0x9c, 0xb4, 0xd6, 0x74, 0xe6, 0xa2,
	0x55, 0xa5, 0xa6, 0xc7, 0x92, 0x7a, 0x1e, 0xf9, 0x3a, 0xff, 0x08, 0x20, 0x7a, 0x52, 0xa5, 0xe6,
	0x33, 0x08, 0x9b, 0x5a, 0x04, 0x59, 0x50, 0x4c, 0x65, 0xd8, 0xd4, 0xfc, 0x14, 0x92, 0x72, 0x6d,
	0x36, 0xba, 0x13, 0x61, 0x16, 0x14, 0x4c, 0x0e, 0x1d, 0xe7, 0x10, 0xe9, 0x72, 0x8d, 0x82, 0x91,
	0x92, 0x6a, 0x9e, 0x41, 0x5a, 0xa3, 0xab, 0x6c, 0xd3, 0x76, 0x8d, 0xd1, 0x22, 0x22, 0x6a, 0x0c,
	0xf1, 0x39, 0x4c, 0x1a, 0x5d, 0xa9, 0x4d, 0x8d, 0x4e, 0xc4, 0x19, 0x2b, 0xa6, 0xf2, 0xbb, 0xe7,
	0x0b, 0x48, 0x14, 0x6e, 0x51, 0x39, 0x91, 0x64, 0xac, 0x48, 0x57, 0xe9, 0x92, 0x52, 0x3e, 0x7a,
	0x4c, 0x0e, 0x54, 0x7e, 0x0d, 0x31, 0x01, 0xfc, 0x3f, 0xc4, 0x04, 0x51, 0xd4, 0x58, 0xf6, 0xcd,
	0xcf, 0x04, 0xe1, 0x5e, 0x82, 0x7c, 0x0a, 0x7f, 0x25, 0xbe, 0x6d, 0xd0, 0x75, 0xf9, 0x15, 0xc4,
	0xf7, 0xd6, 0x1a, 0xeb, 0x77, 0xa9, 0x4c, 0x8d, 0x83, 0x15, 0xd5, 0x47, 0x38, 0xbd, 0x07, 0x30,
	0x91, 0xe8, 0x5a, 0xa3, 0x1d, 0xfe, 0x6a, 0x71, 0x06, 0x74, 0x5b, 0x9a, 0x4d, 0x57, 0xd0, 0xaf,
	0xe3, 0x8f, 0x2c, 0x09, 0xe7, 0x19, 0xc4, 0xfe, 0xed, 0x04, 0xa3, 0x7d, 0xc7, 0x82, 0x9e, 0xf0,
	0x27, 0x41, 0x9f, 0xd0, 0x89, 0x68, 0x7c, 0x12, 0x4a, 0x2d, 0x07, 0x6a, 0xf5, 0x19, 0x40, 0xea,
	0x87, 0x9e, 0xd1, 0x6e, 0x9b, 0x0a, 0xf9, 0x39, 0x24, 0xb7, 0x16, 0xcb, 0x0e, 0xf9, 0xc8, 0x71,
	0x3e, 0xeb, 0xeb, 0x5d, 0xe0, 0xfc, 0x8f, 0xd7, 0xbd, 0xb4, 0xf5, 0x51, 0xba, 0x3b, 0x54, 0x78,
	0x50, 0xb7, 0x00, 0xf6, 0x80, 0xdd, 0x41, 0x33, 0x76, 0xa3, 0x14, 0xff, 0xb7, 0x23, 0xe8, 0x4b,
	0xec, 0xeb, 0x5e, 0x13, 0xfa, 0x39, 0x2f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xe6, 0x4f,
	0xa2, 0xb5, 0x02, 0x00, 0x00,
}
