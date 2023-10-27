// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: trxlocal.proto

package trxLocal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TrxLocalService_AddTrxWithCard_FullMethodName            = "/trx.trxLocalService/AddTrxWithCard"
	TrxLocalService_AddTrxWithoutCard_FullMethodName         = "/trx.trxLocalService/AddTrxWithoutCard"
	TrxLocalService_InquiryTrxWithoutCard_FullMethodName     = "/trx.trxLocalService/InquiryTrxWithoutCard"
	TrxLocalService_InquiryTrxWithCard_FullMethodName        = "/trx.trxLocalService/InquiryTrxWithCard"
	TrxLocalService_ConfirmTrx_FullMethodName                = "/trx.trxLocalService/ConfirmTrx"
	TrxLocalService_ConfirmTrxByPass_FullMethodName          = "/trx.trxLocalService/ConfirmTrxByPass"
	TrxLocalService_ConfirmSyncTrxToCloud_FullMethodName     = "/trx.trxLocalService/ConfirmSyncTrxToCloud"
	TrxLocalService_InquiryPayment_FullMethodName            = "/trx.trxLocalService/InquiryPayment"
	TrxLocalService_InquiryWithCardP3_FullMethodName         = "/trx.trxLocalService/InquiryWithCardP3"
	TrxLocalService_InquiryPaymentP3_FullMethodName          = "/trx.trxLocalService/InquiryPaymentP3"
	TrxLocalService_GetTrxListForDocDate_FullMethodName      = "/trx.trxLocalService/GetTrxListForDocDate"
	TrxLocalService_UpdateStatusManualTrx_FullMethodName     = "/trx.trxLocalService/UpdateStatusManualTrx"
	TrxLocalService_FindTrxOutstandingByIndex_FullMethodName = "/trx.trxLocalService/FindTrxOutstandingByIndex"
	TrxLocalService_UpdateProductPrice_FullMethodName        = "/trx.trxLocalService/UpdateProductPrice"
	TrxLocalService_RegisterMember_FullMethodName            = "/trx.trxLocalService/RegisterMember"
	TrxLocalService_DecryptMKey_FullMethodName               = "/trx.trxLocalService/DecryptMKey"
	TrxLocalService_UpdateAutoClearTrx_FullMethodName        = "/trx.trxLocalService/UpdateAutoClearTrx"
	TrxLocalService_ConfirmTrxP3_FullMethodName              = "/trx.trxLocalService/ConfirmTrxP3"
)

// TrxLocalServiceClient is the client API for TrxLocalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrxLocalServiceClient interface {
	AddTrxWithCard(ctx context.Context, in *RequestTrxCheckin, opts ...grpc.CallOption) (*MyResponse, error)
	AddTrxWithoutCard(ctx context.Context, in *RequestTrxCheckInWithoutCard, opts ...grpc.CallOption) (*MyResponse, error)
	InquiryTrxWithoutCard(ctx context.Context, in *RequestInquiryWithoutCard, opts ...grpc.CallOption) (*MyResponse, error)
	InquiryTrxWithCard(ctx context.Context, in *RequestInquiryWithCard, opts ...grpc.CallOption) (*MyResponse, error)
	ConfirmTrx(ctx context.Context, in *RequestConfirmTrx, opts ...grpc.CallOption) (*MyResponse, error)
	ConfirmTrxByPass(ctx context.Context, in *ConfirmTrxByPassMessage, opts ...grpc.CallOption) (*MyResponse, error)
	ConfirmSyncTrxToCloud(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MyResponse, error)
	InquiryPayment(ctx context.Context, in *RequestInquiryPayment, opts ...grpc.CallOption) (*MyResponse, error)
	InquiryWithCardP3(ctx context.Context, in *RequestInquiryWithCardP3, opts ...grpc.CallOption) (*MyResponse, error)
	InquiryPaymentP3(ctx context.Context, in *RequestInquiryPaymentP3, opts ...grpc.CallOption) (*MyResponse, error)
	GetTrxListForDocDate(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error)
	UpdateStatusManualTrx(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error)
	FindTrxOutstandingByIndex(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error)
	UpdateProductPrice(ctx context.Context, in *RequestUpdateProductPrice, opts ...grpc.CallOption) (*MyResponse, error)
	RegisterMember(ctx context.Context, in *RequestRegistrationMemberLocal, opts ...grpc.CallOption) (*MyResponse, error)
	DecryptMKey(ctx context.Context, in *Decrypt, opts ...grpc.CallOption) (*MyResponse, error)
	UpdateAutoClearTrx(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error)
	ConfirmTrxP3(ctx context.Context, in *RequestConfirmTrx, opts ...grpc.CallOption) (*MyResponse, error)
}

type trxLocalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrxLocalServiceClient(cc grpc.ClientConnInterface) TrxLocalServiceClient {
	return &trxLocalServiceClient{cc}
}

func (c *trxLocalServiceClient) AddTrxWithCard(ctx context.Context, in *RequestTrxCheckin, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_AddTrxWithCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) AddTrxWithoutCard(ctx context.Context, in *RequestTrxCheckInWithoutCard, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_AddTrxWithoutCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) InquiryTrxWithoutCard(ctx context.Context, in *RequestInquiryWithoutCard, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_InquiryTrxWithoutCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) InquiryTrxWithCard(ctx context.Context, in *RequestInquiryWithCard, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_InquiryTrxWithCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) ConfirmTrx(ctx context.Context, in *RequestConfirmTrx, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_ConfirmTrx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) ConfirmTrxByPass(ctx context.Context, in *ConfirmTrxByPassMessage, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_ConfirmTrxByPass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) ConfirmSyncTrxToCloud(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_ConfirmSyncTrxToCloud_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) InquiryPayment(ctx context.Context, in *RequestInquiryPayment, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_InquiryPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) InquiryWithCardP3(ctx context.Context, in *RequestInquiryWithCardP3, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_InquiryWithCardP3_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) InquiryPaymentP3(ctx context.Context, in *RequestInquiryPaymentP3, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_InquiryPaymentP3_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) GetTrxListForDocDate(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_GetTrxListForDocDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) UpdateStatusManualTrx(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_UpdateStatusManualTrx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) FindTrxOutstandingByIndex(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_FindTrxOutstandingByIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) UpdateProductPrice(ctx context.Context, in *RequestUpdateProductPrice, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_UpdateProductPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) RegisterMember(ctx context.Context, in *RequestRegistrationMemberLocal, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_RegisterMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) DecryptMKey(ctx context.Context, in *Decrypt, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_DecryptMKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) UpdateAutoClearTrx(ctx context.Context, in *Param, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_UpdateAutoClearTrx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trxLocalServiceClient) ConfirmTrxP3(ctx context.Context, in *RequestConfirmTrx, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, TrxLocalService_ConfirmTrxP3_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrxLocalServiceServer is the server API for TrxLocalService service.
// All implementations must embed UnimplementedTrxLocalServiceServer
// for forward compatibility
type TrxLocalServiceServer interface {
	AddTrxWithCard(context.Context, *RequestTrxCheckin) (*MyResponse, error)
	AddTrxWithoutCard(context.Context, *RequestTrxCheckInWithoutCard) (*MyResponse, error)
	InquiryTrxWithoutCard(context.Context, *RequestInquiryWithoutCard) (*MyResponse, error)
	InquiryTrxWithCard(context.Context, *RequestInquiryWithCard) (*MyResponse, error)
	ConfirmTrx(context.Context, *RequestConfirmTrx) (*MyResponse, error)
	ConfirmTrxByPass(context.Context, *ConfirmTrxByPassMessage) (*MyResponse, error)
	ConfirmSyncTrxToCloud(context.Context, *Empty) (*MyResponse, error)
	InquiryPayment(context.Context, *RequestInquiryPayment) (*MyResponse, error)
	InquiryWithCardP3(context.Context, *RequestInquiryWithCardP3) (*MyResponse, error)
	InquiryPaymentP3(context.Context, *RequestInquiryPaymentP3) (*MyResponse, error)
	GetTrxListForDocDate(context.Context, *Param) (*MyResponse, error)
	UpdateStatusManualTrx(context.Context, *Param) (*MyResponse, error)
	FindTrxOutstandingByIndex(context.Context, *Param) (*MyResponse, error)
	UpdateProductPrice(context.Context, *RequestUpdateProductPrice) (*MyResponse, error)
	RegisterMember(context.Context, *RequestRegistrationMemberLocal) (*MyResponse, error)
	DecryptMKey(context.Context, *Decrypt) (*MyResponse, error)
	UpdateAutoClearTrx(context.Context, *Param) (*MyResponse, error)
	ConfirmTrxP3(context.Context, *RequestConfirmTrx) (*MyResponse, error)
	mustEmbedUnimplementedTrxLocalServiceServer()
}

// UnimplementedTrxLocalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrxLocalServiceServer struct {
}

func (UnimplementedTrxLocalServiceServer) AddTrxWithCard(context.Context, *RequestTrxCheckin) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrxWithCard not implemented")
}
func (UnimplementedTrxLocalServiceServer) AddTrxWithoutCard(context.Context, *RequestTrxCheckInWithoutCard) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrxWithoutCard not implemented")
}
func (UnimplementedTrxLocalServiceServer) InquiryTrxWithoutCard(context.Context, *RequestInquiryWithoutCard) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryTrxWithoutCard not implemented")
}
func (UnimplementedTrxLocalServiceServer) InquiryTrxWithCard(context.Context, *RequestInquiryWithCard) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryTrxWithCard not implemented")
}
func (UnimplementedTrxLocalServiceServer) ConfirmTrx(context.Context, *RequestConfirmTrx) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTrx not implemented")
}
func (UnimplementedTrxLocalServiceServer) ConfirmTrxByPass(context.Context, *ConfirmTrxByPassMessage) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTrxByPass not implemented")
}
func (UnimplementedTrxLocalServiceServer) ConfirmSyncTrxToCloud(context.Context, *Empty) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSyncTrxToCloud not implemented")
}
func (UnimplementedTrxLocalServiceServer) InquiryPayment(context.Context, *RequestInquiryPayment) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryPayment not implemented")
}
func (UnimplementedTrxLocalServiceServer) InquiryWithCardP3(context.Context, *RequestInquiryWithCardP3) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryWithCardP3 not implemented")
}
func (UnimplementedTrxLocalServiceServer) InquiryPaymentP3(context.Context, *RequestInquiryPaymentP3) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryPaymentP3 not implemented")
}
func (UnimplementedTrxLocalServiceServer) GetTrxListForDocDate(context.Context, *Param) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrxListForDocDate not implemented")
}
func (UnimplementedTrxLocalServiceServer) UpdateStatusManualTrx(context.Context, *Param) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatusManualTrx not implemented")
}
func (UnimplementedTrxLocalServiceServer) FindTrxOutstandingByIndex(context.Context, *Param) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTrxOutstandingByIndex not implemented")
}
func (UnimplementedTrxLocalServiceServer) UpdateProductPrice(context.Context, *RequestUpdateProductPrice) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductPrice not implemented")
}
func (UnimplementedTrxLocalServiceServer) RegisterMember(context.Context, *RequestRegistrationMemberLocal) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMember not implemented")
}
func (UnimplementedTrxLocalServiceServer) DecryptMKey(context.Context, *Decrypt) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptMKey not implemented")
}
func (UnimplementedTrxLocalServiceServer) UpdateAutoClearTrx(context.Context, *Param) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAutoClearTrx not implemented")
}
func (UnimplementedTrxLocalServiceServer) ConfirmTrxP3(context.Context, *RequestConfirmTrx) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTrxP3 not implemented")
}
func (UnimplementedTrxLocalServiceServer) mustEmbedUnimplementedTrxLocalServiceServer() {}

// UnsafeTrxLocalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrxLocalServiceServer will
// result in compilation errors.
type UnsafeTrxLocalServiceServer interface {
	mustEmbedUnimplementedTrxLocalServiceServer()
}

func RegisterTrxLocalServiceServer(s grpc.ServiceRegistrar, srv TrxLocalServiceServer) {
	s.RegisterService(&TrxLocalService_ServiceDesc, srv)
}

func _TrxLocalService_AddTrxWithCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTrxCheckin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).AddTrxWithCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_AddTrxWithCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).AddTrxWithCard(ctx, req.(*RequestTrxCheckin))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_AddTrxWithoutCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTrxCheckInWithoutCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).AddTrxWithoutCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_AddTrxWithoutCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).AddTrxWithoutCard(ctx, req.(*RequestTrxCheckInWithoutCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_InquiryTrxWithoutCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInquiryWithoutCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).InquiryTrxWithoutCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_InquiryTrxWithoutCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).InquiryTrxWithoutCard(ctx, req.(*RequestInquiryWithoutCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_InquiryTrxWithCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInquiryWithCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).InquiryTrxWithCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_InquiryTrxWithCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).InquiryTrxWithCard(ctx, req.(*RequestInquiryWithCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_ConfirmTrx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConfirmTrx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).ConfirmTrx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_ConfirmTrx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).ConfirmTrx(ctx, req.(*RequestConfirmTrx))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_ConfirmTrxByPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTrxByPassMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).ConfirmTrxByPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_ConfirmTrxByPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).ConfirmTrxByPass(ctx, req.(*ConfirmTrxByPassMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_ConfirmSyncTrxToCloud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).ConfirmSyncTrxToCloud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_ConfirmSyncTrxToCloud_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).ConfirmSyncTrxToCloud(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_InquiryPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInquiryPayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).InquiryPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_InquiryPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).InquiryPayment(ctx, req.(*RequestInquiryPayment))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_InquiryWithCardP3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInquiryWithCardP3)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).InquiryWithCardP3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_InquiryWithCardP3_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).InquiryWithCardP3(ctx, req.(*RequestInquiryWithCardP3))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_InquiryPaymentP3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInquiryPaymentP3)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).InquiryPaymentP3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_InquiryPaymentP3_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).InquiryPaymentP3(ctx, req.(*RequestInquiryPaymentP3))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_GetTrxListForDocDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).GetTrxListForDocDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_GetTrxListForDocDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).GetTrxListForDocDate(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_UpdateStatusManualTrx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).UpdateStatusManualTrx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_UpdateStatusManualTrx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).UpdateStatusManualTrx(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_FindTrxOutstandingByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).FindTrxOutstandingByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_FindTrxOutstandingByIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).FindTrxOutstandingByIndex(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_UpdateProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateProductPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).UpdateProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_UpdateProductPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).UpdateProductPrice(ctx, req.(*RequestUpdateProductPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_RegisterMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegistrationMemberLocal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).RegisterMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_RegisterMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).RegisterMember(ctx, req.(*RequestRegistrationMemberLocal))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_DecryptMKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Decrypt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).DecryptMKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_DecryptMKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).DecryptMKey(ctx, req.(*Decrypt))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_UpdateAutoClearTrx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).UpdateAutoClearTrx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_UpdateAutoClearTrx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).UpdateAutoClearTrx(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrxLocalService_ConfirmTrxP3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConfirmTrx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrxLocalServiceServer).ConfirmTrxP3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrxLocalService_ConfirmTrxP3_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrxLocalServiceServer).ConfirmTrxP3(ctx, req.(*RequestConfirmTrx))
	}
	return interceptor(ctx, in, info, handler)
}

// TrxLocalService_ServiceDesc is the grpc.ServiceDesc for TrxLocalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrxLocalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trx.trxLocalService",
	HandlerType: (*TrxLocalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTrxWithCard",
			Handler:    _TrxLocalService_AddTrxWithCard_Handler,
		},
		{
			MethodName: "AddTrxWithoutCard",
			Handler:    _TrxLocalService_AddTrxWithoutCard_Handler,
		},
		{
			MethodName: "InquiryTrxWithoutCard",
			Handler:    _TrxLocalService_InquiryTrxWithoutCard_Handler,
		},
		{
			MethodName: "InquiryTrxWithCard",
			Handler:    _TrxLocalService_InquiryTrxWithCard_Handler,
		},
		{
			MethodName: "ConfirmTrx",
			Handler:    _TrxLocalService_ConfirmTrx_Handler,
		},
		{
			MethodName: "ConfirmTrxByPass",
			Handler:    _TrxLocalService_ConfirmTrxByPass_Handler,
		},
		{
			MethodName: "ConfirmSyncTrxToCloud",
			Handler:    _TrxLocalService_ConfirmSyncTrxToCloud_Handler,
		},
		{
			MethodName: "InquiryPayment",
			Handler:    _TrxLocalService_InquiryPayment_Handler,
		},
		{
			MethodName: "InquiryWithCardP3",
			Handler:    _TrxLocalService_InquiryWithCardP3_Handler,
		},
		{
			MethodName: "InquiryPaymentP3",
			Handler:    _TrxLocalService_InquiryPaymentP3_Handler,
		},
		{
			MethodName: "GetTrxListForDocDate",
			Handler:    _TrxLocalService_GetTrxListForDocDate_Handler,
		},
		{
			MethodName: "UpdateStatusManualTrx",
			Handler:    _TrxLocalService_UpdateStatusManualTrx_Handler,
		},
		{
			MethodName: "FindTrxOutstandingByIndex",
			Handler:    _TrxLocalService_FindTrxOutstandingByIndex_Handler,
		},
		{
			MethodName: "UpdateProductPrice",
			Handler:    _TrxLocalService_UpdateProductPrice_Handler,
		},
		{
			MethodName: "RegisterMember",
			Handler:    _TrxLocalService_RegisterMember_Handler,
		},
		{
			MethodName: "DecryptMKey",
			Handler:    _TrxLocalService_DecryptMKey_Handler,
		},
		{
			MethodName: "UpdateAutoClearTrx",
			Handler:    _TrxLocalService_UpdateAutoClearTrx_Handler,
		},
		{
			MethodName: "ConfirmTrxP3",
			Handler:    _TrxLocalService_ConfirmTrxP3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trxlocal.proto",
}