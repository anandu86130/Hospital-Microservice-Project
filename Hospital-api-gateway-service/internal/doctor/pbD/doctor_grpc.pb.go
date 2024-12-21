// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: doctor.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DoctorService_DoctorSignup_FullMethodName     = "/pbD.DoctorService/DoctorSignup"
	DoctorService_VerifyOTP_FullMethodName        = "/pbD.DoctorService/VerifyOTP"
	DoctorService_DoctorLogin_FullMethodName      = "/pbD.DoctorService/DoctorLogin"
	DoctorService_ViewProfile_FullMethodName      = "/pbD.DoctorService/ViewProfile"
	DoctorService_EditProfile_FullMethodName      = "/pbD.DoctorService/EditProfile"
	DoctorService_ChangePassword_FullMethodName   = "/pbD.DoctorService/ChangePassword"
	DoctorService_BlockDoctor_FullMethodName      = "/pbD.DoctorService/BlockDoctor"
	DoctorService_UnblockDoctor_FullMethodName    = "/pbD.DoctorService/UnblockDoctor"
	DoctorService_IsVerified_FullMethodName       = "/pbD.DoctorService/IsVerified"
	DoctorService_DoctorList_FullMethodName       = "/pbD.DoctorService/DoctorList"
	DoctorService_AddAvailability_FullMethodName  = "/pbD.DoctorService/AddAvailability"
	DoctorService_EditAvailability_FullMethodName = "/pbD.DoctorService/EditAvailability"
	DoctorService_ViewAvailability_FullMethodName = "/pbD.DoctorService/ViewAvailability"
	DoctorService_UserList_FullMethodName         = "/pbD.DoctorService/UserList"
	DoctorService_ViewAppointment_FullMethodName  = "/pbD.DoctorService/ViewAppointment"
	DoctorService_AddPrescription_FullMethodName  = "/pbD.DoctorService/AddPrescription"
)

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	DoctorSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error)
	VerifyOTP(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error)
	DoctorLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error)
	ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DoctorProfile, error)
	EditProfile(ctx context.Context, in *DoctorProfile, opts ...grpc.CallOption) (*DoctorProfile, error)
	ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error)
	BlockDoctor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	UnblockDoctor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	IsVerified(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error)
	DoctorList(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*DoctorListResponse, error)
	AddAvailability(ctx context.Context, in *Availability, opts ...grpc.CallOption) (*Response, error)
	EditAvailability(ctx context.Context, in *Availability, opts ...grpc.CallOption) (*Response, error)
	ViewAvailability(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*AvailabilityListResponse, error)
	UserList(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*UserListResponse, error)
	ViewAppointment(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AppointmentList, error)
	AddPrescription(ctx context.Context, in *Prescription, opts ...grpc.CallOption) (*Response, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) DoctorSignup(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_DoctorSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) VerifyOTP(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_VerifyOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_DoctorLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ViewProfile(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DoctorProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoctorProfile)
	err := c.cc.Invoke(ctx, DoctorService_ViewProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) EditProfile(ctx context.Context, in *DoctorProfile, opts ...grpc.CallOption) (*DoctorProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoctorProfile)
	err := c.cc.Invoke(ctx, DoctorService_EditProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ChangePassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) BlockDoctor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_BlockDoctor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UnblockDoctor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_UnblockDoctor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) IsVerified(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_IsVerified_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorList(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*DoctorListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoctorListResponse)
	err := c.cc.Invoke(ctx, DoctorService_DoctorList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddAvailability(ctx context.Context, in *Availability, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_AddAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) EditAvailability(ctx context.Context, in *Availability, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_EditAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ViewAvailability(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*AvailabilityListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvailabilityListResponse)
	err := c.cc.Invoke(ctx, DoctorService_ViewAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UserList(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*UserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, DoctorService_UserList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ViewAppointment(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AppointmentList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppointmentList)
	err := c.cc.Invoke(ctx, DoctorService_ViewAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddPrescription(ctx context.Context, in *Prescription, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, DoctorService_AddPrescription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility.
type DoctorServiceServer interface {
	DoctorSignup(context.Context, *Signup) (*Response, error)
	VerifyOTP(context.Context, *OTP) (*Response, error)
	DoctorLogin(context.Context, *Login) (*Response, error)
	ViewProfile(context.Context, *ID) (*DoctorProfile, error)
	EditProfile(context.Context, *DoctorProfile) (*DoctorProfile, error)
	ChangePassword(context.Context, *Password) (*Response, error)
	BlockDoctor(context.Context, *ID) (*Response, error)
	UnblockDoctor(context.Context, *ID) (*Response, error)
	IsVerified(context.Context, *ID) (*Response, error)
	DoctorList(context.Context, *NoParam) (*DoctorListResponse, error)
	AddAvailability(context.Context, *Availability) (*Response, error)
	EditAvailability(context.Context, *Availability) (*Response, error)
	ViewAvailability(context.Context, *NoParam) (*AvailabilityListResponse, error)
	UserList(context.Context, *NoParam) (*UserListResponse, error)
	ViewAppointment(context.Context, *ID) (*AppointmentList, error)
	AddPrescription(context.Context, *Prescription) (*Response, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDoctorServiceServer struct{}

func (UnimplementedDoctorServiceServer) DoctorSignup(context.Context, *Signup) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorSignup not implemented")
}
func (UnimplementedDoctorServiceServer) VerifyOTP(context.Context, *OTP) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorLogin(context.Context, *Login) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorLogin not implemented")
}
func (UnimplementedDoctorServiceServer) ViewProfile(context.Context, *ID) (*DoctorProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProfile not implemented")
}
func (UnimplementedDoctorServiceServer) EditProfile(context.Context, *DoctorProfile) (*DoctorProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedDoctorServiceServer) ChangePassword(context.Context, *Password) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedDoctorServiceServer) BlockDoctor(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) UnblockDoctor(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) IsVerified(context.Context, *ID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsVerified not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorList(context.Context, *NoParam) (*DoctorListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorList not implemented")
}
func (UnimplementedDoctorServiceServer) AddAvailability(context.Context, *Availability) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAvailability not implemented")
}
func (UnimplementedDoctorServiceServer) EditAvailability(context.Context, *Availability) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAvailability not implemented")
}
func (UnimplementedDoctorServiceServer) ViewAvailability(context.Context, *NoParam) (*AvailabilityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAvailability not implemented")
}
func (UnimplementedDoctorServiceServer) UserList(context.Context, *NoParam) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedDoctorServiceServer) ViewAppointment(context.Context, *ID) (*AppointmentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAppointment not implemented")
}
func (UnimplementedDoctorServiceServer) AddPrescription(context.Context, *Prescription) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPrescription not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}
func (UnimplementedDoctorServiceServer) testEmbeddedByValue()                       {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	// If the following call pancis, it indicates UnimplementedDoctorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_DoctorSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_DoctorSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorSignup(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).VerifyOTP(ctx, req.(*OTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_DoctorLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorLogin(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ViewProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ViewProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ViewProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ViewProfile(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_EditProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).EditProfile(ctx, req.(*DoctorProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Password)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ChangePassword(ctx, req.(*Password))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_BlockDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).BlockDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_BlockDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).BlockDoctor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UnblockDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UnblockDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_UnblockDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UnblockDoctor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_IsVerified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).IsVerified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_IsVerified_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).IsVerified(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_DoctorList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorList(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_AddAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Availability)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_AddAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddAvailability(ctx, req.(*Availability))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_EditAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Availability)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).EditAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_EditAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).EditAvailability(ctx, req.(*Availability))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ViewAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ViewAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ViewAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ViewAvailability(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UserList(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ViewAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ViewAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ViewAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ViewAppointment(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_AddPrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddPrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_AddPrescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddPrescription(ctx, req.(*Prescription))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbD.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoctorSignup",
			Handler:    _DoctorService_DoctorSignup_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _DoctorService_VerifyOTP_Handler,
		},
		{
			MethodName: "DoctorLogin",
			Handler:    _DoctorService_DoctorLogin_Handler,
		},
		{
			MethodName: "ViewProfile",
			Handler:    _DoctorService_ViewProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _DoctorService_EditProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _DoctorService_ChangePassword_Handler,
		},
		{
			MethodName: "BlockDoctor",
			Handler:    _DoctorService_BlockDoctor_Handler,
		},
		{
			MethodName: "UnblockDoctor",
			Handler:    _DoctorService_UnblockDoctor_Handler,
		},
		{
			MethodName: "IsVerified",
			Handler:    _DoctorService_IsVerified_Handler,
		},
		{
			MethodName: "DoctorList",
			Handler:    _DoctorService_DoctorList_Handler,
		},
		{
			MethodName: "AddAvailability",
			Handler:    _DoctorService_AddAvailability_Handler,
		},
		{
			MethodName: "EditAvailability",
			Handler:    _DoctorService_EditAvailability_Handler,
		},
		{
			MethodName: "ViewAvailability",
			Handler:    _DoctorService_ViewAvailability_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _DoctorService_UserList_Handler,
		},
		{
			MethodName: "ViewAppointment",
			Handler:    _DoctorService_ViewAppointment_Handler,
		},
		{
			MethodName: "AddPrescription",
			Handler:    _DoctorService_AddPrescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor.proto",
}