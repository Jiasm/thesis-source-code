package dao

import (
	"entity"
	"fmt"
	"log"
	"util"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
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

func (p *UserDao) FindOne(username, password string) *entity.User {
	rows, err := util.DB.Query("SELECT * FROM user WHERE username = ? AND password = ? LIMIT 1", username, password)

	if err != nil {
		log.Println(err)
		return nil
	}

	var user entity.User

	rows.Next()

	rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Status, &user.RoleId)

	return &user
}
