package dao

import (
	"log"
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
