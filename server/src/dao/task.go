package dao

import (
	"entity"
	"fmt"
	"log"
	"strings"
	"util"
)

type TaskDao struct {
}

func (p *TaskDao) Insert(task *entity.Task) int64 {
	result, err := util.DB.Exec("INSERT INTO `task` (`title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", task.Title, task.Desc, task.Creator, task.Status, task.ExpireDate, task.TaskProjectId, task.ParentTaskId, task.Type, task.Priority)
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

func (p *TaskDao) FindOne(taskId uint) *entity.Task {
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id = ? LIMIT 1", taskId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var task entity.Task

	rows.Next()

	rows.Scan(&task.ID, &task.Title, task.Desc, &task.Creator, &task.Status, &task.ExpireDate, &task.Type, &task.Priority)

	rows.Close()

	return &task
}

func (p *TaskDao) FindAll(taskIdList []int) []entity.Task {
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id IN (?)", taskIdList)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, task.Desc, &task.Creator, &task.Status, &task.ExpireDate, &task.TaskProjectId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}

func (p *TaskDao) FindByProjectId(projectId uint) []entity.Task {
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id IN (?)", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, task.Desc, &task.Creator, &task.Status, &task.ExpireDate, &task.TaskProjectId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}

func (p *TaskDao) FindByFilter(creator, status, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority uint) []entity.Task {
	var where []string

	if creator > 0 {
		where = append(where, fmt.Sprintf(`creator = %d`, creator))
	}

	if status > 0 {
		where = append(where, fmt.Sprintf(`status = %d`, status))
	}

	if maxExpireDate > 0 {
		where = append(where, fmt.Sprintf(`expire_date <= %d`, maxExpireDate))
	}

	if minExpireDate > 0 {
		where = append(where, fmt.Sprintf(`expire_date >= %d`, minExpireDate))
	}

	if taskProjectId > 0 {
		where = append(where, fmt.Sprintf(`task_project_id = %d`, taskProjectId))
	}

	if parentTaskId > 0 {
		where = append(where, fmt.Sprintf(`parent_task_id = %d`, parentTaskId))
	}

	if taskType > 0 {
		where = append(where, fmt.Sprintf(`type = %d`, taskType))
	}

	if priority > 0 {
		where = append(where, fmt.Sprintf(`priority = %d`, priority))
	}

	var sql string

	if len(where) > 0 {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE " + strings.Join(where, " AND ")
	} else {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `status`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task"
	}

	fmt.Println(sql)

	rows, err := util.DB.Query(sql)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Status, &task.ExpireDate, &task.TaskProjectId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}


