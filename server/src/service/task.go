package service

import (
	"dao"
	"entity"
)

type TaskService struct {
}

var taskDao = new(dao.TaskDao)

func (p *TaskService) FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority, page, size uint) []entity.Task {
	taskList := taskDao.FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority, page, size)

	return taskList
}
