package dao

import (
	"entity"
	"util"
	"log"
)

type ListDao struct {
}

func (p *ListDao) Insert(list *entity.List) int64 {
	result,err := util.DB.Exec("INSERT INTO list(`user_id`,`title`,`content`,`create_time`) VALUE(?,?,?,?)",list.UserID,list.Title,list.Content,list.CreateTime)
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

