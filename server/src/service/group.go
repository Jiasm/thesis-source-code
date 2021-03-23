package service

import (
	"dao"
	"entity"
)

type GroupService struct {
}

type GroupData struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Status uint `json:"status"`
}

var groupDao = new(dao.GroupDao)

func (p *GroupService) FindParticipantProjectList(uid uint) []GroupData {
	groupIdList := groupMemberDao.GetGroupIdByMember(uid)
	groupList := groupDao.FindAll(groupIdList)

	return BuildGroupData(groupList)
}

func (p *GroupService) FindCreatedProjectList(uid uint) []GroupData {
	groupIdList := groupDao.FindGroupIdByCreator(uid)
	groupList := groupDao.FindAll(groupIdList)

	return BuildGroupData(groupList)
}

func BuildGroupData(groupList []entity.Group) []GroupData {
	var dataList []GroupData

	for _, item := range groupList {
		dataList = append(dataList, GroupData{ item.ID, item.Name, item.Status })
	}

	return dataList
}