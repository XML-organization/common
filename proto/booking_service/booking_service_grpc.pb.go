// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: booking_service.proto

package booking

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

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetAll(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*BookingResponse, error)
	GetAllOnPending(ctx context.Context, in *GetAllPendingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	Decline(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	Confirm(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetAllReservations(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	GuestHasReservationInPast(ctx context.Context, in *GuestHasReservationInPastRequest, opts ...grpc.CallOption) (*GuestHasReservationInPastResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAll(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllOnPending(ctx context.Context, in *GetAllPendingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllOnPending", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) Decline(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/Decline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) Confirm(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllReservations(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GuestHasReservationInPast(ctx context.Context, in *GuestHasReservationInPastRequest, opts ...grpc.CallOption) (*GuestHasReservationInPastResponse, error) {
	out := new(GuestHasReservationInPastResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GuestHasReservationInPast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetAll(context.Context, *EmptyRequst) (*BookingResponse, error)
	GetAllOnPending(context.Context, *GetAllPendingRequest) (*BookingResponse, error)
	Decline(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	Confirm(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetAllReservations(context.Context, *ReservationRequest) (*ReservationResponse, error)
	GuestHasReservationInPast(context.Context, *GuestHasReservationInPastRequest) (*GuestHasReservationInPastResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetAll(context.Context, *EmptyRequst) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBookingServiceServer) GetAllOnPending(context.Context, *GetAllPendingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOnPending not implemented")
}
func (UnimplementedBookingServiceServer) Decline(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decline not implemented")
}
func (UnimplementedBookingServiceServer) Confirm(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedBookingServiceServer) GetAllReservations(context.Context, *ReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedBookingServiceServer) GuestHasReservationInPast(context.Context, *GuestHasReservationInPastRequest) (*GuestHasReservationInPastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestHasReservationInPast not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAll(ctx, req.(*EmptyRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllOnPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllOnPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllOnPending",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllOnPending(ctx, req.(*GetAllPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_Decline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Decline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/Decline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Decline(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Confirm(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllReservations(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GuestHasReservationInPast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestHasReservationInPastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GuestHasReservationInPast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GuestHasReservationInPast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GuestHasReservationInPast(ctx, req.(*GuestHasReservationInPastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookingService_GetAll_Handler,
		},
		{
			MethodName: "GetAllOnPending",
			Handler:    _BookingService_GetAllOnPending_Handler,
		},
		{
			MethodName: "Decline",
			Handler:    _BookingService_Decline_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _BookingService_Confirm_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _BookingService_GetAllReservations_Handler,
		},
		{
			MethodName: "GuestHasReservationInPast",
			Handler:    _BookingService_GuestHasReservationInPast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking_service.proto",
}
