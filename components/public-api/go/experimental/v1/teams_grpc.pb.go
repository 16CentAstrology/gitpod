// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gitpod/experimental/v1/teams.proto

package v1

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

// TeamsServiceClient is the client API for TeamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamsServiceClient interface {
	// CreateTeam creates a new Team.
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	// GetTeam retrieves a single Team.
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error)
	// ListTeams lists the caller has access to.
	ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error)
	// DeleteTeam deletes the specified team.
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
	// GetTeamInvitation retrieves the invitation for a Team.
	GetTeamInvitation(ctx context.Context, in *GetTeamInvitationRequest, opts ...grpc.CallOption) (*GetTeamInvitationResponse, error)
	// JoinTeam makes the caller a TeamMember of the Team.
	JoinTeam(ctx context.Context, in *JoinTeamRequest, opts ...grpc.CallOption) (*JoinTeamResponse, error)
	// ResetTeamInvitation resets the invitation_id for a Team.
	ResetTeamInvitation(ctx context.Context, in *ResetTeamInvitationRequest, opts ...grpc.CallOption) (*ResetTeamInvitationResponse, error)
	// ListTeamMembers lists the members of a Team.
	ListTeamMembers(ctx context.Context, in *ListTeamMembersRequest, opts ...grpc.CallOption) (*ListTeamMembersResponse, error)
	// UpdateTeamMember updates team membership properties.
	UpdateTeamMember(ctx context.Context, in *UpdateTeamMemberRequest, opts ...grpc.CallOption) (*UpdateTeamMemberResponse, error)
	// DeleteTeamMember removes a TeamMember from the Team.
	DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...grpc.CallOption) (*DeleteTeamMemberResponse, error)
}

type teamsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsServiceClient(cc grpc.ClientConnInterface) TeamsServiceClient {
	return &teamsServiceClient{cc}
}

func (c *teamsServiceClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	out := new(GetTeamResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error) {
	out := new(ListTeamsResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/ListTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	out := new(DeleteTeamResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) GetTeamInvitation(ctx context.Context, in *GetTeamInvitationRequest, opts ...grpc.CallOption) (*GetTeamInvitationResponse, error) {
	out := new(GetTeamInvitationResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/GetTeamInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) JoinTeam(ctx context.Context, in *JoinTeamRequest, opts ...grpc.CallOption) (*JoinTeamResponse, error) {
	out := new(JoinTeamResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/JoinTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) ResetTeamInvitation(ctx context.Context, in *ResetTeamInvitationRequest, opts ...grpc.CallOption) (*ResetTeamInvitationResponse, error) {
	out := new(ResetTeamInvitationResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/ResetTeamInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) ListTeamMembers(ctx context.Context, in *ListTeamMembersRequest, opts ...grpc.CallOption) (*ListTeamMembersResponse, error) {
	out := new(ListTeamMembersResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/ListTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) UpdateTeamMember(ctx context.Context, in *UpdateTeamMemberRequest, opts ...grpc.CallOption) (*UpdateTeamMemberResponse, error) {
	out := new(UpdateTeamMemberResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/UpdateTeamMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...grpc.CallOption) (*DeleteTeamMemberResponse, error) {
	out := new(DeleteTeamMemberResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.TeamsService/DeleteTeamMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServiceServer is the server API for TeamsService service.
// All implementations must embed UnimplementedTeamsServiceServer
// for forward compatibility
type TeamsServiceServer interface {
	// CreateTeam creates a new Team.
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	// GetTeam retrieves a single Team.
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error)
	// ListTeams lists the caller has access to.
	ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error)
	// DeleteTeam deletes the specified team.
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
	// GetTeamInvitation retrieves the invitation for a Team.
	GetTeamInvitation(context.Context, *GetTeamInvitationRequest) (*GetTeamInvitationResponse, error)
	// JoinTeam makes the caller a TeamMember of the Team.
	JoinTeam(context.Context, *JoinTeamRequest) (*JoinTeamResponse, error)
	// ResetTeamInvitation resets the invitation_id for a Team.
	ResetTeamInvitation(context.Context, *ResetTeamInvitationRequest) (*ResetTeamInvitationResponse, error)
	// ListTeamMembers lists the members of a Team.
	ListTeamMembers(context.Context, *ListTeamMembersRequest) (*ListTeamMembersResponse, error)
	// UpdateTeamMember updates team membership properties.
	UpdateTeamMember(context.Context, *UpdateTeamMemberRequest) (*UpdateTeamMemberResponse, error)
	// DeleteTeamMember removes a TeamMember from the Team.
	DeleteTeamMember(context.Context, *DeleteTeamMemberRequest) (*DeleteTeamMemberResponse, error)
	mustEmbedUnimplementedTeamsServiceServer()
}

// UnimplementedTeamsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamsServiceServer struct {
}

func (UnimplementedTeamsServiceServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamsServiceServer) GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedTeamsServiceServer) ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (UnimplementedTeamsServiceServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamsServiceServer) GetTeamInvitation(context.Context, *GetTeamInvitationRequest) (*GetTeamInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamInvitation not implemented")
}
func (UnimplementedTeamsServiceServer) JoinTeam(context.Context, *JoinTeamRequest) (*JoinTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTeam not implemented")
}
func (UnimplementedTeamsServiceServer) ResetTeamInvitation(context.Context, *ResetTeamInvitationRequest) (*ResetTeamInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetTeamInvitation not implemented")
}
func (UnimplementedTeamsServiceServer) ListTeamMembers(context.Context, *ListTeamMembersRequest) (*ListTeamMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeamMembers not implemented")
}
func (UnimplementedTeamsServiceServer) UpdateTeamMember(context.Context, *UpdateTeamMemberRequest) (*UpdateTeamMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamMember not implemented")
}
func (UnimplementedTeamsServiceServer) DeleteTeamMember(context.Context, *DeleteTeamMemberRequest) (*DeleteTeamMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeamMember not implemented")
}
func (UnimplementedTeamsServiceServer) mustEmbedUnimplementedTeamsServiceServer() {}

// UnsafeTeamsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamsServiceServer will
// result in compilation errors.
type UnsafeTeamsServiceServer interface {
	mustEmbedUnimplementedTeamsServiceServer()
}

func RegisterTeamsServiceServer(s grpc.ServiceRegistrar, srv TeamsServiceServer) {
	s.RegisterService(&TeamsService_ServiceDesc, srv)
}

func _TeamsService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/ListTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).ListTeams(ctx, req.(*ListTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_GetTeamInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).GetTeamInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/GetTeamInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).GetTeamInvitation(ctx, req.(*GetTeamInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_JoinTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).JoinTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/JoinTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).JoinTeam(ctx, req.(*JoinTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_ResetTeamInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetTeamInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).ResetTeamInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/ResetTeamInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).ResetTeamInvitation(ctx, req.(*ResetTeamInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_ListTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).ListTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/ListTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).ListTeamMembers(ctx, req.(*ListTeamMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_UpdateTeamMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).UpdateTeamMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/UpdateTeamMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).UpdateTeamMember(ctx, req.(*UpdateTeamMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_DeleteTeamMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).DeleteTeamMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.TeamsService/DeleteTeamMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).DeleteTeamMember(ctx, req.(*DeleteTeamMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamsService_ServiceDesc is the grpc.ServiceDesc for TeamsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.experimental.v1.TeamsService",
	HandlerType: (*TeamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _TeamsService_CreateTeam_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _TeamsService_GetTeam_Handler,
		},
		{
			MethodName: "ListTeams",
			Handler:    _TeamsService_ListTeams_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _TeamsService_DeleteTeam_Handler,
		},
		{
			MethodName: "GetTeamInvitation",
			Handler:    _TeamsService_GetTeamInvitation_Handler,
		},
		{
			MethodName: "JoinTeam",
			Handler:    _TeamsService_JoinTeam_Handler,
		},
		{
			MethodName: "ResetTeamInvitation",
			Handler:    _TeamsService_ResetTeamInvitation_Handler,
		},
		{
			MethodName: "ListTeamMembers",
			Handler:    _TeamsService_ListTeamMembers_Handler,
		},
		{
			MethodName: "UpdateTeamMember",
			Handler:    _TeamsService_UpdateTeamMember_Handler,
		},
		{
			MethodName: "DeleteTeamMember",
			Handler:    _TeamsService_DeleteTeamMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/experimental/v1/teams.proto",
}
