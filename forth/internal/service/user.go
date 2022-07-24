package service

import (
	"context"
	v1 "geektime_homework/forth/api/user/v1"
)

func (us *UserService) AddUser(ctx context.Context, in *v1.AddUserRequest) (*v1.AddUserResponse, error) {
	id, err := us.uc.Add(in.GetUser())
	if err != nil {
		return nil, err
	}

	return &v1.AddUserResponse{
		Id: id,
	}, nil
}
func (us *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := us.uc.Get(in.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.GetUserResponse{
		User: user,
	}, nil
}
func (us *UserService) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	users, err := us.uc.List(in.GetFilter())
	if err != nil {
		return nil, err
	}

	return &v1.ListUserResponse{
		User: users,
	}, nil
}
