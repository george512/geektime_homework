package biz

import v1 "geektime_homework/forth/api/user/v1"

type User struct {
    Id       string
    Name     string
    Age      string
    Password string
}

type Filter struct {
    Name   string
    MinAge string
}

type UserRepo interface {
    // 添加用户
    Save(user *User) error
    // 通过id查找用户
    Find(id string) (*User, error)
    // 通过条件查找用户
    Search(filter *Filter) ([]*User, error)
}

type UserUseCase struct {
    repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
    return &UserUseCase{
        repo: repo,
    }
}

func (uc *UserUseCase) Add(in *v1.AddUserRequest) (*v1.AddUserResponse, error) {
    return nil, nil
}

func (uc *UserUseCase) Get(in *v1.GetUserRequest) (*v1.GetUserResponse, error) {
    return nil, nil
}

func (uc *UserUseCase) List(in *v1.ListUserRequest) (*v1.ListUserResponse, error) {
    return nil, nil
}
