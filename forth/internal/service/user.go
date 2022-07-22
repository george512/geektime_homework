package service

import (
    "context"
    v1 "geektime_homework/forth/api/user/v1"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (us *UserService) AddUser(context.Context, *v1.AddUserRequest) (*v1.AddUserResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (us *UserService) GetUser(context.Context, *v1.GetUserRequest) (*v1.GetUserResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (us *UserService) ListUser(context.Context, *v1.ListUserRequest) (*v1.ListUserResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}