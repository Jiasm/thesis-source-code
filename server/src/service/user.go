package service

import (
	"dao"
	"entity"
)

type ListService struct {
}

var listDao = new(dao.UserDao)

func (p *ListService) CreateAccount(userName, password string, roleId, status int) int64 {
	id := listDao.Insert(&entity.User{UserName: userName, Password: password, RoleId: roleId, Status: status})
	if id <= 0 {
		return 0
	}
	return id
}

func (p *ListService) FindOne(userName, password string) *entity.User {
	user := listDao.FindOne(userName, password)

	return user
}

func (p *ListService) FindAll(userIdList []int) []entity.User {
	userList := listDao.FindAll(userIdList)

	return userList
}
