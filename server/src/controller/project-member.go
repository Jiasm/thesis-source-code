package controller

import (
	"encoding/json"
	"entity"
	"fmt"
	"net/http"
	"service"
	"strconv"
	"util"
)

type ProjectMemberController struct {
	ApiController
}

type AddProjectMemberRequest struct {
	ProjectId 	uint   `json:"project_id"`
	Uid  		uint   `json:"uid"`
	RoleId  	uint   `json:"role_id"`
	Status  	uint   `json:"status"`
}

type ChangeProjectMemberRoleRequest struct {
	ProjectId 	uint   `json:"project_id"`
	Uid  		uint   `json:"uid"`
	RoleId  	uint   `json:"role_id"`
}

type RemoveProjectMemberRoleRequest struct {
	ProjectId 	uint   `json:"project_id"`
	Uid  		uint   `json:"uid"`
}

type ProjectMemberResponse struct {
	Project []entity.ProjectMember 	`json:"project"`
	Group 	[]entity.GroupMember 	`json:"group"`
}

var projectMemberService = new(service.ProjectMemberService)

func (p *ProjectMemberController) Router(router *util.RouterHandler) {
	router.Router("/project-member/list", p.FindAll)
	router.Router("/project-member/add", p.AddMember)
	router.Router("/project-member/remove", p.RemoveMember)
	router.Router("/project-member/active", p.ActiveMember)
	router.Router("/project-member/change-role", p.ChangeRole)
}

func (p *ProjectMemberController) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	projectId, _ := strconv.Atoi(vars.Get("project_id"))

	projectMemberList := projectMemberService.FindByGroupId(uint(projectId))
	groupMemberList := groupMemberService.FindByProjectId(uint(projectId))

	util.ResultJsonOk(w, ProjectMemberResponse{projectMemberList, groupMemberList })
}

func (p *ProjectMemberController) AddMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &AddProjectMemberRequest{}

	decoder.Decode(&request)

	id := projectMemberService.AddMember(request.ProjectId, request.Uid, request.RoleId, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *ProjectMemberController) ChangeRole(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &ChangeProjectMemberRoleRequest{}

	decoder.Decode(&request)

	fmt.Println(request.ProjectId, request.Uid, request.RoleId)

	id := projectMemberService.ChangeRole(request.ProjectId, request.Uid, request.RoleId)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *ProjectMemberController) RemoveMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &RemoveProjectMemberRoleRequest{}

	decoder.Decode(&request)

	id := projectMemberService.RemoveMember(request.ProjectId, request.Uid)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *ProjectMemberController) ActiveMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &RemoveProjectMemberRoleRequest{}

	decoder.Decode(&request)

	id := projectMemberService.ActiveMember(request.ProjectId, request.Uid)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}