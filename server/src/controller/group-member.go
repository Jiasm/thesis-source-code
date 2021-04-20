package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"strconv"
	"util"
)

type GroupMemberController struct {
	ApiController
}

type AddGroupMemberRequest struct {
	GroupId uint   `json:"group_id"`
	Uid  	uint   `json:"uid"`
	RoleId  uint   `json:"role_id"`
	Status  uint   `json:"status"`
}

type ChangeGroupMemberRoleRequest struct {
	GroupId uint   `json:"group_id"`
	Uid  	uint   `json:"uid"`
	RoleId  uint   `json:"role_id"`
}

type RemoveGroupMemberRoleRequest struct {
	GroupId uint   `json:"group_id"`
	Uid  	uint   `json:"uid"`
}

var groupMemberService = new(service.GroupMemberService)

func (p *GroupMemberController) Router(router *util.RouterHandler) {
	router.Router("/group-member/list", p.FindAll)
	router.Router("/group-member/add", p.AddMember)
	router.Router("/group-member/remove", p.RemoveMember)
	router.Router("/group-member/change-role", p.ChangeRole)
}

func (p *GroupMemberController) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	groupId, _ := strconv.Atoi(vars.Get("group_id"))

	groupMemberList := groupMemberService.FindByGroupId(uint(groupId))

	util.ResultJsonOk(w, groupMemberList)
}

func (p *GroupMemberController) AddMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &AddGroupMemberRequest{}

	decoder.Decode(&request)

	id := groupMemberService.AddMember(request.GroupId, request.Uid, request.RoleId, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *GroupMemberController) ChangeRole(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &ChangeGroupMemberRoleRequest{}

	decoder.Decode(&request)

	id := groupMemberService.ChangeRole(request.GroupId, request.Uid, request.RoleId)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *GroupMemberController) RemoveMember(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &RemoveGroupMemberRoleRequest{}

	decoder.Decode(&request)

	id := groupMemberService.RemoveMember(request.GroupId, request.Uid)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}