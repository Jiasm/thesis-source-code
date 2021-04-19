package controller

import (
	"constant"
	"fmt"
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
	key_user := session.Get(constant.KEY_USER)

	fmt.Println("key_user", key_user)

	if uid,ok := key_user.(uint);ok{
		return uid
	}
	return 0
}
