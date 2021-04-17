package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type TaskTagDao struct {
}

func (p *TaskTagDao) Insert(taskId, tagId uint) uint {
	result, err := util.DB.Exec("INSERT INTO `task_tag` (`task_id`, `tag_id`) VALUES (?, ?)", taskId, tagId)
	if err != nil {
		log.Println(err)
		return 0
	}
	changedCount, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(changedCount)
}

func (p *TaskTagDao) Remove(taskId, tagId uint) uint {
	result, err := util.DB.Exec("DELETE FROM `task_tag` WHERE `task_id` = ? AND `tag_id` = ?", taskId, tagId)
	if err != nil {
		log.Println(err)
		return 0
	}
	changedCount, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(changedCount)
}

func (p *TaskTagDao) FindAll(taskIdList []uint) []entity.TaskTag {
	var taskIdListStr []string

	for _, item := range taskIdList {
		taskIdListStr = append(taskIdListStr, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query(fmt.Sprintf("SELECT task_id, tag_id FROM task_tag WHERE task_id IN (%s)", strings.Join(taskIdListStr, " , ")))

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskTagList []entity.TaskTag

	for rows.Next() {
		fmt.Println("join")
		var project entity.TaskTag
		rows.Scan(&project.TaskId, &project.TagId)
		if err != nil {
			log.Println(err)
			continue
		}
		taskTagList = append(taskTagList, project)
	}
	rows.Close()

	return taskTagList
}
