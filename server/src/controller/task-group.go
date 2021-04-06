package controller

import (
	"dao"
	"encoding/json"
	"net/http"
	"service"
	"util"
)

type TaskGroupController struct {
	ApiController
}

type CreateTaskGroupResponse struct {
	TaskGroupId uint `json:"task_group_id"`
}

var taskGroupService = new(service.TaskGroupService)

func (p *TaskGroupController) Router(router *util.RouterHandler) {
	router.Router("/task-group", p.createTaskGroup)
}

func (p *TaskGroupController) createTaskGroup(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)
	decoder := json.NewDecoder(r.Body)

	request := &dao.NewTaskGroup{}

	decoder.Decode(&request)

	if uid <= 0 {
		uid = 3
	}

	request.Creator = uid

	taskId := taskGroupService.Create(*request)

	util.ResultJsonOk(w, CreateTaskGroupResponse{ taskId })
}
