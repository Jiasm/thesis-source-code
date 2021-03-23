package dao

import (
	"entity"
	"log"
	"strconv"
	"strings"
	"util"
)

type ProjectDao struct {
}

func (p *ProjectDao) Insert(project *entity.Project) int64 {
	result, err := util.DB.Exec("INSERT INTO `project` (`creator`, `group_id`, `name`, `status`) VALUES (?, ?, ?, ?);", project.Creator, project.GroupId, project.Name, project.Status)
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

func (p *ProjectDao) FindOne(projectName string) *entity.Project {
	rows, err := util.DB.Query("SELECT id, creator, group_id, name, status FROM project WHERE name = ? LIMIT 1", projectName)

	if err != nil {
		log.Println(err)
		return nil
	}

	var project entity.Project

	rows.Next()

	rows.Scan(&project.ID, &project.Creator, &project.GroupId, &project.Name, &project.Status)

	rows.Close()

	return &project
}

func (p *ProjectDao) FindAll(projectIdList []uint) []entity.Project {
	var projectIdStrList []string

	for _, item := range projectIdList {
		projectIdStrList = append(projectIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query("SELECT id, creator, group_id, name, status FROM project WHERE id IN (?)", strings.Join(projectIdStrList, " , "))

	if err != nil {
		log.Println(err)
		return nil
	}

	var projectList []entity.Project

	for rows.Next() {
		var project entity.Project
		rows.Scan(&project.ID, &project.Creator, &project.GroupId, &project.Name, &project.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		projectList = append(projectList, project)
	}
	rows.Close()

	return projectList
}

func (p *ProjectDao) FindByCreator(creator uint) []entity.Project {
	rows, err := util.DB.Query("SELECT id, creator, group_id, name, status FROM project WHERE creator = ?", creator)

	if err != nil {
		log.Println(err)
		return nil
	}

	var projectList []entity.Project

	for rows.Next() {
		var project entity.Project
		rows.Scan(&project.ID, &project.Creator, &project.GroupId, &project.Name, &project.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		projectList = append(projectList, project)
	}
	rows.Close()

	return projectList
}