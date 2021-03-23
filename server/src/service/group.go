package service

import (
	"dao"
	"entity"
	"time"
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

func (p *GroupService) CreateGroup(uid uint, name string, status uint) uint {
	id := groupDao.Insert(&(dao.NewGroup{Name: name, Status: status, Creator: uid, CreatedDate: uint(time.Now().Unix()) }))

	groupMemberDao.AddMember(uid, id, 1, 1, uint(time.Now().Unix()))

	return id
}

func BuildGroupData(groupList []entity.Group) []GroupData {
	var dataList []GroupData

	for _, item := range groupList {
		dataList = append(dataList, GroupData{ item.ID, item.Name, item.Status })
	}

	return dataList
}