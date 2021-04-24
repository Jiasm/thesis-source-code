package service

import (
	"dao"
	"entity"
)

type InfoService struct {
}

var infoDao = new(dao.InfoDao)

func (p *InfoService) FindNewCountList(projectId uint) []entity.NewTask {
	list := infoDao.FindNewCountList(projectId)

	return list
}

func (p *InfoService) FindTypedCountList(projectId uint) []entity.TypedTask {
	list := infoDao.FindTypedCountList(projectId)

	return list
}

func (p *InfoService) FindTodoCount(projectId uint) []entity.TodoCount {
	list := infoDao.FindTodoCount(projectId)

	return list
}

func (p *InfoService) FindDoneCount(projectId uint) []entity.DoneCount {
	list := infoDao.FindDoneCount(projectId)

	return list
}
