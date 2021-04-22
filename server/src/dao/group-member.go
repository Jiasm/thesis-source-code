package dao

import (
	"entity"
	"log"
	"util"
)

type GroupMemberDao struct {
}

func (p *GroupMemberDao) GetGroupIdByMember(uid uint) []uint {
	rows, err := util.DB.Query("SELECT group_id FROM `group_member` WHERE uid = ?", uid)

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupIdList []uint

	for rows.Next() {
		var groupMember entity.GroupMember
		rows.Scan(&groupMember.GroupId)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdList = append(groupIdList, groupMember.GroupId)
	}
	rows.Close()

	return groupIdList
}

func (p *GroupMemberDao) AddMember(groupId, uid, roleId, status, createdDate uint) uint {
	result, err := util.DB.Exec("INSERT INTO `group_member` (`group_id`, `uid`, `role_id`, `status`, `created_date`) VALUES (?, ?, ?, ?, ?)", groupId, uid, roleId, status, createdDate)
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

func (p *GroupMemberDao) ChangeRole(groupId, uid, roleId uint) uint {
	result, err := util.DB.Exec("UPDATE `group_member` SET `role_id` = ? WHERE `group_id` = ? AND `uid` = ?", roleId, groupId, uid)
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

func (p *GroupMemberDao) RemoveMember(groupId, uid uint) uint {
	result, err := util.DB.Exec("UPDATE `group_member` SET `status` = 0 WHERE `group_id` = ? AND `uid` = ?", groupId, uid)
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

func (p *GroupMemberDao) ActiveMember(groupId, uid uint) uint {
	result, err := util.DB.Exec("UPDATE `group_member` SET `status` = 1 WHERE `group_id` = ? AND `uid` = ?", groupId, uid)

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

func (p *GroupMemberDao) GetMemberByGroupId(groupId uint) []entity.GroupMember {
	rows, err := util.DB.Query("SELECT `group_id`, `uid`, `role_id`, `status`, `created_date` FROM `group_member` WHERE `group_id` = ?", groupId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupIdList []entity.GroupMember

	for rows.Next() {
		var groupMember entity.GroupMember
		rows.Scan(&groupMember.GroupId, &groupMember.Uid, &groupMember.RoleId, &groupMember.Status, &groupMember.CreatedDate)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdList = append(groupIdList, groupMember)
	}
	rows.Close()

	return groupIdList
}