package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type NewTask struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Creator uint `json:"creator"`
	Executor string `json:"executor"`
	ExecutorId uint `json:"executor_id"`
	Status uint `json:"status"`
	CreatedDate uint `json:"created_date"`
	ExpireDate uint `json:"expire_date"`
	TaskProjectId uint `json:"task_project_id"`
	TaskGroupId uint `json:"task_group_id"`
	ParentTaskId uint `json:"parent_task_id"`
	TaskType uint `json:"task_type"`
	Priority uint `json:"priority"`
}

type TaskDao struct {}

func (p *TaskDao) Insert(request NewTask) uint {
	result, err := util.DB.Exec("INSERT INTO `task` (`title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", request.Title, request.Desc, request.Creator, request.ExecutorId, request.Status, request.CreatedDate, request.ExpireDate, request.TaskProjectId, request.TaskGroupId, request.ParentTaskId, request.TaskType, request.Priority)
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

func (p *TaskDao) FindOne(taskId uint) entity.Task {
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id = ? LIMIT 1", taskId)

	fmt.Println("SELECT " +
		"`id`," +
		"`title`," +
		"`desc`," +
		"`creator`," +
		"`executor`," +
		"`status`," +
		"`created_date`," +
		"`expire_date`," +
		"`task_project_id`," +
		"`task_group_id`," +
		"`type`," +
		"`priority`," +
		"`parent_task_id`" +
		"FROM task WHERE id = ? LIMIT 1", taskId)

	var task entity.Task

	if err != nil {
		log.Println(err)
		return task
	}

	rows.Next()

	rows.Scan(
		&task.ID,
		&task.Title,
		&task.Desc,
		&task.Creator,
		&task.Executor,
		&task.Status,
		&task.CreatedDate,
		&task.ExpireDate,
		&task.TaskProjectId,
		&task.TaskGroupId,
		&task.Type,
		&task.Priority,
		&task.ParentTaskId,
	)

	rows.Close()

	fmt.Println(task.Priority, task.Type)

	return task
}

func (p *TaskDao) FindByParentId(taskId uint) []entity.Task {
	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task WHERE parent_task_id = ?", taskId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.TaskProjectId, &task.TaskGroupId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}

func (p *TaskDao) FindAll(taskIdList []int) []entity.Task {
	var taskIdStrList []string

	for _, item := range taskIdList {
		taskIdStrList = append(taskIdStrList, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query("SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task WHERE id IN (?)", strings.Join(taskIdStrList, " , "))

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskList []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.TaskProjectId, &task.TaskGroupId, &task.ParentTaskId, &task.Type, &task.Priority)
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

func (p *TaskDao) FindByFilter(creator, executor, status, maxCreatedDate, minCreatedDate, maxExpireDate, minExpireDate, taskProjectId, taskGroupId, parentTaskId, taskType, priority, page, size uint) []entity.Task {
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

	if taskGroupId > 0 {
		where = append(where, fmt.Sprintf(`task_group_id = %d`, taskGroupId))
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
	orderStr := " ORDER BY task_project_id DESC, task_group_id DESC, parent_task_id DESC, priority DESC, created_date DESC"

	if len(where) > 0 {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task WHERE " + strings.Join(where, " AND ") + orderStr + limitStr
	} else {
		sql = "SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task" + orderStr + limitStr
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
		rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Creator, &task.Executor, &task.Status, &task.CreatedDate, &task.ExpireDate, &task.TaskProjectId, &task.TaskGroupId, &task.ParentTaskId, &task.Type, &task.Priority)
		if err != nil {
			log.Println(err)
			continue
		}
		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList
}

func (p *TaskDao) ChangeStatus(taskId, status uint) uint {
	result, err := util.DB.Exec("UPDATE `task` SET status = ? WHERE id = ?", status, taskId)
	if err != nil {
		log.Println(err)
		return 0
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(count)
}


func (p *TaskDao) ChangeExecutor(taskId, executor uint) uint {
	result, err := util.DB.Exec("UPDATE `task` SET executor = ? WHERE id = ?", executor, taskId)
	if err != nil {
		log.Println(err)
		return 0
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(count)
}

func (p *TaskDao) ChangeFields(taskId, title, desc, executor, status, expireDate, taskGroupId, taskType, priority string) uint {
	var fields []string

	if title != "" {
		fields = append(fields, fmt.Sprintf(`title = '%s'`, title))
	}

	if desc != "" {
		fields = append(fields, fmt.Sprintf("`desc` = '%s'", desc))
	}

	if executor != "" {
		fields = append(fields, fmt.Sprintf(`executor = %s`, executor))
	}

	if status != "" {
		fields = append(fields, fmt.Sprintf(`status = %s`, status))
	}

	if expireDate != "" {
		fields = append(fields, fmt.Sprintf(`expire_date = %s`, expireDate))
	}

	if taskGroupId != "" {
		fields = append(fields, fmt.Sprintf(`task_group_id = %s`, taskGroupId))
	}

	if taskType != "" {
		fields = append(fields, fmt.Sprintf(`type = %s`, taskType))
	}

	if priority != "" {
		fields = append(fields, fmt.Sprintf(`priority = %s`, priority))
	}

	var sql string

	sql = "UPDATE `task` SET " + strings.Join(fields, " , ") + " WHERE id = ?"

	fmt.Println(sql)

	result, err := util.DB.Exec(sql, taskId)
	if err != nil {
		log.Println(err)
		return 0
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(count)
}
