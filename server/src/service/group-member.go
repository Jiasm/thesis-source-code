package service

import (
	"dao"
	"entity"
	"time"
)

type GroupMemberService struct {
}

var groupMemberDao = new(dao.GroupMemberDao)

func (p *GroupMemberService) FindByGroupId(groupId uint) []entity.GroupMember {
	groupMemberList := groupMemberDao.GetMemberByGroupId(groupId)

	return groupMemberList
}

func (p *GroupMemberService) FindByProjectId(projectId uint) []entity.GroupMember {
	projectInfo := projectDao.FindAll([]uint{projectId})
	groupMemberList := groupMemberDao.GetMemberByGroupId(projectInfo[0].GroupId)

	return groupMemberList
}

func (p *GroupMemberService) AddMember(groupId, uid, roleId, status uint) uint {
	id := groupMemberDao.AddMember(groupId, uid, roleId, status, uint(time.Now().Unix()))

	return id
}

func (p *GroupMemberService) ChangeRole(groupId, uid, roleId uint) uint {
	id := groupMemberDao.ChangeRole(groupId, uid, roleId)

	return id
}

func (p *GroupMemberService) RemoveMember(groupId, uid uint) uint {
	id := groupMemberDao.RemoveMember(groupId, uid)

	return id
}

func (p *GroupMemberService) ActiveMember(groupId, uid uint) uint {
	id := groupMemberDao.ActiveMember(groupId, uid)

	return id
}