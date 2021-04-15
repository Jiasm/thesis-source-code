package dao

import (
	"entity"
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


func (p *TaskCommentDao) FindByTask(taskId uint) []entity.TaskComment {
	rows, err := util.DB.Query("SELECT id, creator, task_id, text, status, created_date FROM task_comment WHERE task_id = ? ORDER BY created_date ASC", taskId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskCommentList []entity.TaskComment

	for rows.Next() {
		var taskComment entity.TaskComment
		rows.Scan(&taskComment.ID, &taskComment.Creator, &taskComment.TaskId, &taskComment.Text, &taskComment.Status, &taskComment.CreatedDate)
		if err != nil {
			log.Println(err)
			continue
		}
		taskCommentList = append(taskCommentList, taskComment)
	}
	rows.Close()

	return taskCommentList
}
