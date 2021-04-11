package controller

import (
	"constant"
	"encoding/json"
	"net/http"
	"service"
	"strconv"
	"strings"
	"util"
)

type UserController struct {
}

type UserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Status   uint   `json:"status"`
}

type UserResponse struct {
	Id uint `json:"id"`
	UserName string `json:"username"`
}

type UserListResponse struct {
	List []UserResponse `json:"list"`
}

var userService = new(service.UserService)

func (p *UserController) Router(router *util.RouterHandler) {
	router.Router("/login", p.login)
	router.Router("/logout", p.logout)
	router.Router("/user/create", p.create)
	router.Router("/user/list", p.GetList)
}

func (p *UserController) logout(w http.ResponseWriter, r *http.Request) {
	//session
	session := util.GlobalSession().SessionStart(w, r)

	session.Delete(constant.KEY_USER)

	util.ResultJsonOk(w, nil)
}

func (p *UserController) login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &UserRequest{}

	decoder.Decode(&request)

	user := userService.FindOne(request.UserName, request.Password)

	if user == nil || user.ID == 0 {
		util.ResultFail(w, "error")
		return
	}

	//session
	session := util.GlobalSession().SessionStart(w, r)
	session.Set(constant.KEY_USER, user.UserName)

	util.ResultJsonOk(w, user)
}

func (p *UserController) create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := &UserRequest{}

	decoder.Decode(&request)

	id := userService.CreateAccount(request.UserName, request.Password, request.RoleId, request.Status)

	if id == 0 {
		util.ResultFail(w, "error")
		return
	}

	util.ResultJsonOk(w, id)
}

func (p *UserController) GetList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	userIdStr := strings.Split(vars.Get("user_ids"), `,`)

	var userIdList []uint

	for _, str := range userIdStr {
		num, _ := strconv.Atoi(str)
		userIdList = append(userIdList, uint(num))
	}

	userList := userService.FindAll(userIdList)

	var responseList []UserResponse

	for _, user := range userList {
		responseList = append(responseList, UserResponse{user.ID, user.UserName})
	}

	util.ResultJsonOk(w, UserListResponse{ responseList })
}
