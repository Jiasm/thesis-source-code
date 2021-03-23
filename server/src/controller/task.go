package controller

import (
	"net/http"
	"service"
	"strconv"
	"util"
)

type TaskController struct {
}

type ListResponse struct {
	List interface{} `json:"list"`
}

var taskService = new(service.TaskService)

func (p *TaskController) Router(router *util.RouterHandler) {
	router.Router("/task/list", p.findByFilter)
}

func (p *TaskController) findByFilter(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	creator, _ := strconv.Atoi(vars.Get("creator"))
	executor, _ := strconv.Atoi(vars.Get("executor"))
	status, _ := strconv.Atoi(vars.Get("status"))
	maxCreatedDate, _ := strconv.Atoi(vars.Get("max_created_date"))
	minCreatedDate, _ := strconv.Atoi(vars.Get("min_created_date"))
	maxExpireDate, _ := strconv.Atoi(vars.Get("max_expire_date"))
	minExpireDate, _ := strconv.Atoi(vars.Get("min_expire_date"))
	taskProjectId, _ := strconv.Atoi(vars.Get("task_project_id"))
	parentTaskId, _ := strconv.Atoi(vars.Get("parent_task_id"))
	taskType, _ := strconv.Atoi(vars.Get("type"))
	priority, _ := strconv.Atoi(vars.Get("priority"))
	page, _ := strconv.Atoi(vars.Get("page"))
	size, _ := strconv.Atoi(vars.Get("size"))

	taskList := taskService.FindByFilter(uint(creator), uint(executor), uint(status), uint(maxCreatedDate), uint(minCreatedDate), uint(maxExpireDate), uint(minExpireDate), uint(taskProjectId), uint(parentTaskId), uint(taskType), uint(priority), uint(page), uint(size))

	if taskList == nil {
		util.ResultJsonOk(w, ListResponse{nil})
		return
	}

	util.ResultJsonOk(w, ListResponse{ taskList })
}