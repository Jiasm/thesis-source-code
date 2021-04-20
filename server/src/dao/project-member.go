package dao

import (
	"entity"
	"log"
	"util"
)

type ProjectMemberDao struct {

}

func (p *ProjectMemberDao) GetGroupIdByMember(uid uint) []uint {
	rows, err := util.DB.Query("SELECT project_id FROM `project_member` WHERE uid = ?", uid)

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupIdList []uint

	for rows.Next() {
		var groupMember entity.ProjectMember
		rows.Scan(&groupMember.ProjectId)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdList = append(groupIdList, groupMember.ProjectId)
	}
	rows.Close()

	return groupIdList
}

func (p *ProjectMemberDao) AddMember(projectId, uid, roleId, status, createdDate uint) uint {
	result, err := util.DB.Exec("INSERT INTO `project_member` (`project_id`, `uid`, `role_id`, `status`, `created_date`) VALUES (?, ?, ?, ?, ?)", projectId, uid, roleId, status, createdDate)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(id)
}

func (p *ProjectMemberDao) ChangeRole(projectId, uid, roleId uint) uint {
	result, err := util.DB.Exec("UPDATE `project_member` SET `role_id` = ? WHERE `project_id` = ? AND `uid` = ?", roleId, projectId, uid)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(id)
}

func (p *ProjectMemberDao) RemoveMember(projectId, uid uint) uint {
	result, err := util.DB.Exec("DELETE FROM `project_member` WHERE `project_id` = ? AND `uid` = ?", projectId, uid)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(id)
}

func (p *ProjectMemberDao) GetMemberByProjectId(projectId uint) []entity.ProjectMember {
	rows, err := util.DB.Query("SELECT `project_id`, `uid`, `role_id`, `status`, `created_date` FROM `project_member` WHERE project_id = ?", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupIdList []entity.ProjectMember

	for rows.Next() {
		var groupMember entity.ProjectMember
		rows.Scan(&groupMember.ProjectId, &groupMember.Uid, &groupMember.RoleId, &groupMember.Status, &groupMember.CreatedDate)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdList = append(groupIdList, groupMember)
	}
	rows.Close()

	return groupIdList
}