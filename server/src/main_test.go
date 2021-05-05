package main

import (
	"constant"
	"entity"
	"fmt"
	"testing"
	"dao"
	"time"
)

func Test_main(t *testing.T) {
	testGroup := entity.Group{1, "Niko", 0, 2, 1 }

	groupIdList := []uint{ 1, 2, 3 }

	projectIdList := []uint{ 1, 2, 3, 3, 4 }

	tagIdList := []uint{ 1, 2 }

	taskTitle := "changeTitle"

	taskDesc := "changeDesc"

	taskExecutor := "1"

	taskStatus := "2"

	taskExpireDate := fmt.Sprintf("%d", time.Now().Unix())

	t.Logf("get session key: %s", constant.KEY_USER)

	t.Logf("build group struct %v", testGroup)

	t.Logf("build group list sql with group id filter %v %s", groupIdList, dao.BuildFindGroupListWithGroupIdList(groupIdList))

	t.Logf("build project list sql with group id filter %v %s", projectIdList, dao.BuildFindProjectListWithProjectId(projectIdList))

	t.Logf("build tag list sql with group id filter %v %s", tagIdList, dao.BuildFindTagListWithTagId(tagIdList))

	t.Logf("build update task sql with title and desc %v %v %s", taskTitle, taskDesc, dao.BuildUpdateTaskFieldsSql(taskTitle, taskDesc, "", "", "", "", "", ""))

	t.Logf("build update task sql with status %v %s", taskStatus, dao.BuildUpdateTaskFieldsSql("", "", "", taskStatus, "", "", "", ""))

	t.Logf("build update task sql with executor %v %s", taskExecutor, dao.BuildUpdateTaskFieldsSql("", "", taskExecutor, "", "", "", "", ""))

	t.Logf("build update task sql with executor and expireDate %v %v %s", taskExecutor, taskExpireDate, dao.BuildUpdateTaskFieldsSql("", "", taskExecutor, "", taskExpireDate, "", "", ""))
}