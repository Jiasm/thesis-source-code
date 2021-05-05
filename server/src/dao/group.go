package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type GroupDao struct {
}

type NewGroup struct {
	Name     	string	`json:"name"`
	Status     	uint   	`json:"status"`
	Creator    	uint	`json:"creator"`
	CreatedDate uint 	`json:"created_date""`
}

func (p *GroupDao) Insert(group *NewGroup) uint {
	result, err := util.DB.Exec("INSERT INTO `group` (`name`, `status`, `creator`, `created_date`) VALUES (?, ?, ?, ?)", group.Name, group.Status, group.Creator, group.CreatedDate)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(id)
}

func (p *GroupDao) FindOne(groupName string) *entity.Group {
	rows, err := util.DB.Query("SELECT id, name, status, creator, created_date FROM `group` WHERE name = ? LIMIT 1", groupName)

	if err != nil {
		log.Println(err)
		return nil
	}

	var group entity.Group

	rows.Next()

	rows.Scan(&group.ID, &group.Name, &group.Status, &group.Creator, &group.CreatedDate)

	rows.Close()

	return &group
}

func (p *GroupDao) FindAll(groupIdList []uint) []entity.Group {
	rows, err := util.DB.Query(BuildFindGroupListWithGroupIdList(groupIdList))

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupList []entity.Group

	for rows.Next() {
		var group entity.Group
		rows.Scan(&group.ID, &group.Name, &group.Status, &group.Creator, &group.CreatedDate)
		if err != nil {
			log.Println(err)
			continue
		}
		groupList = append(groupList, group)
	}
	rows.Close()

	return groupList
}

func BuildFindGroupListWithGroupIdList(groupIdList []uint) string {
	var groupIdStrList []string

	for _, item := range groupIdList {
		groupIdStrList = append(groupIdStrList, strconv.Itoa(int(item)))
	}

	return fmt.Sprintf("SELECT id, name, status, creator, created_date FROM `group` WHERE id IN (%s)", strings.Join(groupIdStrList, " , "))
}

func (p *GroupDao) FindGroupIdByCreator(creator uint) []uint {
	rows, err := util.DB.Query("SELECT id FROM `group` WHERE creator = ?", creator)

	if err != nil {
		log.Println(err)
		return nil
	}

	var groupIdList []uint

	for rows.Next() {
		var group entity.Group
		rows.Scan(&group.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		groupIdList = append(groupIdList, group.ID)
	}
	rows.Close()

	return groupIdList
}
