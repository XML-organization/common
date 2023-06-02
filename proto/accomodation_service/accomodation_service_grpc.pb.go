// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: accomodation_service.proto

package accomodation

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

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccomodationSearchResponse, error)
	GetAllAccomodations(ctx context.Context, in *GetAllAccomodationsRequest, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error)
	GetAutoApprovalForAccommodation(ctx context.Context, in *AutoApprovalRequest, opts ...grpc.CallOption) (*AutoApprovalResponse, error)
	GetAllAvailability(ctx context.Context, in *GetAllAvailabilityRequest, opts ...grpc.CallOption) (*GetAllAvailabilityResponse, error)
	GetOneAccomodation(ctx context.Context, in *GetOneAccomodationRequest, opts ...grpc.CallOption) (*GetOneAccomodationResponse, error)
	GetAllAccomodationIdsByHostId(ctx context.Context, in *GetAllAccomodationIdsByHostIdRequest, opts ...grpc.CallOption) (*GetAllAccomodationIdsByHostIdResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error) {
	out := new(UpdateAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/UpdateAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccomodationSearchResponse, error) {
	out := new(AccomodationSearchResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAccomodations(ctx context.Context, in *GetAllAccomodationsRequest, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error) {
	out := new(GetAllAccomodationsResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/GetAllAccomodations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAutoApprovalForAccommodation(ctx context.Context, in *AutoApprovalRequest, opts ...grpc.CallOption) (*AutoApprovalResponse, error) {
	out := new(AutoApprovalResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/GetAutoApprovalForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAvailability(ctx context.Context, in *GetAllAvailabilityRequest, opts ...grpc.CallOption) (*GetAllAvailabilityResponse, error) {
	out := new(GetAllAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/GetAllAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetOneAccomodation(ctx context.Context, in *GetOneAccomodationRequest, opts ...grpc.CallOption) (*GetOneAccomodationResponse, error) {
	out := new(GetOneAccomodationResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/GetOneAccomodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAccomodationIdsByHostId(ctx context.Context, in *GetAllAccomodationIdsByHostIdRequest, opts ...grpc.CallOption) (*GetAllAccomodationIdsByHostIdResponse, error) {
	out := new(GetAllAccomodationIdsByHostIdResponse)
	err := c.cc.Invoke(ctx, "/accomodation.AccommodationService/GetAllAccomodationIdsByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	UpdateAvailability(context.Context, *UpdateAvailabilityRequest) (*UpdateAvailabilityResponse, error)
	Search(context.Context, *SearchRequest) (*AccomodationSearchResponse, error)
	GetAllAccomodations(context.Context, *GetAllAccomodationsRequest) (*GetAllAccomodationsResponse, error)
	GetAutoApprovalForAccommodation(context.Context, *AutoApprovalRequest) (*AutoApprovalResponse, error)
	GetAllAvailability(context.Context, *GetAllAvailabilityRequest) (*GetAllAvailabilityResponse, error)
	GetOneAccomodation(context.Context, *GetOneAccomodationRequest) (*GetOneAccomodationResponse, error)
	GetAllAccomodationIdsByHostId(context.Context, *GetAllAccomodationIdsByHostIdRequest) (*GetAllAccomodationIdsByHostIdResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdateAvailability(context.Context, *UpdateAvailabilityRequest) (*UpdateAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailability not implemented")
}
func (UnimplementedAccommodationServiceServer) Search(context.Context, *SearchRequest) (*AccomodationSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllAccomodations(context.Context, *GetAllAccomodationsRequest) (*GetAllAccomodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccomodations not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAutoApprovalForAccommodation(context.Context, *AutoApprovalRequest) (*AutoApprovalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoApprovalForAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllAvailability(context.Context, *GetAllAvailabilityRequest) (*GetAllAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAvailability not implemented")
}
func (UnimplementedAccommodationServiceServer) GetOneAccomodation(context.Context, *GetOneAccomodationRequest) (*GetOneAccomodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneAccomodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllAccomodationIdsByHostId(context.Context, *GetAllAccomodationIdsByHostIdRequest) (*GetAllAccomodationIdsByHostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccomodationIdsByHostId not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/UpdateAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAvailability(ctx, req.(*UpdateAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllAccomodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccomodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllAccomodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/GetAllAccomodations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAccomodations(ctx, req.(*GetAllAccomodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAutoApprovalForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutoApprovalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAutoApprovalForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/GetAutoApprovalForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAutoApprovalForAccommodation(ctx, req.(*AutoApprovalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/GetAllAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAvailability(ctx, req.(*GetAllAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetOneAccomodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneAccomodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetOneAccomodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/GetOneAccomodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetOneAccomodation(ctx, req.(*GetOneAccomodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllAccomodationIdsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccomodationIdsByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllAccomodationIdsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accomodation.AccommodationService/GetAllAccomodationIdsByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAccomodationIdsByHostId(ctx, req.(*GetAllAccomodationIdsByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accomodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccommodationService_Create_Handler,
		},
		{
			MethodName: "UpdateAvailability",
			Handler:    _AccommodationService_UpdateAvailability_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _AccommodationService_Search_Handler,
		},
		{
			MethodName: "GetAllAccomodations",
			Handler:    _AccommodationService_GetAllAccomodations_Handler,
		},
		{
			MethodName: "GetAutoApprovalForAccommodation",
			Handler:    _AccommodationService_GetAutoApprovalForAccommodation_Handler,
		},
		{
			MethodName: "GetAllAvailability",
			Handler:    _AccommodationService_GetAllAvailability_Handler,
		},
		{
			MethodName: "GetOneAccomodation",
			Handler:    _AccommodationService_GetOneAccomodation_Handler,
		},
		{
			MethodName: "GetAllAccomodationIdsByHostId",
			Handler:    _AccommodationService_GetAllAccomodationIdsByHostId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accomodation_service.proto",
}
