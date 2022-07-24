package data

import (
	"geektime_homework/forth/internal/biz"
	"strings"
	"sync"
)

type User struct {
	Id       string
	Name     string
	Age      int
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

func (ur *UserRepoInMemory) Save(user *biz.User) error {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if ur.data[user.Id] != nil {
		return biz.ErrAlreadyExists
	}

	// deep copy
	cUser := &User{
		Id:       user.Id,
		Name:     user.Name,
		Age:      user.Age,
		Password: user.Password,
	}

	ur.data[cUser.Id] = cUser
	return nil
}

func (ur *UserRepoInMemory) Find(id string) (*biz.User, error) {

	ur.mutex.RLock()
	defer ur.mutex.RUnlock()
	user := ur.data[id]
	if user == nil {
		return nil, biz.ErrNotFound
	}

	// deep copy
	return &biz.User{
		Id:       user.Id,
		Name:     user.Name,
		Age:      user.Age,
		Password: user.Password,
	}, nil
}

func (ur *UserRepoInMemory) Search(filter *biz.Filter) ([]*biz.User, error) {
	ur.mutex.RLock()
	defer ur.mutex.RUnlock()

	users := make([]*biz.User, 0)
	for _, user := range ur.data {
		if isPassed(filter, user) {
			users = append(users, &biz.User{Id: user.Id, Name: user.Name, Age: user.Age, Password: user.Password})
		}
	}

	return users, nil
}

func isPassed(filter *biz.Filter, user *User) bool {
	if !strings.Contains(user.Name, filter.Name) {
		return false
	}
	if user.Age < filter.MinAge {
		return false
	}
	return true
}
