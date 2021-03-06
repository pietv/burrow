// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keys.proto

/*
Package pbkeys is a generated protocol buffer package.

It is generated from these files:
	keys.proto

It has these top-level messages:
	Name
	Empty
	GenRequest
	GenResponse
	PubRequest
	PubResponse
	ImportJSONRequest
	ImportResponse
	ImportRequest
	ExportRequest
	ExportResponse
	SignRequest
	SignResponse
	VerifyRequest
	HashRequest
	HashResponse
	Key
	ListResponse
	AddNameRequest
*/
package pbkeys

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Name struct {
	Keyname string `protobuf:"bytes,1,opt,name=keyname" json:"keyname,omitempty"`
}

func (m *Name) Reset()                    { *m = Name{} }
func (m *Name) String() string            { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()               {}
func (*Name) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Name) GetKeyname() string {
	if m != nil {
		return m.Keyname
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GenRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	Curvetype  string `protobuf:"bytes,2,opt,name=curvetype" json:"curvetype,omitempty"`
	Keyname    string `protobuf:"bytes,3,opt,name=keyname" json:"keyname,omitempty"`
}

func (m *GenRequest) Reset()                    { *m = GenRequest{} }
func (m *GenRequest) String() string            { return proto.CompactTextString(m) }
func (*GenRequest) ProtoMessage()               {}
func (*GenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GenRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *GenRequest) GetCurvetype() string {
	if m != nil {
		return m.Curvetype
	}
	return ""
}

func (m *GenRequest) GetKeyname() string {
	if m != nil {
		return m.Keyname
	}
	return ""
}

type GenResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *GenResponse) Reset()                    { *m = GenResponse{} }
func (m *GenResponse) String() string            { return proto.CompactTextString(m) }
func (*GenResponse) ProtoMessage()               {}
func (*GenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PubRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PubRequest) Reset()                    { *m = PubRequest{} }
func (m *PubRequest) String() string            { return proto.CompactTextString(m) }
func (*PubRequest) ProtoMessage()               {}
func (*PubRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PubRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PubRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PubResponse struct {
	Pub       []byte `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	Curvetype string `protobuf:"bytes,2,opt,name=curvetype" json:"curvetype,omitempty"`
}

func (m *PubResponse) Reset()                    { *m = PubResponse{} }
func (m *PubResponse) String() string            { return proto.CompactTextString(m) }
func (*PubResponse) ProtoMessage()               {}
func (*PubResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PubResponse) GetPub() []byte {
	if m != nil {
		return m.Pub
	}
	return nil
}

func (m *PubResponse) GetCurvetype() string {
	if m != nil {
		return m.Curvetype
	}
	return ""
}

type ImportJSONRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	JSON       string `protobuf:"bytes,2,opt,name=JSON" json:"JSON,omitempty"`
}

func (m *ImportJSONRequest) Reset()                    { *m = ImportJSONRequest{} }
func (m *ImportJSONRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportJSONRequest) ProtoMessage()               {}
func (*ImportJSONRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ImportJSONRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ImportJSONRequest) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

type ImportResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *ImportResponse) Reset()                    { *m = ImportResponse{} }
func (m *ImportResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportResponse) ProtoMessage()               {}
func (*ImportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ImportResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ImportRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Curvetype  string `protobuf:"bytes,3,opt,name=curvetype" json:"curvetype,omitempty"`
	Keybytes   []byte `protobuf:"bytes,4,opt,name=keybytes,proto3" json:"keybytes,omitempty"`
}

func (m *ImportRequest) Reset()                    { *m = ImportRequest{} }
func (m *ImportRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()               {}
func (*ImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ImportRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ImportRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImportRequest) GetCurvetype() string {
	if m != nil {
		return m.Curvetype
	}
	return ""
}

func (m *ImportRequest) GetKeybytes() []byte {
	if m != nil {
		return m.Keybytes
	}
	return nil
}

type ExportRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ExportRequest) Reset()                    { *m = ExportRequest{} }
func (m *ExportRequest) String() string            { return proto.CompactTextString(m) }
func (*ExportRequest) ProtoMessage()               {}
func (*ExportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ExportRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ExportRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExportRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ExportResponse struct {
	Export string `protobuf:"bytes,1,opt,name=export" json:"export,omitempty"`
}

func (m *ExportResponse) Reset()                    { *m = ExportResponse{} }
func (m *ExportResponse) String() string            { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()               {}
func (*ExportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ExportResponse) GetExport() string {
	if m != nil {
		return m.Export
	}
	return ""
}

type SignRequest struct {
	Passphrase string `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Message    []byte `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *SignRequest) Reset()                    { *m = SignRequest{} }
func (m *SignRequest) String() string            { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()               {}
func (*SignRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SignRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *SignRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SignRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type SignResponse struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Curvetype string `protobuf:"bytes,2,opt,name=curvetype" json:"curvetype,omitempty"`
}

func (m *SignResponse) Reset()                    { *m = SignResponse{} }
func (m *SignResponse) String() string            { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()               {}
func (*SignResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SignResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignResponse) GetCurvetype() string {
	if m != nil {
		return m.Curvetype
	}
	return ""
}

type VerifyRequest struct {
	Curvetype string `protobuf:"bytes,1,opt,name=curvetype" json:"curvetype,omitempty"`
	Pub       []byte `protobuf:"bytes,2,opt,name=pub,proto3" json:"pub,omitempty"`
	Message   []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *VerifyRequest) Reset()                    { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string            { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()               {}
func (*VerifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *VerifyRequest) GetCurvetype() string {
	if m != nil {
		return m.Curvetype
	}
	return ""
}

func (m *VerifyRequest) GetPub() []byte {
	if m != nil {
		return m.Pub
	}
	return nil
}

func (m *VerifyRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *VerifyRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type HashRequest struct {
	Hashtype string `protobuf:"bytes,1,opt,name=hashtype" json:"hashtype,omitempty"`
	Message  []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HashRequest) Reset()                    { *m = HashRequest{} }
func (m *HashRequest) String() string            { return proto.CompactTextString(m) }
func (*HashRequest) ProtoMessage()               {}
func (*HashRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *HashRequest) GetHashtype() string {
	if m != nil {
		return m.Hashtype
	}
	return ""
}

func (m *HashRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type HashResponse struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *HashResponse) Reset()                    { *m = HashResponse{} }
func (m *HashResponse) String() string            { return proto.CompactTextString(m) }
func (*HashResponse) ProtoMessage()               {}
func (*HashResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *HashResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Key struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Keyname string `protobuf:"bytes,2,opt,name=keyname" json:"keyname,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Key) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Key) GetKeyname() string {
	if m != nil {
		return m.Keyname
	}
	return ""
}

type ListResponse struct {
	Key []*Key `protobuf:"bytes,1,rep,name=key" json:"key,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ListResponse) GetKey() []*Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type AddNameRequest struct {
	Keyname string `protobuf:"bytes,1,opt,name=keyname" json:"keyname,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *AddNameRequest) Reset()                    { *m = AddNameRequest{} }
func (m *AddNameRequest) String() string            { return proto.CompactTextString(m) }
func (*AddNameRequest) ProtoMessage()               {}
func (*AddNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *AddNameRequest) GetKeyname() string {
	if m != nil {
		return m.Keyname
	}
	return ""
}

func (m *AddNameRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Name)(nil), "pbkeys.Name")
	proto.RegisterType((*Empty)(nil), "pbkeys.Empty")
	proto.RegisterType((*GenRequest)(nil), "pbkeys.GenRequest")
	proto.RegisterType((*GenResponse)(nil), "pbkeys.GenResponse")
	proto.RegisterType((*PubRequest)(nil), "pbkeys.PubRequest")
	proto.RegisterType((*PubResponse)(nil), "pbkeys.PubResponse")
	proto.RegisterType((*ImportJSONRequest)(nil), "pbkeys.ImportJSONRequest")
	proto.RegisterType((*ImportResponse)(nil), "pbkeys.ImportResponse")
	proto.RegisterType((*ImportRequest)(nil), "pbkeys.ImportRequest")
	proto.RegisterType((*ExportRequest)(nil), "pbkeys.ExportRequest")
	proto.RegisterType((*ExportResponse)(nil), "pbkeys.ExportResponse")
	proto.RegisterType((*SignRequest)(nil), "pbkeys.SignRequest")
	proto.RegisterType((*SignResponse)(nil), "pbkeys.SignResponse")
	proto.RegisterType((*VerifyRequest)(nil), "pbkeys.VerifyRequest")
	proto.RegisterType((*HashRequest)(nil), "pbkeys.HashRequest")
	proto.RegisterType((*HashResponse)(nil), "pbkeys.HashResponse")
	proto.RegisterType((*Key)(nil), "pbkeys.Key")
	proto.RegisterType((*ListResponse)(nil), "pbkeys.ListResponse")
	proto.RegisterType((*AddNameRequest)(nil), "pbkeys.AddNameRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Keys service

type KeysClient interface {
	GenerateKey(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (*GenResponse, error)
	PublicKey(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubResponse, error)
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*Empty, error)
	Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error)
	ImportJSON(ctx context.Context, in *ImportJSONRequest, opts ...grpc.CallOption) (*ImportResponse, error)
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	RemoveName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *Name, opts ...grpc.CallOption) (*ListResponse, error)
	AddName(ctx context.Context, in *AddNameRequest, opts ...grpc.CallOption) (*Empty, error)
}

type keysClient struct {
	cc *grpc.ClientConn
}

func NewKeysClient(cc *grpc.ClientConn) KeysClient {
	return &keysClient{cc}
}

func (c *keysClient) GenerateKey(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (*GenResponse, error) {
	out := new(GenResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/GenerateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) PublicKey(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubResponse, error) {
	out := new(PubResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/PublicKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/Sign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error) {
	out := new(ImportResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/Import", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) ImportJSON(ctx context.Context, in *ImportJSONRequest, opts ...grpc.CallOption) (*ImportResponse, error) {
	out := new(ImportResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/ImportJSON", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/Export", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/Hash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) RemoveName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/RemoveName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) List(ctx context.Context, in *Name, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) AddName(ctx context.Context, in *AddNameRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pbkeys.Keys/AddName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Keys service

type KeysServer interface {
	GenerateKey(context.Context, *GenRequest) (*GenResponse, error)
	PublicKey(context.Context, *PubRequest) (*PubResponse, error)
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	Verify(context.Context, *VerifyRequest) (*Empty, error)
	Import(context.Context, *ImportRequest) (*ImportResponse, error)
	ImportJSON(context.Context, *ImportJSONRequest) (*ImportResponse, error)
	Export(context.Context, *ExportRequest) (*ExportResponse, error)
	Hash(context.Context, *HashRequest) (*HashResponse, error)
	RemoveName(context.Context, *Name) (*Empty, error)
	List(context.Context, *Name) (*ListResponse, error)
	AddName(context.Context, *AddNameRequest) (*Empty, error)
}

func RegisterKeysServer(s *grpc.Server, srv KeysServer) {
	s.RegisterService(&_Keys_serviceDesc, srv)
}

func _Keys_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/GenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).GenerateKey(ctx, req.(*GenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_PublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).PublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/PublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).PublicKey(ctx, req.(*PubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).Import(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/Import",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).Import(ctx, req.(*ImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_ImportJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportJSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).ImportJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/ImportJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).ImportJSON(ctx, req.(*ImportJSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).Export(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_RemoveName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).RemoveName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/RemoveName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).RemoveName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).List(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_AddName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).AddName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkeys.Keys/AddName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).AddName(ctx, req.(*AddNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Keys_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbkeys.Keys",
	HandlerType: (*KeysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKey",
			Handler:    _Keys_GenerateKey_Handler,
		},
		{
			MethodName: "PublicKey",
			Handler:    _Keys_PublicKey_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Keys_Sign_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Keys_Verify_Handler,
		},
		{
			MethodName: "Import",
			Handler:    _Keys_Import_Handler,
		},
		{
			MethodName: "ImportJSON",
			Handler:    _Keys_ImportJSON_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _Keys_Export_Handler,
		},
		{
			MethodName: "Hash",
			Handler:    _Keys_Hash_Handler,
		},
		{
			MethodName: "RemoveName",
			Handler:    _Keys_RemoveName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Keys_List_Handler,
		},
		{
			MethodName: "AddName",
			Handler:    _Keys_AddName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keys.proto",
}

func init() { proto.RegisterFile("keys.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x55, 0x9b, 0x7c, 0xdd, 0x7a, 0xda, 0x4e, 0x1f, 0x06, 0xa6, 0x52, 0x0d, 0x34, 0xf9, 0x85,
	0x31, 0x89, 0x8a, 0x0d, 0xc4, 0x04, 0x12, 0x42, 0x08, 0xa6, 0xc1, 0x86, 0xc6, 0xd4, 0x49, 0xbc,
	0xf1, 0x90, 0xae, 0x97, 0xb5, 0x2a, 0x69, 0x43, 0x9c, 0x0c, 0xf2, 0xc0, 0xdf, 0xe3, 0x77, 0x21,
	0x3b, 0x76, 0x6d, 0x47, 0xeb, 0xa8, 0xc4, 0x9b, 0xef, 0xcd, 0xbd, 0xf7, 0x1c, 0xdf, 0x9c, 0x9c,
	0x00, 0x53, 0x2a, 0x44, 0x3f, 0x49, 0xe7, 0xd9, 0x9c, 0x35, 0x92, 0xa1, 0x8c, 0xf8, 0x36, 0xc2,
	0xd3, 0x28, 0x26, 0xd6, 0xc5, 0xda, 0x94, 0x8a, 0x59, 0x14, 0x53, 0xb7, 0xb6, 0x5d, 0xdb, 0x69,
	0x0e, 0x4c, 0xc8, 0xd7, 0xf0, 0xdf, 0x61, 0x9c, 0x64, 0x05, 0x1f, 0x01, 0x47, 0x34, 0x1b, 0xd0,
	0xf7, 0x9c, 0x44, 0xc6, 0x1e, 0x00, 0x49, 0x24, 0x44, 0x32, 0x4e, 0x23, 0x61, 0x7a, 0x9c, 0x0c,
	0xdb, 0x42, 0xf3, 0x22, 0x4f, 0xaf, 0x28, 0x2b, 0x12, 0xea, 0xd6, 0xd5, 0x63, 0x9b, 0x70, 0xe1,
	0x02, 0x1f, 0xee, 0x21, 0x5a, 0x0a, 0x45, 0x24, 0xf3, 0x99, 0x50, 0x85, 0xd1, 0x68, 0x94, 0x92,
	0x10, 0x86, 0x97, 0x0e, 0xf9, 0x4b, 0xe0, 0x2c, 0x1f, 0x1a, 0x3a, 0x4b, 0xeb, 0x18, 0x43, 0xa8,
	0x70, 0x4a, 0x0e, 0xea, 0xcc, 0x5f, 0xa1, 0xa5, 0x7a, 0x35, 0xc8, 0xff, 0x08, 0x92, 0x7c, 0xa8,
	0x1a, 0xdb, 0x03, 0x79, 0xbc, 0x99, 0x3d, 0x3f, 0xc2, 0xad, 0x0f, 0x71, 0x32, 0x4f, 0xb3, 0xe3,
	0xf3, 0x4f, 0xa7, 0xab, 0x2e, 0x84, 0x21, 0x94, 0xe5, 0x86, 0x87, 0x3c, 0xf3, 0x5d, 0x6c, 0x94,
	0x83, 0x56, 0xb8, 0xef, 0x2f, 0x74, 0x4c, 0xed, 0xca, 0x80, 0xd5, 0x8b, 0xfb, 0xf7, 0x0a, 0xaa,
	0x6f, 0xa5, 0x87, 0xf5, 0x29, 0x15, 0xc3, 0x22, 0x23, 0xd1, 0x0d, 0xd5, 0x32, 0x16, 0x31, 0xff,
	0x82, 0xce, 0xe1, 0xcf, 0x7f, 0x85, 0x77, 0x6e, 0x17, 0xf8, 0xb7, 0xdb, 0xc1, 0x86, 0x19, 0xaf,
	0x37, 0xb1, 0x89, 0x06, 0xa9, 0x8c, 0x9e, 0xad, 0x23, 0x9e, 0xa3, 0x75, 0x3e, 0xb9, 0x5c, 0x59,
	0x87, 0x0e, 0x64, 0xfd, 0x7a, 0x61, 0x04, 0x3e, 0xc1, 0x98, 0x84, 0x88, 0x2e, 0x49, 0x2f, 0xc0,
	0x84, 0xfc, 0x18, 0xed, 0x12, 0x56, 0xd3, 0xdb, 0x42, 0x53, 0x4c, 0x2e, 0x67, 0x51, 0x96, 0xa7,
	0xa4, 0x95, 0x63, 0x13, 0x7f, 0xd1, 0xcf, 0x0f, 0x74, 0x3e, 0x53, 0x3a, 0xf9, 0x5a, 0x98, 0x4b,
	0x78, 0xe5, 0xb5, 0xea, 0x6b, 0xd1, 0xf2, 0xac, 0x5b, 0x79, 0x3a, 0x34, 0x03, 0x8f, 0xa6, 0x4f,
	0x2b, 0xac, 0xd0, 0xe2, 0x6f, 0xd1, 0x7a, 0x1f, 0x89, 0xb1, 0x81, 0xed, 0x61, 0x7d, 0x1c, 0x89,
	0xb1, 0x83, 0xba, 0x88, 0x5d, 0x88, 0xba, 0xbf, 0x09, 0x8e, 0x76, 0x39, 0x44, 0x6f, 0x82, 0x21,
	0x94, 0x5d, 0x7a, 0x82, 0x3a, 0xf3, 0x17, 0x08, 0x4e, 0xa8, 0xb8, 0xe1, 0xab, 0x74, 0x0c, 0xa0,
	0xee, 0x1b, 0xc0, 0x63, 0xb4, 0x3f, 0x4e, 0x84, 0xd5, 0xc1, 0x7d, 0x04, 0x53, 0x2a, 0xba, 0xb5,
	0xed, 0x60, 0xa7, 0xb5, 0xdf, 0xea, 0x97, 0xbe, 0xd5, 0x3f, 0xa1, 0x62, 0x20, 0xf3, 0xfc, 0x1d,
	0x36, 0xde, 0x8c, 0x46, 0xd2, 0xc3, 0x1c, 0x2b, 0xb8, 0xde, 0xca, 0x96, 0x6b, 0x61, 0xff, 0x77,
	0x88, 0xf0, 0x84, 0x0a, 0xc1, 0x9e, 0x2b, 0xfb, 0xa1, 0x34, 0xca, 0x48, 0x5e, 0x80, 0x19, 0x3c,
	0xeb, 0x7c, 0xbd, 0xdb, 0x5e, 0x4e, 0xb3, 0x7c, 0x86, 0xe6, 0x59, 0x3e, 0xfc, 0x36, 0xb9, 0xf0,
	0xba, 0xac, 0x41, 0xd9, 0x2e, 0xd7, 0x78, 0xf6, 0x10, 0x4a, 0x51, 0xb1, 0xc5, 0x43, 0x47, 0xd9,
	0xbd, 0x3b, 0x7e, 0x52, 0xb7, 0xf4, 0xd1, 0x28, 0xb5, 0xc3, 0xee, 0x9a, 0xe7, 0x9e, 0x96, 0x7a,
	0x1d, 0x93, 0x56, 0xae, 0xcd, 0x0e, 0xd0, 0x28, 0x6d, 0xc3, 0xd6, 0x7b, 0x36, 0xd2, 0xdb, 0xac,
	0xa6, 0x35, 0xd0, 0x6b, 0xc0, 0x9a, 0x1c, 0xbb, 0xe7, 0x57, 0x39, 0xc6, 0xb7, 0x74, 0xc0, 0x01,
	0x1a, 0xe5, 0x27, 0x6d, 0x91, 0x3d, 0x07, 0xb1, 0x8d, 0x95, 0x2f, 0x7f, 0x0f, 0xa1, 0x14, 0x98,
	0xdd, 0x8a, 0xa3, 0x59, 0xbb, 0x15, 0x4f, 0x83, 0x8f, 0x80, 0x01, 0xc5, 0xf3, 0x2b, 0x52, 0x3f,
	0xb3, 0xb6, 0xa9, 0x91, 0x51, 0x75, 0x21, 0xbb, 0x08, 0xa5, 0xbe, 0x2a, 0x45, 0x8b, 0xb1, 0x9e,
	0xf6, 0x9e, 0x60, 0x4d, 0x8b, 0x8b, 0x2d, 0xc8, 0xfa, 0x6a, 0xab, 0x4c, 0x1f, 0x36, 0xd4, 0xef,
	0xf5, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xed, 0x32, 0x21, 0x6c, 0x07, 0x00, 0x00,
}
