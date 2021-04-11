package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type TaskGroupDao struct {
}

type NewTaskGroup struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Creator uint   `json:"creator"`
	Status  uint   `json:"status"`
}

func (p *TaskGroupDao) Insert(task NewTaskGroup) uint {
	result, err := util.DB.Exec("INSERT INTO `task_group` (`title`, `desc`, `creator`, `status`) VALUES (?, ?, ?, ?)", task.Title, task.Desc, task.Creator, task.Status)
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

func (p *TaskGroupDao) FindAll(taskGroupIdList []uint) []entity.TaskGroup {
	var taskGroupIdStrList []string

	for _, item := range taskGroupIdList {
		taskGroupIdStrList = append(taskGroupIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query(fmt.Sprintf("SELECT id, title, `desc`, creator, status FROM task_group WHERE id IN (%s)", strings.Join(taskGroupIdStrList, " , ")))

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskGroupList []entity.TaskGroup

	for rows.Next() {
		fmt.Println("join")
		var project entity.TaskGroup
		rows.Scan(&project.ID, &project.Title, &project.Desc, &project.Creator, &project.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		taskGroupList = append(taskGroupList, project)
	}
	rows.Close()

	return taskGroupList
}
