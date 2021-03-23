package dao

import (
	"entity"
	"log"
	"util"
)

type TaskParticipantDao struct {
}


func (p *TaskParticipantDao) FindTaskIdByParticipant(uid uint) []uint {
	rows, err := util.DB.Query("SELECT task_id FROM task_participant WHERE uid = ?", uid)

	if err != nil {
		log.Println(err)
		return nil
	}

	var taskIdList []uint

	for rows.Next() {
		var taskParticipant entity.TaskParticipant
		rows.Scan(&taskParticipant.TaskId)
		if err != nil {
			log.Println(err)
			continue
		}
		taskIdList = append(taskIdList, taskParticipant.TaskId)
	}
	rows.Close()

	return taskIdList
}
