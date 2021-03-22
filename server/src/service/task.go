package service

import (
	"dao"
	"entity"
)

type TaskService struct {
}

var taskDao = new(dao.TaskDao)

func (p *TaskService) FindByFilter(creator, status, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority uint) []entity.Task {
	taskList := taskDao.FindByFilter(creator, status, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority)

	return taskList
}
