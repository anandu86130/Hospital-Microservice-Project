package service

import (
	"context"

	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
	userpb "github.com/anandu86130/Hospital-doctor-service/internal/user/pbU"
)

func (a *DoctorService) GetUserListService(p *pb.NoParam) (*pb.UserListResponse, error) {
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
