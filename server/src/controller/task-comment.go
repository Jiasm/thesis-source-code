package controller

import (
	"net/http"
	"service"
	"strconv"
	"util"
)

type TaskCommentController struct {
	ApiController
}

var taskCommentService = new(service.TaskCommentService)

func (p *TaskCommentController) Router(router *util.RouterHandler) {
	router.Router("/task/comment/list", p.findCommentByTaskId)
}

func (p *TaskCommentController) findCommentByTaskId(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	taskIdU, _ := strconv.ParseUint(vars.Get("task_id"), 10, 32)

	taskId := uint(taskIdU)

	taskCommentList := taskCommentService.FindByTaskId(taskId)

	util.ResultJsonOk(w, taskCommentList)
}