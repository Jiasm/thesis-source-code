package service

import (
	"dao"
	"entity"
)

type TaskTagService struct {
}

var taskTagDao = new(dao.TaskTagDao)

func (p *TaskTagService) FindAll(taskIdList []uint) []entity.TaskTag {
	taskTagList := taskTagDao.FindAll(taskIdList)

	return taskTagList
}
