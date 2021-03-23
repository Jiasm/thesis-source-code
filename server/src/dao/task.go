package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type TaskDao struct {
}

func (p *TaskDao) Insert(task *entity.Task) int64 {
	result, err := util.DB.Exec("INSERT INTO `task` (`title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", task.Title, task.Desc, task.Creator, task.Status, task.CreatedDate, task.ExpireDate, task.TaskProjectId, task.ParentTaskId, task.Type, task.Priority)
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
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id = ? LIMIT 1", taskId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var task entity.Task

	rows.Next()

	rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.Type, &task.Priority)

	rows.Close()

	return &task
}

func (p *TaskDao) FindAll(taskIdList []int) []entity.Task {
	var taskIdStrList []string

	for _, item := range taskIdList {
		taskIdStrList = append(taskIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id IN (?)", strings.Join(taskIdStrList, " , "))

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.TaskProjectId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}

func (p *TaskDao) FindProjectId(taskIdList []uint) []uint {
	var taskIdStrList []string

	for _, item := range taskIdList {
		taskIdStrList = append(taskIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query("SELECT`task_project_id` FROM task WHERE id IN (?) GROUP BY task_project_id", strings.Join(taskIdStrList, " , "))

	if err != nil {
		log.Println(err)
		return nil
	}

	var projectIdList []uint

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.TaskProjectId)
		if err != nil {
			log.Println(err)
			continue
		}
		projectIdList = append(projectIdList, task.TaskProjectId)
	}
	rows.Close()

	return projectIdList
}

func (p *TaskDao) FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, parentTaskId, taskType, priority, page, size uint) []entity.Task {
	var where []string

	if creator > 0 {
		where = append(where, fmt.Sprintf(`creator = %d`, creator))
	}

	if executor > 0 {
		where = append(where, fmt.Sprintf(`executor = %d`, executor))
	}

	if status > 0 {
		where = append(where, fmt.Sprintf(`status = %d`, status))
	}

	if maxCreatedDate > 0 {
		where = append(where, fmt.Sprintf(`created_date <= %d`, maxExpireDate))
	}

	if minCreatedDate > 0 {
		where = append(where, fmt.Sprintf(`created_date >= %d`, minExpireDate))
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

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 10
	}

	var sql string

	limitStr := " LIMIT " + strconv.Itoa(int(page - 1)) + " , " + strconv.Itoa(int(size * page))
	orderStr := " ORDER BY task_project_id DESC, parent_task_id DESC, created_date DESC"

	if len(where) > 0 {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task WHERE " + strings.Join(where, " AND ") + orderStr + limitStr
	} else {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `parent_task_id`, `type`, `priority` FROM task" + orderStr + limitStr
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
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.TaskProjectId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}


