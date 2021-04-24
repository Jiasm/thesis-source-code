package dao

import (
	"entity"
	"fmt"
	"log"
	"util"
)

type InfoDao struct {
}

func (p *InfoDao) FindNewCountList(projectId uint) []entity.NewTask {
	rows, err := util.DB.Query("SELECT COUNT(*) AS new_count, FROM_UNIXTIME(created_date,'%Y%m%d') AS `date` FROM task WHERE task_project_id = ? GROUP BY `date` ORDER BY `date` DESC", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var list []entity.NewTask

	for rows.Next() {
		var item entity.NewTask
		rows.Scan(&item.NewCount, &item.Date)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, item)
	}
	rows.Close()

	return list
}

func (p *InfoDao) FindTypedCountList(projectId uint) []entity.TypedTask {
	rows, err := util.DB.Query("SELECT COUNT(*) AS type_count, type FROM task WHERE task_project_id = ? GROUP BY type", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var list []entity.TypedTask

	for rows.Next() {
		var item entity.TypedTask
		rows.Scan(&item.TypeCount, &item.Type)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, item)
	}
	rows.Close()

	return list
}


func (p *InfoDao) FindTodoCount(projectId uint) []entity.TodoCount {
	rows, err := util.DB.Query("SELECT COUNT(*) AS todo_count, executor, status FROM task WHERE task_project_id = ? AND (status = 0 OR status = 1) GROUP BY executor, status", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var list []entity.TodoCount

	for rows.Next() {
		fmt.Println("join")
		var item entity.TodoCount
		rows.Scan(&item.TodoCount, &item.Executor, &item.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, item)
	}
	rows.Close()

	return list
}


func (p *InfoDao) FindDoneCount(projectId uint) []entity.DoneCount {
	rows, err := util.DB.Query("SELECT COUNT(*) AS done_count, executor FROM task WHERE task_project_id = ? AND status = 2 GROUP BY executor", projectId)

	if err != nil {
		log.Println(err)
		return nil
	}

	var list []entity.DoneCount

	for rows.Next() {
		fmt.Println("join")
		var item entity.DoneCount
		rows.Scan(&item.DoneCount, &item.Executor)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, item)
	}
	rows.Close()

	return list
}

