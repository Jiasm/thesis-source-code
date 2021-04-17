package controller

import (
	"dao"
	"encoding/json"
	"entity"
	"fmt"
	"net/http"
	"service"
	"strconv"
	"util"
)

type TaskController struct {
	ApiController
}

type ListResponse struct {
	List interface{} `json:"list"`
}

type CreateTaskResponse struct {
	TaskId uint `json:"task_id"`
}

type ChangeStatusRequest struct {
	TaskId uint `json:"task_id"`
	Status uint `json:"status"`
}

type SetExecutorRequest struct {
	TaskId 		uint `json:"task_id"`
	Executor 	uint `json:"executor"`
}

type AppendMemberRequest struct {
	TaskId 	uint `json:"task_id"`
	Uid 	uint `json:"uid"`
}

type AddCommentRequest struct {
	TaskId 	uint 	`json:"task_id"`
	Text 	string 	`json:"text"`
}

type AddTagRequest struct {
	TaskId 	uint 	`json:"task_id"`
	Tag 	string 	`json:"tag"`
}

type RemoveTagRequest struct {
	TaskId 	uint 	`json:"task_id"`
	TagId 	uint 	`json:"tag_id"`
}

type ChangeStatusResponse struct {
	ChangedCount uint `json:"changed_count"`
}

type TaskDetailResponse struct {
	entity.Task
	ChildTask []entity.Task `json:"child_task"`
}

type AddCommentResponse struct {
	CommentId uint `json:"comment_id"`
}

type ChangeFieldsResponse struct {
	TaskId string `json:"task_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Executor string `json:"executor"`
	Status string `json:"status"`
	ExpireDate string `json:"expire_date"`
	TaskGroupId string `json:"task_group_id"`
	TaskType string `json:"task_type"`
	Priority string `json:"priority"`
}

var taskService = new(service.TaskService)

func (p *TaskController) Router(router *util.RouterHandler) {
	router.Router("/task/list", p.findByFilter)
	router.Router("/task", p.createTask)
	router.Router("/child-task", p.createChildTask)
	router.Router("/task/update-status", p.changeStatus)
	router.Router("/task/set-executor", p.setExecutor)
	router.Router("/task/append-member", p.appendMember)
	router.Router("/task/add-comment", p.addComment)
	router.Router("/task/add-tag", p.addTag)
	router.Router("/task/remove-tag", p.RemoveTag)
	router.Router("/task/detail", p.FindByTaskId)
	router.Router("/task/update", p.ChangeTask)
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
	taskGroupId, _ := strconv.Atoi(vars.Get("task_group_id"))
	parentTaskId, _ := strconv.Atoi(vars.Get("parent_task_id"))
	taskType, _ := strconv.Atoi(vars.Get("type"))
	priority, _ := strconv.Atoi(vars.Get("priority"))
	page, _ := strconv.Atoi(vars.Get("page"))
	size, _ := strconv.Atoi(vars.Get("size"))

	taskList := taskService.FindByFilter(uint(creator), uint(executor), uint(status), uint(maxCreatedDate), uint(minCreatedDate), uint(maxExpireDate), uint(minExpireDate), uint(taskProjectId), uint(taskGroupId), uint(parentTaskId), uint(taskType), uint(priority), uint(page), uint(size))

	if taskList == nil {
		util.ResultJsonOk(w, ListResponse{nil})
		return
	}

	util.ResultJsonOk(w, ListResponse{ taskList })
}

func (p *TaskController) createTask(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)
	decoder := json.NewDecoder(r.Body)

	request := &dao.NewTask{}

	decoder.Decode(&request)

	if uid <= 0 {
		uid = 3
	}

	request.Creator = uid

	taskId := taskService.Create(*request)

	util.ResultJsonOk(w, CreateTaskResponse{ taskId })
}

func (p *TaskController) createChildTask(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)
	decoder := json.NewDecoder(r.Body)

	request := &dao.NewTask{}

	decoder.Decode(&request)

	if uid <= 0 {
		uid = 3
	}

	request.Creator = uid

	taskId := taskService.CreateChildTask(*request)

	util.ResultJsonOk(w, CreateTaskResponse{ taskId })
}

func (p *TaskController) changeStatus(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &ChangeStatusRequest{}

	decoder.Decode(&request)

	changeCount := taskService.ChangeStatus(request.TaskId, request.Status)

	util.ResultJsonOk(w, ChangeStatusResponse{ changeCount })
}

func (p *TaskController) setExecutor(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &SetExecutorRequest{}

	decoder.Decode(&request)

	changeCount := taskService.ChangeExecutor(request.TaskId, request.Executor)

	util.ResultJsonOk(w, ChangeStatusResponse{ changeCount })
}

func (p *TaskController) appendMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &AppendMemberRequest{}

	decoder.Decode(&request)

	changeCount := taskService.AppendMember(request.TaskId, request.Uid)

	util.ResultJsonOk(w, ChangeStatusResponse{ changeCount })
}

func (p *TaskController) addComment(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)

	if uid <= 0 {
		uid = 3
	}
	decoder := json.NewDecoder(r.Body)

	request := &AddCommentRequest{}

	decoder.Decode(&request)

	commentId := taskService.AddComment(request.TaskId, request.Text, uid)

	util.ResultJsonOk(w, AddCommentResponse{ commentId })
}

func (p *TaskController) addTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &AddTagRequest{}

	decoder.Decode(&request)

	commentId := taskService.AddTag(request.TaskId, request.Tag)

	util.ResultJsonOk(w, ChangeStatusResponse{ commentId })
}

func (p *TaskController) RemoveTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &RemoveTagRequest{}

	decoder.Decode(&request)

	commentId := taskService.RemoveTag(request.TaskId, request.TagId)

	util.ResultJsonOk(w, ChangeStatusResponse{ commentId })
}

func (p *TaskController) FindByTaskId(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	taskId, _ := strconv.Atoi(vars.Get("task_id"))

	task := taskService.FindByTaskId(uint(taskId))

	childTaskList := taskService.FindByParentTaskId(uint(taskId))

	fmt.Println(task)

	util.ResultJsonOk(w, TaskDetailResponse{task, childTaskList})
}

func (p *TaskController) ChangeTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &ChangeFieldsResponse{}

	decoder.Decode(&request)

	changedRows := taskService.Change(request.TaskId, request.Title, request.Desc, request.Executor, request.Status, request.ExpireDate, request.TaskGroupId, request.TaskType, request.Priority)

	util.ResultJsonOk(w, ChangeStatusResponse{ changedRows })
}