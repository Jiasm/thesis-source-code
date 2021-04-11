package service

import (
	"dao"
	"entity"
)

type UserService struct {
}

var listDao = new(dao.UserDao)

func (p *UserService) CreateAccount(userName, password string, roleId, status uint) int64 {
	id := listDao.Insert(&entity.User{UserName: userName, Password: password, RoleId: roleId, Status: status})
	if id <= 0 {
		return 0
	}
	return id
}

func (p *UserService) FindOne(userName, password string) *entity.User {
	user := listDao.FindOne(userName, password)

	return user
}

func (p *UserService) FindAll(userIdList []uint) []entity.User {
	userList := listDao.FindAll(userIdList)

	return userList
}
