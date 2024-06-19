package service

import "github.com/lrayt/small-sparrow/example/app/dao"

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s UserService) AddUser() {

}

func (s UserService) UserList() {

}
