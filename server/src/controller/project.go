package controller

import (
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

var projectService = new(service.ProjectService)

func (p *ProjectController) Router(router *util.RouterHandler) {
	router.Router("/project/list", p.findAll)
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