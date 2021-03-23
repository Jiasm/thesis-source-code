package dao

import (
	"entity"
	"log"
	"strconv"
	"strings"
	"util"
)

type GroupDao struct {
}

func (p *GroupDao) Insert(group *entity.Group) int64 {
	result, err := util.DB.Exec("INSERT INTO `group` (`name`, `status`, `creator`) VALUES (?, ?, ?)", group.Name, group.Status, group.Creator)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *GroupDao) FindOne(groupName string) *entity.Group {
	rows, err := util.DB.Query("SELECT id, name, status, creator FROM `group` WHERE name = ? LIMIT 1", groupName)

	if err != nil {
		log.Println(err)
		return nil
	}

	var group entity.Group

	rows.Next()

	rows.Scan(&group.ID, &group.Name, &group.Status, &group.Creator)

	rows.Close()

	return &group
}

func (p *GroupDao) FindAll(groupIdList []uint) []entity.Group {
	var groupIdStrList []string

	for _, item := range groupIdList {
		groupIdStrList = append(groupIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query("SELECT id, name, status, creator FROM `group` WHERE id IN (?)", strings.Join(groupIdStrList, " , "))

	if err != nil {
		log.Println(err)
		return nil
	}

	var projectList []entity.Group

	for rows.Next() {
		var group entity.Group
		rows.Scan(&group.ID, &group.Name, &group.Status, &group.Creator)
		if err != nil {
			log.Println(err)
			continue
		}
		projectList = append(projectList, group)
	}
	rows.Close()

	return projectList
}
