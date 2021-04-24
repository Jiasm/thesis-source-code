package controller

import (
	"constant"
	"net/http"
	"util"
)

type ApiController struct {}

func (p *ApiController) GetUserId(w http.ResponseWriter,r *http.Request) uint {
	uid := p.GetUser(w,r)
	if uid == 0 {
		return 3
	}

	return uid
}

func (p *ApiController) GetUser(w http.ResponseWriter,r *http.Request) uint {
	session := util.GlobalSession().SessionStart(w,r)
	if session == nil {
		return 0
	}
	keyUid := session.Get(constant.KEY_USER)

	if uid,ok := keyUid.(uint);ok{
		return uid
	}
	return 0
}
