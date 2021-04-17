package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
	result, err := util.DB.Exec("INSERT INTO user(`username`,`password`,`role_id`,`status`) VALUE(?,?,?,?)", user.UserName, user.Password, user.RoleId, user.Status)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *UserDao) FindOne(username, password string) *entity.User {
	rows, err := util.DB.Query("SELECT id, username, password, status, role_id FROM user WHERE username = ? AND password = ? LIMIT 1", username, password)

	if err != nil {
		log.Println(err)
		return nil
	}

	var user entity.User

	rows.Next()

	rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Status, &user.RoleId)

	rows.Close()

	return &user
}

func (p *UserDao) FindAll(userIdList []uint) []entity.User {
	var userIdListStr []string

	for _, item := range userIdList {
		userIdListStr = append(userIdListStr, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query(fmt.Sprintf("SELECT id, username, password, status, role_id FROM user WHERE id IN(%s)", strings.Join(userIdListStr, " , ")))

	if err != nil {
		log.Println(err)
		return nil
	}

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Status, &user.RoleId)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()

	return users
}

func (p *UserDao) FindAllList() []entity.User {
	rows, err := util.DB.Query("SELECT id, username, password, status, role_id FROM user")

	if err != nil {
		log.Println(err)
		return nil
	}

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Status, &user.RoleId)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()

	return users
}

func (p *UserDao) FindByName(username string) *entity.User {
	rows, err := util.DB.Query("SELECT id, username, password, status, role_id FROM `user` WHERE username = ? LIMIT 1", username)

	if err != nil {
		log.Println(err)
		return nil
	}

	var user entity.User

	rows.Next()

	rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Status, &user.RoleId)

	rows.Close()

	return &user
}