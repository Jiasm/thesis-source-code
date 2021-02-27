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
	userId := uint(123)

	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	id := listService.Insert(userId,title,content)
	if id <= 0{
		util.ResultFail(w,"feedback fail")
		return
	}
	util.ResultOk(w,"feedback success")
}
