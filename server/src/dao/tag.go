package dao

import (
	"entity"
	"log"
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
