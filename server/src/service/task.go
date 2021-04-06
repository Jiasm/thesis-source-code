package service

import (
	"dao"
	"entity"
	"time"
)

type TaskService struct {
}

var taskDao = new(dao.TaskDao)

func (p *TaskService) FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, taskGroupId, parentTaskId, taskType, priority, page, size uint) []entity.Task {
	taskList := taskDao.FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, taskGroupId, parentTaskId, taskType, priority, page, size)

	return taskList
}

func (p *TaskService) Create(request dao.NewTask) uint {
	request.CreatedDate = uint(time.Now().Unix())
	taskId := taskDao.Insert(request)

	return taskId
}