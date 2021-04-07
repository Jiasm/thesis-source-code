package dao

import (
	"log"
	"util"
)

type TaskTagDao struct {
}

func (p *TaskTagDao) Insert(taskId, tagId uint) uint {
	result, err := util.DB.Exec("INSERT INTO `task_tag` (`task_id`, `tag_id`) VALUES (?, ?)", taskId, tagId)
	if err != nil {
		log.Println(err)
		return 0
	}
	changedCount, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(changedCount)
}
