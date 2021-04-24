package controller

import (
	"net/http"
	"service"
	"strconv"
	"util"
)


type InfoController struct {
}

var infoService = new(service.InfoService)

func (p *InfoController) Router(router *util.RouterHandler) {
	router.Router("/info/new-count-list", p.FindNewCountList)
	router.Router("/info/typed-count-list", p.FindTypedCountList)
	router.Router("/info/todo-count-list", p.FindTodoCount)
	router.Router("/info/done-count-list", p.FindDoneCount)
}

func (p *InfoController) FindNewCountList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	projectIdStr := vars.Get("project_id")
	projectId, _ := strconv.Atoi(projectIdStr)

	list := infoService.FindNewCountList(uint(projectId))

	util.ResultJsonOk(w, list)
}

func (p *InfoController) FindTypedCountList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	projectIdStr := vars.Get("project_id")
	projectId, _ := strconv.Atoi(projectIdStr)

	list := infoService.FindTypedCountList(uint(projectId))

	util.ResultJsonOk(w, list)
}

func (p *InfoController) FindTodoCount(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	projectIdStr := vars.Get("project_id")
	projectId, _ := strconv.Atoi(projectIdStr)

	list := infoService.FindTodoCount(uint(projectId))

	util.ResultJsonOk(w, list)
}

func (p *InfoController) FindDoneCount(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	projectIdStr := vars.Get("project_id")
	projectId, _ := strconv.Atoi(projectIdStr)

	list := infoService.FindDoneCount(uint(projectId))

	util.ResultJsonOk(w, list)
}
