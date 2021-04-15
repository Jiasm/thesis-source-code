package main

import (
	"controller"
	"fmt"
	"log"
	"net/http"
	_ "session"
	"time"
	"util"
)

func main() {

	util.InitDB()

	fmt.Println("server run as http://127.0.0.1:8080")

	server := &http.Server{
		Addr:        ":8080",
		Handler:     util.Router,
		ReadTimeout: 5 * time.Second,

	}
	RegiterRouter(util.Router)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func RegiterRouter(handler *util.RouterHandler) {
	new(controller.UserController).Router(handler)
	new(controller.TaskController).Router(handler)
	new(controller.ProjectController).Router(handler)
	new(controller.GroupController).Router(handler)
	new(controller.TaskGroupController).Router(handler)
	new(controller.TaskTagController).Router(handler)
	new(controller.TagController).Router(handler)
	new(controller.TaskCommentController).Router(handler)
}
