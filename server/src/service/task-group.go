package service

import (
	"dao"
	"entity"
)

type TaskGroupService struct {
}

var taskGroupDao = new(dao.TaskGroupDao)

func (p *TaskGroupService) Create(request dao.NewTaskGroup) uint {
	taskGroupId := taskGroupDao.Insert(request)

	return taskGroupId
}

func (p *TaskGroupService) FindAll(taskGroupIdList []uint) []entity.TaskGroup {
	taskGroupList := taskGroupDao.FindAll(taskGroupIdList)

	return taskGroupList
}

func (p *TaskGroupService) FindByProjectId(projectId uint) []entity.TaskGroup {
	taskGroupList := taskGroupDao.FindByProjectId(projectId)

	return taskGroupList
}
