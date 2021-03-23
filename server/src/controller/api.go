package controller

import (
	"constant"
	"entity"
	"net/http"
	"util"
)

type ApiController struct {}

func (p *ApiController) GetUserId(w http.ResponseWriter,r *http.Request) uint {
	user := p.GetUser(w,r)
	if user == nil {
		return 3
	}
	return user.ID
}

func (p *ApiController) GetUser(w http.ResponseWriter,r *http.Request) *entity.User {
	session := util.GlobalSession().SessionStart(w,r)
	if session == nil {
		return nil
	}
	key_user := session.Get(constant.KEY_USER)
	if user,ok := key_user.(*entity.User);ok{
		return user
	}
	return nil
}
