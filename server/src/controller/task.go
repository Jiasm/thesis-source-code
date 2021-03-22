package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"strconv"
	"util"
)

type TaskController struct {
}

type FilterRequest struct {
	Creator uint `json:"creator"`
	Status uint `json:"status"`
	ExpireDate uint `json:"expire_date"`
	TaskProjectId uint `json:"task_project_id"`
	ParentTaskId uint `json:"parent_task_id"`
	TaskType uint `json:"type"`
	Priority uint `json:"creator"`
}

var taskService = new(service.TaskService)

func (p *TaskController) Router(router *util.RouterHandler) {
	router.Router("/task/list", p.findByFilter)
}

func (p *TaskController) findByFilter(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	decoder := json.NewDecoder(r.Body)

	request := &FilterRequest{}

	decoder.Decode(&request)

	creator, _ := strconv.Atoi(vars.Get("creator"))
	status, _ := strconv.Atoi(vars.Get("status"))
	maxExpireDate, _ := strconv.Atoi(vars.Get("max_expire_date"))
	minExpireDate, _ := strconv.Atoi(vars.Get("min_expire_date"))
	taskProjectId, _ := strconv.Atoi(vars.Get("task_project_id"))
	parentTaskId, _ := strconv.Atoi(vars.Get("parent_task_id"))
	taskType, _ := strconv.Atoi(vars.Get("type"))
	priority, _ := strconv.Atoi(vars.Get("priority"))

	taskList := taskService.FindByFilter(uint(creator), uint(status), uint(maxExpireDate), uint(minExpireDate), uint(taskProjectId), uint(parentTaskId), uint(taskType), uint(priority))

	if taskList == nil {
		util.ResultJsonOk(w, nil)
		return
	}

	util.ResultJsonOk(w, taskList)
}