package controller

import (
	"dao"
	"encoding/json"
	"net/http"
	"service"
	"util"
)

type ProjectController struct {
	ApiController
}

type ProjectResponse struct {
	Created interface{} `json:"created"`
	Participant interface{} `json:"participant"`
}

type CreateProjectResponse struct {
	ProjectId uint `json:"project_id"`
}

var projectService = new(service.ProjectService)

func (p *ProjectController) Router(router *util.RouterHandler) {
	router.Router("/project/list", p.findAll)
	router.Router("/project", p.createProject)
}

func (p *ProjectController) findAll(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)

	if uid <= 0 {
		uid = 3
	}

	createdProjectList := projectService.FindCreatedProjectList(uid)
	participantProjectList := projectService.FindParticipantProjectList(uid)

	util.ResultJsonOk(w, ProjectResponse{ createdProjectList, participantProjectList })
}

func (p *ProjectController) createProject(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)
	decoder := json.NewDecoder(r.Body)

	request := &dao.NewProject{}

	decoder.Decode(&request)

	if uid <= 0 {
		uid = 3
	}

	request.Creator = uid

	taskId := projectService.Create(*request)

	util.ResultJsonOk(w, CreateProjectResponse{ taskId })
}
