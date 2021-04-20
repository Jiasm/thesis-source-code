package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"strconv"
	"strings"
	"util"
)

type GroupController struct {
	ApiController
}

type GroupResponse struct {
	Created interface{} `json:"created"`
	Participant interface{} `json:"participant"`
}

type GroupRequest struct {
	Name 	string `json:"name"`
	Status  uint   `json:"status"`
}

type CreateGroupResponse struct {
	GroupId uint `json:"group_id"`
}

var groupService = new(service.GroupService)

func (p *GroupController) Router(router *util.RouterHandler) {
	router.Router("/group/list/all", p.findAll)
	router.Router("/group/list", p.FindByGroupId)
	router.Router("/group", p.create)
}

func (p *GroupController) findAll(w http.ResponseWriter, r *http.Request) {
	uid := p.GetUserId(w, r)

	if uid <= 0 {
		uid = 3
	}

	createdGroupList := groupService.FindCreatedGroupList(uid)
	participantGroupList := groupService.FindParticipantGroupList(uid)

	util.ResultJsonOk(w, GroupResponse{ createdGroupList, participantGroupList })
}

func (p *GroupController) FindByGroupId(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	groupIdStr := strings.Split(vars.Get("group_ids"), `,`)

	var groupIdList []uint

	for _, str := range groupIdStr {
		num, _ := strconv.Atoi(str)
		groupIdList = append(groupIdList, uint(num))
	}

	list := groupService.FindByGroupId(groupIdList)

	util.ResultJsonOk(w, list)
}

func (p *GroupController) create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &GroupRequest{}

	decoder.Decode(&request)

	uid := p.GetUserId(w, r)

	id := groupService.CreateGroup(uid, request.Name, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, CreateGroupResponse{id})
}
