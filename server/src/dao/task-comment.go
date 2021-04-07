package dao

import (
	"log"
	"util"
)

type TaskCommentDao struct {
}

func (p *TaskCommentDao) Insert(taskId, creator, status, createdDate uint, text string) uint {
	result, err := util.DB.Exec("INSERT INTO `task_comment` (`task_id`, `creator`, `text`, `status`, `created_date`) VALUES (?, ?, ?, ?, ?)", taskId, creator, text, status, createdDate)
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
