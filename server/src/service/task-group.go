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

func (p *TaskGroupService) FindByGroupId(taskGroupIdList []uint) []entity.TaskGroup {
	taskGroupList := taskGroupDao.FindByGroupId(taskGroupIdList)

	return taskGroupList
}

func (p *TaskGroupService) FindAll() []entity.TaskGroup {
	taskGroupList := taskGroupDao.FindAll()

	return taskGroupList
}

func (p *TaskGroupService) FindByProjectId(projectId uint) []entity.TaskGroup {
	taskGroupList := taskGroupDao.FindByProjectId(projectId)

	return taskGroupList
}
