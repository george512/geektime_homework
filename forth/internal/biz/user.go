package biz

import (
	"errors"
	v1 "geektime_homework/forth/api/user/v1"
)

var ErrAlreadyExists = errors.New("user already exists")
var ErrNotFound = errors.New("user not found")

type User struct {
	Id       string
	Name     string
	Age      int
	Password string
}

type Filter struct {
	Name   string
	MinAge int
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

func (uc *UserUseCase) Add(in *v1.User) (string, error) {

	// deep copy DTO -> DO
	doUser := &User{
		Id:       in.Id,
		Name:     in.Name,
		Age:      int(in.Age),
		Password: in.Password,
	}

	if err := uc.repo.Save(doUser); err != nil {
		return "", err
	}

	return doUser.Id, nil
}

func (uc *UserUseCase) Get(in string) (*v1.User, error) {

	doUser, err := uc.repo.Find(in)
	if err != nil {
		return nil, err
	}

	// deep copy, DO -> DTO
	rsp := &v1.User{
		Id:       doUser.Id,
		Name:     doUser.Name,
		Age:      uint32(doUser.Age),
		Password: doUser.Password,
	}

	return rsp, nil
}

func (uc *UserUseCase) List(in *v1.Filter) ([]*v1.User, error) {

	// deep copy
	doFilter := &Filter{
		Name:   in.Name,
		MinAge: int(in.MinAge),
	}

	doUsers, err := uc.repo.Search(doFilter)
	if err != nil {
		return nil, err
	}

	// deep copy, DO -> DTO
	dtoUsers := make([]*v1.User, 0)
	for i := range doUsers {
		dtoUser := &v1.User{
			Id:       doUsers[i].Id,
			Name:     doUsers[i].Name,
			Age:      uint32(doUsers[i].Age),
			Password: doUsers[i].Password,
		}
		dtoUsers = append(dtoUsers, dtoUser)
	}

	return dtoUsers, nil
}
