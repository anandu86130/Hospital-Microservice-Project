package services

import (
	"context"

	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
	userpb "github.com/anandu86130/Hospital-admin-service/internal/user/pbU"
)

// BlockUserService handle the admin to block the users using the provided information
func (a *AdminService) BlockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}
	_, err := a.UserClient.BlockUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User blocked successfully",
	}, nil
}

// UnblockUserService implements interfaces.AdminServiceInter.
func (a *AdminService) UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}
	_, err := a.UserClient.UnblockUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User unblocked successfully",
	}, nil
}

func (a *AdminService) GetUserListService(p *pb.NoParam) (*pb.UserListResponse, error) {
	ctx := context.Background()
	user := &userpb.NoParam{}

	// Call the UserList method from UserClient
	response, err := a.UserClient.UserList(ctx, user)
	if err != nil {
		return nil, err
	}

	// Prepare the UserListResponse
	var profiles []*pb.Profile
	for _, userProfile := range response.Profiles {
		profiles = append(profiles, &pb.Profile{
			User_ID: userProfile.User_ID,
			Name:    userProfile.Name,
			Email:   userProfile.Email,
			Gender:  userProfile.Gender,
			Age:     userProfile.Age,
			Number:  userProfile.Number,
			Address: userProfile.Address,
		})
	}

	return &pb.UserListResponse{
		Profiles: profiles,
	}, nil
}
