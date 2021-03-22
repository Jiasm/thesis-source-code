package controller

import (
	"constant"
	"encoding/json"
	"net/http"
	"service"
	"util"
)

type ListController struct {
}

type UserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	Status   int    `json:"status"`
}

var listService = new(service.ListService)

func (p *ListController) Router(router *util.RouterHandler) {
	router.Router("/login", p.login)
	router.Router("/logout", p.logout)
	router.Router("/user/create", p.create)
}

func (p *ListController) logout(w http.ResponseWriter, r *http.Request) {
	//session
	session := util.GlobalSession().SessionStart(w, r)

	session.Delete(constant.KEY_USER)

	util.ResultJsonOk(w, nil)
}

func (p *ListController) login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &UserRequest{}

	decoder.Decode(&request)

	user := listService.FindOne(request.UserName, request.Password)

	if user == nil || user.ID == 0 {
		util.ResultFail(w, "error")
		return
	}

	//session
	session := util.GlobalSession().SessionStart(w, r)
	session.Set(constant.KEY_USER, user.UserName)

	util.ResultJsonOk(w, user)
}

func (p *ListController) create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &UserRequest{}

	decoder.Decode(&request)

	id := listService.CreateAccount(request.UserName, request.Password, request.RoleId, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}
