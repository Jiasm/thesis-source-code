package service

import (
	"dao"
	"entity"
	"time"
)

type ProjectMemberService struct {
}

var projectMemberDao = new(dao.ProjectMemberDao)

func (p *ProjectMemberService) FindByGroupId(projectId uint) []entity.ProjectMember {
	projectMemberList := projectMemberDao.GetMemberByProjectId(projectId)

	return projectMemberList
}

func (p *ProjectMemberService) AddMember(projectId, uid, roleId, status uint) uint {
	id := projectMemberDao.AddMember(projectId, uid, roleId, status, uint(time.Now().Unix()))

	return id
}

func (p *ProjectMemberService) ChangeRole(projectId, uid, roleId uint) uint {
	id := projectMemberDao.ChangeRole(projectId, uid, roleId)

	return id
}

func (p *ProjectMemberService) RemoveMember(projectId, uid uint) uint {
	id := projectMemberDao.RemoveMember(projectId, uid)

	return id
}