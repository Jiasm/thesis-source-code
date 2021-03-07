package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"util"
)

type ListController struct {
}

type CreateAccountRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	RoleId int `json:"role_id"`
	Status int `json:"status"`
}

var listService = new(service.ListService)

func (p *ListController)Router(router *util.RouterHandler)  {
	router.Router("/login",p.login)
	router.Router("/user/create",p.create)
}

func (p *ListController)login(w http.ResponseWriter,r *http.Request)  {
	decoder := json.NewDecoder(r.Body)

	var params map[string]string

	decoder.Decode(&params)

	user := listService.FindOne(params["username"], params["password"])

	if user == nil || user.ID == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, user)
}

func (p *ListController)create(w http.ResponseWriter,r *http.Request)  {
	decoder := json.NewDecoder(r.Body)

	request := &CreateAccountRequest{}

	decoder.Decode(&request)

	id := listService.CreateAccount(request.UserName, request.Password, request.RoleId, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}
