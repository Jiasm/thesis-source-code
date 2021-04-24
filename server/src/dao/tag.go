package dao

import (
	"entity"
	"fmt"
	"log"
	"strconv"
	"strings"
	"util"
)

type TagDao struct {
}

func (p *TagDao) Insert(tagName string, addDate uint) uint {
	result, err := util.DB.Exec("INSERT INTO `tag` (`text`, `add_date`) VALUES (?, ?)", tagName, addDate)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return uint(id)
}

func (p *TagDao) FindOne(tagName string) *entity.Tag {
	rows, err := util.DB.Query("SELECT id FROM `tag` WHERE text = ? LIMIT 1", tagName)

	if err != nil {
		log.Println(err)
		return nil
	}

	var tag entity.Tag

	if rows.Next() {
		rows.Scan(&tag.Id)

		rows.Close()
	}

	return &tag
}

func (p *TagDao) FindAll(tagIdList []uint) []entity.Tag {
	var tagIdListStr []string

	for _, item := range tagIdList {
		tagIdListStr = append(tagIdListStr, strconv.Itoa(int(item)))
	}

	rows, err := util.DB.Query(fmt.Sprintf("SELECT id, text FROM tag WHERE id IN (%s)", strings.Join(tagIdListStr, " , ")))

	if err != nil {
		log.Println(err)
		return nil
	}

	var tagList []entity.Tag

	for rows.Next() {
		var project entity.Tag
		rows.Scan(&project.Id, &project.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		tagList = append(tagList, project)
	}
	rows.Close()

	return tagList
}
