package controller

import (
	"net/http"
	"service"
	"util"
)

type GroupController struct {
	ApiController
}

type GroupResponse struct {
	Created interface{} `json:"created"`
	Participant interface{} `json:"participant"`
}

var groupService = new(service.GroupService)

func (p *GroupController) Router(router *util.RouterHandler) {
	router.Router("/group/list", p.findAll)
}

func (p *GroupController) findAll(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)

	if uid <= 0 {
		uid = 3
	}

	createdGroupList := groupService.FindCreatedProjectList(uid)
	participantGroupList := groupService.FindParticipantProjectList(uid)

	util.ResultJsonOk(w, GroupResponse{ createdGroupList, participantGroupList })
}