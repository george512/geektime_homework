package service

import (
    v1 "geektime_homework/forth/api/user/v1"
    "geektime_homework/forth/internal/biz"
    "github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewUserService)

type UserService struct {
    v1.UnimplementedUserServiceServer
    uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
    return &UserService{
        uc: uc,
    }
}
