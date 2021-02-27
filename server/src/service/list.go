package service

import (
	"dao"
	"entity"
	"time"
)

type ListService struct {
}

var listDao = new(dao.ListDao)

func (p *ListService) Insert(userId uint,title string,content string) int64 {
	id := feedbackDao.Insert(&entity.List{UserID:userId,Title:title,Content:content,CreateTime:time.Now()})
	if id <= 0 {
		return 0
	}
	return id
}
