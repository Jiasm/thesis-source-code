package service

import (
	"dao"
	"entity"
	"time"
)

type ProjectService struct {
}

type ProjectData struct {
	ID 			uint 	`json:"id"`
	GroupName 	string 	`json:"group_name"`
	GroupId 	uint 	`json:"group_id"`
	Name 		string 	`json:"name"`
	Status 		uint 	`json:"status"`
}

var projectDao = new(dao.ProjectDao)

func (p *ProjectService) FindParticipantProjectList(uid uint) []ProjectData {
	participantTaskIdList := taskParticipantDao.FindTaskIdByParticipant(uid)
	projectIdList := taskDao.FindProjectId(participantTaskIdList)
	projectList := projectDao.FindAll(projectIdList)

	return BuildProjectData(projectList)
}

func (p *ProjectService) FindCreatedProjectList(uid uint) []ProjectData {
	projectList := projectDao.FindByCreator(uid)

	return BuildProjectData(projectList)
}

func (p *ProjectService) Create(request dao.NewProject) uint {
	request.CreatedDate = uint(time.Now().Unix())
	projectId := projectDao.Insert(request)

	return projectId
}

func (p *ProjectService) FindAll(projectIdList []uint) []entity.Project {
	projectList := projectDao.FindAll(projectIdList)

	return projectList
}

func (p *ProjectService) FindGroupList(projectId uint) []entity.TaskGroup {
	groupList := taskGroupDao.FindByProjectId(projectId)

	return groupList
}

func BuildProjectData(projectList []entity.Project) []ProjectData {
	// get group id list
	var groupIdList []uint

	for _, item := range projectList {
		groupIdList = append(groupIdList, item.GroupId)
	}

	groupList := groupDao.FindAll(groupIdList)

	var dataList []ProjectData

	for _, item := range projectList {
		for _, groupItem := range groupList {
			if item.GroupId == groupItem.ID {
				dataList = append(dataList, ProjectData{ item.ID, groupItem.Name, groupItem.ID, item.Name, item.Status })
			}
		}
	}

	return dataList
}
