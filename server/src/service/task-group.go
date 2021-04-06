package service

import (
	"dao"
)

type TaskGroupService struct {
}

var taskGroupDao = new(dao.TaskGroupDao)

func (p *TaskGroupService) Create(request dao.NewTaskGroup) uint {
	taskGroupId := taskGroupDao.Insert(request)

	return taskGroupId
}