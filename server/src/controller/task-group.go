package controller

import (
	"dao"
	"encoding/json"
	"entity"
	"net/http"
	"service"
	"strconv"
	"strings"
	"util"
)

type TaskGroupController struct {
	ApiController
}

type CreateTaskGroupResponse struct {
	TaskGroupId uint `json:"task_group_id"`
}

type TaskGroupListResponse struct {
	List []entity.TaskGroup `json:"list"`
}

var taskGroupService = new(service.TaskGroupService)

func (p *TaskGroupController) Router(router *util.RouterHandler) {
	router.Router("/task-group", p.createTaskGroup)
	router.Router("/task/group/list", p.findTaskGroup)
	router.Router("/task/group/list/all", p.GetAllTaskGroup)
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

func (p *TaskGroupController) findTaskGroup(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	taskGroupIdStr := strings.Split(vars.Get("group_ids"), `,`)

	var taskGroupIdList []uint

	for _, str := range taskGroupIdStr {
		num, _ := strconv.Atoi(str)
		taskGroupIdList = append(taskGroupIdList, uint(num))
	}

	taskGroupList := taskGroupService.FindByGroupId(taskGroupIdList)

	util.ResultJsonOk(w, TaskGroupListResponse{ taskGroupList })
}

func (p *TaskGroupController) GetAllTaskGroup(w http.ResponseWriter, r *http.Request) {
	taskGroupList := taskGroupService.FindAll()

	util.ResultJsonOk(w, TaskGroupListResponse{ taskGroupList })
}
