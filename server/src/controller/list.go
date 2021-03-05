package controller

import (
	"net/http"
	"service"
	"util"
)

type ListController struct {
}

var listService = new(service.ListService)

func (p *ListController)Router(router *util.RouterHandler)  {
	router.Router("/list",p.list)
}

func (p *ListController)list(w http.ResponseWriter,r *http.Request)  {
	id := listService.Insert("test insert","test password", 1, 1)
	if id <= 0{
		util.ResultFail(w,"feedback fail")
		return
	}
	util.ResultOk(w,"feedback success")
}
