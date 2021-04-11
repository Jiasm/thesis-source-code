package controller

import (
	"entity"
	"net/http"
	"service"
	"strconv"
	"strings"
	"util"
)


type TaskTagController struct {
}

type TaskTagListResponse struct {
	List []entity.TaskTag `json:"list"`
}

var taskTagService = new(service.TaskTagService)

func (p *TaskTagController) Router(router *util.RouterHandler) {
	router.Router("/task/tag/list", p.FindTaskTagList)
}

func (p *TaskTagController) FindTaskTagList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	TaskIdStr := strings.Split(vars.Get("task_ids"), `,`)

	var taskTagIdList []uint

	for _, str := range TaskIdStr {
		num, _ := strconv.Atoi(str)
		taskTagIdList = append(taskTagIdList, uint(num))
	}

	taskTagList := taskTagService.FindAll(taskTagIdList)

	util.ResultJsonOk(w, TaskTagListResponse{ taskTagList })
}