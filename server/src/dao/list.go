package dao

import (
	"entity"
	"util"
	"log"
)

type ListDao struct {
}

func (p *ListDao) Insert(user *entity.User) int64 {
	result,err := util.DB.Exec("INSERT INTO user(`username`,`password`,`role_id`,`status`) VALUE(?,?,?,?)",user.UserName,user.Password,user.RoleId,user.Status)
	if err != nil {
		log.Println(err)
		return 0
	}
	id,err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

