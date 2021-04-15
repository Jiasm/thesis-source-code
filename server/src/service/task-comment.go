package service

import (
	"dao"
	"entity"
)

type TaskCommentService struct {
}

var taskCommentDao = new(dao.TaskCommentDao)

func (p *TaskCommentService) FindByTaskId(taskId uint) []entity.TaskComment {
	taskCommentList := taskCommentDao.FindByTask(taskId)

	return taskCommentList
}
