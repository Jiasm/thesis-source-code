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
