package service

import (
	"dao"
	"entity"
	"time"
)

type TaskService struct {
}

var taskDao = new(dao.TaskDao)

func (p *TaskService) FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, taskGroupId, parentTaskId, taskType, priority, page, size uint) []entity.Task {
	taskList := taskDao.FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, taskGroupId, parentTaskId, taskType, priority, page, size)

	return taskList
}

func (p *TaskService) Create(request dao.NewTask) uint {
	request.CreatedDate = uint(time.Now().Unix())
	//executorInfo := userDao.FindByName(request.Executor)
	//request.ExecutorId = executorInfo.ID
	taskId := taskDao.Insert(request)

	return taskId
}

func (p *TaskService) Change(taskId, title, desc, executor, status, expireDate, taskGroupId, taskType, priority string) uint {
	//var executorId string
	//if executor != "" {
	//	executorInfo := userDao.FindByName(executor)
	//	executorId = strconv.Itoa(int(executorInfo.ID))
	//} else {
	//	executorId = executor
	//}
	changedRows := taskDao.ChangeFields(taskId, title, desc, executor, status, expireDate, taskGroupId, taskType, priority)

	return changedRows
}

func (p *TaskService) CreateChildTask(request dao.NewTask) uint {
	request.CreatedDate = uint(time.Now().Unix())
	//executorInfo := userDao.FindByName(request.Executor)
	//request.ExecutorId = executorInfo.ID
	taskId := taskDao.Insert(request)

	return taskId
}

func (p *TaskService) ChangeStatus(taskId, status uint) uint {
	changedCount := taskDao.ChangeStatus(taskId, status)

	return changedCount
}

func (p *TaskService) ChangeExecutor(taskId, executor uint) uint {
	changedCount := taskDao.ChangeExecutor(taskId, executor)

	return changedCount
}

func (p *TaskService) AppendMember(taskId, executor uint) uint {
	changedCount := taskParticipantDao.Insert(taskId, executor, uint(time.Now().Unix()))

	return changedCount
}

func (p *TaskService) AddComment(taskId uint, text string, uid uint) uint {
	commentId := taskCommentDao.Insert(taskId, uid, 1, uint(time.Now().Unix()), text)

	return commentId
}

func (p *TaskService) AddTag(taskId uint, text string) uint {
	var tagId uint
	tag := tagDao.FindOne(text)

	if tag.Id <= 0 {
		tagId = tagDao.Insert(text, uint(time.Now().Unix()))
	} else {
		tagId = tag.Id
	}

	changedCount := taskTagDao.Insert(taskId, tagId)

	return changedCount
}

func (p *TaskService) RemoveTag(taskId uint, tagId uint) uint {
	changedCount := taskTagDao.Remove(taskId, tagId)

	return changedCount
}

func (p *TaskService) FindByTaskId(taskId uint) entity.Task {
	task := taskDao.FindOne(taskId)

	return task
}

func (p *TaskService) FindByParentTaskId(taskId uint) []entity.Task {
	taskList := taskDao.FindByParentId(taskId)

	return taskList
}