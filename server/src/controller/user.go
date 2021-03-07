package controller

import (
	"encoding/json"
	"net/http"
	"service"
	"util"
)

type ListController struct {
}

var listService = new(service.ListService)

func (p *ListController)Router(router *util.RouterHandler)  {
	router.Router("/login",p.login)
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
