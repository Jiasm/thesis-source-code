package service

import (
	"dao"
	"entity"
)

type ListService struct {
}

var listDao = new(dao.ListDao)

func (p *ListService) Insert(userName, password string, roleId, status uint) int64 {
	id := listDao.Insert(&entity.User{ UserName: userName, Password: password, RoleId: roleId, Status: status})
	if id <= 0 {
		return 0
	}
	return id
}
