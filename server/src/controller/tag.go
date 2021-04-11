package controller

import (
	"entity"
	"net/http"
	"service"
	"strconv"
	"strings"
	"util"
)


type TagController struct {
}

type TagListResponse struct {
	List []entity.Tag `json:"list"`
}

var tagService = new(service.TagService)

func (p *TagController) Router(router *util.RouterHandler) {
	router.Router("/tag/list", p.FindTagList)
}

func (p *TagController) FindTagList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	tagIdStr := strings.Split(vars.Get("tag_ids"), `,`)

	var tagIdList []uint

	for _, str := range tagIdStr {
		num, _ := strconv.Atoi(str)
		tagIdList = append(tagIdList, uint(num))
	}

	tagList := tagService.FindAll(tagIdList)

	util.ResultJsonOk(w, TagListResponse{ tagList })
}