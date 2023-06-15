// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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

const (
	AccommodationService_Create_FullMethodName                           = "/accomodation.AccommodationService/Create"
	AccommodationService_UpdateAvailability_FullMethodName               = "/accomodation.AccommodationService/UpdateAvailability"
	AccommodationService_Search_FullMethodName                           = "/accomodation.AccommodationService/Search"
	AccommodationService_GetAllAccomodations_FullMethodName              = "/accomodation.AccommodationService/GetAllAccomodations"
	AccommodationService_GetAccomodations_FullMethodName                 = "/accomodation.AccommodationService/GetAccomodations"
	AccommodationService_GetAutoApprovalForAccommodation_FullMethodName  = "/accomodation.AccommodationService/GetAutoApprovalForAccommodation"
	AccommodationService_GetAllAvailability_FullMethodName               = "/accomodation.AccommodationService/GetAllAvailability"
	AccommodationService_GetOneAccomodation_FullMethodName               = "/accomodation.AccommodationService/GetOneAccomodation"
	AccommodationService_CheckIfGuestHasReservationInPast_FullMethodName = "/accomodation.AccommodationService/CheckIfGuestHasReservationInPast"
	AccommodationService_GradeHost_FullMethodName                        = "/accomodation.AccommodationService/GradeHost"
	AccommodationService_GetAccommodationRecommendations_FullMethodName  = "/accomodation.AccommodationService/GetAccommodationRecommendations"
	AccommodationService_GetGradesByAccomodationId_FullMethodName        = "/accomodation.AccommodationService/GetGradesByAccomodationId"
	AccommodationService_EditGrade_FullMethodName                        = "/accomodation.AccommodationService/EditGrade"
	AccommodationService_DeleteGrade_FullMethodName                      = "/accomodation.AccommodationService/DeleteGrade"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccomodationSearchResponse, error)
	GetAllAccomodations(ctx context.Context, in *GetAllAccomodationsRequest, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error)
	GetAccomodations(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error)
	GetAutoApprovalForAccommodation(ctx context.Context, in *AutoApprovalRequest, opts ...grpc.CallOption) (*AutoApprovalResponse, error)
	GetAllAvailability(ctx context.Context, in *GetAllAvailabilityRequest, opts ...grpc.CallOption) (*GetAllAvailabilityResponse, error)
	GetOneAccomodation(ctx context.Context, in *GetOneAccomodationRequest, opts ...grpc.CallOption) (*GetOneAccomodationResponse, error)
	CheckIfGuestHasReservationInPast(ctx context.Context, in *CheckIfGuestHasReservationInPastRequest, opts ...grpc.CallOption) (*CheckIfGuestHasReservationInPastResponse, error)
	GradeHost(ctx context.Context, in *GradeHostRequest, opts ...grpc.CallOption) (*GradeHostResponse, error)
	GetAccommodationRecommendations(ctx context.Context, in *GetAccommodationRecommendationsRequest, opts ...grpc.CallOption) (*GetAccommodationRecommendationsResponse, error)
	GetGradesByAccomodationId(ctx context.Context, in *GradesByAccomodationIdRequest, opts ...grpc.CallOption) (*GradesByAccomodationIdResponse, error)
	EditGrade(ctx context.Context, in *EditGradeRequest, opts ...grpc.CallOption) (*EditGradeResponse, error)
	DeleteGrade(ctx context.Context, in *DeleteGradeRequest, opts ...grpc.CallOption) (*DeleteGradeResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error) {
	out := new(UpdateAvailabilityResponse)
	err := c.cc.Invoke(ctx, AccommodationService_UpdateAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccomodationSearchResponse, error) {
	out := new(AccomodationSearchResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAccomodations(ctx context.Context, in *GetAllAccomodationsRequest, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error) {
	out := new(GetAllAccomodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllAccomodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccomodations(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*GetAllAccomodationsResponse, error) {
	out := new(GetAllAccomodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAccomodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAutoApprovalForAccommodation(ctx context.Context, in *AutoApprovalRequest, opts ...grpc.CallOption) (*AutoApprovalResponse, error) {
	out := new(AutoApprovalResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAutoApprovalForAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAvailability(ctx context.Context, in *GetAllAvailabilityRequest, opts ...grpc.CallOption) (*GetAllAvailabilityResponse, error) {
	out := new(GetAllAvailabilityResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetOneAccomodation(ctx context.Context, in *GetOneAccomodationRequest, opts ...grpc.CallOption) (*GetOneAccomodationResponse, error) {
	out := new(GetOneAccomodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetOneAccomodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CheckIfGuestHasReservationInPast(ctx context.Context, in *CheckIfGuestHasReservationInPastRequest, opts ...grpc.CallOption) (*CheckIfGuestHasReservationInPastResponse, error) {
	out := new(CheckIfGuestHasReservationInPastResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CheckIfGuestHasReservationInPast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GradeHost(ctx context.Context, in *GradeHostRequest, opts ...grpc.CallOption) (*GradeHostResponse, error) {
	out := new(GradeHostResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GradeHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationRecommendations(ctx context.Context, in *GetAccommodationRecommendationsRequest, opts ...grpc.CallOption) (*GetAccommodationRecommendationsResponse, error) {
	out := new(GetAccommodationRecommendationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAccommodationRecommendations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetGradesByAccomodationId(ctx context.Context, in *GradesByAccomodationIdRequest, opts ...grpc.CallOption) (*GradesByAccomodationIdResponse, error) {
	out := new(GradesByAccomodationIdResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetGradesByAccomodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) EditGrade(ctx context.Context, in *EditGradeRequest, opts ...grpc.CallOption) (*EditGradeResponse, error) {
	out := new(EditGradeResponse)
	err := c.cc.Invoke(ctx, AccommodationService_EditGrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteGrade(ctx context.Context, in *DeleteGradeRequest, opts ...grpc.CallOption) (*DeleteGradeResponse, error) {
	out := new(DeleteGradeResponse)
	err := c.cc.Invoke(ctx, AccommodationService_DeleteGrade_FullMethodName, in, out, opts...)
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
	GetAccomodations(context.Context, *EmptyRequst) (*GetAllAccomodationsResponse, error)
	GetAutoApprovalForAccommodation(context.Context, *AutoApprovalRequest) (*AutoApprovalResponse, error)
	GetAllAvailability(context.Context, *GetAllAvailabilityRequest) (*GetAllAvailabilityResponse, error)
	GetOneAccomodation(context.Context, *GetOneAccomodationRequest) (*GetOneAccomodationResponse, error)
	CheckIfGuestHasReservationInPast(context.Context, *CheckIfGuestHasReservationInPastRequest) (*CheckIfGuestHasReservationInPastResponse, error)
	GradeHost(context.Context, *GradeHostRequest) (*GradeHostResponse, error)
	GetAccommodationRecommendations(context.Context, *GetAccommodationRecommendationsRequest) (*GetAccommodationRecommendationsResponse, error)
	GetGradesByAccomodationId(context.Context, *GradesByAccomodationIdRequest) (*GradesByAccomodationIdResponse, error)
	EditGrade(context.Context, *EditGradeRequest) (*EditGradeResponse, error)
	DeleteGrade(context.Context, *DeleteGradeRequest) (*DeleteGradeResponse, error)
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
func (UnimplementedAccommodationServiceServer) GetAccomodations(context.Context, *EmptyRequst) (*GetAllAccomodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccomodations not implemented")
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
func (UnimplementedAccommodationServiceServer) CheckIfGuestHasReservationInPast(context.Context, *CheckIfGuestHasReservationInPastRequest) (*CheckIfGuestHasReservationInPastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfGuestHasReservationInPast not implemented")
}
func (UnimplementedAccommodationServiceServer) GradeHost(context.Context, *GradeHostRequest) (*GradeHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GradeHost not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationRecommendations(context.Context, *GetAccommodationRecommendationsRequest) (*GetAccommodationRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationRecommendations not implemented")
}
func (UnimplementedAccommodationServiceServer) GetGradesByAccomodationId(context.Context, *GradesByAccomodationIdRequest) (*GradesByAccomodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGradesByAccomodationId not implemented")
}
func (UnimplementedAccommodationServiceServer) EditGrade(context.Context, *EditGradeRequest) (*EditGradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditGrade not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteGrade(context.Context, *DeleteGradeRequest) (*DeleteGradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGrade not implemented")
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
		FullMethod: AccommodationService_Create_FullMethodName,
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
		FullMethod: AccommodationService_UpdateAvailability_FullMethodName,
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
		FullMethod: AccommodationService_Search_FullMethodName,
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
		FullMethod: AccommodationService_GetAllAccomodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAccomodations(ctx, req.(*GetAllAccomodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccomodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccomodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAccomodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccomodations(ctx, req.(*EmptyRequst))
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
		FullMethod: AccommodationService_GetAutoApprovalForAccommodation_FullMethodName,
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
		FullMethod: AccommodationService_GetAllAvailability_FullMethodName,
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
		FullMethod: AccommodationService_GetOneAccomodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetOneAccomodation(ctx, req.(*GetOneAccomodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CheckIfGuestHasReservationInPast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfGuestHasReservationInPastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CheckIfGuestHasReservationInPast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CheckIfGuestHasReservationInPast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CheckIfGuestHasReservationInPast(ctx, req.(*CheckIfGuestHasReservationInPastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GradeHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GradeHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GradeHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GradeHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GradeHost(ctx, req.(*GradeHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAccommodationRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationRecommendations(ctx, req.(*GetAccommodationRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetGradesByAccomodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GradesByAccomodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetGradesByAccomodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetGradesByAccomodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetGradesByAccomodationId(ctx, req.(*GradesByAccomodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_EditGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).EditGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_EditGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).EditGrade(ctx, req.(*EditGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_DeleteGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteGrade(ctx, req.(*DeleteGradeRequest))
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
			MethodName: "GetAccomodations",
			Handler:    _AccommodationService_GetAccomodations_Handler,
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
			MethodName: "CheckIfGuestHasReservationInPast",
			Handler:    _AccommodationService_CheckIfGuestHasReservationInPast_Handler,
		},
		{
			MethodName: "GradeHost",
			Handler:    _AccommodationService_GradeHost_Handler,
		},
		{
			MethodName: "GetAccommodationRecommendations",
			Handler:    _AccommodationService_GetAccommodationRecommendations_Handler,
		},
		{
			MethodName: "GetGradesByAccomodationId",
			Handler:    _AccommodationService_GetGradesByAccomodationId_Handler,
		},
		{
			MethodName: "EditGrade",
			Handler:    _AccommodationService_EditGrade_Handler,
		},
		{
			MethodName: "DeleteGrade",
			Handler:    _AccommodationService_DeleteGrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accomodation_service.proto",
}
