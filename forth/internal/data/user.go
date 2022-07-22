package data

import (
    "geektime_homework/forth/internal/biz"
    "github.com/google/wire"
    "sync"
)

var ProviderSet = wire.NewSet(NewUserRepoInMemory())

type User struct{
    Id       string
    Name     string
    Age      string
    Password string
}

type UserRepoInMemory struct {
    mutex sync.RWMutex
    data  map[string]*User
}

func NewUserRepoInMemory() biz.UserRepo {
    return &UserRepoInMemory{
        data: make(map[string]*User),
    }
}

func (ur *UserRepoInMemory)Save(user *biz.User) error{
return nil
}

func (ur *UserRepoInMemory)Find(id string) (*biz.User, error){
    return nil, nil
}

func (ur *UserRepoInMemory)Search(filter *biz.Filter) ([]*biz.User, error){
    return nil, nil
}
