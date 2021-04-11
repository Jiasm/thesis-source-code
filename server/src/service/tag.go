package service

import (
	"dao"
	"entity"
)

type TagService struct {
}

var tagDao = new(dao.TagDao)

func (p *TagService) FindAll(tagIdList []uint) []entity.Tag {
	tagList := tagDao.FindAll(tagIdList)

	return tagList
}
